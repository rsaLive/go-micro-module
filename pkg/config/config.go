package config

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"fmt"
	"github.com/oa-meeting/pkg/msg"
	"github.com/spf13/viper"
	"os"
)

var Data = new(AppConfig)

type AppConfig struct {
	System struct {
		Mode string
	}
	MealOrder struct {
		Host     string
		Port     int32
		User     string
		Password string
		DbName   string `mapstructure:"db_name"`
	}
	Redis struct {
		DB       int
		Addr     string
		Password string
	}
	ZapLog struct {
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	}
	SnowFlake struct {
		NodeNum   int32  `mapstructure:"node_num"`
		StartTime string `mapstructure:"start_time"`
	}
	Jaeger struct {
		Addr string `mapstructure:"host"`
		Open bool   `mapstructure:"open"`
	}
	RabbitMq struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int32  `mapstructure:"port"`
		Vhost    string `mapstructure:"vhost"`
	}
	WxMini struct {
		AppId  string `mapstructure:"appid"`
		Secret string `mapstructure:"secret"`
	}
	WxPay struct {
		MchID                      string `mapstructure:"mchid"`
		MchAPIv3Key                string `mapstructure:"apiv3key"`
		MchCertificateSerialNumber string `mapstructure:"mch_certificate_serial_number"`
		ApiClientCerPenPath        string `mapstructure:"api_client_cert_pem_path"`
		NotifyUrl                  string `mapstructure:"notify_url"`
	}
	HttpServer struct {
		AdminMainServer string `mapstructure:"admin_main_server"`
		MinusAppoint    string `mapstructure:"minus_appoint"`
	}
	Feie struct {
		Sn string `mapstructure:"sn"`
	}
}

func GetConf() (iniConf string, err error) {
	if os.Getenv(msg.MODE_ENV) != "" {
		if err = os.Setenv(constant.ConfigFileEnvKey, fmt.Sprintf("./conf/%s/%s", os.Getenv(msg.MODE_ENV), msg.SERVER_DUBBOGO_CONFIG)); err != nil {
			return
		}
	}
	if os.Getenv(msg.MODE_ENV) == "" {
		iniConf = fmt.Sprintf("../conf/%s", msg.SERVER_CONFIG)
	} else {
		iniConf = fmt.Sprintf("./conf/%s/%s", os.Getenv(msg.MODE_ENV), msg.SERVER_CONFIG)
	}
	return
}

func GetOptions() {
	iniConf, err := GetConf()
	if err != nil {
		panic("GetOptions err" + err.Error())
	}
	if err = Viper(iniConf); err != nil {
		return
	}
}

func Viper(iniConf string) (err error) {
	viper.SetConfigFile(iniConf)
	err = viper.ReadInConfig()
	if err != nil {
		//panic("viper.ReadInConfig failed" + err.Error())
		return
	}
	if err = viper.Unmarshal(Data); err != nil {
		//panic("viper.Unmarshal failed" + err.Error())
		return
	}
	return
}
