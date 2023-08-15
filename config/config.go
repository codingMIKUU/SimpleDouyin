package config

const Ip_address = "10.135.95.126" // （填入本机 IP 地址）

const Video_quantity_limit = 30 // 每次获取视频流的数量

const DateTime = "2006-01-02 15:04:05" // 固定的时间格式

// MysqlConfig 用于存储 MySQL 数据库的配置信息
type MysqlConfig struct {
	Host     string `json:"host" yaml:"host"`               // 主机
	Port     string `json:"port" yaml:"port"`               // 端口
	Dbname   string `json:"db_name" mapstructure:"db_name"` // 数据库名
	Username string `json:"username" yaml:"username"`       // 用户名
	Password string `json:"password" yaml:"password"`       // 密码
	Charset  string `json:"charset" yaml:"charset"`         // 字符集
	Prefix   string `json:"prefix" yaml:"prefix"`           // 前缀
}

// 返回一个DSN字符串，它是连接 MySQL 数据库所需的参数字符串
func (m *MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}

type ServerConfig struct {
	Port   int    `json:"port,omitempty" yaml:"port"`
	DbType string `json:"db_type,omitempty" yaml:"db_type"`
}

// 包含配置信息的结构体
type Conf struct {
	Mysql  *MysqlConfig  // MysqlConfig 结构体定义了 MySQL 数据库的配置信息
	Server *ServerConfig // ServerConfig 结构体定义了一些服务器配置的信息。
}

const ResourceServerURL = "http://" + Ip_address + ":8080/public/"                    // 服务端地址
const AvatarURL = "http://" + Ip_address + ":8080/public/Initdata/initavater.jpg"     // 默认头像地址
const BackgroundURL = "http://" + Ip_address + ":8080/public/Initdata/background.jpg" // 默认背景图地址
const SignatureStr = "这个人很懒，什么也没有留下..."                                               // 默认简介内容
