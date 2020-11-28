package main

import (
	"flag"
	"fmt"
	"github.com/raspi/fidas/pkg/client"
	"os"
)

func statusToStr(s client.Status) string {
	switch s {
	case client.Active:
		return `active`
	case client.Available:
		return `available`
	case client.Invalid:
		return `invalid`
	default:
		return `???`
	}
}

func main() {
	flag.Usage = func() {
		f := os.Args[0] // Executable name
		_, _ = fmt.Fprintf(os.Stdout, `Usage:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s <domain>`+"\n", f)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		_, _ = fmt.Fprintf(os.Stdout, `See --help`)
		os.Exit(0)
	}

	domain := flag.Arg(flag.NArg() - 1)

	c := client.New(client.DefaultHost, client.DefaultPort, client.DefaultProtocol)
	ans, err := c.Request(domain)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`Domain %q is %s`+"\n", ans.Domain, statusToStr(ans.Status))
}
