//
// main.go
// Copyright (C) 2015 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"os"
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/codegangsta/cli"
	log "github.com/Sirupsen/logrus"
)

const (
	URL	= "https://api.douban.com/v2/book/isbn/"
	VERSION = "v0.0.1"
)

// This function add Indent to json strings
func parseBody(body []byte) ([]byte, error) {
	var	out bytes.Buffer
	err := json.Indent(&out, body, "", "\t")
	return out.Bytes(), err
}

func getBookInfo(c *cli.Context) {
	arg := c.String("isbn")
	if arg == "" {
		log.Fatal("The isbn must be set")
	}

	// The douban url is : https://api.douban.com/v2/book/isbn/:name
	url := URL + arg
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res, err := parseBody(body)
	if err != nil {
		log.Fatal(err)
	}

	output := c.String("output")
	if output == "" {
		fmt.Printf("%s\n", res)
	} else {
		err := ioutil.WriteFile(output, res, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}


// The MAIN entry
func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "fetch book data from douban.com"
	app.Author = "Datawolf"
	app.Email = "wanglong@laoqinren.net"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:	"isbn, i",
			Usage:	"the isbn13 of the book",
		},
		cli.StringFlag{
			Name:	"output, o",
			Usage:	"file to which to save",
		},
	}
	app.Action = getBookInfo

	app.Run(os.Args)
}
