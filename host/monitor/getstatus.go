package monitor

import (
	"deployment/host/connent"
	"strings"
	"log"
	"sync"
)

var m *sync.Mutex

//AuditComponent 监视组件运行状态
func AuditComponent(user string,password string,port string,componentname string) (status bool,err error) {
	m.Lock()
	defer m.Unlock()
	var ComponentState bool //组件正常运行为true，关闭为false
	command := "docker ps"
	result,err := connent.ConnentHost(user,password,port,command)
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

