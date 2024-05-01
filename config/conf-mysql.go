package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log-level"`
	Config   string `yaml:"config"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Db + m.Config
}
