package publisher

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func Init() {

	fmt.Println("Iniciando conexão ao rabbitmq...")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Conexão estabelecida")

	//Aqui tamo abrindo um canal concorrente para processar as mensagens que forem enviadas
	ch, err = conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("mensagem-teste", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	endProgram := make(chan bool)

	readInput(endProgram)
	<-endProgram

}

func publicar(msg string) {

	info := map[string]interface{}{
		"message": msg,
		"header":  "Json from go",
	}

	js, err := json.Marshal(info)

	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",
		"mensagem-teste",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(js),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Mensagem publicada")

}

func readInput(endProgram chan<- bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Envie uma mensagem para o queue: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if text == "" {
		endProgram <- true
	}
	publicar(text)
	readInput(endProgram)
}
