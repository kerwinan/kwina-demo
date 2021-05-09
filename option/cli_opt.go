package option

type ClientOpt struct {
	DemoClientOpt *CliConf `yaml:"rpc_client_opt"`
}

// 客户端配置文件
type CliConf struct {
	// 调用方名称, 格式为{biz}.{app}
	Client string `yaml:"client"`
	// 程序所在节点
	NodeID int `yaml:"node_id"`
	// ETCD集群地址
	ETCDHost []string `yaml:"etcd_host"`
	// ETCD用户名
	Username string `yaml:"username"`
	// ETCD密码
	Password string `yaml:"password"`
	// 网卡名称, 默认为eth0
	NetworkInterface string `yaml:"network_interface"`
	//调用方的IP, 与NetworkInterface二选一
	ClientIP string `yaml:"client_ip"`
	// 调用超时时间, 默认10秒
	CallTimeout int `yaml:"call_timeout"`
}
