package main

import (
	"fmt"
	"strings"
	"os"

	"log"
	"time"
)

func main() {
	name := os.Args[1]
	fmt.Println(name)
	content := "hello how are you"
	b:= []byte("hi how")
	reader := strings.NewReader(content)

	fi,err:=os.Create("hello.txt")
	if err != nil {
		fmt.Println("cannot create file",err)
	}
	log.Fatal(time.Now(),"hello")
	defer fi.Close()
	h:=string(b)
	i,_:=reader.Read(b)


	//io.Copy(fi,reader)

	fmt.Println(b,reader.Len(),i,h)
}