package main

//From https://www.thepolyglotdeveloper.com/2017/05/network-sockets-with-the-go-programming-language/
import (
	"fmt"
	"net"
	"os"
)

// Client is struct to hold socket and data channel
type Client struct {
	socket net.Conn
	data   chan []byte
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
			client.socket.Write(message)
		}
	}
	fmt.Println("Socket Closed")
}

func startServer(port string) {
	listener, error := net.Listen("tcp", ":"+port)
	if error != nil {
		fmt.Println(error)
	}
	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		go client.receive()
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) >= 1 {
		startServer(argsWithoutProg[0])
	} else {
		startServer("1234")
	}
}
