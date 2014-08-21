package proxy

import (
	"log"
	"net"
	"time"
	"github.com/bamdadd/gopheroxy/configuration"
)

type TCPProxy struct{
	config *configuration.Configuration
}


func ProxyTCP(config *configuration.Configuration) {
	p:= &TCPProxy{config}
	p.proxy()
}


func (p *TCPProxy) proxy() {
	server, err := net.Listen("tcp", p.config.GetFrontend())

	if err != nil {
		log.Fatal(err)
	}

	// The channel of connections which are waiting to be processed.
	waiting := make(chan net.Conn, p.config.GetMaxWaitConn())
	// The booleans representing the free active connection spaces.
	spaces := make(chan bool, p.config.GetMaxConn())
	// Initialize the spaces
	for i := 0; i < p.config.GetMaxConn(); i++ {
		spaces <- true
	}

	// Start the connection matcher.
	go matchConnections(waiting, spaces, p.config.GetBackend())

	// Loop indefinitely, accepting connections and handling them.
	for {
		connection, err := server.Accept()
		if err != nil {
			// Log the error.
			log.Print(err)
		} else {
			// Create a goroutine to handle the conn
			log.Printf("Received connection from %s.\r\n",
				connection.RemoteAddr())
			waiting <- connection
		}
	}

}



func matchConnections(waiting chan net.Conn, spaces chan bool, toHost string) {
	// Iterate over each connection in the waiting channel
	for connection := range waiting {
		// Block until we have a space.
		<-spaces
		// Create a new goroutine which will call the connection handler and
		// then free up the space.
		go func(connection net.Conn) {
			handleConnection(connection, toHost)
			spaces <- true
			log.Printf("Closed connection from %s.\r\n", connection.RemoteAddr())
		}(connection)

	}
}

func handleConnection(connection net.Conn, toHost string) {
	// Always close our connection.
	defer connection.Close()

	// Try to connect to remote server.
	remote, err := net.Dial("tcp", toHost)
	if err != nil {
		// Exit out when an error occurs
		log.Print(err)
		return
	}
	defer remote.Close()

	// Create our channel which waits for completion, and our two channels to
	// signal that a goroutine is done.
	complete := make(chan bool, 2)
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	go copyContent(connection, remote, complete, ch1, ch2)
	go copyContent(remote, connection, complete, ch2, ch1)
	// Block until we've completed both goroutines!
	<- complete
	<- complete
}

func copyContent(from net.Conn, to net.Conn, complete chan bool, done chan bool, otherDone chan bool) {
	var err error = nil
	var bytes []byte = make([]byte, 256)
	var read int = 0
	for {
		select {
			// If we received a done message from the other goroutine, we exit.
		case <- otherDone:
		complete <- true
			return
		default:

			// Read data from the source connection.
			from.SetReadDeadline(time.Now().Add(time.Second * 5))
			read, err = from.Read(bytes)
			// If any errors occured, write to complete as we are done (one of the
			// connections closed.)
			if err != nil {
				complete <- true
				done <- true
				return
			}
			// Write data to the destination.
			to.SetWriteDeadline(time.Now().Add(time.Second * 5))
			_, err = to.Write(bytes[:read])
			// Same error checking.
			if err != nil {
				complete <- true
				done <- true
				return
			}
		}
	}
}
