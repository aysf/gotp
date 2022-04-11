# Connecting to Server via SFTP with Go

## How to use:
1. create folder in server home dir -- for receiving file from client `mkdir filefromclient`,
2. create folder in server home dir -- for client downloading `mkdir filefromserver`, then `touch filefromserver/server1` 
2. create `server.conf` in this app dir,
3. run `go run main.go`


## How to set configuration file:

put this into `server.conf` and change the username, password, address with your server credentials:

```sh
USERNAME=admin
PASSWORD=Golang2022
ADDRESS=103.172.204.217
PORT=22
```

## How to set the server
1. use paid server - i used [idcloudhost](https://console.idcloudhost.com/)
2. use vm ware - i used [fedora server](https://getfedora.org/id/server/download/)
    - to get ipadress from fedora run `ifconfig`
    - pick *inet* value as ip address setting for config file

## source: 
- [inanzzz](http://www.inanzzz.com/index.php/post/tjp9/golang-sftp-client-server-example-to-upload-and-download-files-over-ssh-connection-streaming)