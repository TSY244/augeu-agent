package windowsLog

var (
	LoginEvent = map[int64]string{
		4624: "成功登陆",
		4625: "登陆失败",
	}
	RdpEvent = map[int64]string{
		4778: "成功登陆",
		4777: "登陆失败",
	}
	ServiceEvent = map[int64]string{
		7045: "服务创建", // 只关注服务创建的日志
	}
	// 记录日志基本没有
	CreateProcessEvent = map[int64]string{
		4688: "进程创建",
	}
	//记录日志太少
	PowershellEvent = map[int64]string{
		//4103: "powershell 命令执行",
		//4104: "远程 powershell 命令执行",
		//400: "powershell 5.0 及以上 引擎启动",
	}
	// 测试的时候没有
	ReadLsassEvent = map[int64]string{
		4663: "对象访问审计",
	}
	// 记录日志太少
	SystemEvent = map[int64]string{
		4608: "windows 启动",
		4609: "windows 关闭",
	}
	// 用户日志
	UserEvent = map[int64]string{
		4726: "删除用户",
		4732: "改变用户组",
		4724: "修改用户密码",
		4722: "修改用户属性",
		4720: "创建用户",
	}
	DnsEvent = map[int64]string{
		22: "DNS 查询",
	}
	// 注册表
	RegistryEvent = map[int64]string{
		12: "注册表值创建",
		13: "注册表值重命名",
		14: "注册表值删除",
	}
)
