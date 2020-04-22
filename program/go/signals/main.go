package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 10)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)

	s := <-sig

	// シグナルを無視する
	signal.Ignore(syscall.SIGHUP)

	// シグナルを受取を停止する
	signal.Stop(sig)

	// シグナルをデフォルトに戻す
	signal.Reset(syscall.SIGHUP)

	switch s {
	case syscall.SIGHUP:
		fmt.Println("hungup")
	case syscall.SIGINT:
		fmt.Println("Warikomi")
	case syscall.SIGTERM:
		fmt.Println("force stop")
	case syscall.SIGKILL:
		fmt.Println("Kill")
	default:
		fmt.Println("Unknown signal.")
	}
}
