# ping-pong
Echo request server/client

Here is an example of a ping-pong server in Go that uses a TOML configuration file to store public keys and only responds to pings from clients with the correct private key:

To run the client, you need to pass the path of the private key file as the first argument. For example:

go run ping.go private_key.pem

This client will connect to the server, send its public key and the "ping" message, and read the response from the server. If the server is configured to only respond to clients with the correct private key, this client should receive a "pong" response.
