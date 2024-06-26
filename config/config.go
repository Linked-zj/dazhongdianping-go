package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Dbname   string `mapstructure:"dbname" json:"dbname"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	JWTConfig   JWTConfig   `mapstructure:"jwt"`
	MysqlConfig MysqlConfig `mapstructure:"mysql" json:"mysql"`
}
