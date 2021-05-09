package option

import (
	conf "github.com/basharal/config"
	"github.com/kerwinan/go/kit/log"
)

// Options program config
type Options struct {
	LogOpt      LogOpt      `yaml:"log_opt"`
	ClientOpt   ClientOpt   `yaml:"client_opt"`
	ServiceOpt  ServiceOpt  `yaml:"service_opt"`
	ProviderOpt ProviderOpt `yaml:"provider_opt"`
}

/* TODO conf 解析做成 kit 工具包，解析指定目录下所有yaml文件 */
var confPath string = ".."

func SetConfigPath(path string) {
	confPath = path
}

// NewOptions config
func NewOptions() (opts *Options, err error) {
	opts = new(Options)
	/* parse log yaml */
	err = conf.Parse(confPath+"/config/log.yaml", &opts.LogOpt)
	if err != nil {
		return nil, err
	}
	l, err := log.GetLoggerByConf((*log.Config)(&opts.LogOpt))
	if err != nil {
		return nil, err
	}
	log.X = l
	/* parse provider yaml */
	err = conf.Parse(confPath+"/config/provider.yaml", &opts.ProviderOpt)
	if err != nil {
		return nil, err
	}
	/* parse service yaml */
	err = conf.Parse(confPath+"/config/service.yaml", &opts.ServiceOpt)
	if err != nil {
		return nil, err
	}
	/* parse client yaml */
	err = conf.Parse(confPath+"/config/client.yaml", &opts.ClientOpt)
	if err != nil {
		return nil, err
	}
	return
}
