package main

import (
	"fmt"
	"flag"
	"log"
)

func main() {
	var str string
	flag.StringVar(&str,"s","some string","")
	flag.StringVar(&str,"string","some string","")
	var help bool
	flag.BoolVar(&help,"h",false,"")
	flag.BoolVar(&help,"help",false,"")
	setflag(flag.CommandLine)
	flag.Parse()
	if help {
		showHelp()
		return
	}
	if str != "" {
		log.Printf("This is string input %s \n",str)
	}

}
func showHelp() {
	fmt.Print(`
Usage:CLI Template [OPTIONS]

Options:

	-s, --string	Prints string input	(default='some string')
	-i, --int 		Prints int			(default=0)
	-e, --error		Prints a custom error.
	-w, --warning	Prints a warning

	-h,	--help 		prints this help info



`)


}

func setflag(flag *flag.FlagSet) {
	flag.Usage = func() {
		showHelp()
	}

}
