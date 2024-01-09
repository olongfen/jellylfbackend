package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"

	"jellylfbackend/global"
	"jellylfbackend/model"
)

func InitApp(configPath string) {
	initConf(configPath)
	initLog()
	initDB()
}

func initConf(configPath string) {
	// 设置默认值
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("RESOURCE", "resource")

	// 读取环境变量
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	// 读取终端变量
	_ = viper.BindPFlags(pflag.CommandLine)

	if err := viper.Unmarshal(&global.Conf); err != nil {
		log.Fatalln(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err := viper.Unmarshal(&global.Conf); err != nil {
			log.Fatal(err)
		}
	})
}

func initLog() {
	global.Log = global.NewProduceLogger()
}

func initDB() {
	open, err := gorm.Open(postgres.Open(global.Conf.DBDsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = open.AutoMigrate(&model.Person{}, &model.Work{}, &model.JobExperience{})
	if err != nil {
		log.Fatalln(err)
	}
	global.DB = open
}
