package initial

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wqh66886/past-present-future/common/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() {
	c := &define.Config{}
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(workDir + "/resources")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	define.Cfg = c
}

func InitMysql(c *define.Config) {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database)
	log.Println(url)
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	define.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	setPool(c)
}

func setPool(conf *define.Config) {
	sqlDb, err := define.DB.DB()
	if err != nil {
		log.Fatalln("连接池设置错误", err.Error())
	}
	sqlDb.SetMaxIdleConns(conf.Mysql.MaxPoolSize)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.Mysql.MaxLifeTime) * time.Second)
}
