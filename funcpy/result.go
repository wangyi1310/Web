package funcpy

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"sync"
)

var (
	err error
	conn net.Conn

 	mutex sync.Mutex
)
func init() {
	conn, err = net.Dial("tcp", "localhost:10000")
	if err != nil {
		logs.Error("server connet failuer")
		return
	}
}
func Close(){
	conn.Close()
}
// http get reslut
func GetEmResult(str string,res *string)  {
	mutex.Lock()
	defer mutex.Unlock()
	buf := make([]byte, 1024)
	var length = len(str)
	if length >= 1024{
		length = 1024
	}

	if length == 0{
		return
	}
	conn.Write([]byte(str[0:length]))
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Fail to read data, %s\n", err)
	}
	*res =string(buf[0:cnt])
}


