# Connecting to Server via SFTP with Go

## How to use:
1. create folder in the server in home dir `mkdir filefromclient`,
2. create `server.conf` in this app dir,
3. run `go run main.go`


## How to set configuration file:

put this into `server.conf` and set the username, password, address as your server credentials:

```sh
USERNAME=foo
PASSWORD=Abcd1234
ADDRESS=127.0.0.1
PORT=22
```

## source: 
- [inanzzz](http://www.inanzzz.com/index.php/post/tjp9/golang-sftp-client-server-example-to-upload-and-download-files-over-ssh-connection-streaming)