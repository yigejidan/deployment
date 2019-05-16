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

	var ComponentState bool //组件正常运行为true，关闭为false
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
	var IsHave bool //组件正常运行为true，关闭为false
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

