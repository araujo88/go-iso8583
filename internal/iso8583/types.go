package iso8583

import (
	"errors"
)

// Message represents an ISO8583 message with a selection of fields and a bitmap
type Message struct {
	MTI    string
	Bitmap [128]bool // Indicates presence of fields; true for present
	Fields map[int]string
}

type FieldDescriptor struct {
	Length   int
	Variable bool
	Type     string
}

const (
	TypeNumeric      = "n"
	TypeAlpha        = "a"
	TypeBinary       = "b"
	TypeAlphanumeric = "an"
)

var (
	ErrInvalidBitmapLength = errors.New("invalid bitmap length")
	ErrInvalidMessageType  = errors.New("invalid message type indicator")
)

var fieldDescriptors = map[int]FieldDescriptor{
	2:  {Length: 19, Variable: true, Type: "n"},   // Primary account number, variable length up to 19 digits
	3:  {Length: 6, Variable: false, Type: "n"},   // Processing code, 6 digits
	4:  {Length: 12, Variable: false, Type: "n"},  // Transaction amount, 12 digits
	7:  {Length: 10, Variable: false, Type: "n"},  // Transmission Date and Time, 10 digits
	11: {Length: 6, Variable: false, Type: "n"},   // Systems Trace Audit Number (STAN), 6 digits
	12: {Length: 6, Variable: false, Type: "n"},   // Local Transaction Time, 6 digits
	13: {Length: 4, Variable: false, Type: "n"},   // Local Transaction Date, 4 digits
	43: {Length: 40, Variable: false, Type: "an"}, // Card Acceptor Name/Location, 40 characters
	63: {Length: 11, Variable: true, Type: "an"},  // Field 63, variable length alphanumeric, up to 11 characters
}
