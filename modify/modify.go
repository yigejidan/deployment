package modify
import (
	"deployment/host/connent"
	"log"
)

// //Component : 变更功能组件的参数结构体 
// type Component struct {
// 	Name     string
// 	Port1    string
// 	Port2    string
// 	Path     string
// 	Logfile  string
// }

/*
CloseOldComponent : 关闭旧的组件
*/
func CloseOldComponent(host *connent.Host,componentname string) {
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

//OpenNewComponent : 开启新的组件
func OpenNewComponent(newhost connent.Host) {
	tidb2host,tidb2component := ReadConfigToHost(path.TIDB2path)
	err = component.OpenDocker(tidb2host)
	if err != nil {
		log.Fatal("tidb2host Open Docker err",err)
	}
	err = component.DeployComponent(tidb2component)
	if err != nil {
		log.Fatal("start tidb2 component failed",err)
	}
}