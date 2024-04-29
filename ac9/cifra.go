package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func encrypt(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return ((c-'A'+22)%26) + 'A'
	}
	return c
}

func decrypt(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return ((c-'A'+4)%26) + 'A'
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		text := scanner.Text()
		scanner.Scan()
		shift, _ := strconv.Atoi(scanner.Text())
		encryptedText := strings.Map(decrypt, text)
		fmt.Println(encryptedText)
	}
}
