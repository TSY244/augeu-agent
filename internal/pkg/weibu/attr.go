package weibu

import "augeu-agent/internal/pkg/param"

type Config struct {
	Jwt      string
	Header   map[string]string
	Conf     *param.Config
	ClientId string
}

// weiBuResponse 根结构体
type weiBuResponse struct {
	ResponseCode int       `json:"response_code"`
	Data         weiBuData `json:"data"`
	VerboseMsg   string    `json:"verbose_msg"`
}

// weiBuData 包含各种信息的结构体
type weiBuData struct {
	Summary      weiBuSummary      `json:"summary"`
	Multiengines weiBuMultiengines `json:"multiengines"`
	Static       Static            `json:"static"`
	Signature    []weiBuSignature  `json:"signature"`
	Dropped      []weiBuDropped    `json:"dropped"`
	Pstree       weiBuPstree       `json:"pstree"`
	Network      weiBuNetwork      `json:"network"`
	Strings      weiBuStrings      `json:"strings"`
	Permalink    string            `json:"permalink"`
}

// weiBuSummary 概要信息结构体
type weiBuSummary struct {
	ThreatLevel     string   `json:"threat_level"`
	MalwareType     string   `json:"malware_type"`
	MalwareFamily   string   `json:"malware_family"`
	IsWhitelist     bool     `json:"is_whitelist"`
	SubmitTime      string   `json:"submit_time"`
	FileName        string   `json:"file_name"`
	FileType        string   `json:"file_type"`
	SampleSha256    string   `json:"sample_sha256"`
	MD5             string   `json:"md5"`
	SHA1            string   `json:"sha1"`
	Tag             Tag      `json:"tag"`
	ThreatScore     int      `json:"threat_score"`
	SandboxType     string   `json:"sandbox_type"`
	SandboxTypeList []string `json:"sandbox_type_list"`
	MultiEngines    string   `json:"multi_engines"`
}

// Tag 标签结构体
type Tag struct {
	S []string `json:"s"`
	X []string `json:"x"`
}

// weiBuMultiengines 反病毒扫描引擎检测结果结构体
type weiBuMultiengines struct {
	Result   Result `json:"result"`
	ScanTime string `json:"scan_time"`
}

// Result 反病毒扫描引擎具体结果结构体
type Result struct {
	Kaspersky string `json:"Kaspersky"`
	Microsoft string `json:"Microsoft"`
}

// Static 静态信息结构体
type Static struct {
	Details weiBuStaticDetails `json:"details"`
	Basic   Basic              `json:"basic"`
}

// weiBuStaticDetails PE 文件静态信息结构体
type weiBuStaticDetails struct {
	PeVersionInfo []interface{}          `json:"pe_version_info"`
	PeSections    []interface{}          `json:"pe_sections"`
	PeSignatures  map[string]interface{} `json:"pe_signatures"`
	PeImports     []interface{}          `json:"pe_imports"`
	PeResources   []interface{}          `json:"pe_resources"`
	Tag           []interface{}          `json:"tag"`
	PeDetect      map[string]interface{} `json:"pe_detect"`
	PeBasic       map[string]interface{} `json:"pe_basic"`
	PeExports     []interface{}          `json:"pe_exports"`
}

// Basic 文件基本信息结构体
type Basic struct {
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
	SSDeep   string `json:"ssdeep"`
	FileSize int    `json:"file_size"`
	MD5      string `json:"md5"`
}

// weiBuSignature 行为签名结构体
type weiBuSignature struct {
	Severity    int                    `json:"severity"`
	References  []interface{}          `json:"references"`
	SigClass    string                 `json:"sig_class"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Markcount   int                    `json:"markcount"`
	Marks       []interface{}          `json:"marks"`
	Families    []interface{}          `json:"families"`
	AttckID     string                 `json:"attck_id"`
	AttckInfo   map[string]interface{} `json:"attck_info"`
}

// weiBuDropped 释放行为结构体
type weiBuDropped struct {
	SHA1     string        `json:"sha1"`
	URLs     []string      `json:"urls"`
	SHA256   string        `json:"sha256"`
	Size     int           `json:"size"`
	Filepath string        `json:"filepath"`
	Name     string        `json:"name"`
	CRC32    string        `json:"crc32"`
	SSDeep   string        `json:"ssdeep"`
	Type     string        `json:"type"`
	Yara     []interface{} `json:"yara"`
	MD5      string        `json:"md5"`
}

// weiBuPstree 进程行为结构体
type weiBuPstree struct {
	Children    []Children       `json:"children"`
	ProcessName weiBuProcessName `json:"process_name"`
}

// Children 进程子结构体
type Children struct {
	Track       bool       `json:"track"`
	PID         int        `json:"pid"`
	ProcessName string     `json:"process_name"`
	CommandLine string     `json:"command_line"`
	FirstSeen   string     `json:"first_seen"`
	PPID        int        `json:"ppid"`
	Children    []Children `json:"children"`
}

// weiBuProcessName 进程名称信息结构体
type weiBuProcessName struct {
	En string `json:"en"`
	Cn string `json:"cn"`
}

// weiBuNetwork 网络行为结构体
type weiBuNetwork struct {
	Fingerprint []interface{} `json:"fingerprint"`
	TLS         []interface{} `json:"tls"`
	UDP         []interface{} `json:"udp"`
	DnsServers  []interface{} `json:"dns_servers"`
	HTTP        []interface{} `json:"http"`
	IRC         []interface{} `json:"irc"`
	SMTP        []interface{} `json:"smtp"`
	TCP         []interface{} `json:"tcp"`
	SMTPEx      []interface{} `json:"smtp_ex"`
	MITM        []interface{} `json:"mitm"`
	Hosts       []interface{} `json:"hosts"`
	DNS         []interface{} `json:"dns"`
	HTTPEx      []interface{} `json:"http_ex"`
	Domains     []interface{} `json:"domains"`
	DeadHosts   []interface{} `json:"dead_hosts"`
	ICMP        []interface{} `json:"icmp"`
	HTTPSEx     []interface{} `json:"https_ex"`
}

// weiBuStrings 字符串信息结构体
type weiBuStrings struct {
	SHA256 []string `json:"sha256"`
	Pcap   []string `json:"pcap"`
}
