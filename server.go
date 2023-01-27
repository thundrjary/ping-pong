package main

import (
    "fmt"
    "net"

    "github.com/BurntSushi/toml"
    "golang.org/x/crypto/ssh"
)

type Config struct {
    PublicKeys []ssh.PublicKey
}

func main() {
    // Load the config file
    var config Config
    _, err := toml.DecodeFile("keys.toml", &config)
    if err != nil {
        fmt.Println("Error loading config file:", err)
        return
    }

    // Listen on port 8080
    ln, _ := net.Listen("tcp", ":7980")
    defer ln.Close()

    // Accept incoming connections
    for {
        conn, _ := ln.Accept()
        go handleConnection(conn, config.PublicKeys)
    }
}

func handleConnection(conn net.Conn, publicKeys []ssh.PublicKey) {
    defer conn.Close()

    // Read the client's public key
    clientPublicKey, _, _, _, err := ssh.ParseAuthorizedKey(conn.RemoteAddr().(*net.TCPAddr).IP)
    if err != nil {
        fmt.Println("Error parsing client public key:", err)
        return
    }

    // Check if the client's public key is in the config file
    var authorized bool
    for _, key := range publicKeys {
        if ssh.KeysEqual(clientPublicKey, key) {
            authorized = true
            break
        }
    }

    if !authorized {
        fmt.Println("Unauthorized client:", conn.RemoteAddr())
        return
    }

    // Read data from the connection
    buf := make([]byte, 1024)
    n, _ := conn.Read(buf)
    message := string(buf[:n])

    // Check for "ping" message and respond with "pong"
    if message == "ping" {
        conn.Write([]byte("pong"))
    } else {
        fmt.Println("Received message:", message)
    }
}
