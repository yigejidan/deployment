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
├── deploy                       
│   ├── component             
│   ├── ├── componentstart.go   //部署组件  
│   ├── restart             
│   ├── ├── restart.go //重启组件   
├── getconf           
│   ├── config.go       //解析配置文件  
├── host                        
│   ├── connent             
│   ├── ├── connent.go   //连接主机并执行命令       
│   ├── monitor             
│   ├── ├── getstatus.go //查看主机tidb组件运行状态  
├── main                         
│   ├── main.go      //主程序  
├── modify                   
│   ├── modify.go                //提供变更功能  
├── path                        
│   ├── path.go     //配置文件路径  
├── README.md   

           
