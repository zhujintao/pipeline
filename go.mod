module github.com/banzaicloud/pipeline

replace github.com/qor/auth => github.com/banzaicloud/auth v0.0.0-20181025135209-1bff4585d99a

require (
	cloud.google.com/go v0.33.1
	contrib.go.opencensus.io/exporter/ocagent v0.2.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.8.0 // indirect
	github.com/Azure/azure-pipeline-go v0.0.0-20180507050906-098e490af5dc // indirect
	github.com/Azure/azure-sdk-for-go v22.2.2+incompatible
	github.com/Azure/azure-storage-blob-go v0.0.0-20180507052152-66ba96e49ebb
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest v11.2.8+incompatible
	github.com/BurntSushi/toml v0.3.0 // indirect
	github.com/DataDog/datadog-go v0.0.0-20180822151419-281ae9f2d895 // indirect
	github.com/Jeffail/gabs v1.1.1 // indirect
	github.com/Knetic/govaluate v3.0.0+incompatible // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Masterminds/semver v1.4.0
	github.com/Masterminds/sprig v2.14.1+incompatible
	github.com/Microsoft/go-winio v0.4.11 // indirect
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/PuerkitoBio/purell v1.1.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/SAP/go-hdb v0.13.2 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20180822052843-1c5a1c93a9c1
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20180615125516-36bf7aa2f916
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/aokoli/goutils v1.0.1
	github.com/araddon/gou v0.0.0-20190110011759-c797efecbb61 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/EventBus v0.0.0-20180315140547-d46933a94f05
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/aws/aws-sdk-go v1.16.11
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/banzaicloud/anchore-image-validator v0.0.0-20181204185657-bf9806201a4e
	github.com/banzaicloud/azure-aks-client v0.0.0-20181129100229-ddd6cb920f47
	github.com/banzaicloud/bank-vaults v0.0.0-20181228154154-8f1348d7821a
	github.com/banzaicloud/cicd-go v0.0.0-20181203135328-de6f3f91e590
	github.com/banzaicloud/go-gin-prometheus v0.0.0-20181204122313-8145dbf52419
	github.com/banzaicloud/logrus-runtime-formatter v0.0.0-20180617171254-12df4a18567f
	github.com/banzaicloud/prometheus-config v0.0.0-20181214142820-fc6ae4756a29
	github.com/bitly/go-hostpool v0.0.0-20171023180738-a3a6125de932 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/boombuler/barcode v1.0.0 // indirect
	github.com/briankassouf/jose v0.9.1 // indirect
	github.com/casbin/casbin v1.4.0
	github.com/casbin/gorm-adapter v0.0.0-20171006093545-e56c6daebd5e
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/census-instrumentation/opencensus-proto v0.1.0 // indirect
	github.com/centrify/cloud-golang-sdk v0.0.0-20180119173102-7c97cc6fde16 // indirect
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/chrismalek/oktasdk-go v0.0.0-20181212195951-3430665dfaa0 // indirect
	github.com/circonus-labs/circonus-gometrics v2.2.5+incompatible // indirect
	github.com/circonus-labs/circonusllhist v0.1.3 // indirect
	github.com/containerd/continuity v0.0.0-20181203112020-004b46473808 // indirect
	github.com/coreos/bbolt v1.3.0 // indirect
	github.com/coreos/go-oidc v2.0.0+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20181031085051-9002847aa142 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dancannon/gorethink v4.0.0+incompatible // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190111225525-2fea367d496d // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.0+incompatible
	github.com/dimchansky/utfbom v1.0.0 // indirect
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/docker/distribution v0.0.0-20180327202408-83389a148052 // indirect
	github.com/docker/docker v0.0.0-20170731201938-4f3616fb1c11 // indirect
	github.com/docker/go-connections v0.3.0 // indirect
	github.com/docker/go-units v0.3.2 // indirect
	github.com/docker/libcompose v0.4.0
	github.com/docker/spdystream v0.0.0-20170912183627-bc6354cbbc29 // indirect
	github.com/duosecurity/duo_api_golang v0.0.0-20190107154727-539434bf0d45 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20181111060418-2ce16c963a8a // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/evanphx/json-patch v3.0.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fullsailor/pkcs7 v0.0.0-20180613152042-8306686428a5 // indirect
	github.com/gammazero/deque v0.0.0-20180920172122-f6adf94963e4 // indirect
	github.com/gammazero/workerpool v0.0.0-20181230203049-86a96b5d5d92 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/cors v0.0.0-20170318125340-cf4846e6a636
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-ldap/ldap v3.0.0+incompatible // indirect
	github.com/go-openapi/jsonpointer v0.0.0-20180322222829-3a0015ad55fa // indirect
	github.com/go-openapi/jsonreference v0.0.0-20180322222742-3fb327e6747d // indirect
	github.com/go-openapi/spec v0.0.0-20180326232708-9acd88844bc1 // indirect
	github.com/go-openapi/swag v0.0.0-20180302192843-ceb469cb0fdf // indirect
	github.com/go-sql-driver/mysql v1.3.0 // indirect
	github.com/go-stomp/stomp v2.0.2+incompatible // indirect
	github.com/go-test/deep v1.0.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocql/gocql v0.0.0-20181124151448-70385f88b28b // indirect
	github.com/gogo/protobuf v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20180203143532-66deaeb636df // indirect
	github.com/golang/protobuf v1.2.0
	github.com/google/btree v0.0.0-20180124185431-e89373fe6b4a // indirect
	github.com/google/go-github v15.0.0+incompatible
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/googleapis/gnostic v0.1.0 // indirect
	github.com/goph/emperror v0.16.0
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v0.0.0-20160922145804-ca9ada445741
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/gosimple/slug v1.1.1 // indirect
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/consul v1.4.0 // indirect
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce // indirect
	github.com/hashicorp/go-cleanhttp v0.0.0-20171218145408-d5fe4b57a186 // indirect
	github.com/hashicorp/go-gcp-common v0.0.0-20180425173946-763e39302965 // indirect
	github.com/hashicorp/go-hclog v0.0.0-20190109152822-4783caec6f2e // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-memdb v0.0.0-20181108192425-032f93b25bec // indirect
	github.com/hashicorp/go-msgpack v0.0.0-20150518234257-fa3f63826f7c // indirect
	github.com/hashicorp/go-multierror v0.0.0-20171204182908-b7773ae21874 // indirect
	github.com/hashicorp/go-plugin v0.0.0-20181212150838-f444068e8f5a // indirect
	github.com/hashicorp/go-retryablehttp v0.0.0-20180531211321-3b087ef2d313 // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/go-sockaddr v0.0.0-20180320115054-6d291a969b86 // indirect
	github.com/hashicorp/go-version v1.1.0 // indirect
	github.com/hashicorp/memberlist v0.1.0 // indirect
	github.com/hashicorp/nomad v0.8.7 // indirect
	github.com/hashicorp/raft v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/hashicorp/vault v1.0.1
	github.com/hashicorp/vault-plugin-auth-alicloud v0.0.0-20181109180636-f278a59ca3e8 // indirect
	github.com/hashicorp/vault-plugin-auth-azure v0.0.0-20181207232528-4c0b46069a22 // indirect
	github.com/hashicorp/vault-plugin-auth-centrify v0.0.0-20180816201131-66b0a34a58bf // indirect
	github.com/hashicorp/vault-plugin-auth-gcp v0.0.0-20181210200133-4d63bbfe6fcf // indirect
	github.com/hashicorp/vault-plugin-auth-jwt v0.0.0-20181031195942-f428c7791733 // indirect
	github.com/hashicorp/vault-plugin-auth-kubernetes v0.0.0-20181130162533-091d9e5d5fab // indirect
	github.com/hashicorp/vault-plugin-secrets-ad v0.0.0-20181109182834-540c0b6f1f11 // indirect
	github.com/hashicorp/vault-plugin-secrets-alicloud v0.0.0-20181109181453-2aee79cc5cbf // indirect
	github.com/hashicorp/vault-plugin-secrets-azure v0.0.0-20181207232500-0087bdef705a // indirect
	github.com/hashicorp/vault-plugin-secrets-gcp v0.0.0-20180921173200-d6445459e80c // indirect
	github.com/hashicorp/vault-plugin-secrets-gcpkms v0.0.0-20181212182553-6cd991800a6d // indirect
	github.com/hashicorp/vault-plugin-secrets-kv v0.0.0-20181219175933-9dbe04db0e34 // indirect
	github.com/heptio/ark v0.9.3
	github.com/huandu/xstrings v1.0.0 // indirect
	github.com/imdario/mergo v0.3.6
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jeffchao/backoff v0.0.0-20140404060208-9d7fd7aa17f2 // indirect
	github.com/jefferai/jsonx v0.0.0-20160721235117-9cc31c3135ee // indirect
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/jinzhu/gorm v1.9.1
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/jinzhu/now v0.0.0-20180314132004-b7dfa9a24504
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/keybase/go-crypto v0.0.0-20181127160227-255a5089e85a // indirect
	github.com/lib/pq v0.0.0-20180327071824-d34b9ff171c2 // indirect
	github.com/mailru/easyjson v0.0.0-20180323154445-8b799c424f57 // indirect
	github.com/marstr/guid v1.1.0 // indirect
	github.com/mattbaird/elastigo v0.0.0-20170123220020-2fe47fd29e4b // indirect
	github.com/mattn/go-isatty v0.0.3 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/michaelklishin/rabbit-hole v1.4.0 // indirect
	github.com/microcosm-cc/bluemonday v0.0.0-20180327211928-995366fdf961
	github.com/miekg/dns v1.1.3 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/oracle/oci-go-sdk v2.0.0+incompatible
	github.com/ory-am/common v0.4.0 // indirect
	github.com/ory/dockertest v3.3.3+incompatible // indirect
	github.com/pascaldekloe/goe v0.0.0-20180627143212-57f6aae5913c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.2.0
	github.com/petar/GoLLRB v0.0.0-20130427215148-53be0d36a84c // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pierrec/lz4 v0.0.0-20181005164709-635575b42742 // indirect
	github.com/pkg/errors v0.8.0
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/pquerna/otp v1.1.0 // indirect
	github.com/prometheus/client_golang v0.9.0-pre1
	github.com/prometheus/common v0.0.0-20181126121408-4724e9255275
	github.com/qor/assetfs v0.0.0-20170713023933-ff57fdc13a14 // indirect
	github.com/qor/auth v0.0.0-20190103025640-46aae9fa92fa
	github.com/qor/mailer v0.0.0-20170814094430-1e6ac7106955 // indirect
	github.com/qor/middlewares v0.0.0-20170822143614-781378b69454 // indirect
	github.com/qor/qor v0.0.0-20180313040854-1fb0672178f1
	github.com/qor/redirect_back v0.0.0-20170907030740-b4161ed6f848 // indirect
	github.com/qor/render v0.0.0-20171201033449-63566e46f01b // indirect
	github.com/qor/responder v0.0.0-20160314063933-ecae0be66c1a // indirect
	github.com/qor/session v0.0.0-20170907035918-8206b0adab70
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/russross/blackfriday v1.5.1 // indirect
	github.com/ryanuber/go-glob v0.0.0-20160226084822-572520ed46db // indirect
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/sirupsen/logrus v1.1.0
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.0-20180319062004-c439c4fa0937 // indirect
	github.com/spf13/viper v1.3.0
	github.com/streadway/amqp v0.0.0-20181205114330-a314942b2fd9 // indirect
	github.com/stretchr/testify v1.2.2
	github.com/technosophos/moniker v0.0.0-20180509230615-a5dbd03a2245
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/tv42/httpunix v0.0.0-20150427012821-b75d8614f926 // indirect
	github.com/xiang90/probing v0.0.0-20160813154853-07dd2e8dfe18 // indirect
	github.com/xlab/handysort v0.0.0-20150421192137-fb3537ed64a1 // indirect
	go.opencensus.io v0.18.0 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/net v0.0.0-20181106065722-10aee1819953
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/time v0.0.0-20180314180208-26559e0f760e // indirect
	google.golang.org/api v0.0.0-20180910000450-7ca32eb868bf
	google.golang.org/genproto v0.0.0-20181109154231-b5d43981345b
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/gomail.v2 v2.0.0-20150902115704-41f357289737 // indirect
	gopkg.in/gorethink/gorethink.v4 v4.1.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
	gopkg.in/ory-am/dockertest.v2 v2.2.3 // indirect
	gopkg.in/square/go-jose.v2 v2.1.6 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	gotest.tools v2.2.0+incompatible // indirect
	k8s.io/api v0.0.0-20180712090710-2d6f90ab1293
	k8s.io/apiextensions-apiserver v0.0.0-20180328075702-9beab23b2663 // indirect
	k8s.io/apimachinery v0.0.0-20180621070125-103fd098999d
	k8s.io/apiserver v0.0.0-20180327065226-f4a9d3132586 // indirect
	k8s.io/client-go v0.0.0-20180718001006-59698c7d9724
	k8s.io/helm v2.11.0+incompatible
	k8s.io/kube-openapi v0.0.0-20180216212618-50ae88d24ede // indirect
	k8s.io/kubernetes v0.0.0-20181108064615-3290824d1c7b
	k8s.io/utils v0.0.0-20180208044234-258e2a2fa645 // indirect
	labix.org/v2/mgo v0.0.0-20140701140051-000000000287 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	layeh.com/radius v0.0.0-20190109000448-e6d9fd7a048a // indirect
	vbom.ml/util v0.0.0-20170409195630-256737ac55c4 // indirect
)
