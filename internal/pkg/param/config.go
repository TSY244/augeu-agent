package param

type Config struct {
	Mode          string // 启动方式
	ConfigPath    string // 配置文件路径
	Secret        string // 密钥
	RemoteAddr    string // 远程地址
	WebsocketAddr string // websocket地址
	WeiBuApiKey   string // 微步api密钥
}

var (
	BaseConfig = Config{}
)
