package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)

	defer connection.Close()

	fmt.Println("Enter your username :")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	logFatal(err)

	username = strings.Trim(username, " \r\n")

	welcomeMessage := fmt.Sprintf("Welcome %s, say hi to your friends.\n", username)

	connection.Write([]byte(welcomeMessage))

}
