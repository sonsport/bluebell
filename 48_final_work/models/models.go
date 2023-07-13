package models

// Config 为接收yaml文件的结构体
type Config struct {
	Name  string       `mapstructure:"name"`
	Mode  string       `mapstructure:"mode"`
	Port  int          `mapstructure:"port"`
	Mysql *MysqlStruct `mapstructure:"mysql"`
	Redis *RedisStruct `mapstructure:"redis"`
	Log   *LogStruct   `mapstructure:"log"`
}

type MysqlStruct struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Maxopenconns int    `mapstructure:"maxOpenConns"`
	Maxidleconns int    `mapstructure:"maxIdleConns"`
}

type RedisStruct struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type LogStruct struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}
