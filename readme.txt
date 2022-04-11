
Visual Studio Code was used as text editor. 
The only requirements are the Go standard library. 

How to run:

Assuming Go is installed and configured,

1) Start the server by running `go run main.go` in server directory
2) Start the client by running `go run main.go` in client directory

The clients sends "helloworld" encrypted, and the server decrypts
the message. You can change the message in main of client should you choose.

Note: The server must be started before the client.

