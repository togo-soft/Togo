package utils

//该处定义了如何从根目录下获取并解析配置信息

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

//读取配置文件信息
type E struct {
	Environments `yaml:"environments"`
}

//请手动添加结构来实现yaml的解析
type Environments struct {
	Debug  bool   `yaml:"debug"`   //是否开启debug模式
	Server string `yaml:"server"`  //服务运行的host:port
	Jwt    Jwt    `yaml:"jwt"`     //jwt签名配置
	Mysql  Driver `yaml:"mysql"`   //mysql数据库配置
	Mongo  Driver `yaml:"mongodb"` //mongodb数据库配置
}

type Driver struct {
	DSN   string `yaml:"dsn"`   //数据库连接的源名称
	Debug bool   `yaml:"debug"` //开启数据库debug模式
}

type Jwt struct {
	SignKey     string `yaml:"sign_key"`
	SignMethod  string `yaml:"sign_method"`
	SignIssuer  string `yaml:"sign_issuer"`
	SignSubject string `yaml:"sign_subject"`
	SignExpires int64  `yaml:"sign_expires"`
}

var conf *E
var once sync.Once

func GetConfig() *E {
	once.Do(func() {
		conf = new(E)
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			//读取配置文件失败,停止执行
			panic("read config file error:" + err.Error())
		}
	})
	return conf
}
