package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

//below random string functions are based on Jon Calhoun code
const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

func producer(length int, ch chan string) string{
	for{
		ch <- RandString(length)
		time.Sleep(time.Second)
	}
}

func consumer(ch chan string){
	for s := range ch{
		fmt.Println("recebido: ", s, "apenas letras?", isLetter(s))
	}
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan string)
	
	go producer(5, ch)
	go consumer(ch)

	select{}
}
