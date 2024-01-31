package iso8583

import (
	"fmt"
	"strconv"
)

// parseBitmap converts a hexadecimal string representation of a bitmap
// into a boolean array indicating the presence of fields.
func parseBitmap(bitmapHex string) ([128]bool, error) {
	var bitmap [128]bool

	// Check if the bitmapHex length is exactly 16 characters (64 bits)
	if len(bitmapHex) != 16 {
		return bitmap, fmt.Errorf("invalid bitmap length: got %d, want 16", len(bitmapHex))
	}

	for i, hexChar := range bitmapHex {
		// Convert hex character to its 4-bit binary equivalent
		val, err := strconv.ParseUint(string(hexChar), 16, 4)
		if err != nil {
			return bitmap, fmt.Errorf("failed to parse hex: %v", err)
		}

		// Map each bit to the boolean array
		for j := 0; j < 4; j++ {
			// Calculate the bit's position in the boolean array
			// (i * 4) shifts the bit to its correct quartet, +j adjusts for the bit's position within the quartet
			position := (i * 4) + j
			bitmap[position] = val&(1<<(3-j)) != 0
		}
	}

	return bitmap, nil
}

// generateBitmap converts a boolean array representing the presence of fields
// into a hexadecimal string representation of the bitmap.
func generateBitmap(bitmap [128]bool) string {
	var bitmapHex string

	for i := 0; i < 64; i += 4 {
		// Accumulate 4 bits to form a single hex digit
		var val byte
		for j := 0; j < 4; j++ {
			val <<= 1
			if bitmap[i+j] {
				val |= 1
			}
		}
		// Convert the 4-bit value to a hex character and append it to the result string
		bitmapHex += fmt.Sprintf("%X", val)
	}

	return bitmapHex
}
