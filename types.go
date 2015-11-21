//
// types.go
// Copyright (C) 2015 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

type	Rate struct {
	Max			int			`json:"max"`
	NumRater	int			`json:"numRater"`
	Average		string		`json:"average"`
	Min			int			`json:"min"`
}

type	Tag struct {
	Count		int		`json:"count"`
	Name		string	`json:"name"`
	Title		string	`json:"title"`
}

type	Book struct {
	Rating			Rate		`json:"rating"`
	Subtitle		string		`json:"subtitle"`
	Author			[]string	`json:"author"`
	Pubdate			string		`json:"pubdate"`
	Tags			[]Tag		`json:"tags"`
	Origin_title	string		`json:"origin_title"`
	Image			string		`json:"image"`
	Binding			string		`json:"binding"`
	Translator		[]string	`json:"translator"`
	Catalog			string		`json:"catalog"`
	Pages			string		`json:"pages"`
	Images			map[string]string			`json:"images"`
	Alt				string			`json:"alt"`
	Id				string			`json:"id"`
	Publisher		string			`json:"publisher"`
	Isbn10			string			`json:"isbn10"`
	Isbn13			string			`json:"isbn13"`
	Title			string			`json:"title"`
	Url				string			`json:"url"`
	Alt_title		string			`json:"alt_title"`
	Author_intro	string			`json:"author_intro"`
	Summary			string			`json:"summary"`
	Series			map[string]string		`json:"series"`
	Price			string			`json:"price"`
}

type	BookEntry struct {
	Id			string			`json:"id"`
	Title		string			`json:"title"`
	Author		[]string		`json:"author"`
	Image		string			`json:"image"`
	Rating		string			`json:"rating"`
	Publisher	string			`json:"publisher"`
}

type	Summary struct {
	Name	string			`json:"library"`
	Books	[]BookEntry		`json:"books"`
}
