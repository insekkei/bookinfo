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
	"strings"

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

// Read the file
func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// generate the summary json file
func generate(dirPath string) ([]byte, error) {
	files := make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	PathSep := string(os.PathSeparator)
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(strings.ToUpper(file.Name()), "JSON") {
			files = append(files, dirPath + PathSep + file.Name())
		}
	}

	var book Book
	summary := Summary {
		Name:	"My Libaray",
		Books:	[]BookEntry{},
	}
	for _, file := range files {
		fileData, err := ReadFile(file)
		if err != nil {
			log.Error(err)
		}

		err = json.Unmarshal(fileData, &book)
		if err != nil {
			log.Error(err)
		}
		if book.Isbn13 == "" {
			continue
		}
		bookentry := BookEntry{}
		bookentry.Id		= book.Isbn13
		bookentry.Title		= book.Title
		bookentry.Image		= book.Image
		bookentry.Rating	= book.Rating.Average
		bookentry.Publisher	= book.Publisher
		for _, name := range book.Author {
			bookentry.Author = append(bookentry.Author, name)
		}
		summary.Books = append(summary.Books, bookentry)
	}

	for _, b := range summary.Books {
		log.Infof("%s %s %s", b.Id, b.Title, b.Author)
	}
	log.Infof("RESULT: The number of books: %d", len(summary.Books))

	res, err := json.Marshal(summary.Books)
	if err != nil {
		log.Error(err)
	}

	return res, err
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


func generateSummary(c *cli.Context) {
	dir := c.String("dir")
	if dir == "" {
		dir = "books"
		log.Info("The dir does not set, use the default directory \"books\"")
	}
	output := c.String("output")
	if output == "" {
		log.Fatal("The output file must be set")
	}

	body, err := generate(dir)
	if err != nil {
		log.Fatal(err)
	}

	res, err := parseBody(body)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(output, res, 0644)
	if err != nil {
		log.Fatal(err)
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
	app.Commands = []cli.Command{
		{
			Name:		"get",
			Usage:		"get book data from douban.com",
			HideHelp:	true,
			Action:		getBookInfo,
			Flags:		[]cli.Flag{
							cli.StringFlag{
								Name:	"isbn, i",
								Usage:	"the isbn13 of the book",
							},
							cli.StringFlag{
								Name:	"output, o",
								Usage:	"file to which to save",
							},
			},
		},
		{
			Name:		"generate",
			Usage:		"generate book summary",
			HideHelp:	true,
			Action:		generateSummary,
			Flags:		[]cli.Flag{
							cli.StringFlag{
								Name:	"dir, d",
								Usage:	"the json file directory",
							},
							cli.StringFlag{
								Name:	"output, o",
								Usage:	"file to which to save",
							},
			},

		},
	}

	app.Run(os.Args)
}
