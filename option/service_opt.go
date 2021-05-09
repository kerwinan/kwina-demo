package option

type ServiceOpt struct {
	HelloServiceOpt *SrvConf `yaml:"hello_service_opt"`
}

/* TODO 实现服务注册与发现工具时移植 */
// 服务端配置文件
type SrvConf struct {
	// 服务禁用
	Disable bool
	// 服务版本号
	Version string `yaml:""`
	// 服务名
	Service string `yaml:"service"`
	// 服务监听端口
	Port int `yaml:"port"`
	// 程序所在节点
	NodeID string `yaml:"node_id"`
	// ETCD集群地址
	ETCDHost []string `yaml:"etcd_host"`
	// ETCD用户名
	Username string `yaml:"username"`
	// ETCD密码
	Password string `yaml:"password"`
	// 网卡名称, 默认为eth0
	NetworkInterface string `yaml:"network_interface"`
	// 是否使用SSL
	UseSSL bool `yaml:"use_ssl"`
	// 证书地址
	CertFile string `yaml:"cert_file"`
	// 证书秘钥地址
	KeyFile string `yaml:"key_file"`
}
