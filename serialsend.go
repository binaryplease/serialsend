package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

// Read command-line arguments
var (
	debug = kingpin.Flag("debug", "Enable debug mode. Also prints data to the console ").Short('v').Bool()
	sPort = kingpin.Flag("port", "Serialport to use e.g. /dev/ttyUSB1").Short('p').Required().String()
	sBaud = kingpin.Flag("baud", "Baudrate to send with").Default("115200").Short('b').Uint()
	sData = kingpin.Arg("data", "Data to send to the port").Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	// Set up options.
	options := serial.OpenOptions{
		PortName:        *sPort,
		BaudRate:        *sBaud,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	// Write bytes to the port.
	n, err := port.Write([]byte(*sData))
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	if *debug {
		fmt.Println("Wrote", n, "bytes.")
		fmt.Println(*sData)
	}
}
