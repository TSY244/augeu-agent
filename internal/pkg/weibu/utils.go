package weibu

import (
	"fmt"
)

// convert2WeiBuDataEnglish 函数将 weiBuResponse 的信息以 WeiBuData 中的字段为键进行打印，并返回信息字符串
func convert2WeiBuDataEnglish(src *weiBuResponse) string {
	var info string

	// 打印 ResponseCode 和 VerboseMsg
	info += fmt.Sprintf("ResponseCode:%s -> %d\n", "ResponseCode", src.ResponseCode)
	info += fmt.Sprintf("VerboseMsg:%s -> %s\n", "VerboseMsg", src.VerboseMsg)

	// 打印 Summary 信息
	summary := src.Data.Summary
	info += fmt.Sprintf("Summary.ThreatLevel:%s -> %s\n", "ThreatLevel", summary.ThreatLevel)
	info += fmt.Sprintf("Summary.MalwareType:%s -> %s\n", "MalwareType", summary.MalwareType)
	info += fmt.Sprintf("Summary.MalwareFamily:%s -> %s\n", "MalwareFamily", summary.MalwareFamily)
	info += fmt.Sprintf("Summary.IsWhitelist:%s -> %v\n", "IsWhitelist", summary.IsWhitelist)
	info += fmt.Sprintf("Summary.SubmitTime:%s -> %s\n", "SubmitTime", summary.SubmitTime)
	info += fmt.Sprintf("Summary.FileName:%s -> %s\n", "FileName", summary.FileName)
	info += fmt.Sprintf("Summary.FileType:%s -> %s\n", "FileType", summary.FileType)
	info += fmt.Sprintf("Summary.SampleSha256:%s -> %s\n", "SampleSha256", summary.SampleSha256)
	info += fmt.Sprintf("Summary.MD5:%s -> %s\n", "MD5", summary.MD5)
	info += fmt.Sprintf("Summary.SHA1:%s -> %s\n", "SHA1", summary.SHA1)
	info += fmt.Sprintf("Summary.ThreatScore:%s -> %d\n", "ThreatScore", summary.ThreatScore)
	info += fmt.Sprintf("Summary.SandboxType:%s -> %s\n", "SandboxType", summary.SandboxType)
	info += fmt.Sprintf("Summary.MultiEngines:%s -> %s\n", "MultiEngines", summary.MultiEngines)

	// 打印 Summary.Tag 信息
	for i, s := range summary.Tag.S {
		info += fmt.Sprintf("Summary.Tag.S[%d]:%s -> %s\n", i, "S", s)
	}
	for i, x := range summary.Tag.X {
		info += fmt.Sprintf("Summary.Tag.X[%d]:%s -> %s\n", i, "X", x)
	}

	// 打印 Multiengines 信息
	multiengines := src.Data.Multiengines
	info += fmt.Sprintf("Multiengines.ScanTime:%s -> %s\n", "ScanTime", multiengines.ScanTime)
	info += fmt.Sprintf("Multiengines.Result.Kaspersky:%s -> %s\n", "Kaspersky", multiengines.Result.Kaspersky)
	info += fmt.Sprintf("Multiengines.Result.Microsoft:%s -> %s\n", "Microsoft", multiengines.Result.Microsoft)

	// 打印 Static 信息
	static := src.Data.Static
	info += fmt.Sprintf("Static.Basic.SHA1:%s -> %s\n", "SHA1", static.Basic.SHA1)
	info += fmt.Sprintf("Static.Basic.SHA256:%s -> %s\n", "SHA256", static.Basic.SHA256)
	info += fmt.Sprintf("Static.Basic.FileType:%s -> %s\n", "FileType", static.Basic.FileType)
	info += fmt.Sprintf("Static.Basic.FileName:%s -> %s\n", "FileName", static.Basic.FileName)
	info += fmt.Sprintf("Static.Basic.SSDeep:%s -> %s\n", "SSDeep", static.Basic.SSDeep)
	info += fmt.Sprintf("Static.Basic.FileSize:%s -> %d\n", "FileSize", static.Basic.FileSize)
	info += fmt.Sprintf("Static.Basic.MD5:%s -> %s\n", "MD5", static.Basic.MD5)

	// 打印 Signature 信息
	for i, sig := range src.Data.Signature {
		info += fmt.Sprintf("Signature[%d].Severity:%s -> %d\n", i, "Severity", sig.Severity)
		info += fmt.Sprintf("Signature[%d].SigClass:%s -> %s\n", i, "SigClass", sig.SigClass)
		info += fmt.Sprintf("Signature[%d].Name:%s -> %s\n", i, "Name", sig.Name)
		info += fmt.Sprintf("Signature[%d].Description:%s -> %s\n", i, "Description", sig.Description)
		info += fmt.Sprintf("Signature[%d].Markcount:%s -> %d\n", i, "Markcount", sig.Markcount)
		info += fmt.Sprintf("Signature[%d].AttckID:%s -> %s\n", i, "AttckID", sig.AttckID)
	}

	// 打印 Dropped 信息
	for i, dropped := range src.Data.Dropped {
		info += fmt.Sprintf("Dropped[%d].SHA1:%s -> %s\n", i, "SHA1", dropped.SHA1)
		info += fmt.Sprintf("Dropped[%d].SHA256:%s -> %s\n", i, "SHA256", dropped.SHA256)
		info += fmt.Sprintf("Dropped[%d].Size:%s -> %d\n", i, "Size", dropped.Size)
		info += fmt.Sprintf("Dropped[%d].Filepath:%s -> %s\n", i, "Filepath", dropped.Filepath)
		info += fmt.Sprintf("Dropped[%d].Name:%s -> %s\n", i, "Name", dropped.Name)
		info += fmt.Sprintf("Dropped[%d].CRC32:%s -> %s\n", i, "CRC32", dropped.CRC32)
		info += fmt.Sprintf("Dropped[%d].SSDeep:%s -> %s\n", i, "SSDeep", dropped.SSDeep)
		info += fmt.Sprintf("Dropped[%d].Type:%s -> %s\n", i, "Type", dropped.Type)
		info += fmt.Sprintf("Dropped[%d].MD5:%s -> %s\n", i, "MD5", dropped.MD5)
		for j, url := range dropped.URLs {
			info += fmt.Sprintf("Dropped[%d].URLs[%d]:%s -> %s\n", i, j, "URLs", url)
		}
	}

	// 打印 Pstree 信息
	pstree := src.Data.Pstree
	info += fmt.Sprintf("Pstree.ProcessName.En:%s -> %s\n", "En", pstree.ProcessName.En)
	info += fmt.Sprintf("Pstree.ProcessName.Cn:%s -> %s\n", "Cn", pstree.ProcessName.Cn)
	printChildren(&info, pstree.Children, 0)

	// 打印 Network 信息
	network := src.Data.Network
	info += printInterfaceSlice(network.Fingerprint, "Network.Fingerprint", "Fingerprint")
	info += printInterfaceSlice(network.TLS, "Network.TLS", "TLS")
	info += printInterfaceSlice(network.UDP, "Network.UDP", "UDP")
	info += printInterfaceSlice(network.DnsServers, "Network.DnsServers", "DnsServers")
	info += printInterfaceSlice(network.HTTP, "Network.HTTP", "HTTP")
	info += printInterfaceSlice(network.IRC, "Network.IRC", "IRC")
	info += printInterfaceSlice(network.SMTP, "Network.SMTP", "SMTP")
	info += printInterfaceSlice(network.TCP, "Network.TCP", "TCP")
	info += printInterfaceSlice(network.SMTPEx, "Network.SMTPEx", "SMTPEx")
	info += printInterfaceSlice(network.MITM, "Network.MITM", "MITM")
	info += printInterfaceSlice(network.Hosts, "Network.Hosts", "Hosts")
	info += printInterfaceSlice(network.DNS, "Network.DNS", "DNS")
	info += printInterfaceSlice(network.HTTPEx, "Network.HTTPEx", "HTTPEx")
	info += printInterfaceSlice(network.Domains, "Network.Domains", "Domains")
	info += printInterfaceSlice(network.DeadHosts, "Network.DeadHosts", "DeadHosts")
	info += printInterfaceSlice(network.ICMP, "Network.ICMP", "ICMP")
	info += printInterfaceSlice(network.HTTPSEx, "Network.HTTPSEx", "HTTPSEx")

	// 打印 Strings 信息
	for i, sha256 := range src.Data.Strings.SHA256 {
		info += fmt.Sprintf("Strings.SHA256[%d]:%s -> %s\n", i, "SHA256", sha256)
	}
	for i, pcap := range src.Data.Strings.Pcap {
		info += fmt.Sprintf("Strings.Pcap[%d]:%s -> %s\n", i, "Pcap", pcap)
	}

	// 打印 Permalink 信息
	info += fmt.Sprintf("Permalink:%s -> %s\n", "Permalink", src.Data.Permalink)

	return info
}

