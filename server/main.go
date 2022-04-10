package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/dannymbrown/project1/cipher"
)

func main() {
	fmt.Println("Starting server...")
	listener, _ := net.Listen("tcp", ":4400")
	i := 0
	k := 0
	for {
		conn, _ := listener.Accept()
		encryptedMessage, _ := bufio.NewReader(conn).ReadString('\n')
		if i == 0 {
			k, _ = strconv.Atoi(encryptedMessage)
			kprime := cipher.TransformK(k)
			fmt.Printf("%s%d\n", "Server K: ", k)
			fmt.Printf("%s%d\n", "Server K-Prime: ", kprime)
			i++
			continue
		}
		fmt.Println("Encrypted Message Received (ciphertext): ", encryptedMessage)
		decryptedMessage := cipher.DecryptMessage(encryptedMessage, k)
		fmt.Println("Message Decrypted (plaintext): ", decryptedMessage)
		conn.Close()
		i++
	}
}
