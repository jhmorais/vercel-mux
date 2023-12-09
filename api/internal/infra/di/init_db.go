package di

import (
	"fmt"
	"myapp/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitGormMysqlDB() (*gorm.DB, error) {
	config.LoadServerEnvironmentVars()

	dsn := fmt.Sprintf("%s:%s@%s", config.GetMysqlUser(), config.GetMysqlPassword(), config.GetMysqlConnectionString())

	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cardbycash.",
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return mysqlDb, err
}
