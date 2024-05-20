package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var h1, m1, h2, m2 int
		scanner.Scan()
		line := scanner.Text()

		if line == "0 0 0 0" {
			break
		}

		parts := strings.Fields(line)
		h1, _ = strconv.Atoi(parts[0])
		m1, _ = strconv.Atoi(parts[1])
		h2, _ = strconv.Atoi(parts[2])
		m2, _ = strconv.Atoi(parts[3])


		minutesUntilAlarm := h2*60 + m2 - (h1*60 + m1)
		if minutesUntilAlarm < 0 {
			minutesUntilAlarm += 24 * 60 
		}

		fmt.Println(minutesUntilAlarm)
	}
}
