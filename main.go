package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
)

var (
	host  = flag.String("host", "localhost", "hostname to listen to")
	token = flag.String("token", "myToken", "token to access mini monkey")
	port  = flag.String("port", "1773", "which port to listen to")
	room  = flag.String("room", "temperatures", "which room to use")
	msg   = flag.String("msg", "it is cold", "what message to publish")
	tag   = flag.String("tag", "myTAG", "tag for subscription")
)

func check(err error, t string) {
	if err != nil {
		fmt.Println(t)
		fmt.Println(err.Error())
		panic(err)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func response(conn net.Conn) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println("got: ", string(buf))
}

func subscribeLoop(conn net.Conn) {
	rem := []byte{}

	for {
		buf := make([]byte, 128)
		n, err := conn.Read(buf)
		check(err, "failed to read in subscribe")

		payload := append(rem, buf[0:n]...)
		ok, _, data, left := Decode(payload)

		if ok {
			fmt.Println("Data: ", string(data))
		}

		rem = left
	}
}

func main() {
	flag.Parse()

	l, err := net.Dial("tcp", *host+":"+*port)
	check(err, "error listening")
	defer l.Close()

	fmt.Println("connected to " + *host + " on port " + *port)

	_, err = l.Write(auth(*token))
	check(err, "write problem")
	response(l)

	err = binary.Write(l, binary.LittleEndian, enter(*room))
	check(err, "room problem")
	response(l)

	if isFlagPassed("msg") {
		err = binary.Write(l, binary.LittleEndian, publish(*msg))
		check(err, "publish problem")
		response(l)
	} else {
		err = binary.Write(l, binary.LittleEndian, subscribe(*tag))
		check(err, "subscribe problem")
		response(l)
		subscribeLoop(l)
	}
}
