package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	upgrade bool
	ln      net.Listener
)

func init() {
	flag.BoolVar(&upgrade, "upgrade", false, "graceful restart")
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3*time.Second)
	fmt.Fprintf(w, "hello world, Tang; pid: %d, ppid: %d\n", os.Getpid(), os.Getppid())
}

func main() {
	fmt.Printf("start.\n")
	flag.Parse()
	http.HandleFunc("/", hello)

	server := &http.Server{Addr: ":9999"}

	var err error
	if upgrade {
		fmt.Printf("upgrade.\n")
		// 3，表示从父进程继承过来的 socket fd；由于已经有了 stdin 0,stdout 1, stderr 2，因此 fd 的序号从3开始
		fd := os.NewFile(3, "")
		ln, err = net.FileListener(fd)
		if err != nil {
			fmt.Printf("file listener fail, err: %s\n", err)
			os.Exit(1)
		}
		fd.Close()
	} else {
		fmt.Printf("not upgrade.\n")
		if ln, err = net.Listen("tcp", server.Addr); err != nil {
			fmt.Printf("listen fail, addr: %s, err: %s\n", server.Addr, err)
			time.Sleep(1 * time.Second)
			os.Exit(1)
		}
	}
	go func() {
		fmt.Printf("go serve.\n")
		if err := server.Serve(ln); err != nil {
			fmt.Printf("serve fail, %s\n", err)
		}
	}()
	setupSignal(server)
}

func setupSignal(server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	fmt.Printf("sig: %#v\n", sig)
	switch sig {
	case syscall.SIGUSR2:
		fmt.Printf("sig: USR2\n")
		err := forkProcess()
		if err != nil {
			fmt.Printf("fork process fail: %s\n", err)
		}
		close(ch)
		fmt.Printf("sig: USR2, fork process ok\n")
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("server shutdown fail after fork process, err: %s\n", err)
		}
		fmt.Printf("sig: USR2, server shutdown ok\n")
	case syscall.SIGTERM, syscall.SIGINT:
		signal.Stop(ch)
		close(ch)
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("server shutdown fail, err: %s\n", err)
		}
		fmt.Printf("sig: term, server shutdown ok\n")
	}
}

func forkProcess() error {
	fmt.Printf("os.Args[0]: %s\n", os.Args[0])
	// 子进程可以默认继承父进程绝大多数的文件描述符
	// golang 标准库os/exec 只默认继承 stdin stdout stderr 这三个。
	cmd := exec.Command(os.Args[0], "-upgrade")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	l, ok := ln.(*net.TCPListener)
	if !ok {
		fmt.Printf("ln assert not ok\n")
	}
	lfd, err := l.File()
	if err != nil {
		return err
	}
	// fork之前，先把子进程继承的 fd 放到 ExtraFiles 中
	cmd.ExtraFiles = []*os.File{lfd}
	return cmd.Start()
}
