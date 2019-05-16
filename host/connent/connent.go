package connent

import (
    "bytes"
    "golang.org/x/crypto/ssh"
    // "log"
	"net"
	// "sync"
)

//Host : 主机
type Host struct {   
	User string
	Password string 
	Port string 
}

// var m *sync.Mutex



//RunCommand 连接主机主程序,发送命令并执行
func RunCommand(host *Host,command string) (result string,err error) {
	// m.Lock()
	// defer m.Unlock()

	clientconfig := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),	
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
	}
	client, err := ssh.Dial("tcp",host.Port,clientconfig)
	if err != nil {
		// log.Fatal("connent post err ",err)
		return "",err
	}
	session, err := client.NewSession()
	if err != nil {
		// log.Fatal(err)
		return "",err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		// log.Fatal("run command err: ",err)
		return "",err
	}
	return b.String(),nil
}
