package main

import (
	"log"

	"github.com/tarm/serial"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Enable debug mode. Also prints data to the console ").Short('v').Bool()
	sPort = kingpin.Flag("port", "Serialport to use e.g. /dev/ttyUSB1").Short('p').Required().String()
	sBaud = kingpin.Flag("baud", "Baudrate to send with").Default("115200").Short('b').Int()
	sData = kingpin.Arg("data", "Data to send to the port").Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	c := &serial.Config{Name: *sPort, Baud: *sBaud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte(*sData))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q", []byte(*sData)[:n])
}
