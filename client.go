package main

import (
	"fmt"
	"net"
	"net/http"
	"io/ioutil"
	"flag"
)

func parseFlag() (*string, *string, *string) {
	var (
		protocol = flag.String("proto", "tcp", "Specify protocol (tcp/udp)")
		host     = flag.String("host", "localhost", "Specify destination address")
		port     = flag.String("port", "9002", "Specify destination port")
	)
	flag.Parse()
	return protocol, host, port
}

func getPublicIPv4() (string) {
	req, _ := http.NewRequest("GET", "http://169.254.169.254/latest/meta-data/public-ipv4", nil)
	res, err := new(http.Client).Do(req)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	byteArray, _ := ioutil.ReadAll(res.Body)
	return string(byteArray)
}

func sendMessage(conn net.Conn, text string) {
	_, err := conn.Write([]byte(text))

	if err != nil {
		panic(err)
	}

	res := make([]byte, 4 * 1024);
	conn.Read(res);
	fmt.Printf("Server> %s \n", res)
}

func main() {

	protocol, host, port := parseFlag()

	conn, err := net.Dial(*protocol, *host+":"+*port)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sendMessage(conn, "public_ipv4: "+getPublicIPv4())
}