// printChildren 递归打印进程子结构体信息
func printChildren(info *string, children []Children, level int) {
	for i, child := range children {
		indent := ""
		for j := 0; j < level; j++ {
			indent += "  "
		}
		*info += fmt.Sprintf("%sPstree.Children[%d].Track:%s -> %v\n", indent, i, "Track", child.Track)
		*info += fmt.Sprintf("%sPstree.Children[%d].PID:%s -> %d\n", indent, i, "PID", child.PID)
		*info += fmt.Sprintf("%sPstree.Children[%d].ProcessName:%s -> %s\n", indent, i, "ProcessName", child.ProcessName)
		*info += fmt.Sprintf("%sPstree.Children[%d].CommandLine:%s -> %s\n", indent, i, "CommandLine", child.CommandLine)
		*info += fmt.Sprintf("%sPstree.Children[%d].FirstSeen:%s -> %s\n", indent, i, "FirstSeen", child.FirstSeen)
		*info += fmt.Sprintf("%sPstree.Children[%d].PPID:%s -> %d\n", indent, i, "PPID", child.PPID)
		printChildren(info, child.Children, level+1)
	}
}

// printInterfaceSlice 打印 interface{} 切片信息
func printInterfaceSlice(slice []interface{}, fieldName, key string) string {
	var info string
	for i, item := range slice {
		info += fmt.Sprintf("%s[%d]:%s -> %v\n", fieldName, i, key, item)
	}
	return info
}

