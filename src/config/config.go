package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Configs ConfigSetting




type ConfigSetting struct {
	TestAdmin struct{
		Port int 			`yaml:"port"`
		AuthTimeout int 	`yaml:"auth_timeout"`
		KeepOptLogDay int	`yaml:"keep_opt_log_day"`
		PageSize int 		`yaml:"page_size"`
	}`yaml:"test_admin"`

	Database struct{
		Mysql struct{
			MaxIdleConnection int			`yaml:"maxIdleConnection"`
			MaxOpenConnection int			`yaml:"maxOpenConnection"`
		}`yaml:"database"`
	}
}

type MysqlSetting struct {
	Name string 			`yaml:"name"`
	Write databaseSetting 	`yaml:"write"`
	Read []databaseSetting 	`yaml:"read"`
}


type databaseSetting struct {
	Host string 			`yaml:"host"`
	Port int 				`yaml:"port"`
	Username string 		`yaml:"username"`
	Password string 		`yaml:"password"`
	Database string 		`yaml:"database"`
	Charset string 			`yaml:"charset"`
	MaxIdleConnection int	`yaml:"maxIdleConnection"`
	MaxOpenConnection int	`yaml:"maxOpenConnection"`
}


func Setup() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		panic("Read config faild : " + err.Error())
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Read config faild: " + err.Error())
	}
	yaml.Unmarshal(byteValue, &Configs)
}