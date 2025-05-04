package param

import "flag"

const (
	ModeUsage = `加载方式: basic, remoteFile, remoteApi, local, monitor
	basic: 从程序内部加载默认规则
	remoteFile: 加载远端提供的配置文件, -cp 参数需要提供一个远程url
	remoteApi: 加载远端api 的配置文件
	local: 从本地加载一个配置文件, -cp 需要提供一个本地文件路径
	monitor: 监听指定的网络请求，tcp和udp
	`
	ConfigPathUsage = `配置文件路径`

	EnvUsage = `环境变量，格式为：key:value
				用法通过-env 多次添加
				-env key1:value1 -env key2:value2
	`
	IsStrongUsage = `是否开启强检测，默认不开启，如果开启当有重复的连接则会多输出。假设使用了socket 那么将会一直提示。`
)

func Init() {
	flag.StringVar(&BaseConfig.Mode, "mode", "basic", ModeUsage)
	flag.StringVar(&BaseConfig.ConfigPath, "cp", "", ConfigPathUsage)
	flag.StringVar(&BaseConfig.Secret, "s", "", "server 提供的 secret")
	flag.StringVar(&BaseConfig.RemoteAddr, "r", "", "server api 地址，请注意加上/api/v1，举个栗子：http://127.0.0.1/api/v1")
	flag.StringVar(&BaseConfig.WebsocketAddr, "ws", "", "server 提供的websocket 地址，注意使用ws://开头")
	flag.StringVar(&BaseConfig.WeiBuApiKey, "wbapi", "", "微步api密钥")
	flag.StringVar(&BaseConfig.Target, "target", "", "监听的ip或者是域名")
	flag.Var(&BaseConfig.Env, "env", EnvUsage)
	flag.BoolVar(&BaseConfig.IsStrong, "strong", false, IsStrongUsage)
	flag.Parse()
}

func Help() {
	flag.Usage()
}
