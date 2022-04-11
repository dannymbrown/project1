package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/dannymbrown/project1/cipher"
)

func main() {
	log.Println("Starting client...")
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 1000000
	for true {
		k := rand.Intn(max-min+1) + min
		kprime := cipher.TransformK(k)
		log.Printf("%s%d\n", "K: ", k)
		log.Printf("%s%d\n", "K-Prime: ", kprime)
		if kprime != 0 {
			log.Println("Establishing connection to server")
			conn, err := net.Dial("tcp", "127.0.0.1:4400")
			if err != nil {
				log.Panicln(err)
			}
			fmt.Fprintf(conn, strconv.Itoa(k))
			conn.Close()
			message := []byte("helloworld")
			for _, plaintext := range message {
				encrypted := cipher.EncryptMessage(string(plaintext), k)
				conn, err := net.Dial("tcp", "127.0.0.1:4400")
				if err != nil {
					log.Panicln(err)
				}
				log.Printf("\n%s:  %s\n%s:  %s\n", "Plaintext", string(plaintext), "Ciphertext", encrypted)
				fmt.Fprintf(conn, encrypted)
				conn.Close()
			}
			return
		}
	}
}
