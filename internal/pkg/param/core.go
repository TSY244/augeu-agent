package param

import "flag"

const (
	ModeUsage = `加载方式: basic, remote, local
	basic: 从程序内部加载默认规则
	remote: 加载远端提供的配置文件, -cp 参数需要提供一个远程url
	local: 从本地加载一个配置文件, -cp 需要提供一个本地文件路径
	`
	ConfigPathUsage = `配置文件路径`
)

func Init() {
	flag.StringVar(&BaseConfig.Mode, "mode", "basic", ModeUsage)
	flag.StringVar(&BaseConfig.ConfigPath, "cp", "", ConfigPathUsage)
	flag.StringVar(&BaseConfig.Secret, "s", "", "server 提供的 secret")
	flag.StringVar(&BaseConfig.RemoteAddr, "r", "", "server api 地址，请注意加上/api/v1，举个栗子：http://127.0.0.1/api/v1")
	flag.StringVar(&BaseConfig.WebsocketAddr, "ws", "", "server 提供的websocket 地址，注意使用ws://开头")
	flag.StringVar(&BaseConfig.WeiBuApiKey, "wbapi", "", "微步api密钥")
	flag.Parse()
}

func Help() {
	flag.Usage()
}
