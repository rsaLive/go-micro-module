package config

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"fmt"
	"github.com/oa-meeting/pkg/msg"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

var Data = new(AppConfig)

type AppConfig struct {
	System struct {
		Mode         string
		ConfigSource string `mapstructure:"config_source"`
	}
	OaMeeting struct {
		Host     string
		Port     int32
		User     string
		Password string
		DbName   string `mapstructure:"db_name"`
	}
	AliYun struct {
		AccessKeyId     string `mapstructure:"access_key_id"`
		AccessKeySecret string `mapstructure:"access_key_secret"`
	}
	AliYunRtc struct {
		Appid  string `mapstructure:"appid"`
		AppKey string `mapstructure:"app_key"`
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
		panic("viper.ReadInConfig failed" + err.Error())
		return
	}
	if err = viper.Unmarshal(Data); err != nil {
		panic("viper.Unmarshal failed" + err.Error())
		return
	}
	// 如果是configmap模式再修改
	fmt.Println(Data.System)
	if Data.System.ConfigSource == "configmap" {
		traverseFields(reflect.ValueOf(*Data), "", Data)
	}
	return
}
func traverseFields(value reflect.Value, prefix string, configPtr interface{}) {
	valueType := value.Type()
	prefixEnv := "${"
	suffixEnv := "}"
	// 遍历结构体的字段
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldValue := value.Field(i)
		// 拼接字段名（带有前缀）
		fieldName := prefix + field.Name
		// 判断字段的类型
		if fieldValue.Kind() == reflect.Struct {
			// 递归遍历嵌套结构体字段
			traverseFields(fieldValue, fieldName+".", configPtr)
		} else {
			// 获取字段的值
			fieldValueStr := fmt.Sprintf("%v", fieldValue.Interface())
			// 判断是不是需要通过环境变量获取
			if len(fieldValueStr) > 3 && strings.HasPrefix(fieldValueStr, prefixEnv) && strings.HasSuffix(fieldValueStr, suffixEnv) {
				end := len(fieldValueStr) - len(suffixEnv)
				var hasDefault bool
				if strings.Index(fieldValueStr, "|") > 0 {
					hasDefault = true
					end = strings.Index(fieldValueStr, "|")
				}
				envStr := fieldValueStr[len(prefixEnv):end]
				getValue := os.Getenv(envStr)
				if getValue == "" && hasDefault {
					getValue = fieldValueStr[end+1 : len(fieldValueStr)-len(suffixEnv)]
				}
				setSubFieldValue(configPtr, fieldName, getValue)
			}
		}
	}
}

func setSubFieldValue(configPtr interface{}, fieldPath string, newValue interface{}) {
	value := reflect.ValueOf(configPtr).Elem()
	fields := strings.Split(fieldPath, ".")
	for _, field := range fields {
		value = value.FieldByName(field)
		if !value.IsValid() {
			return // 字段不存在，直接返回
		}
		if value.Kind() == reflect.Ptr {
			value = value.Elem() // 解引用指针类型的字段
		}
	}
	// 检查字段是否可设置
	if value.CanSet() {
		// 根据字段类型，将新值转换为对应类型并设置字段的值
		newValue := reflect.ValueOf(newValue).Convert(value.Type())
		value.Set(newValue)
	}
}
