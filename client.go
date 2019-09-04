package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var url string
	flag.StringVar(&url, "url", "", "usage : IP:PORT ")
	var file string
	flag.StringVar(&file, "file", "test.text", "the file name")
	flag.Parse()

	msg, err := ioutil.ReadFile(file)
	check(err)
	fmt.Println(string(msg))

	//res, err := sendTCP("127.0.0.1:8000", "hi")
	res, err := sendTCP(url, string(msg))
	check(err)
	fmt.Println(res)
	/*
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	*/
}

func sendTCP(addr, msg string) (string, error) {
	// connect to this socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send to socket
	conn.Write([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
