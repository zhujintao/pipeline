// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"crypto/x509/pkix"
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/banzaicloud/bank-vaults/pkg/tls"
	"github.com/banzaicloud/pipeline/auth"
	pipConfig "github.com/banzaicloud/pipeline/config"
	"github.com/banzaicloud/pipeline/internal/global"
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgHelm "github.com/banzaicloud/pipeline/pkg/helm"
	pkgSecret "github.com/banzaicloud/pipeline/pkg/secret"
	"github.com/banzaicloud/pipeline/secret"
	"github.com/ghodss/yaml"
	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"k8s.io/api/core/v1"
)

type ingressControllerValues struct {
	Traefik traefikValues `json:"traefik"`
}

type traefikValues struct {
	SSL         sslTraefikValues `json:"ssl"`
	Affinity    v1.Affinity      `json:"affinity"`
	Tolerations []v1.Toleration  `json:"tolerations"`
	Service     traefikService   `json:"service,omitempty"`
}

type traefikService struct {
	Annotations traefikAnnotations `json:"annotations"`
}

type traefikAnnotations map[string]string

type sslTraefikValues struct {
	Enabled        bool     `json:"enabled"`
	GenerateTLS    bool     `json:"generateTLS"`
	DefaultCN      string   `json:"defaultCN"`
	DefaultSANList []string `json:"defaultSANList"`
	DefaultCert    string   `json:"defaultCert"`
	DefaultKey     string   `json:"defaultKey"`
}

const (
	DefaultCertSecretName          = "default-ingress-cert"
	askLoadBalancerSvcASAnnotation = "service.beta.kubernetes.io/azure-load-balancer-mode"
)

// InstallIngressControllerPostHook post hooks can't return value, they can log error and/or update state?
func InstallIngressControllerPostHook(cluster CommonCluster) error {
	defaultCertSecret, err := secret.Store.GetByName(cluster.GetOrganizationId(), DefaultCertSecretName)
	if err == secret.ErrSecretNotExists {
		certGenerator := global.GetCertGenerator()

		orgID := cluster.GetOrganizationId()
		organization, err := auth.GetOrganizationById(orgID)
		if err != nil {
			return emperror.WrapWith(err, "failed to get organization", "organizationId", orgID)
		}

		orgDomainName := fmt.Sprintf("%s.%s", organization.Name, viper.GetString(pipConfig.DNSBaseDomain))
		wildcardOrgDomainName := fmt.Sprintf("*.%s", orgDomainName)

		certRequest := tls.ServerCertificateRequest{
			Subject: pkix.Name{
				CommonName: wildcardOrgDomainName,
			},
			DNSNames: []string{orgDomainName, wildcardOrgDomainName},
		}

		rootCA, cert, key, err := certGenerator.GenerateServerCertificate(certRequest)
		if err != nil {
			return errors.Wrap(err, "failed to generate certificate")
		}

		defaultCertSecretRequest := &secret.CreateSecretRequest{
			Name: DefaultCertSecretName,
			Type: pkgSecret.TLSSecretType,
			Values: map[string]string{
				pkgSecret.CACert:     string(rootCA),
				pkgSecret.ServerCert: string(cert),
				pkgSecret.ServerKey:  string(key),
			},
			Tags: []string{
				pkgSecret.TagBanzaiReadonly,
			},
		}

		secretId, err := secret.Store.Store(cluster.GetOrganizationId(), defaultCertSecretRequest)
		if err != nil {
			return errors.Wrap(err, "failed to save generated certificate")
		}

		defaultCertSecret, err = secret.Store.Get(cluster.GetOrganizationId(), secretId)
		if err != nil {
			return errors.Wrap(err, "failed to load generated certificate")
		}
	} else if err != nil {
		return errors.Wrap(err, "failed to check default ingress cert existence")
	}

	ingressValues := ingressControllerValues{
		Traefik: traefikValues{
			SSL: sslTraefikValues{
				Enabled:     true,
				DefaultCert: base64.StdEncoding.EncodeToString([]byte(defaultCertSecret.Values[pkgSecret.ServerCert])),
				DefaultKey:  base64.StdEncoding.EncodeToString([]byte(defaultCertSecret.Values[pkgSecret.ServerKey])),
			},
			Affinity:    getHeadNodeAffinity(cluster),
			Tolerations: getHeadNodeTolerations(),
		},
	}

	if cluster.GetDistribution() == pkgCluster.AKS {
		asName, err := getAKSIngressServiceAvailabilitySetName(cluster)
		if err != nil {
			return emperror.Wrap(err, "could not get AKS availability set for ingress affinity")
		}
		if asName != "" {
			log.WithField("asName", asName).Debug("use specific availabilitySet for ingress on AKS")
			ingressValues.Traefik.Service = traefikService{
				Annotations: traefikAnnotations{
					askLoadBalancerSvcASAnnotation: asName,
				},
			}
		}
	}

	ingressValuesJson, err := yaml.Marshal(ingressValues)
	if err != nil {
		return emperror.Wrap(err, "converting ingress config to json failed")
	}

	namespace := viper.GetString(pipConfig.PipelineSystemNamespace)

	return installDeployment(cluster, namespace, pkgHelm.BanzaiRepository+"/pipeline-cluster-ingress", "ingress", ingressValuesJson, "", false)
}

func getAKSIngressServiceAvailabilitySetName(cluster CommonCluster) (string, error) {
	var asName string

	headNodePoolName := viper.GetString(pipConfig.PipelineHeadNodePoolName)

	r, err := regexp.Compile("(\\d+)-\\d+$")
	if err != nil {
		return asName, emperror.Wrap(err, "could not compile regex")
	}

	var matches []string
	nodeNames, err := cluster.ListNodeNames()
	if err != nil {
		return asName, emperror.Wrap(err, "cloud not get node names")
	}
	if len(nodeNames[headNodePoolName]) > 0 {
		matches = r.FindStringSubmatch(nodeNames[headNodePoolName][0])
		if len(matches) == 2 {
			asName = fmt.Sprintf("%s-availabilitySet-%s", headNodePoolName, matches[1])
			return asName, nil
		}
	}

	return asName, nil
}
