package equationqueue

import (
	"log"
	"github.com/jjeffery/stomp"
	"equationmessage"
)

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

// SendMessage sends given EquationMessage into given queueName
func (queueWriter *QueueWriter) SendMessage(equationMessage *equationmessage.EquationMessage, queueName string) error {
	if !queueWriter.isConnected {
		log.Panic("Could not send to queue since queue writer is not connected")
	}

	serializedMessage, err := equationMessage.SerializeToJson()

	if err != nil {
		log.Printf("Could not serialize message of Id %s", equationMessage.CorrelationId)
		return err
	}

	err = queueWriter.connection.Send(queueName, contentType, []byte(serializedMessage))

	if err != nil {
		log.Printf("Could not send message to queue %s", queueName)
		return err
	}

	log.Printf("Sent message '%s' to queue '%s'", serializedMessage, queueName)
	return nil
}

// Disconnect sends given equation into given queueName
func (queueWriter *QueueWriter) Disconnect() {
	queueWriter.isConnected = false
	queueWriter.connection.Disconnect()

	log.Printf("Disconnected")
}