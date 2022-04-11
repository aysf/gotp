package main

import (
	"fmt"
	"gotp/sftp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("error read home dir")
		log.Fatalln(err)
	}

	pk, err := ioutil.ReadFile(home + "/.ssh/id_rsa") // required only if private key authentication is to be used
	if err != nil {
		log.Println("error read private key")
		log.Fatalln(err)
	}

	config := sftp.Config{
		Username:   "aysf",
		Password:   "Geniox2022",
		PrivateKey: string(pk),
		Server:     "103.172.205.180:22",
		Timeout:    time.Second * 30,
	}

	client, err := sftp.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Open local file for reading
	source, err := os.Open("filetobeuploaded/file.html")
	if err != nil {
		log.Println("error open local file")
		log.Fatalln(err)
	}
	defer source.Close()

	// Create remote file for writing
	destination, err := client.Create("filefromclient/filecopied.html")
	if err != nil {
		log.Println("error create remote file")
		log.Fatalln(err)
	}
	defer destination.Close()

	// Upload local file to a remote location as in 1MB (byte) chunks.
	if err := client.Upload(source, destination, 1000000); err != nil {
		log.Println("error upload local file")
		log.Fatalln(err)
	}

	// Download remote file
	file, err := client.Download("filefromserver/server1")
	if err != nil {
		log.Println("error download local file")
		log.Fatalln(err)
	}
	defer file.Close()

	// Read downloaded file.
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))

	// Get remote file stats.
	info, err := client.Info("filefromclient/filecopied.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", info)
}
