package keyboard

import (
	"bufio"
	"strconv"
	"os"
	"strings"
	)

func GetFloat() (float64, error) {
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