func convert2WeiBuDataChinese(src *weiBuResponse) string {
	var info string

	// 打印 Summary 信息
	summary := src.Data.Summary
	info += fmt.Sprintf("威胁等级:%s -> %s\n", "threat_level", summary.ThreatLevel)
	info += fmt.Sprintf("威胁分类:%s -> %s\n", "malware_type", summary.MalwareType)
	info += fmt.Sprintf("病毒家族:%s -> %s\n", "malware_family", summary.MalwareFamily)
	//info += fmt.Sprintf("是否白名单文件:%s -> %v\n", "is_whitelist", summary.IsWhitelist)
	//info += fmt.Sprintf("文件提交时间:%s -> %s\n", "submit_time", summary.SubmitTime)
	info += fmt.Sprintf("文件名称:%s -> %s\n", "file_name", summary.FileName)
	info += fmt.Sprintf("文件类型:%s -> %s\n", "file_type", summary.FileType)
	info += fmt.Sprintf("文件的sha256值:%s -> %s\n", "sample_sha256", summary.SampleSha256)
	info += fmt.Sprintf("文件的MD5值:%s -> %s\n", "md5", summary.MD5)
	info += fmt.Sprintf("文件的SHA1值:%s -> %s\n", "sha1", summary.SHA1)
	info += fmt.Sprintf("威胁分数:%s -> %d\n", "threat_score", summary.ThreatScore)
	//info += fmt.Sprintf("沙箱运行分析环境:%s -> %s\n", "sandbox_type", summary.SandboxType)
	info += fmt.Sprintf("反病毒扫描引擎检出率:%s -> %s\n", "multi_engines", summary.MultiEngines)

	// 打印 Summary.Tag 信息
	for i, s := range summary.Tag.S {
		info += fmt.Sprintf("标签.静态标签[%d]:%s -> %s\n", i, "s", s)
	}
	for i, x := range summary.Tag.X {
		info += fmt.Sprintf("标签.检测标签[%d]:%s -> %s\n", i, "x", x)
	}

	//// 打印 Multiengines 信息
	//multiengines := src.Data.Multiengines
	//info += fmt.Sprintf("反病毒扫描引擎检测结果.扫描时间:%s -> %s\n", "scan_time", multiengines.ScanTime)
	//info += fmt.Sprintf("反病毒扫描引擎检测结果.卡巴斯基检测结果:%s -> %s\n", "Kaspersky", multiengines.Result.Kaspersky)
	//info += fmt.Sprintf("反病毒扫描引擎检测结果.微软检测结果:%s -> %s\n", "Microsoft", multiengines.Result.Microsoft)

	// 打印 Static 信息
	//static := src.Data.Static
	//info += fmt.Sprintf("静态信息.文件基本信息.SHA1:%s -> %s\n", "sha1", static.Basic.SHA1)
	//info += fmt.Sprintf("静态信息.文件基本信息.SHA256:%s -> %s\n", "sha256", static.Basic.SHA256)
	//info += fmt.Sprintf("静态信息.文件基本信息.文件类型:%s -> %s\n", "file_type", static.Basic.FileType)
	//info += fmt.Sprintf("静态信息.文件基本信息.文件名称:%s -> %s\n", "file_name", static.Basic.FileName)
	//info += fmt.Sprintf("静态信息.文件基本信息.SSDeep:%s -> %s\n", "ssdeep", static.Basic.SSDeep)
	//info += fmt.Sprintf("静态信息.文件基本信息.文件大小:%s -> %d\n", "file_size", static.Basic.FileSize)
	//info += fmt.Sprintf("静态信息.文件基本信息.MD5:%s -> %s\n", "md5", static.Basic.MD5)

	// 打印 Signature 信息
	for i, sig := range src.Data.Signature {
		info += fmt.Sprintf("行为签名[%d].严重等级:%s -> %d\n", i, "severity", sig.Severity)
		info += fmt.Sprintf("行为签名[%d].签名分类:%s -> %s\n", i, "sig_class", sig.SigClass)
		info += fmt.Sprintf("行为签名[%d].签名名称:%s -> %s\n", i, "name", sig.Name)
		info += fmt.Sprintf("行为签名[%d].行为描述:%s -> %s\n", i, "description", sig.Description)
		info += fmt.Sprintf("行为签名[%d].标记计数:%s -> %d\n", i, "markcount", sig.Markcount)
		info += fmt.Sprintf("行为签名[%d].ATT&CK ID:%s -> %s\n", i, "attck_id", sig.AttckID)
	}

	// 打印 Dropped 信息
	//for i, dropped := range src.Data.Dropped {
	//	info += fmt.Sprintf("释放行为[%d].SHA1:%s -> %s\n", i, "sha1", dropped.SHA1)
	//	info += fmt.Sprintf("释放行为[%d].SHA256:%s -> %s\n", i, "sha256", dropped.SHA256)
	//	info += fmt.Sprintf("释放行为[%d].大小:%s -> %d\n", i, "size", dropped.Size)
	//	info += fmt.Sprintf("释放行为[%d].文件路径:%s -> %s\n", i, "filepath", dropped.Filepath)
	//	info += fmt.Sprintf("释放行为[%d].名称:%s -> %s\n", i, "name", dropped.Name)
	//	info += fmt.Sprintf("释放行为[%d].CRC32:%s -> %s\n", i, "crc32", dropped.CRC32)
	//	info += fmt.Sprintf("释放行为[%d].SSDeep:%s -> %s\n", i, "ssdeep", dropped.SSDeep)
	//	info += fmt.Sprintf("释放行为[%d].类型:%s -> %s\n", i, "type", dropped.Type)
	//	info += fmt.Sprintf("释放行为[%d].MD5:%s -> %s\n", i, "md5", dropped.MD5)
	//	for j, url := range dropped.URLs {
	//		info += fmt.Sprintf("释放行为[%d].URLs[%d]:%s -> %s\n", i, j, "urls", url)
	//	}
	//}

	// 打印 Pstree 信息
	pstree := src.Data.Pstree
	info += fmt.Sprintf("进程行为.进程名称.英文:%s -> %s\n", "en", pstree.ProcessName.En)
	info += fmt.Sprintf("进程行为.进程名称.中文:%s -> %s\n", "cn", pstree.ProcessName.Cn)
	printChildrenChinese(&info, pstree.Children, 0)

	// 打印 Network 信息
	network := src.Data.Network
	info += printInterfaceSliceChinese(network.Fingerprint, "网络行为.指纹信息", "fingerprint")
	info += printInterfaceSliceChinese(network.TLS, "网络行为.TLS信息", "tls")
	info += printInterfaceSliceChinese(network.UDP, "网络行为.UDP信息", "udp")
	info += printInterfaceSliceChinese(network.DnsServers, "网络行为.DNS服务器信息", "dns_servers")
	info += printInterfaceSliceChinese(network.HTTP, "网络行为.HTTP信息", "http")
	info += printInterfaceSliceChinese(network.IRC, "网络行为.IRC信息", "irc")
	info += printInterfaceSliceChinese(network.SMTP, "网络行为.SMTP信息", "smtp")
	info += printInterfaceSliceChinese(network.TCP, "网络行为.TCP信息", "tcp")
	info += printInterfaceSliceChinese(network.SMTPEx, "网络行为.SMTP扩展信息", "smtp_ex")
	info += printInterfaceSliceChinese(network.MITM, "网络行为.MITM信息", "mitm")
	info += printInterfaceSliceChinese(network.Hosts, "网络行为.主机信息", "hosts")
	info += printInterfaceSliceChinese(network.DNS, "网络行为.DNS信息", "dns")
	info += printInterfaceSliceChinese(network.HTTPEx, "网络行为.HTTP扩展信息", "http_ex")
	info += printInterfaceSliceChinese(network.Domains, "网络行为.域名信息", "domains")
	info += printInterfaceSliceChinese(network.DeadHosts, "网络行为.无效主机信息", "dead_hosts")
	info += printInterfaceSliceChinese(network.ICMP, "网络行为.ICMP信息", "icmp")
	info += printInterfaceSliceChinese(network.HTTPSEx, "网络行为.HTTPS扩展信息", "https_ex")

	// 打印 Strings 信息
	for i, sha256 := range src.Data.Strings.SHA256 {
		info += fmt.Sprintf("字符串信息.SHA256提取的字符串[%d]:%s -> %s\n", i, "sha256", sha256)
	}
	for i, pcap := range src.Data.Strings.Pcap {
		info += fmt.Sprintf("字符串信息.Pcap提取的字符串[%d]:%s -> %s\n", i, "pcap", pcap)
	}

	//// 打印 Permalink 信息
	//info += fmt.Sprintf("网页沙箱报告页网址:%s -> %s\n", "permalink", src.Data.Permalink)

	return info
}

