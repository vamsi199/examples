package main

import (
	"text/template"
	//"fmt"
	"log"
	"os"
)
var tpl *template.Template
func init(){
	tpl:=template.Must(template.ParseFiles("hello.gohtml"))
}
func main () {
	//tpl,err:=template.ParseFiles("hello.gohtml")
	//if err!=nil {
	//	log.Fatal("cannot parse files",err)
	//}
	err:=tpl.Execute(os.Stdout,43)
	if err!=nil {
		log.Fatal("execte",err)
	}
	os.Stdin()
}