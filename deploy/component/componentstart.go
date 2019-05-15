package component

import (
	"deployment/host/connent"
	"log"
	"deployment/getconf"
)

//OpenDocker ： 开启主机docker
func OpenDocker(host connent.Host) {
	start := "systemctl start docker.service"
	enable := "systemctl enable docker.service"
	_,err := connent.RunCommand(host,enable)
	if err != nil {
		log.Fatal("docker enable err",err)
	}
	_,err = connent.RunCommand(host,start)
	if err != nil {
		log.Fatal("docker start err",err)
	}
}

//DeployComponent ； 部署组件
func DeployComponent(path string) error{
	config,err := getconf.ReadConfig(path)   //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	host := connent.Host {
		User : config.User,
		Password : config.Password,
		Port : config.Port,
	}
	_,err = connent.RunCommand(host,config.Command)
	if err != nil {
		log.Fatal("depoly component err",err)
	}
	return nil 
}

