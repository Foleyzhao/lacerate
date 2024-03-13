package main

import (
	"flag"
	"github.com/fatih/color"
	"lacerate/core/command"
	"lacerate/core/common"
	"lacerate/core/config"
	"lacerate/core/service"
	"os"
	"strconv"
)

var (
	// 命令行参数
	args []string
)

// 入口函数
func main() {
	_, _ = color.New(color.FgGreen).Fprintln(os.Stdout, common.Banner)

	flag.Parse()
	args = flag.Args()
	if len(args) == 0 || len(args) > 3 {
		command.PrintHelp()
		os.Exit(1)
	}

	switch args[0] {

	default:
		command.PrintHelp()
		os.Exit(1)
	case "init":
		command.Initialize()
	case "new":
		if len(args) == 2 {
			name := args[1]
			service.CreateMarkdown(name)
		} else {
			panic("the file name is missing.")
		}
	case "compile", "c":
		service.Compile()
	case "watch", "w":
		service.NewWatch(config.Config().Paths, config.Config().Suffix).Watcher()
		done := make(chan bool)
		<-done
	case "run":
		service.Compile()
		service.NewWatch(config.Config().Paths, config.Config().Suffix).Watcher()
		var port = 8090
		if len(args) == 2 {
			p, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			port = p
		}
		command.ListenHttpServer(port)
	case "http", "web":
		var port = 8090
		if len(args) == 2 {
			p, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			port = p
		}
		command.ListenHttpServer(port)
	}
}
