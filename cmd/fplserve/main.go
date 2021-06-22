package main

import (
	"fmt"
	"net"
	"net/http"

	server "github.com/florian74/assignement/server"
)

func main() {

	var controller server.Controller
	http.HandleFunc("/flight/", controller.HandleSearch)
	http.ListenAndServe(":8081", nil)

	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	sock, _ := net.ListenUDP("udp", addr)

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