// printChildrenChinese 递归打印进程子结构体信息（中文）
func printChildrenChinese(info *string, children []Children, level int) {
	for i, child := range children {
		indent := ""
		for j := 0; j < level; j++ {
			indent += "  "
		}
		*info += fmt.Sprintf("%s进程行为.子进程[%d].跟踪:%s -> %v\n", indent, i, "track", child.Track)
		*info += fmt.Sprintf("%s进程行为.子进程[%d].进程ID:%s -> %d\n", indent, i, "pid", child.PID)
		*info += fmt.Sprintf("%s进程行为.子进程[%d].进程名称:%s -> %s\n", indent, i, "process_name", child.ProcessName)
		*info += fmt.Sprintf("%s进程行为.子进程[%d].进程命令符:%s -> %s\n", indent, i, "command_line", child.CommandLine)
		*info += fmt.Sprintf("%s进程行为.子进程[%d].首次出现时间:%s -> %s\n", indent, i, "first_seen", child.FirstSeen)
		*info += fmt.Sprintf("%s进程行为.子进程[%d].父进程ID:%s -> %d\n", indent, i, "ppid", child.PPID)
		printChildrenChinese(info, child.Children, level+1)
	}
}

// printInterfaceSliceChinese 打印 interface{} 切片信息（中文）
func printInterfaceSliceChinese(slice []interface{}, fieldName, key string) string {
	var info string
	for i, item := range slice {
		info += fmt.Sprintf("%s[%d]:%s -> %v\n", fieldName, i, key, item)
	}
	return info
}
