package monitor

import (
	"deployment/host/connent"
	"strings"
	"log"
	// "sync"
)

// var m *sync.Mutex

//AuditComponent 监视组件运行状态
func AuditComponent(host *connent.Host,componentname string) (status bool,err error) {

	var ComponentState bool 
	command := "docker ps"
	result,err := connent.RunCommand(host,command)
	if err != nil {
		log.Fatal("docker ps run err",err)
	}
	isCon := strings.Contains(result,componentname)
	if isCon == true {
		ComponentState = true
		return ComponentState,nil
	}else {
		ComponentState = false 
		return ComponentState,nil
	}
}

//IsHaveContainer : 查看docker里是否有该组件的container
func IsHaveContainer(host *connent.Host,componentname string) (bool,error) {
	var IsHave bool 
	command := "docker ps -a"
	result,err := connent.RunCommand(host,command)
	if err != nil {
		log.Fatal("docker ps -a run err",err)
	}
	isCon := strings.Contains(result,componentname)
	if isCon == true {
		IsHave = true
		return IsHave,nil
	}else {
		IsHave = false 
		return IsHave,nil
	}
}

//IsOpenDocker ： 查看docker服务是否开启
func IsOpenDocker(host *connent.Host) (bool,error) {
	var IsOpen bool 
	command := "systemctl is-active docker.service"
	result,err := connent.RunCommand(host,command)
	if err != nil {
		log.Fatal("cmd :systemctl is-active docker.service run err : ",err)
	}
	isCon := strings.Contains(result,"active")
	if isCon == true {
		IsOpen = true
		return IsOpen,nil
	}else {
		IsOpen = false 
		return IsOpen,nil
	}
}

