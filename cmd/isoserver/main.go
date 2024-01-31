package main

import (
	"fmt"
	"log"

	"github.com/araujo88/go-iso8583/tree/main/internal/iso8583"
)

func main() {
	messageStr := "08002038000000200002810000000001084909052253415630305A303537363331202020205341564E47583130303131303032303030302020202020200011010008B9F3F723CA3CD2F8"

	// Parse the message
	parsedMessage, err := iso8583.ParseMessage(messageStr)
	if err != nil {
		log.Fatalf("Failed to parse message: %v", err)
	}

	fmt.Printf("Parsed Message:\nMTI: %s\n", parsedMessage.MTI)
	for i, present := range parsedMessage.Bitmap {
		if present {
			fmt.Printf("Field %d: %s\n", i+1, parsedMessage.Fields[i+1])
		}
	}

	responseMsg := iso8583.Message{
		MTI: "0210",
		Fields: map[int]string{
			39: "00", // Response code field (39) set to "00" indicating approval
		},
	}
	responseMsgStr := iso8583.GenerateMessage(&responseMsg)
	fmt.Println(responseMsgStr)

}
