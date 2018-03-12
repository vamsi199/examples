package main

import (
	"fmt"
	"strings"
)

func main(){
	s:="hello my name is vamsi"
	m:=map[string]int{}
//
	for _,val:=range s{
		i:=1
		_,exists:=m[string(val)]
		if exists{
			i=m[string(val)]
			fmt.Println("here")
			i++
			m[string(val)]=i
			continue
		}
		m[string(val)]=i
	}
	fmt.Println(m)
	b:=[]byte(s)
	s3:=" "
	b3:=[]byte(s3)
	fmt.Println(b3)
	j:=1
	c:=make([]byte,len(b),cap(b))
	for i,val:=range b{
		if val== b3[0]{
			c[i]=val
			c[i+1]=strings.ToUpper(string(b[i+1]))[0]
			j++
			fmt.Println(string(c[i+1]))
			continue
		}
		if j==2{
			j--
			continue
		}
		if i==0 {
			fmt.Println(i)
			c[i] = strings.ToUpper(string(b[i]))[0]
			continue
		}
		c[i]=b[i]

	}

	fmt.Println("final",string(c))




}
