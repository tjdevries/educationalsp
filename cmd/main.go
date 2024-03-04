package main

import (
	"bufio"
	"fmt"
	"golsp"
	"os"
	"strconv"
	"time"
)

type Message struct {
	method   string
	contents interface{}
}

func main() {
	file, err := os.OpenFile("/home/tjdevries/git/go-lsp/log.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0775)
	if err != nil {
		panic("COULDNT OPEN THE LOG")
	}
	defer file.Close()

	stdinReader := bufio.NewScanner(os.Stdin)
	stdinReader.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		chunk, ok := golsp.Scan(data)
		if !ok {
			return 0, nil, nil
		}

		return chunk, data[:chunk], nil
	})

	stdoutWriter := os.Stdout

	file.WriteString("starting LSP server: " + strconv.Itoa(int(time.Now().Unix())) + "\n")

	for stdinReader.Scan() {
		message := stdinReader.Bytes()
		msg, err := golsp.DecodeMessage(message)
		if err != nil {
			file.WriteString("got err: " + fmt.Sprintf("%+v", err) + "\n")
			return
		}

		switch v := msg.(type) {
		case golsp.InitializeMessage:
			file.WriteString(fmt.Sprintf("got initialize message: %d\n", v.ID))
			response := golsp.EncodeMessageStruct(golsp.NewInitializeResponse(v.ID))
			stdoutWriter.WriteString(response)
			file.WriteString("... sent initialize response")
		default:
			file.WriteString("I DONT KNOW:" + string(message) + "\n")

		}
	}

}
