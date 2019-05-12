# depolyment
部署TiDB集群的go语言程序

## 目录结构描述  
├── conf                    // 配置文件  
│   ├── modify.conf  
│   ├── pd.conf               
│   ├── tidb1.conf         
│   ├── tidb2.conf               
│   ├── tidb3.conf              
│   ├── tikv.conf        
├── getconf           
│   ├── config.go       //解析配置文件  
├── host                        
│   ├── connent.go             //连接主机并部署tidb组件    
├── main                         
│   ├── main.go      //主程序  
├── modify                   
│   ├── modify.go                //提供变更功能  
├── path                        
│   ├── path.go     //配置文件路径  
├── README.md   

## 程序各部分实现  
config.go : 
           利用"gopkg.in/ini.v1"包实现对配置文件的解析并封装成结构体
connent.go :
           调用"golang.org/x/crypto/ssh"包向连接主机并发送运行tiDB组件部署的命令   
path.go :
           封装各个配置文件的路径  
change.go :
           组件变更程序，首先关闭原来的组件，再在新的主机上开启新的组件  
main.go :
           调用client包中的方法实现tidb集群的部署，同时也可以调用change.go中的方法实现组件的变更
