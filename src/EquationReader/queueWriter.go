package equationreader

import (
	"log"

	"github.com/jjeffery/stomp"
)

const contentType = "text/plain"

// QueueWriter connects and write messages to queue
type QueueWriter struct {
	isConnected bool
	connection  *stomp.Conn
}

// Connect Creates a connection to given address
func (queueWriter *QueueWriter) Connect(connectionAddress string) {
	connection, err := stomp.Dial("tcp", connectionAddress)

	if err != nil {
		log.Fatalf("Could not connect to %s", connectionAddress)
		return
	}

	queueWriter.isConnected = true
	queueWriter.connection = connection

	log.Printf("Connected to address: %s", connectionAddress)
}

// SendEquetion sends given equation into given queueName
func (queueWriter *QueueWriter) SendEquetion(queueName string, equations string) {
	if !queueWriter.isConnected {
		log.Panic("Could not send to queue since queue writer is not connected")
	}

	err := queueWriter.connection.Send(queueName, contentType, []byte(equations))

	if err != nil {
		log.Printf("Could not send message to queue %s", queueName)
		return
	}

	log.Printf("Send message '%s' to queue '%s'", equations, queueName)
}

// Disconnect sends given equation into given queueName
func (queueWriter *QueueWriter) Disconnect() {
	queueWriter.isConnected = false
	queueWriter.connection.Disconnect()

	log.Printf("Disconnected")
}