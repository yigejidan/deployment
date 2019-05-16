package main

import(
	"deployment/host/connent"
	"deployment/path"
	"deployment/getconf"
	"log"
	"deployment/deploy/component"
)


func main() {
	pdhost,pdcomponent := ReadConfigToHost(path.PDpath)
	err := component.OpenDocker(pdhost)
	if err != nil {
		log.Fatal("pdhost Open Docker err",err)
	}
	err = component.DeployComponent(pdcomponent)
	if err != nil {
		log.Fatal("start pd component failed",err)
	}
	
	tidb1host,tidb1component := ReadConfigToHost(path.TIDB1path)
	err = component.OpenDocker(tidb1host)
	if err != nil {
		log.Fatal("tidb1host Open Docker err",err)
	}
	err = component.DeployComponent(tidb1component)
	if err != nil {
		log.Fatal("start tidb1 component failed",err)
	}

	tidb2host,tidb2component := ReadConfigToHost(path.TIDB2path)
	err = component.OpenDocker(tidb2host)
	if err != nil {
		log.Fatal("tidb2host Open Docker err",err)
	}
	err = component.DeployComponent(tidb2component)
	if err != nil {
		log.Fatal("start tidb2 component failed",err)
	}

	tikvhost,tikvcomponent := ReadConfigToHost(path.TIKVpath)
	err = component.OpenDocker(tikvhost)
	if err != nil {
		log.Fatal("tikvhost Open Docker err",err)
	}
	err = component.DeployComponent(tikvcomponent)
	if err != nil {
		log.Fatal("start tikv component failed",err)
	}

}


//ReadConfigToHost ：读取配置文件
func ReadConfigToHost(path string) (host *connent.Host,component *getconf.Config) {
	component,err := getconf.ReadConfig(path)
	if err != nil {
		log.Fatal("Read Config err",err)
	}
	host = &connent.Host{
		User : component.User,
		Password : component.Password,
		Port : component.Port,
	}
	return host,component
}

