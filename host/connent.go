package host

import (
    "fmt"
    "bytes"
    "golang.org/x/crypto/ssh"
    "deployment/getconf"
    "log"
    "net"
)
//ConnentHost 连接主机主程序
func ConnentHost(path string) error {
    config,err := getconf.ReadConfig(path)   //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
		return err
	}
	//log.Println(config)

	clientconfig := &ssh.ClientConfig{
		User: config.Uesr,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),	
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
	}
	client, err := ssh.Dial("tcp",config.Port,clientconfig)
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
	if err := session.Run(config.Command); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(b.String()) 
	return nil
}
