package restart

import (
	"deployment/host/monitor"
	"deployment/host/connent"
	"log"
	"strings"
)

//ReStartComponent ：重启意外关闭的组件 
func ReStartComponent (host connent.Host,componentname string) error {
	isStart,err := monitor.AuditComponent(host,componentname)
	if err != nil {
		log.Fatal("auditcomponent err",err)
	}
	if isStart == true {
		log.Fatal(componentname,"Component is running")
	}
	result,err := connent.RunCommand(host, "docker ps -a")
	isCon := strings.Contains(result,componentname)
	if isCon == true {
		cmd := "docker container start " + componentname
		_,err := connent.RunCommand(host, cmd)
		if err != nil {
			log.Fatal("restart component err",err)
		}
	}else {
		log.Fatal("docker no container for this component")
	}
	return nil 
}

