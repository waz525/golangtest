package main

import (
    "flag"
    "fmt"
    "./Utils"  //通用类
)


func printUsage() {
	fmt.Println("  Usage: sshagent -f hostlist -c CmdStr")
	fmt.Println("           hostlist format: ip port password [username]")
}



func main() {
	var hostlistfile string
	var cmdstr string
	flag.StringVar(&hostlistfile, "f" , "./hostlist", "Host List File Path ." )
	flag.StringVar(&cmdstr, "c" , "", "Remote Command .")
	h := flag.Bool("h", false, "show help info")
	flag.Parse()

	if *h || cmdstr == "" {
		printUsage()
		return
	}

	if !Utils.IsFileExist(hostlistfile) {
		fmt.Println("  Error: "+hostlistfile+" is not exist !!!")
		return 
	}

	hostlist := Utils.Str2List( Utils.GetFileContent( hostlistfile ), "\n" , " ")
	for ind := 0 ; ind<len(hostlist) ;ind++   {
		host1 := hostlist[ind]
		var host,password,user string
		var port int
		user = "root"
		port = 22
		password = "123456"
		if len(host1) > 0 {
			host = host1[0]
		}
		if len(host1) > 1 {
			port = Utils.Atoi(host1[1])
		}
		if len(host1) > 2 {
			password = host1[2]
		}
		if len(host1) > 3 {
			user = host1[3]
		}
		if host != "" {
			rst := ""
			fmt.Println("==> Login "+user+"@"+host+":"+Utils.Itoa(port)+"("+password+") run "+cmdstr+" ...")
			var nSSHMethod Utils.SSHMethod
			if nSSHMethod.SSHConnect(user, password, host, port) {
				rst = nSSHMethod.RunShellCmd(cmdstr)
				nSSHMethod.Close()
			} else {
				rst = "SSH Failed!"
			}
			fmt.Println(rst)
		}
	}
	fmt.Println("Run Over!!!")
}

