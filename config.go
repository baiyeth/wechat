package wechat

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

//Configuration ...
type Configuration struct {
	Official    OfficialConfiguration    `mapstructure:"official" json:"official"`
	MiniProgram MiniProgramConfiguration `mapstructure:"miniprogram" json:"miniprogram"`
}

//OfficialConfiguration ...
type OfficialConfiguration struct {
	AppId              string                                        `mapstructure:"app-id" json:"appId"`
	AppSecret          string                                        `mapstructure:"app-secret" json:"appSecret"`
	Encoding           string                                        `mapstructure:"encoding" json:"encoding"`
	TplMessageCronTask OfficialTplMessageCronTaskConfiguration `mapstructure:"tpl-message-cron-task" json:"tplMessageCronTask"`
}

//OfficialTplMessageCronTaskConfiguration ...
type OfficialTplMessageCronTaskConfiguration struct {
	Expr                string `mapstructure:"expr" json:"expr"`
	Users               string `mapstructure:"users" json:"users"`
	TemplateId          string `mapstructure:"template-id" json:"templateId"`
	MiniProgramAppId    string `mapstructure:"mini-program-app-id" json:"miniProgramAppId"`
	MiniProgramPagePath string `mapstructure:"mini-program-page-path" json:"miniProgramPagePath"`
}

//MiniProgramConfiguration ...
type MiniProgramConfiguration struct {
	AppId     string `mapstructure:"app-id" json:"appId"`
	AppSecret string `mapstructure:"app-secret" json:"appSecret"`
	Encoding  string `mapstructure:"encoding" json:"encoding"`
}

var (
	Conf Configuration
)

const (
	configType   = "yml"
	wechatConfig = "../conf/wechat.yml"
)

// Config 初始化配置文件
func init() {
	// 获取viper实例(可创建多实例读取多个配置文件, 这里不做演示)
	v := viper.New()
	readConfig(v, wechatConfig)
	// 将default中的配置全部以默认配置写入
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}
	// 转换为结构体
	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Sprintf("初始化配置文件失败: %v, wechat configure: %s", err, Conf))
	}
	fmt.Println("初始化配置文件完成, wechat configure: ", Conf)
}

func readConfig(v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v, wechat configure: %s", err, configFile))
	}
	// 加载配置
	if err = v.ReadConfig(bytes.NewReader(config)); err != nil {
		panic(fmt.Sprintf("加载配置文件失败: %v, wechat configure: %s", err, configFile))
	}
}