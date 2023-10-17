package main

import (
	"fmt"
	"os"
	"strings"
)

type Options struct {
	action     string // "print" "convert" "view"
	sourcePath string
	destPath   string
	sourceKind string // "md" "csv"
}

func main() {
	opts := parseArgs(os.Args)

	data, alignment, err := Open(opts.sourcePath, opts.sourceKind)
	if err != nil {
		fmt.Printf("\nerror: %s\n", err.Error())
		return
	}

	switch opts.action {
	case "print":
		t := NewTable(data, alignment)
		fmt.Println(t)
	case "convert":
		//converted := Convert(data, opts.sourceKind)
		var err error
		if opts.sourceKind == "md" {
			err = Write(opts.destPath, "csv", data)
		} else if opts.sourceKind == "csv" {
			err = Write(opts.destPath, "md", data)
		}
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	case "view":
		fmt.Println("View has not been implemented.")
	}
}

func parseArgs(args []string) Options {
	if len(args) < 3 {
		fmt.Printf("\nerror: too few args\n%s", help())
	}
	var opts Options

	switch args[1] {
	case "print", "convert", "view":
		opts.action = args[1]
		opts.sourcePath = args[2]
		if args[1] == "convert" {
			if len(args) > 3 {
				opts.destPath = args[3]
			} else {
				fmt.Printf("\nerror: too few args\n%s", help())
			}

		}

	case "help":
		fmt.Println(help())
		os.Exit(0)
	default:
		fmt.Println(help())
		os.Exit(1)
	}

	opts.sourceKind = opts.sourcePath[strings.LastIndex(opts.sourcePath, "."):]
	if len(opts.sourceKind) > 1 {
		opts.sourceKind = opts.sourceKind[1:]
	} else {
		opts.sourceKind = ""
	}
	return opts
}

func help() string {
	return "\ncsview help\n" +
		"print\t\tDisplays table\n" +
		"\t\t\tUsage: csview print filename\n" +
		"convert\t\tConverts between CSV files and MarkDown files\n" +
		"\t\t\tUsage: csview convert sourcename destname\n" +
		"view\t\tDisplays scrollable table\n" +
		"\t\t\tUsage: csview view filename\n" +
		"help\t\tDisplays this help\n" +
		"\t\t\tUsage: csview help\n"
}
