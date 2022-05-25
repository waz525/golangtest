/*
	go env | grep GOROOT
	cd $GOROOT
	mkdir -p src/golang.org/x/
	cd src/golang.org/x/
	git clone https://github.com/golang/crypto.git

*/
package main

import (
    "net"
    "log"
    "fmt"
    "bytes"
    "golang.org/x/crypto/ssh"
)

func SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
    var (
        auth         []ssh.AuthMethod
        addr         string
        clientConfig *ssh.ClientConfig
        client       *ssh.Client
        session      *ssh.Session
        err          error
    )
    // get auth method
    auth = make([]ssh.AuthMethod, 0)
    auth = append(auth, ssh.Password(password))

    hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
    }

    clientConfig = &ssh.ClientConfig{
        User:               user,
        Auth:               auth,
        // Timeout:             30 * time.Second,
        HostKeyCallback:    hostKeyCallbk,
    }

    // connet to ssh
    addr = fmt.Sprintf( "%s:%d", host, port )

    if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
        return nil, err
    }

    // create session
    if session, err = client.NewSession(); err != nil {
        return nil, err
    }

    return session, nil
}

func runSsh(user, password, host string, port int, cmd string) string {

    var stdOut, stdErr bytes.Buffer

    session, err := SSHConnect( user, password, host, port )
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()

    session.Stdout = &stdOut
    session.Stderr = &stdErr

    //session.Run("if [ -d liujx/project ]; then echo 0; else echo 1; fi")
    session.Run(cmd)
    return stdOut.String() 

}

func main() {
	rst := runSsh("root", "85894259@j", "112.124.0.8", 22, "uptime")
	fmt.Print(rst)
}
