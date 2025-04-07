package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Init() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	fmt.Println("Conexão estabelecida com o rabbitMQ")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("mensagem-teste", "", true, false, false, false, nil)

	forever := make(chan bool)
	go readStuff(msgs)
	fmt.Println("Esperando as mensagens da queue")

	<-forever

}

func readStuff(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		fmt.Print(string(d.Body))
	}
}
