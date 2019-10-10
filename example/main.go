package main

import (
	"flag"

	"github.com/caiwp/syslogserver"
)

var (
	Addr       string
	Path       string
	MaxSize    int
	MaxAge     int
	MaxBackups int
)

func init() {
	flag.StringVar(&Addr, "addr", ":9889", "udp server address")
	flag.StringVar(&Path, "path", "data", "file write to path")
	flag.IntVar(&MaxSize, "maxSize", 30, "file max size (M)")
	flag.IntVar(&MaxAge, "maxAge", 30, "file max age")
	flag.IntVar(&MaxBackups, "maxBackups", 30, "file max backups")
}

func main() {
	flag.Parse()

	if err := syslogserver.ListenUDP(Addr, NewHandler(Path, MaxSize, MaxAge, MaxBackups)); err != nil {
		panic(err)
	}
}
