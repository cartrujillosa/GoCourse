package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil && err.Error() == "EOF" {
				log.Println("chat closed")
				return
			} else if err != nil {
				log.Println(err)
				return
			}
			io.WriteString(os.Stdout, msg)

		}
	}()

	for {
		msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
		_, err = io.WriteString(conn, msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	// TODO: when exit delete user
}
