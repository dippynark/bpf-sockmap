package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"

	"github.com/dippynark/bpf-sockmap/pkg/sockmap"
)

const (
	defaultPort = 12345
)

var portFlag = flag.Int("port", defaultPort, "listen port")

func main() {

	flag.Parse()
	port := *portFlag

	// setup sockmap module
	sockmap, err := sockmap.New()
	if err != nil {
		log.Fatalf("failed to create new sockmap module: %s", err)
	}
	log.Print("created new sockmap module")
	defer func() {
		err := sockmap.Close()
		if err != nil {
			log.Fatalf("failed to close sockmap module: %s", err)
		}
	}()

	// listen
	listenAddress := fmt.Sprintf("0.0.0.0:%d", port)
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen on port %d: %s", port, err)
	}
	defer func() {
		err := l.Close()
		if err != nil {
			log.Fatalf("failed to close socket: %s", err)
		}
	}()
	log.Printf("listening on address: %s", listenAddress)

	// add accepted connections to sockmap
	for {
		// accept
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error accepting: %s", err)
		}

		// retrieve copy of connection file descriptor
		tcpConn, ok := conn.(*net.TCPConn)
		if !ok {
			log.Fatalf("failed to cast connection to TCP connection")
		}
		f, err := tcpConn.File()
		if err != nil {
			log.Fatalf("failed to retrieve copy of the underlying TCP connection file")
		}
		d := f.Fd()

		// update element
		err = sockmap.UpdateSocketDescriptor(d)
		if err != nil {
			log.Fatalf("failed to update socket descriptor: %s", err)
		}

		// we don't need two copies of the connection file descriptor so close the copy
		// https://stackoverflow.com/questions/28967701/golang-tcp-socket-cant-close-after-get-file#answer-28968431
		err = syscall.SetNonblock(int(d), true)
		if err != nil {
			log.Fatalf("failed to put file descriptor in non-blocking mode: %s", err)
		}
		err = f.Close()
		if err != nil {
			log.Fatalf("failed to close file descriptor copy: %s", err)
		}
	}
}

/*
func waitForCloseByClient(conn net.Conn) {
	fmt.Println("Accepted connection from", conn.RemoteAddr())

	defer func() {
		fmt.Println("Closing connection from", conn.RemoteAddr())
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("Read error", err.Error())
			return
		}
	}
}
*/
