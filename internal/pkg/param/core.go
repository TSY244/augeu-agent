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
	flag.Parse()
}

func Help() {
	flag.Usage()
}
