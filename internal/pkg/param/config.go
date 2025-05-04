package param

type Config struct {
	Mode          string    // 启动方式
	ConfigPath    string    // 配置文件路径
	Secret        string    // 密钥
	RemoteAddr    string    // 远程地址
	WebsocketAddr string    // websocket地址
	WeiBuApiKey   string    // 微步api密钥
	Env           StrSplice // 传给rule 的环境变量
	Target        string    // ip或domain，监控网络信息
	IsStrong      bool      // 是否开启强监听模式
}

var (
	BaseConfig = Config{
		Env: NewStrSplice(),
	}
)
