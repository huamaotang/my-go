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
	fmt.Fprintf(w, "hello world, Tang; pid: %d, ppid: %d\n", os.Getpid(), os.Getppid())
}

func main() {
	flag.Parse()
	http.HandleFunc("/", hello)

	server := &http.Server{Addr: ":9999"}

	var err error
	if upgrade {
		fd := os.NewFile(3, "")
		ln, err = net.FileListener(fd)
		if err != nil {
			fmt.Printf("file listener fail, err: %s\n", err)
			os.Exit(1)
		}
		fd.Close()
	} else {
		if ln, err = net.Listen("tcp", server.Addr); err != nil {
			fmt.Printf("listen fail, addr: %s, err: %s\n", server.Addr, err)
			time.Sleep(1 * time.Second)
			os.Exit(1)
		}
	}
	go func() {
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
	switch sig {
	case syscall.SIGUSR2:
		err := forkProcess()
		if err != nil {
			fmt.Printf("fork process fail: %s\n", err)
		}
		close(ch)
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("server shutdown fail after fork process, err: %s\n", err)
		}
	case syscall.SIGTERM, syscall.SIGINT:
		signal.Stop(ch)
		close(ch)
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("server shutdown fail, err: %s\n", err)
		}
	}
}

func forkProcess() error {
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
	cmd.ExtraFiles = []*os.File{lfd}
	return cmd.Start()
}
