package main

func RabbitMQCfg() {

	if Init_RabbitMQ {
		FolderCheck("util/rabbitmq", "util/rabbitmq", "[RABBITMQ] ")
		// ReadWriteContentFromRabbitMQCfgTxt("[RABBITMQ] ", "kafkaCfg.txt")
		WriteContentToConfigYaml(RabbitMQ_Init_Library, "util/rabbitmq/init.go", "[RABBITMQ] ")
		WriteContentToConfigYaml(RabbitMQ_Init_Producer, "util/rabbitmq/producer.go", "[RABBITMQ] ")
		WriteContentToConfigYaml(RabbitMQ_Init_Consumer, "util/rabbitmq/consumer.go", "[RABBITMQ] ")

	}

}

// func ReadWriteContentFromRabbitMQCfgTxt(LogParameter, FileName string) {
// 	log.SetPrefix(green.Render(LogParameter))
// 	fileTxt, err := os.ReadFile("./" + FileName)
// 	if err != nil {
// 		log.Println(red.Render("%s file open error. \n")+FileName, err)
// 		zap.L().Info("file open error: "+FileName, zap.Error(err))
// 	}
// 	fileGo, err := os.OpenFile("./util/rabbitmq/"+"rabbitmq.go", os.O_RDWR|os.O_CREATE, 0666)
// 	fileGoName := strings.Split("rabbitmq.go", ".")
// 	if err != nil {
// 		log.Println(red.Render("%s file open error. \n")+fileGoName[0]+".go", err)
// 		zap.L().Info("file open error: "+FileName, zap.Error(err))
// 	}
// 	defer fileGo.Close()
// 	fileGo.Write([]byte(fileTxt))
// 	log.Printf("Write %s configuration to %s success. \n", LogParameter, fileGoName[0]+".go")
// }

var (
	RabbitMQ_Init_Library = `package rabbitMq

import (
	"log"

	"github.com/streadway/amqp"
)

// MQURL format amqp://username：password@rabbitmq server address：port/vhost (default is 5672 port)
// port can view from  /etc/rabbitmq/rabbitmq-env.conf，or use command netstat -tlnp check
const MQURL = "amqp://admin:admin@172.21.138.131:5672/"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// queue name
	QueueName string
	// exchange name
	Exchange string
	// routing Key
	RoutingKey string
	//MQ connection string
	Mqurl string
}

// create rabbitmq instance
func NewRabbitMQ(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      MQURL,
	}
	var err error
	//create rabbitmq connection
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	checkErr(err, "Failure to create connection")

	//creat channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "Failure to create channel")

	return &rabbitMQ

}

// release resource
func (mq *RabbitMQ) ReleaseRes() {
	mq.Conn.Close()
	mq.Channel.Close()
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
`
	RabbitMQ_Init_Producer = `package main

import (
	"fmt"
	"mq/rabbitMq"

	"github.com/streadway/amqp"
)

// mainProducer demonstrates the producer workflow for RabbitMQ
// It shows how to declare a queue, exchange, bind them, and publish messages
func mainProducer() {
	// Initialize RabbitMQ connection with queue, exchange, and routing key
	mq := rabbitMq.NewRabbitMQ("queue_publisher", "exchange_publisher", "key1")
	defer mq.ReleaseRes() // Release resources after completion

	// 1. Declare a queue
	// Declaring from both producer and consumer sides is recommended to:
	// - Prevent consumers from failing to subscribe to non-existent queues
	// - Prevent messages from being dropped when no matching queue exists
	// Note: RabbitMQ ignores attempts to declare existing queues and returns success
	_, err := mq.Channel.QueueDeclare(
		mq.QueueName, // Queue name
		true,         // Durable - survive broker restart
		false,        // Auto-delete when no consumers - only if consumers were connected
		false,        // Exclusive - only accessible via this connection
		false,        // No-wait - don't wait for server confirmation
		nil,          // Additional arguments (not used here)
	)
	if err != nil {
		fmt.Println("Failed to declare queue:", err)
		return
	}

	// 2. Declare an exchange
	err = mq.Channel.ExchangeDeclare(
		mq.Exchange, // Exchange name
		"topic",     // Exchange type (fanout, direct, topic are common)
		true,        // Durable - survive broker restart
		false,       // Auto-delete when no bindings remain
		false,       // Internal - if true, clients can't publish directly
		false,       // No-wait - don't wait for server confirmation
		nil,         // Additional arguments
	)
	if err != nil {
		fmt.Println("Failed to declare exchange:", err)
		return
	}

	// 3. Bind queue to exchange
	// Multiple bindings can be created as needed
	err = mq.Channel.QueueBind(
		mq.QueueName,  // Queue to bind
		mq.RoutingKey, // Routing key for message distribution
		mq.Exchange,   // Exchange to bind to
		false,         // No-wait - don't wait for server confirmation
		nil,           // Additional arguments
	)
	if err != nil {
		fmt.Println("Failed to bind queue to exchange:", err)
		return
	}

	// 4. Publish a message
	err = mq.Channel.Publish(
		mq.Exchange,   // Exchange to publish to
		mq.RoutingKey, // Routing key
		false,         // Mandatory - if true, return unroutable messages
		false,         // Immediate - if true, return messages with no consumers
		amqp.Publishing{ // Message content and properties
			ContentType: "text/plain",          // MIME type of content
			Body:        []byte("hello world"), // Message payload
		},
	)
	if err != nil {
		fmt.Println("Failed to publish message:", err)
	}
}

`
	RabbitMQ_Init_Consumer = `package main

import (
	"fmt"
	"mq/rabbitMq"
)

func mainConsumer() {

	mq := rabbitMq.NewRabbitMQ("queue_publisher", "exchange_publisher", "key1")
	defer mq.ReleaseRes()

	_, err := mq.Channel.QueueDeclare(
		mq.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failure to declare queen.", err)
		return
	}

	msgChanl, err := mq.Channel.Consume(
		mq.QueueName,
		"",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		fmt.Println("Failure to get message.", err)
		return
	}

	for msg := range msgChanl {

		fmt.Println(string(msg.Body))

	}

}
`
)
