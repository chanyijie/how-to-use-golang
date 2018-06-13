package sshutils

import (
	_ "fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	_ "os"
)

// func main() {
// 	if len(os.Args) != 4 {
// 		log.Fatalf("Usage: %s <user> <host:port> <command>", os.Args[0])
// 	}

// 	client, session, err := connectToHost(os.Args[1], os.Args[2])
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer client.Close()
// 	defer session.Close()
// 	out, err := session.CombinedOutput(os.Args[3])
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(out))
// }

func createSSHConfigFromKey(user string) *ssh.ClientConfig {

	key, err := ioutil.ReadFile("/Users/david/do-key/do-pri")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}
func ConnectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	/*
		var pass string

		fmt.Print("Password: ")
		fmt.Scanf("%s\n", &pass)

		sshConfig := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{ssh.Password(pass)},
		}
		sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	*/
	sshConfig := createSSHConfigFromKey(user)
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
