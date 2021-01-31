package queuereader

import (
	"log"

	"github.com/jjeffery/stomp"
)

const contentType = "text/plain"

// QueueWriter connects and write messages to queue
type QueueWriter struct {
	isConnected 	bool
	connection  	*stomp.Conn
	// subscription 	// TODO
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

func // TODO suscription

// ReadMessage reads message from given queueName
func (queueWriter *QueueWriter) ReadMessage(queueName string) {
	if !queueWriter.isConnected {
		log.Panic("Could not read from queue since queue reader is not connected")
	}

	sub, _ := conn.Subscribe("queueName", stomp.AckAuto)
 
  	msg := <- sub.C

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

	// tODO maybe unsubscribe

	log.Printf("Disconnected")
}