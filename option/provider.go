package option

type ProviderOpt struct {
	HelloProviderOpt  HelloProviderOpt  `yaml:"hello_provider_opt"`
}


type HelloProviderOpt struct {
	Enable bool   `yaml:"enable"`
	Addr   string `yaml:"addr"`
}

