package main

import (
	"flag"
	"fmt"
	"github.com/skysoft-atm/assignement/adexp"
	"github.com/skysoft-atm/assignement/transport"
)

var destAddr = flag.String("addrdest", "0.0.0.0:6669", "ip and port to send messages to (ie 127.0.0.5:6669)")

func main() {
	flag.Parse()
	fmt.Println("random ifpl udp message sender")
	connect, err := transport.NewConnection(*destAddr)
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	fmt.Printf("%s \n", connect.String())

	gen := adexp.NewGenerator()

	var n int64
	for n = 0; ; n++ {
		// send on the wire
		_, err = connect.Write(gen.Next())
		if err != nil {
			panic(err)
		}
		fmt.Printf("sent %d messages \r", n)
	}

}
