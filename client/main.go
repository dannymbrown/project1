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

// public class Client {
//     /**
//      * Runs the client as an application.  First it displays a dialog
//      * box asking for the IP address or hostname of a host running
//      * the date server, then connects to it and displays the date that
//      * it serves.
//      */
//     public static void main(String[] args) throws IOException {
//         String serverAddress = JOptionPane.showInputDialog(
//             "Enter IP Address of a machine that is\n" +
//             "running the date service on port 9090:");
//         Socket s = new Socket(serverAddress, 9090);
//         BufferedReader input =
//             new BufferedReader(new InputStreamReader(s.getInputStream()));
//         String answer = input.readLine();
//         JOptionPane.showMessageDialog(null, answer);
//         System.exit(0);
//     }
// }

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
			log.Println("Sending message to server")
			fmt.Fprintf(conn, strconv.Itoa(k))
			conn.Close()
			// y = ek’(x) ≡ (x + k’) mod 52 (where x is a letter in the plaintext and y is a letter in
			// ciphertext transmitted to the Server)
			//Y = (X + K) mod 26
			message := []byte("abcdefg")
			// a := (message)
			for _, char := range message {
				cipherChar := cipher.ToCipher(string(char))
				conn, err := net.Dial("tcp", "127.0.0.1:4400")
				if err != nil {
					log.Panicln(err)
				}
				log.Printf("\n%s:  %s\n%s:  %s\n", "Plaintext", string(char), "Ciphertext", string(cipherChar))
				fmt.Fprintf(conn, cipherChar)
				conn.Close()
			}
			return
		}
	}
}
