package iso8583

import (
	"regexp"
	"strings"
)

// PadRight pads a string on the right with a specified pad character up to a total length
func PadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) >= length {
			return str[0:length]
		}
	}
}

// TrimLeft removes all leading occurrences of a set of characters specified in a cutset
func TrimLeft(str, cutset string) string {
	return strings.TrimLeft(str, cutset)
}

// HexToBCD converts a hexadecimal string to a BCD-encoded string
func HexToBCD(hexStr string) (string, error) {
	var bcdStr string
	// Implementation of conversion logic
	return bcdStr, nil
}

// BCDToHex converts a BCD-encoded string to a hexadecimal string
func BCDToHex(bcdStr string) (string, error) {
	var hexStr string
	// Implementation of conversion logic
	return hexStr, nil
}

// ValidateNumericField checks if a string contains only numeric characters
func ValidateNumericField(field string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(field)
}

// ValidateFieldLength checks if a field's length is within allowed limits
func ValidateFieldLength(field string, maxLength int) bool {
	return len(field) <= maxLength
}

// IsFieldPresent checks if a field is present in the bitmap
func IsFieldPresent(bitmap [128]bool, field int) bool {
	if field < 1 || field > 128 {
		return false
	}
	return bitmap[field-1]
}
