package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gihub.com/dalton02/rabbitTutorial/consumer"
	"gihub.com/dalton02/rabbitTutorial/publisher"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Você é um publisher (0) ou consumer (1): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print(text)
	if strings.TrimSpace(text) == "0" {
		publisher.Init()
	} else if strings.TrimSpace(text) == "1" {
		consumer.Init()
	} else {
		fmt.Println("Operação invalida, encerrando aplicação")
	}

}
