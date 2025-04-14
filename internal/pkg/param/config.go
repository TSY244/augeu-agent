package param

type Config struct {
	Mode       string // 启动方式
	ConfigPath string // 配置文件路径
}

var (
	BaseConfig = Config{}
)
