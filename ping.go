package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"

    "golang.org/x/crypto/ssh"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: client [private_key_file]")
        return
    }

    // Read the private key from the file
    privateKeyBytes, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error reading private key file:", err)
        return
    }

    // Parse the private key
    privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
    if err != nil {
        fmt.Println("Error parsing private key:", err)
        return
    }

    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:7980")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    // Send the client's public key to the server
    conn.Write(ssh.MarshalAuthorizedKey(privateKey.PublicKey()))

    // Send the "ping" message
    conn.Write([]byte("ping"))

    // Read the response from the server
    buf := make([]byte, 1024)
    n, _ := conn.Read(buf)
    response := string(buf[:n])

    fmt.Println("Server response:", response)
}
