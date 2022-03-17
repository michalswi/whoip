package whoip

import (
	"log"
	"net"
)

func WHOip(domainName, NSserver string) string {
	conn, err := net.Dial("tcp", NSserver+":43")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer conn.Close()
	conn.Write([]byte(domainName + "\r\n"))
	buf := make([]byte, 1024)
	result := []byte{}
	for {
		numBytes, err := conn.Read(buf)
		sbuf := buf[0:numBytes]
		result = append(result, sbuf...)
		if err != nil {
			break
		}
	}
	return string(result)
}
