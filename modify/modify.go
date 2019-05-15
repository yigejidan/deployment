package modify
import (
	"deployment/host/connent"
	"log"
)

//
type Component struct {
	Name     string
	Port1    string
	Port2    string
	Path     string
	Logfile  string
}

/*
CloseOldComponent : 关闭旧的组件
*/
func CloseOldComponent(host connent.Host,componentname string) {
	cmd := "docker container stop " + componentname
	_,err := connent.RunCommand(host, cmd)
	if err != nil {
		log.Fatal("close component failed",err)
	}
	_,err = connent.RunCommand(host, "systemctl stop docker.service")
	if err != nil {
		log.Fatal("stop docker failed",err)
	}
	_,err = connent.RunCommand(host, "systemctl disable doccker.service")
	if err != nil {
		log.Fatal("disable docker failed",err)
	}
}

//ChangeComponentHost : 变更功能实现
func ChangeComponentHost(oldhost connent.Host,newhost connent.Host) {
	
}