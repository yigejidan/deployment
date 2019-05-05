package main

import(
	"deployment/host"
	"deployment/path"
	"deployment/modify"
	"log"
)


func main() {
	err := host.ConnentHost(path.PDpath)
	if err != nil {
		log.Fatal("deploy PD err",err)
		return
	}
	err = host.ConnentHost(path.TIDB1path)
	if err != nil {
		log.Fatal("deploy TIDB1 err",err)
		return
	}
	err = host.ConnentHost(path.TIDB2path)
	if err != nil {
		log.Fatal("deploy TIDB2 err",err)
		return
	}
	err = host.ConnentHost(path.TIKVpath)
	if err != nil {
		log.Fatal("deploy TIKV err",err)
		return
	}
	err = modify.Modify(path.Changepath,path.TIDB3path)
	if err != nil {
		log.Fatal("modify component err",err)
		return
	}
}