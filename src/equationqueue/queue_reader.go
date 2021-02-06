package equationqueue

import (
	"log"
	"github.com/jjeffery/stomp"
	"equationmessage"
)

// QueueReader connects and subscribe to a given channel
type QueueReader struct {
	isConnected 			bool
	connection  			*stomp.Conn
	subscription 			*stomp.Subscription
	subscriptionQueueName	string
}

// Connect Creates a connection to given address
func (queueReader *QueueReader) Connect(connectionAddress string) {
	connection, err := stomp.Dial("tcp", connectionAddress)

	if err != nil {
		log.Fatalf("Could not connect to %s", connectionAddress)
		return
	}

	queueReader.isConnected = true
	queueReader.connection = connection

	log.Printf("Connected to address: %s", connectionAddress)
}

// ReadMessage reads message from given queueName
func (queueReader *QueueReader) Subscribe(queueName string) {
	if !queueReader.isConnected {
		log.Panicf("Could not subscribe to queue %s queue since queue reader is not connected", queueName)
	}

	sub, err := queueReader.connection.Subscribe(queueName, stomp.AckAuto)

	if err != nil {
		log.Printf("Could not subscribe to queue %s", queueName)
		return
	}

	queueReader.subscription = sub
	queueReader.subscriptionQueueName = queueName
	log.Printf("Subscribe to queue %s", queueName)
}

func (queueReader *QueueReader) ReadMessage() (*equationmessage.EquationMessage, error) {
	if !queueReader.isConnected {
		log.Panic("Could not read from queue since queue reader is not connected")
	}

  	message := <- queueReader.subscription.C
  	messageString := string(message.Body)

  	deserializedEquationMessage, err := equationmessage.DeserializeFromJson(messageString)

  	if err != nil {
		return nil, err
	}

  	return deserializedEquationMessage, nil
}

// Disconnect sends given equation into given queueName
func (queueReader *QueueReader) Disconnect() {

	queueReader.subscription.Unsubscribe()

	queueReader.isConnected = false
	queueReader.connection.Disconnect()

	log.Printf("Disconnected")
}