package model

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/banzaicloud/bank-vaults/database"
	"github.com/banzaicloud/pipeline/config"
	"github.com/jinzhu/gorm"
	// blank import is used here for simplicity
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

var dbOnce sync.Once
var db *gorm.DB
var logger *logrus.Logger

// Simple init for logging
func init() {
	logger = config.Logger()
}

type logrusAdapter struct {
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func (*logrusAdapter) Print(values ...interface{}) {
	fields := logrus.Fields{}
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
		)

		//fields["source"] = values[1]
		fields["duration"] = fmt.Sprintf("%fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0)

		if level == "sql" {
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			fields["sql"] = sql
			logger.WithFields(fields).Print(strconv.FormatInt(values[5].(int64), 10) + " rows affected or returned")
		}
	}
}

func initDatabase() {
	dbName := viper.GetString("database.dbname")
	db = ConnectDB(dbName)
	db.SetLogger(&logrusAdapter{})
}

// GetDataSource returns with datasource by database name
func GetDataSource(dbName string) string {
	log := logger.WithFields(logrus.Fields{"action": "GetDataSource"})
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	role := viper.GetString("database.role")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dataSource := "@tcp(" + host + ":" + port + ")/" + dbName
	if role != "" {
		var err error
		dataSource, err = database.DynamicSecretDataSource("mysql", role+dataSource)
		if err != nil {
			log.Error("Database dynamic secret acquisition failed")
			panic(err.Error())
		}
	} else {
		dataSource = user + ":" + password + dataSource
	}
	return dataSource
}

// ConnectDB connects to GORM
func ConnectDB(dbName string) *gorm.DB {
	log := logger.WithFields(logrus.Fields{"action": "ConnectDB"})
	dataSource := GetDataSource(dbName)
	dataSource += "?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Error("Database connection failed")
		panic(err.Error())
	}
	database.LogMode(true)
	return database
}

//GetDB returns an initialized DB
func GetDB() *gorm.DB {
	dbOnce.Do(initDatabase)
	return db
}

//IsErrorGormNotFound returns gorm.ErrRecordNotFound
func IsErrorGormNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
