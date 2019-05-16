package main

import(
	"deployment/path"
	"deployment/getconf"
	"log"
	"deployment/deploy/component"
	"deployment/host/monitor"
	"deployment/deploy/restart"
)


func main() {
	for {
		pdhost,pdcomponent := getconf.ReadConfigToHost(path.PDpath)
		isOpen,err := monitor.IsOpenDocker(pdhost)
		if err != nil {
			log.Fatal("pd IsOpenDocker err : ",err)
		}
		if !isOpen {
			err = component.OpenDocker(pdhost)
			if err != nil {
				log.Fatal("pdhost Open Docker err",err)
			}
		}
		isHave,err := monitor.IsHaveContainer(pdhost,"pd")
		if err != nil {
			log.Fatal("IsHaveContainer err : ",err)
		}
		if !isHave {
			err = component.DeployComponent(pdcomponent)
			if err != nil {
				log.Fatal("Deploy pd component failed",err)
			}
		}
		isStart,err := monitor.AuditComponent(pdhost, "pd")
		if !isStart {
			err = restart.ReStartComponent(pdhost,"pd")
			if err != nil {
				log.Fatal("restart pd component failed",err)
			}
		}
		
		tidb1host,tidb1component := getconf.ReadConfigToHost(path.TIDB1path)
		isOpen,err = monitor.IsOpenDocker(tidb1host)
		if err != nil {
			log.Fatal("tidb1 IsOpenDocker err : ",err)
		}
		if !isOpen {
			err = component.OpenDocker(tidb1host)
			if err != nil {
				log.Fatal("tidb1host Open Docker err",err)
			}
		}
		isHave,err = monitor.IsHaveContainer(tidb1host,"tidb1")
		if err != nil {
			log.Fatal("tidb1 IsHaveContainer err : ",err)
		}
		if !isHave {
			err = component.DeployComponent(tidb1component)
			if err != nil {
				log.Fatal("Deploy tidb1 component failed",err)
			}
		}
		isStart,err = monitor.AuditComponent(tidb1host, "tidb1")
		if !isStart {
			err = restart.ReStartComponent(tidb1host,"tidb1")
			if err != nil {
				log.Fatal("restart tidb1 component failed",err)
			}
		}

		tidb2host,tidb2component := getconf.ReadConfigToHost(path.TIDB2path)
		isOpen,err = monitor.IsOpenDocker(tidb2host)
		if err != nil {
			log.Fatal("tidb2 IsOpenDocker err : ",err)
		}
		if !isOpen {
			err = component.OpenDocker(tidb2host)
			if err != nil {
				log.Fatal("tidb2host Open Docker err",err)
			}
		}
		isHave,err = monitor.IsHaveContainer(tidb2host,"tidb2")
		if err != nil {
			log.Fatal("tidb2 IsHaveContainer err : ",err)
		}
		if !isHave {
			err = component.DeployComponent(tidb2component)
			if err != nil {
				log.Fatal("Deploy tidb2 component failed",err)
			}
		}
		isStart,err = monitor.AuditComponent(tidb2host, "tidb2")
		if !isStart {
			err = restart.ReStartComponent(tidb2host,"tidb2")
			if err != nil {
				log.Fatal("restart tidb2 component failed",err)
			}
		}

		tikvhost,tikvcomponent := getconf.ReadConfigToHost(path.TIKVpath)
		isOpen,err = monitor.IsOpenDocker(tikvhost)
		if err != nil {
			log.Fatal("tikv IsOpenDocker err : ",err)
		}
		if !isOpen {
			err = component.OpenDocker(tikvhost)
			if err != nil {
				log.Fatal("tikvhost Open Docker err",err)
			}
		}
		isHave,err = monitor.IsHaveContainer(tikvhost,"tikv")
		if err != nil {
			log.Fatal("tikv IsHaveContainer err : ",err)
		}
		if !isHave {
			err = component.DeployComponent(tikvcomponent)
			if err != nil {
				log.Fatal("Deploy tikv component failed",err)
			}
		}
		isStart,err = monitor.AuditComponent(tikvhost, "tikv")
		if !isStart {
			err = restart.ReStartComponent(tikvhost,"tikv")
			if err != nil {
				log.Fatal("restart tikv component failed",err)
			}
		}
	}
}




