package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	ftp "github.com/jlaffaye/ftp"
	godotenv "github.com/joho/godotenv"
)

var path_file *string

func init() {

	path_file = flag.String("path", "", "path /path/to/file")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	host := os.Getenv("SERVER")
	port := os.Getenv("PORT")

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}

	c, err := ftp.Dial(fmt.Sprintf("%s:%s", host, port), ftp.DialWithTimeout(5*time.Second), ftp.DialWithExplicitTLS(&tlsConfig))
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

	os.Exit(0)
}
