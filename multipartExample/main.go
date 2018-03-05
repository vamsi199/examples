package main

import (
	"bytes"
	"fmt"
	"github.com/craigivy/dalog"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"strconv"
	"time"
)

type Attachment struct {
	ID          string    `db:"id"`
	Created     time.Time `db:"created"`
	LastUpdated time.Time `db:"last_updated"`
	EventID   string `db:"event_id"`
	EventType string `db:"event_type"`
	FileType  string `db:"file_type"`
	Encoding  string `db:"encoding"`
	Filename  string `db:"filename"`
	Size      int    `db:"size"`
	Payload   []byte `db:"payload"`
}

const (
	id          = "id"
	filename    = "filename"
	fileType    = "file_type"
	encoding    = "encoding"
	size        = "size"
	lastUpdated = "last_updated"
	created     = "created"
)

type server struct {
	Log dalog.Log
}
type values struct {
	boundary    string
	body        []byte
	contentType string
}

var hs = server{dalog.NoContext()}

func main() {
	fmt.Println("listening on 8086 connect with /upload")
	router := mux.NewRouter()
	router.HandleFunc("/upload", handler).Methods("POST")
	err := http.ListenAndServe(":8086", router)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	attachments, err := getMultipartFiles(r)
	if err != nil {
		hs.Log.Error(err)
		return
	}

	form, err := DocumentMIMESerialize(attachments)
	if err != nil {
		hs.Log.Error(err)
		return
	}

	attachs, err := DocumentMIMEDeSerialize(form)
	if err != nil {
		hs.Log.Error(err)
		return
	}

	fmt.Println(attachs)

}

func getMultipartFiles(r *http.Request) (attachments []Attachment, err error) {
	var attachs = []Attachment{}
	multiPartReader, err := r.MultipartReader()
	if err != nil {
		hs.Log.Error(err)
		return attachs, err
	}
	form, err := multiPartReader.ReadForm(32 << 30)
	if err != nil {
		hs.Log.Error(err)
		return attachs, err
	}
	fileHeaders := form.File
	for _, val := range fileHeaders {
		attach :=Attachment{}
		for _, fileHeader := range val {
			file, err := fileHeader.Open()
			if err != nil {
				hs.Log.Error(err)
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				hs.Log.Error(err)
			}
			attach.Created=time.Now()
			attach.LastUpdated=time.Now()
			attach.Payload = b
			fn := strings.SplitN(fileHeader.Filename, ".", 2)
			attach.FileType = fn[1]
			attach.Encoding = "" //TODO figure out about the encoding
			attach.Filename = fileHeader.Filename
			attach.Size = int(fileHeader.Size)
			attach.Payload = b
		}
		attachs = append(attachs, attach)
	}
	return attachs, nil
}
func DocumentMIMESerialize(attachs []Attachment) (form *multipart.Form, err error) {
	//create a multipart file
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	for _, attach := range attachs {
		mw.WriteField(id, attach.ID)
		mw.WriteField(fileType, attach.FileType)
		mw.WriteField(lastUpdated, attach.LastUpdated.Format("Mon, 02 Jan 2006 15:04:05 -0700")) //TODO confirm the time format
		mw.WriteField(created, attach.Created.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
		mw.WriteField(encoding, attach.Encoding)
		mw.WriteField(size, strconv.Itoa(attach.Size))
		mw.WriteField(filename, attach.Filename)
		writer, err := mw.CreateFormFile(attach.Filename, attach.Filename)
		if err != nil {
			return nil, errors.Wrap(err, "error in CreateFormFile")
		}
		_, err = writer.Write(attach.Payload)
		if err != nil {
			return nil, errors.Wrap(err, "error in write file")
		}

	}
	boundary:=mw.Boundary()
	err = mw.Close()
	if err != nil {
		return nil, errors.Wrap(err, "error in closing the multipart writer")
	}
	reader := multipart.NewReader(bytes.NewReader(body.Bytes()), boundary)
	form, err = reader.ReadForm(32 << 30)
	if err != nil {
		return nil, errors.Wrap(err, "error in ReadForm")
	}
	return form, nil
}

func DocumentMIMEDeSerialize(form *multipart.Form) (attachs []Attachment, err error) {
	values := form.Value
	attachs = make([]Attachment, 4) //TODO discuss about length
	for key, value := range values {
		for i, val := range value {
			switch key {
			case id:
				attachs[i].ID = val
			case filename:
				attachs[i].Filename = val
			case fileType:
				attachs[i].FileType = val

			case encoding:
				attachs[i].Encoding = val

			case size:
				attachs[i].Size, err = strconv.Atoi(val)
				if err != nil {
					hs.Log.Error(err)
					return
				}
			case lastUpdated:
				attachs[i].LastUpdated, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", val)
				if err != nil {
					hs.Log.Error(err)
					return
				}
			case created:
				attachs[i].Created, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", val)
				if err != nil {
					hs.Log.Error(err)
					return
				}

			}

		}

	}
	fileHeaders := form.File
	for _, val := range fileHeaders {
		for i, fileHeader := range val {
			attach := Attachment{}
			file, err := fileHeader.Open()
			if err != nil {
				return nil, errors.Wrap(err, "error in file header open")
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, errors.Wrap(err, "error in file readall")
			}
			attachs[i].Payload = b
			fmt.Println("finally", string(attachs[i].Payload))
			attachs = append(attachs, attach)
		}
	}
	return attachs, nil
}