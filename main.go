package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"io"
	"log"
	"time"

	ftp "github.com/jlaffaye/ftp"
)

var path_file *string

func init() {

	path_file = flag.String("path", "", "path /path/to/file")
	flag.Parse()
}

func main() {
	tlsConfig := tls.Config{
		InsecureSkipVerify: true, 
	}

	c, err := ftp.Dial("137.184.13.150:21", ftp.DialWithTimeout(5*time.Second), ftp.DialWithExplicitTLS(&tlsConfig))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("ftp_karga", "413321337")
	if err != nil {
		log.Fatal(err)
	}

	

	data := bytes.NewReader(reader(*path_file))

	err = c.Stor("./upload/my_test_44.txt", io.Reader(data)) 
	if err != nil {
		log.Fatal(err)
	}



	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
