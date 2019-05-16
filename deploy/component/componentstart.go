package component

import (
	"deployment/host/connent"
	"log"
	"deployment/getconf"
	"sync"
)

var m *sync.Mutex

//OpenDocker ： 开启主机docker
func OpenDocker(host *connent.Host) error{
	m.Lock()
	defer m.Unlock()
	start := "systemctl start docker.service"
	enable := "systemctl enable docker.service"
	_,err := connent.RunCommand(host,enable)
	if err != nil {
		log.Fatal("docker enable err",err)
		return err
	}
	_,err = connent.RunCommand(host,start)
	if err != nil {
		log.Fatal("docker start err",err)
		return err
	}
	return nil
}

//DeployComponent ； 部署组件
func DeployComponent(component *getconf.Config) error{
	m.Lock()
	defer m.Unlock()
	// config,err := getconf.ReadConfig(path)   //也可以通过os.arg或flag从命令行指定配置文件路径
	// if err != nil {
	// 	log.Fatal(err)
	// }
	host := connent.Host {
		User : component.User,
		Password : component.Password,
		Port : component.Port,
	}
	_,err := connent.RunCommand(&host,component.Command)
	if err != nil {
		log.Fatal("depoly component err",err)
	}
	return nil 
}

