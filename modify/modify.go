package modify
import (
	"deployment/host"
	"log"
)
/*
Modify : 提供变更功能
*/
func Modify(path1 string,path2 string) error {
	err := host.ConnentHost(path1)
	if err != nil {
		log.Fatal("stop The old components err",err)
		return err
	}
	err = host.ConnentHost(path2)
	if err != nil {
		log.Fatal("start new component err ",err)
		return err 
	}
	return nil 
}