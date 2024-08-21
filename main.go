package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	ftp "github.com/jlaffaye/ftp"
	godotenv "github.com/joho/godotenv"
)

var path_file *string
var name_file *string

func init() {

	path_file = flag.String("path", "", "path /path/to/file")
	name_file = flag.String("name", "", "name name_you_file")
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
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS12,
	}

	var c *ftp.ServerConn
	var err error

	for i := 0; i < 3; i++ {
		c, err = ftp.Dial(fmt.Sprintf("%s:%s", host, port), ftp.DialWithTimeout(20*time.Second), ftp.DialWithExplicitTLS(&tlsConfig))
		if err == nil {
			break
		}
		log.Printf("Dial attempt %d failed: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Dial: ", err)
	}

	err = c.Login("ftp_karga", "413321337")
	if err != nil {
		log.Fatal("Login: ", err)
	}

	data := bytes.NewReader(reader(*path_file))

	err = c.Stor(fmt.Sprintf("./upload/%s", *name_file), data)
	if err != nil {
		log.Fatal("Stor: ", err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal("Quit: ", err)
	}

	os.Exit(0)
}
