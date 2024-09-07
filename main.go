package main

import (
	"bufio"
	"log"
	"os"

	"github.com/agent-e11/educationalsp/rpc"
)

func main() {
	logger := getLogger("/home/eamon/coding/go/src/educationalsp/lsp.log")
	logger.Println("Hey, I started!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
