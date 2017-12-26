package main

import (
	"encoding/gob"
	"fmt"
	"bytes"
	"time"
	"github.com/prometheus/common/log"
)
type gobvar struct {
	Name string
	Id int
	Date time.Time
}
func main()  {
	var g gobvar
	//g:=gobvar{"vamsi",123,time.Now()}
	fmt.Println("got the input structure and starting marshaling")
	b,err:=Marshal(&g)
	if err!= nil {
		log.Error("error in marshaling",err)
	}
	fmt.Println("got the input bytes and starting unmarshaling")
	err=g.Unmarshal(b)
	if err!= nil {
		log.Error("error in unmarshaling",err)
	}
	fmt.Println(g)

}

func (v gobvar) Unmarshal(b []byte) error {
	fmt.Println("Unmarshal: I am here", v)
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	return dec.Decode(&v)
}

func Marshal(v *gobvar) ([]byte, error) {
	*v = gobvar{"vamsi",123,time.Now()}
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

