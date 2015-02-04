// go run udp.go
package main

import(
	"fmt"
	"net"
	"strconv"
	"time"
)

func read(connection *net.UDPConn){
	msg := make([]byte,1024)
	for{
		msgSize,_,_ := connection.ReadFromUDP(msg)
		fmt.Println("Received: " +string(msg[0:msgSize]))
		
	}
}

func send(connection *net.UDPConn){
	for{
		time.Sleep(1000*time.Millisecond)
		msg := "Hello from workspace 24"
		connection.Write([]byte(msg))
		fmt.Println("Msg sent to server")
	}
}
func main(){
	workspace := 24;
	serverIP := "129.241.187.132";
	serverPort := 20000 + workspace;
	serverAddr,_ := net.ResolveUDPAddr("udp",serverIP+ ":" +strconv.Itoa(serverPort));
	
	readPort := 20000 + workspace;
	readAddr,_ := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(readPort));

	fmt.Println(readAddr);
	fmt.Println(serverAddr);

	readConn,_ := net.ListenUDP("udp",readAddr);
	sendConn,_ := net.DialUDP("udp",nil,serverAddr);
	go read(readConn);
	go send(sendConn);
	
	deadChan :=make(chan bool,1);
	<- deadChan;
	
	
	

}
