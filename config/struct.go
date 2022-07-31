package config

type Config struct {
	ServerConfig   `yaml:"Server"`
	BlogInfoConfig `yaml:"BlogInfo"`
	RedisConfig    `yaml:"Redis"`
	SendByConfig   `yaml:"SendBy"`
	SmtpConfig     `yaml:"Smtp"`
	AliyunConfig   `yaml:"Aliyun"`
}

type ServerConfig struct {
	Addr  string `yaml:"Address"`
	Debug bool   `yaml:"Debug"`
}

type BlogInfoConfig struct {
	Title       string `yaml:"Title"`
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

type SendByConfig struct {
	Owner string `yaml:"Owner"`
	Guest string `yaml:"Guest"`
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

type AliyunConfig struct {
	AccessKey    string `yaml:"AccessKey"`
	AccessSecret string `yaml:"AccessSecret"`
	Email        string `yaml:"Email"`
	Region       string `yaml:"Region"`
	Domain       string `yaml:"Domain"`
}
