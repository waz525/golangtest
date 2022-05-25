package main

import (
		"os"
		"fmt"
		"time"
		logs "log"

)



func main() {
	  // when os.Getppid() != 1, not in daemon
	   if os.Getppid() != 1 {

		// create log file: when run in daemon success, but log file has not init, stdout/stderr will dup2 the log
		// 这里为什么这样做是因为： 如果我们daemon启动成功了， 父进程会返回， 子进程运行时脱离控制台， 当我们直接向控制台再次打印日志时， 会被定向到/dev/null， 所以如果有什么错误信息， 我们是看不到的。所以最好的做法时， 先创建一个文件， 将子进程输出到控制台的日志dup2 到该日志文件
			createLogFile := func (fileName string) (fd *os.File, err error) {
			dir := path.Dir(fileName)
			if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
				if err = os.MkdirAll(dir, 0755); err != nil {
					logs.Error("Start-Daemon: create dir: %s failed, err is: %v", dir, err)
					return
				}
			}

			if fd, err = os.Create(fileName); err != nil {
				logs.Error("Start-Daemon: create log file: %s failed, err is: %v", fileName, err)
				return
			}

			return
		}

		// logFd 的close放在子进程， 其实这里不知道会不会有问题。 fork时， 肯定会传递logFd到子进程， 但是通过StartProcess不晓得会不会？
		logFd, err := createLogFile(config.FileName)
		if err != nil {
			return
		}
		defer logFd.Close()

		// start Daemon
		cmdName := os.Args[0]
	   // 为什么这样直接就能创建守护进程？ 按理说， 应该 设置os.ProcAttr中Sys（*syscall.SysProcAttr）的Setsid为true才对？
	   /*
		type SysProcAttr struct {
			Chroot	   string		 // Chroot.
			Credential   *Credential	// Credential.
			Ptrace	   bool		   // Enable tracing.
			Setsid	   bool		   // Create session.
			Setpgid	  bool		   // Set process group ID to Pgid, or, if Pgid == 0, to new pid.
			Setctty	  bool		   // Set controlling terminal to fd Ctty (only meaningful if Setsid is set)
			Noctty	   bool		   // Detach fd 0 from controlling terminal
			Ctty		 int			// Controlling TTY fd
			Foreground   bool		   // Place child's process group in foreground. (Implies Setpgid. Uses Ctty as fd of controlling TTY)
			Pgid		 int			// Child's process group ID if Setpgid.
			Pdeathsig	Signal		 // Signal that the process will get when its parent dies (Linux only)
			Cloneflags   uintptr		// Flags for clone calls (Linux only)
			Unshareflags uintptr		// Flags for unshare calls (Linux only)
			UidMappings  []SysProcIDMap // User ID mappings for user namespaces.
			GidMappings  []SysProcIDMap // Group ID mappings for user namespaces.
			// GidMappingsEnableSetgroups enabling setgroups syscall.
			// If false, then setgroups syscall will be disabled for the child process.
			// This parameter is no-op if GidMappings == nil. Otherwise for unprivileged
			// users this should be set to false for mappings work.
			GidMappingsEnableSetgroups bool
			AmbientCaps				[]uintptr // Ambient capabilities (Linux only)
		}
		*/
		newProc, err := os.StartProcess(cmdName, os.Args, &os.ProcAttr{Files: []*os.File{logFd, logFd, logFd}})
		if err != nil {
			logs.Error("Start-Deamon: start process: %s failed, err is: %v", cmdName, err)
			return
		}

		logs.Trace("Start-Deamon: run in daemon success, pid: %v", newProc.Pid)
		return
	}

	for {
		// will write "test" to file 
		fmt.Prinln("test")
		time.Sleep(time.Duration(3) * time.Second)
	}
}
