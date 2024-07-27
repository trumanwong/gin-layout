package conf

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"time"
)

type Http struct {
	Port uint64 `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Server struct {
	Http Http `yaml:"http"`
}

type Other struct {
	Robot        string   `yaml:"robot" mapstructure:"robot"`
	JwtKey       string   `yaml:"jwt_key" mapstructure:"jwt_key"`
	TemplatePath string   `yaml:"template_path" mapstructure:"template_path"`
	StoragePath  string   `yaml:"storage_path" mapstructure:"storage_path"`
	AllowOrigins []string `yaml:"allow_origins" mapstructure:"allow_origins"`
}

type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

type Redis struct {
	Addr         string        `yaml:"addr"`
	Password     string        `yaml:"password"`
	Prefix       string        `yaml:"prefix"`
	Db           uint8         `yaml:"db"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

type Data struct {
	Database Database `yaml:"database" mapstructure:"database"`
	Redis    Redis    `yaml:"redis"`
}

type RabbitMQ struct {
	Url string `yaml:"url"`
}

type Crypto struct {
	Enable         bool   `yaml:"enable"`
	AesKey         string `yaml:"aes_key" mapstructure:"aes_key"`
	AesIv          string `yaml:"aes_iv" mapstructure:"aes_iv"`
	PlainHeaderKey string `yaml:"plain_header_key" mapstructure:"plain_header_key"`
	PlainHeaderVal string `yaml:"plain_header_val" mapstructure:"plain_header_val"`
}

type Configs struct {
	Server   Server   `yaml:"server"`
	Data     Data     `yaml:"data"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq" mapstructure:"rabbitmq"`
	Crypto   Crypto   `yaml:"crypto"`
	Other    Other    `yaml:"other" mapstructure:"other"`
}

var c = flag.String("config", "", "配置文件路径")
var v *viper.Viper

func NewConfig() *Configs {
	flag.Parse()
	fileName := "config.yaml"
	if len(*c) > 0 {
		fileName = *c
	}

	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(fileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config := new(Configs)
	//监听配置文件的修改和变动
	v.OnConfigChange(func(e fsnotify.Event) {

	})
	v.WatchConfig()
	//将配置解析为Struct对象
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConfig)
