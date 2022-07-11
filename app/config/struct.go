package config

type Config struct {
	ServerConfig   `yaml:"Server"`
	SmtpConfig     `yaml:"Smtp"`
	BlogInfoConfig `yaml:"BlogInfo"`
	RedisConfig    `yaml:"Redis"`
}

type ServerConfig struct {
	Addr  string `yaml:"Address"`
	Debug bool   `yaml:"Debug"`
}

type SmtpConfig struct {
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
	Secure      bool   `yaml:"Secure"`
	User        string `yaml:"Username"`
	Pwd         string `yaml:"Password"`
	SenderEmail string `yaml:"SenderEmail"`
	SenderName  string `yaml:"SenderName"`
}

type BlogInfoConfig struct {
	Tittle      string `yaml:"Tittle"`
	Addr        string `yaml:"BlogAddress"`
	AuthorEmail string `yaml:"AuthorEmail"`
}

type RedisConfig struct {
	Enable      bool   `yaml:"Enable"`
	Addr        string `yaml:"Address"`
	Pwd         string `yaml:"Password"`
	Db          int    `yaml:"Database"`
	Prefix      string `yaml:"Prefix"`
	Concurrency int    `yaml:"Concurrency"`
}
