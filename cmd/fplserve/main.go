package main

import (
	"fmt"
	"net"
	"net/http"

	server "github.com/florian74/assignement/server"
)

func main() {

	fmt.Println("server start TCP")
	var controller server.Controller
	go asyncListenAndServe(controller)

	fmt.Println("server start UDP")
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	sock, _ := net.ListenUDP("udp", addr)
	fmt.Println("server start listening udp on 8080")

	i := 0
	for {
		i++
		buf := make([]byte, 1048576)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buf[0:rlen]))
		fmt.Println(i)

		controller.HandlePut(buf, rlen, i)
	}

}

func asyncListenAndServe(controller server.Controller) {
	fmt.Println("server start listening http on 8081")
	http.HandleFunc("/flight/", controller.HandleSearch)
	http.ListenAndServe(":8081", nil)
	fmt.Println("server stop listening http on 8081")
}
