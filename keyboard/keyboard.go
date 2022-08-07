//Package keyboard reads input from user through keyboard
package keyboard

import (
	"bufio"
	"strconv"
	"os"
	"strings"
	)

// GetFloat reads a floating-point number from the keyboard 
// it returns the number read and any error encountred.
func GetFloat() (float64, error) {
	// GetFloat code here
	reader:= bufio.NewReader(os.Stdin);
	input, err := reader.ReadString('\n');
	if err != nil {
		return 0, err;
	}
	input = strings.TrimSpace(input);
	in, err := strconv.ParseFloat(input, 64);
	if err != nil {
		return 0, err;
	}
	return in, nil;
}

