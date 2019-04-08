package funcpy
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func GetEmResult(){
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		buf := make([]byte, 1024)
		conn.Write([]byte(input))
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
		}
		fmt.Print(string(buf[0:cnt]))
	}
}