package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // : ：a@qq.com b@qq.com
	From     string `mapstructure:"from" json:"from" yaml:"from"`             //
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             //   smtp.qq.com  QQsmtp
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       //       smtp
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` //
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             //      QQsmtp  465
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // SSL   SSL
}
