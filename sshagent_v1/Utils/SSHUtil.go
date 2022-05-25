package Utils

import (
    "net"
    "fmt"
    "bytes"
	"time"
    "golang.org/x/crypto/ssh"
	seelog "github.com/cihub/seelog"
)

type SSHMethod struct {
		session		*ssh.Session
}

//创建ssh链接
func (this *SSHMethod) SSHConnect( user, password, host string, port int ) bool {
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
         Timeout:           5 * time.Second,
        HostKeyCallback:    hostKeyCallbk,
    }

    // connet to ssh
    addr = fmt.Sprintf( "%s:%d", host, port )

    if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
        return false
    }

    // create session
    if session, err = client.NewSession(); err != nil {
		seelog.Critical("SSHConnect to ",host," Failed, ",err); seelog.Flush()
        return false
    }
	this.session = session
	return true
}

//执行命令并返回结果
func (this *SSHMethod) RunShellCmd(cmd string) string {

	if this.session == nil {
			seelog.Critical("session is null .") ; seelog.Flush()
			return ""
	}

    var stdOut, stdErr bytes.Buffer

    this.session.Stdout = &stdOut
    this.session.Stderr = &stdErr

    this.session.Run(cmd)
    return stdOut.String() 

}

//关闭链接
func (this *SSHMethod) Close() {
		if this.session != nil {
				//seelog.Info("Close session") ; seelog.Flush()
				this.session.Close()
		}
}

