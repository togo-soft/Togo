package utils

//该处定义了如何从根目录下获取配置信息

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//读取配置文件信息
type E struct {
	Environments `yaml:"environments"`
}

//请手动添加结构来实现yaml的解析
type Environments struct {
	Debug    bool   `yaml:"debug"`     //是否开启debug模式
	Server   string `yaml:"server"`    //服务运行的host:port
	MysqlDSN string `yaml:"mysql_dsn"` //MySQL数据库连接的DSN
}

func GetConfig() *E {
	conf := new(E)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		//读取配置文件失败,停止执行
		panic(err)
	}
	return conf
}
