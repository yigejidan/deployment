package main

import(
	"deployment/path"
	"deployment/getconf"
	"log"
	"deployment/deploy/component"
	// "fmt"
	// "deployment/host/monitor"
)


func main() {
	pdhost,pdcomponent := getconf.ReadConfigToHost(path.PDpath)
	err := component.OpenDocker(pdhost)
	if err != nil {
		log.Fatal("pdhost Open Docker err",err)
	}
	err = component.DeployComponent(pdcomponent)
	if err != nil {
		log.Fatal("start pd component failed",err)
	}
	
	tidb1host,tidb1component := getconf.ReadConfigToHost(path.TIDB1path)
	err = component.OpenDocker(tidb1host)
	if err != nil {
		log.Fatal("tidb1host Open Docker err",err)
	}
	err = component.DeployComponent(tidb1component)
	if err != nil {
		log.Fatal("start tidb1 component failed",err)
	}

	tidb2host,tidb2component := getconf.ReadConfigToHost(path.TIDB2path)
	err = component.OpenDocker(tidb2host)
	if err != nil {
		log.Fatal("tidb2host Open Docker err",err)
	}
	err = component.DeployComponent(tidb2component)
	if err != nil {
		log.Fatal("start tidb2 component failed",err)
	}

	tikvhost,tikvcomponent := getconf.ReadConfigToHost(path.TIKVpath)
	err = component.OpenDocker(tikvhost)
	if err != nil {
		log.Fatal("tikvhost Open Docker err",err)
	}
	err = component.DeployComponent(tikvcomponent)
	if err != nil {
		log.Fatal("start tikv component failed",err)
	}

}




