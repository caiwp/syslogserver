package syslogserver

import (
	"log"
	"net"
)

const maxBufferSize = 1024

func ListenUDP(address string, handler Handler) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}

	defer pc.Close()

	buffer := make([]byte, maxBufferSize)

	for {
		n, _, err := pc.ReadFrom(buffer)
		if err != nil {
			return err
		}

		parser := NewParser(buffer[:n])
		if err = parser.Parse(); err != nil {
			log.Println(err)
			continue
		}

		handler.Handle(parser.Dump())
	}
}
