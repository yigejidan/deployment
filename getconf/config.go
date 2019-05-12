package getconf
 
import (
  "log"
  "gopkg.in/ini.v1"
  "sync"
)

//Config ： 部署集群结构体
type Config struct {   
	Uesr string  `ini:"uesr"`
	Password string `ini:"password"`
	Port string `ini:"port"`
	Command string   `ini:"command"`
}


var m *sync.RWMutex

//ReadConfig 读取配置文件并转成结构体
func ReadConfig(path string) (Config, error) {
  m.RLock()
  defer m.RUnlock()
  var config Config
  conf, err := ini.Load(path)   //加载配置文件
  if err != nil {
    log.Println("load config file fail!")
    return config, err
  }
  conf.BlockMode = false
  err = conf.MapTo(&config)   //解析成结构体
  if err != nil {
    log.Println("mapto config file fail!")
    return config, err
  }
  return config, nil
}
