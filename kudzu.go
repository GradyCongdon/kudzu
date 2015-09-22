package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	// fmt.Printf("Hey, let's start:\n")
	const playlist = `Ghost Writing Pt.2 - Tim Hecker - Haunt Me, Haunt Me Do It Again		
		Ghost Writing Pt.3 - Tim Hecker - Haunt Me, Haunt Me Do It Again`
	// Previously had \n as 1st [token?] after ` to algin,
	//  this caused the \n if stagement in loop to fire, printing 'Song: '
	var info string = "Song"
	// var and type required outside func def
	// := can be used within func defin and does not have type cast
	// fmt.Println(info)
	var s scanner.Scanner
	// scanner package wtih Scanner method
	var artist bytes.Buffer
	// initalize bytes buffer to hold artist
	s.Init(strings.NewReader(playlist))
	// init scanner with stirngs package.NewReader( string const as param)
	s.Whitespace = 1 << '\t'
	// whitespace is set as a uint64 after initalinz to a var
	// so we take 1 to the power of tab (technciallt the byte equiv)
	// Previously with the OR operator to mask out both tab and space as our whitespace
	//fmt.Println(s.Whitespace)

	var token rune
	// TODO look up runes
	// One byte symbol? only not a actual letter repersentation on print tho
	for token != scanner.EOF {
		// loop until scanner reaches End Of File token

		// defer fmt.Println(info, ": ", artist.String(), "\n byebye")
		// print the final info once the EOF is reached
		// !!NOPE printed every scanner word parse [I think]
		// OUTPUT:
		// Album :  HauntMe,HauntMeDoItAgain
		//  byebye
		// Album :  HauntMe,HauntMeDoIt
		//  byebye
		// Album :  HauntMe,HauntMeDo
		//  byebye
		// Album :  HauntMe,HauntMe

		token = s.Scan()
		// sets token to return of scanner Scan [ a rune]
		// artist buffer expands as much as needed
		// s.TokenText retuns string for most recent scanned token [from s.Scan()]
		// Possibly scanner not returning ' ' (space), all output has no spaces
		// SWITCHED s.Whitespace to only tab as whitespace to print out spaces in output

		if s.TokenText() == "-" {
			fmt.Println(info, ": ", artist.String())
			artist.Reset()
			info = nextStage(info)
			continue
		}
		//My code probably wrong, possible var init on 1st tokenText() call then  compare that value
		// Does output split line on '-'

		if s.TokenText() == "\n" {
			fmt.Println(info, ": ", artist.String(), "\n")
			artist.Reset()
			info = "Song"
			continue
		}
		// Not 2 tokens ?
		// Works though

		artist.WriteString(s.TokenText())
		// write at the end so contiue can run to remove dash
	}
	fmt.Println(info, ": ", artist.String())
	// Runs after EOF breaks loops to print final line
}

func nextStage(info string) string {
	if info == "Song" {
		return "Artist"
	}
	if info == "Artist" {
		return "Album"
	}
	if info == "Album" {
		return "Song"
	}
	return "fuck"
	// Really dumb way to do states
}
