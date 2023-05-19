package nacos

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	remote "github.com/yoyofxteam/nacos-viper-remote"
)

var RemoteViper = &viper.Viper{}

func Init(url string, port uint64, namespace, group, dataId, auth, user, password string, configType string) error {
	viper.AutomaticEnv()
	config := initOption(url, port, namespace, group, dataId, auth, user, password)

	// init nacos
	return initNacos(config, configType)
}

func initNacos(conf *remote.Option, configType string) error {
	// check whether in nacos
	if conf == nil {
		return errors.New("empty nacos config")
	}
	RemoteViper = viper.New()
	// 配置 Viper for Nacos 的远程仓库参数
	remote.SetOptions(conf)

	err := RemoteViper.AddRemoteProvider("nacos", conf.Url, "nacos")
	if err != nil {
		return err
	}
	RemoteViper.SetConfigType(configType)

	if err := RemoteViper.ReadRemoteConfig(); err != nil {
		return err
	} //sync get remote configs to RemoteViper instance memory . for example , RemoteViper.GetString(key)
	if err := RemoteViper.WatchRemoteConfigOnChannel(); err != nil {
		return err
	} //异步监听Nacos中的配置变化，如发生配置更改，会直接同步到 viper实例中。
	return nil
}

// initOption：初始化option
func initOption(url string, port uint64, namespace, group, dataId, auth, user, password string) *remote.Option {
	return &remote.Option{
		Url:         url,                           // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
		Port:        port,                          // nacos server端口号
		NamespaceId: namespace,                     // nacos namespace
		GroupName:   group,                         // nacos group
		Config:      remote.Config{DataId: dataId}, // nacos DataID
		Auth: func() *remote.Auth {
			if auth == "false" {
				return &remote.Auth{
					Enable: false,
				}
			}
			return &remote.Auth{ // 如果需要验证登录,需要此参数
				Enable:   true,
				User:     user,
				Password: password,
			}
		}(),
	}
}

func Get(key string) string {
	return RemoteViper.GetString(key)
}

func GetInt(key string) int {
	return RemoteViper.GetInt(key)
}

func GetBool(key string) bool {
	return RemoteViper.GetBool(key)
}
