package getconf
 
import (
  "log"
  "gopkg.in/ini.v1"
  // "sync"
  "deployment/host/connent"
)

//Config ： 部署集群结构体
type Config struct {   
	User string  `ini:"user"`
	Password string `ini:"password"`
	Port string `ini:"port"`
	Command string   `ini:"command"`
}


// var m *sync.RWMutex

//ReadConfig 读取配置文件并转成结构体
func ReadConfig(path string) (*Config, error) {
  // m.RLock()
  // defer m.RUnlock()
  var config Config
  conf, err := ini.Load(path)   //加载配置文件
  if err != nil {
    log.Println("load config file fail!")
  }
  conf.BlockMode = false
  err = conf.MapTo(&config)   //解析成结构体
  if err != nil {
    log.Println("mapto config file fail!")
  }
  return &config, nil
}

//ReadConfigToHost ：读取配置文件返回主机结构体和组件结构体
func ReadConfigToHost(path string) (host *connent.Host,component *Config) {
	component,err := ReadConfig(path)
	if err != nil {
		log.Fatal("Read Config err",err)
	}
	host = &connent.Host{
		User : component.User,
		Password : component.Password,
		Port : component.Port,
	}
	return host,component
}
