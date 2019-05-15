package host

import (
    "fmt"
    "bytes"
    "golang.org/x/crypto/ssh"
    "deployment/getconf"
    "log"
	"net"
	"sync"
)

var m *sync.Mutex
config,err := getconf.ReadConfig(path)   //也可以通过os.arg或flag从命令行指定配置文件路径
if err != nil {
	log.Fatal(err)
	return err
}

//ConnentHost 连接主机主程序
func ConnentHost(user string,password string,port string,command string) (result string,err error) {
	m.Lock()
	defer m.Unlock()

	clientconfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),	
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
	}
	client, err := ssh.Dial("tcp",port,clientconfig)
	if err != nil {
		log.Fatal(err)
		return err
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		log.Fatal(err)
		return err
	}
	return b.string(),nil
}
