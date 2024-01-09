package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"jellylfbackend/ctrl"
	"jellylfbackend/global"
	"jellylfbackend/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pflag.Parse()
	initialize.InitApp(config)

	app := gin.Default()
	ctrl.Setup(app)
	_ = app.Run(fmt.Sprintf(`:%v`, global.Conf.HTTPPort))

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigterm
	}()
}

var (
	config string
)

func init() {
	pflag.StringVar(&config, "c", "conf/config.yaml", "choose config file.")
	// 配置文件路径
	pflag.Int("HTTP_PORT", 8818, "")
	// 数据库
	pflag.String("DB_DRIVER", "postgresql", "")
	pflag.String("DB_DSN", "postgres://postgres:123456@localhost:5432/public?sslmode=disable&TimeZone=Asia/Shanghai", "")
	pflag.Bool("DB_AUTO_MIGRATE", false, "")
}
