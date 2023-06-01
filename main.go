package main

import (
	"context"
	"encoding/json"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"os"
)

type Config struct {
	Username  string
	Password  string
	Ipaddress string
	Port      uint16
	Filename  string
}

func main() {
	configStr := os.Getenv("CONFIG")
	var config Config

	err := json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		fmt.Println("unmarshal failed with ", err)
		return
	}
	clientConfig, _ := auth.PasswordKey(config.Username, config.Password, ssh.InsecureIgnoreHostKey())
	remote := fmt.Sprintf("%s:%d", config.Ipaddress, config.Port)
	client := scp.NewClient(remote, &clientConfig)
	err = client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	f, _ := os.Open(config.Filename)

	defer client.Close()
	defer f.Close()
	//remoteFilepath := fmt.Sprintf("~/%s", config.Filename)
	//fmt.Println("remote filepath ", remoteFilepath)
	err = client.CopyFromFile(context.Background(), *f, config.Filename, "0777")
	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}
