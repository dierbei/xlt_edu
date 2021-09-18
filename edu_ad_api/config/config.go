package config

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataID    string `mapstructure:"data_id" json:"data_id"`
	Group     string `mapstructure:"group" json:"group"`
}

type ServerConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	Port         int          `mapstructure:"port" json:"port"`
	Host         string       `mapstructure:"host" json:"host"`
	Tags         []string     `mapstructure:"tags" json:"tags"`
	Mode         string       `mapstructure:"mode" json:"mode"`
	JwtConfig    JwtConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulConfig ConsulConfig `mapstructure:"consul" json:"consul"`
	SpaceConfig  SpaceConfig  `mapstructure:"space_srv" json:"space_srv"`
}

type JwtConfig struct {
	SignatureKey string `mapstructure:"signature_key" json:"signature_key"`
	ExpireSecond int    `mapstructure:"expire_second" json:"expire_second"`
	ExpireCount  int    `mapstructure:"expire_count" json:"expire_count"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type SpaceConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
