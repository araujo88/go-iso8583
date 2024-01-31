package iso8583

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseMessage parses an ISO8583 message string into a structured Message.
func ParseMessage(messageStr string) (*Message, error) {
	if len(messageStr) < 20 { // Ensure there's enough data for MTI and bitmap
		return nil, fmt.Errorf("message too short")
	}

	mti := messageStr[:4]
	bitmapHex := messageStr[4:20]
	bitmap, err := parseBitmap(bitmapHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bitmap: %v", err)
	}

	message := &Message{
		MTI:    mti,
		Bitmap: bitmap,
		Fields: make(map[int]string),
	}

	messageData := messageStr[20:]
	position := 0

	for i := 1; i <= 64; i++ {
		if bitmap[i-1] {
			descriptor, exists := fieldDescriptors[i]
			if !exists {
				continue
			}

			if descriptor.Variable {
				if i == 63 {
					// Special handling for field 63 to capture the entire remaining messageData
					fieldValue := messageData[position:]
					message.Fields[i] = fieldValue
					break // Field 63 captures the remaining data, so we can exit the loop
				}

				if position+2 > len(messageData) {
					return nil, fmt.Errorf("out of range when attempting to read length prefix for field %d", i)
				}
				lengthPrefixStr := messageData[position : position+2]
				lengthPrefix, err := strconv.Atoi(lengthPrefixStr)
				if err != nil || lengthPrefix < 0 {
					return nil, fmt.Errorf("invalid length prefix for field %d: %v", i, err)
				}
				position += 2

				if position+lengthPrefix > len(messageData) {
					return nil, fmt.Errorf("out of range when attempting to read data for field %d", i)
				}
				fieldValue := messageData[position : position+lengthPrefix]

				// Check if the field is of type 'an' and remove trailing spaces
				if descriptor.Type == "an" {
					fieldValue = strings.TrimSpace(fieldValue)
				}

				message.Fields[i] = fieldValue
				position += lengthPrefix
			} else {
				if i == 43 {
					// Special handling for field 43 to capture exactly 40 characters
					if position+40 > len(messageData) {
						return nil, fmt.Errorf("out of range when attempting to read fixed-length field %d", i)
					}
					fieldValue := messageData[position : position+40]
					message.Fields[i] = fieldValue
					position += 40
				} else {
					if position+descriptor.Length > len(messageData) {
						return nil, fmt.Errorf("out of range when attempting to read fixed-length field %d", i)
					}
					fieldValue := messageData[position : position+descriptor.Length]
					message.Fields[i] = fieldValue
					position += descriptor.Length
				}
			}
		}
	}

	return message, nil
}

// GeneratingMessage involves taking a Message structure and converting it into a string
func GenerateMessage(msg *Message) string {
	result := msg.MTI
	bitmap := generateBitmap(msg.Bitmap)
	result += bitmap

	for i := 1; i <= 64; i++ {
		if msg.Bitmap[i] {
			value := msg.Fields[i]
			descriptor := fieldDescriptors[i]

			if descriptor.Variable {
				// Prefix the value with its length. Assume two digits for simplicity.
				lengthPrefix := fmt.Sprintf("%02d", len(value))
				result += lengthPrefix + value
			} else {
				result += value
			}
		}
	}

	return result
}
