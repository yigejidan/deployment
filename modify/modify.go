package modify
import (
	"deployment/host/connent"
	"log"
)
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

//
func ChangeComponentHost(oldhost connent.Host,newhost connent.Host,) {

}