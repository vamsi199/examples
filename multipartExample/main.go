package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/craigivy/dalog"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"gitlab.com/skyrepublic/sky/pkg/backendevent"
	"gitlab.com/skyrepublic/sky/pkg/event"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Error holds the Message required for json response
type Error struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}

var (

	// errLimitExeeded will be returened by the getMultipartFiles
	// when we get the attachements more than the limit
	errLimitExeeded = errors.New("attachment limit exeeded")

	// errNotExists will be returned by the getMultipartFiles
	//when we doesn't get the payload
	errNotExists = errors.New("payload doesn't exists")

	// errDuplication will be returned by the getMultipartFiles
	//when we get the duplicate files
	errDuplication = errors.New("file duplication not allowed")

	// errInRead will be returned by the getMultipartFiles
	//when we get the error while reading the multipart file
	errInRead = errors.New("error in reading the documents")
)

type Attachment struct {
	ID          string    `db:"id"`
	Created     time.Time `db:"created"`
	LastUpdated time.Time `db:"last_updated"`
	EventID     string    `db:"event_id"`
	EventType   string    `db:"event_type"`
	FileType    string    `db:"file_type"`
	Encoding    string    `db:"encoding"`
	Filename    string    `db:"filename"`
	Size        int64     `db:"size"`
	Payload     []byte    `db:"payload"`
}

const (
	id                        = "id"
	filename                  = "filename"
	fileType                  = "file_type"
	encoding                  = "encoding"
	size                      = "size"
	lastUpdated               = "last_updated"
	created                   = "created"
	skyType                   = "sky_type"
	version                   = "version"
	documentLevelCMSFieldName = "document_level_cms"
	documentLevelCMSFileName  = "document_level_cms_file"
)

type server struct {
	Log dalog.Log
}

var hs = server{dalog.NoContext()}

func main() {
	fmt.Println("listening on 8087")
	router := mux.NewRouter()
	router.HandleFunc("/app/{app_id}/ref/{user_ref}/events", handler).Methods("POST")
	err := http.ListenAndServe(":8087", router)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var backendevent backendevent.BackendEvent
	params := mux.Vars(r)
	if userRef, exists := params["user_ref"]; exists {
		backendevent.UserRef = userRef
		fmt.Println(userRef)
	}
	if appID, exists := params["app_id"]; exists {
		backendevent.AppID = appID
		fmt.Println(appID)
	}
	if backendevent.UserRef == "" || backendevent.AppID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "userref or app id is missing"})
		return
	}

	attachments, err := getMultipartFiles(r, w)
	if err != nil {
		//http.Error(w,err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Code: http.StatusBadRequest})

		return
	}

	data, boundary, err := DocumentMIMESerialize(attachments)
	if err != nil {
		hs.Log.Error(err)
		return
	}

	attachs, err := DocumentMIMEDeSerialize(data, boundary)
	if err != nil {
		hs.Log.Error(err)
		return
	}

	fmt.Println(attachs)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		hs.Log.Error(errors.Wrap(err, "error in read all"))
		return
	}
	ev := event.Header{ID: uuid.NewV4().String(), SkyType: created, Version: 1, Created: time.Now()}
	serviceLevelMIME, boundary, err := ServiceLevelMIMESerialize(b, ev)
	if err != nil {
		hs.Log.Error(errors.Wrap(err, "error service serialize"))
		return
	}
	doc, eve, err := ServiceLevelMIMEDeSerialize(serviceLevelMIME, boundary)
	if err != nil {
		hs.Log.Error(errors.Wrap(err, "error service deserialize"))
		return
	}
	fmt.Println(doc, eve)

}

func getMultipartFiles(r *http.Request, w http.ResponseWriter) (attachments []Attachment, err error) {
	var attachs = []Attachment{}
	multiPartReader, err := r.MultipartReader()
	if err != nil {
		hs.Log.Error(errors.Wrap(err, "error in multiapart reader"))
		return []Attachment{}, errInRead
	}
	form, err := multiPartReader.ReadForm(32 << 30) //TODO make sure about the max memory
	if err != nil {
		hs.Log.Error(errors.Wrap(err, "error in read form"))
		return []Attachment{}, errInRead
	}
	fileHeaders := form.File
	attachementLimit := 0
	var keys []string
	for key, val := range fileHeaders {
		if attachementLimit > 3 {
			hs.Log.Error(errors.Wrap(errLimitExeeded, "received more than the attachement limit"))
			return []Attachment{}, errLimitExeeded
		}
		attachementLimit++
		duplicationCheck := 0
		attach := Attachment{}
		for _, fileHeader := range val {
			if duplicationCheck > 0 {
				hs.Log.Error(errors.Wrap(errDuplication, "file received multipale times"))
				return []Attachment{}, errDuplication
			}
			duplicationCheck++
			file, err := fileHeader.Open()
			if err != nil {
				hs.Log.Error(errors.Wrap(err, "error in open file"))
				return []Attachment{}, errInRead
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				hs.Log.Error(errors.Wrap(err, "error in file readall"))
				return []Attachment{}, errInRead
			}
			attach.Payload = b
			fn := strings.SplitN(fileHeader.Filename, ".", 2)
			attach.FileType = fn[1]
			attach.Encoding = "" //TODO figure out about the encoding
			attach.Filename = fileHeader.Filename
			attach.Size = fileHeader.Size
			attach.Payload = b
		}
		attachs = append(attachs, attach)
		keys = append(keys, key)
	}
	exists := ContainsString(keys, "toto")
	if !exists {
		hs.Log.Error(errors.Wrap(errNotExists, "payload not received"))
		return []Attachment{}, errNotExists
	}
	return attachs, nil
}

func DocumentMIMESerialize(attachs []Attachment) (data []byte, boundary string, err error) {
	//create a multipart file
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	for _, attach := range attachs {
		mw.WriteField(id, attach.ID)
		mw.WriteField(fileType, attach.FileType)
		mw.WriteField(lastUpdated, attach.LastUpdated.Format("Mon, 02 Jan 2006 15:04:05 -0700")) //TODO confirm the time format
		mw.WriteField(created, attach.Created.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
		mw.WriteField(encoding, attach.Encoding)
		mw.WriteField(size, strconv.Itoa(int(attach.Size)))
		mw.WriteField(filename, attach.Filename)
		writer, err := mw.CreateFormFile(attach.Filename, attach.Filename)
		if err != nil {
			return nil, "", errors.Wrap(err, "error in CreateFormFile")
		}
		_, err = writer.Write(attach.Payload)
		if err != nil {
			return nil, "", errors.Wrap(err, "error in write file")
		}
	}
	boundary = mw.Boundary()
	err = mw.Close()
	if err != nil {
		return nil, "", errors.Wrap(err, "error in closing the multipart writer")
	}
	fmt.Println(body.String())
	return body.Bytes(), boundary, nil
}

func DocumentMIMEDeSerialize(data []byte, boundary string) (attachs []Attachment, err error) {
	reader := multipart.NewReader(bytes.NewReader(data), boundary)
	form, err := reader.ReadForm(32 << 30)
	if err != nil {
		return nil, errors.Wrap(err, "error in ReadForm")
	}
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
				attachs[i].Size, err = strconv.ParseInt(val, 10, 64)
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

// ServiceLevelMIMESerialize will parse all the attachements and returns multipart form
func ServiceLevelMIMESerialize(documentLevelCMS []byte, eventHeader event.Header) (serviceLevelMIME []byte, boundary string, err error) {
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	mw.WriteField(id, eventHeader.ID)
	mw.WriteField(skyType, eventHeader.SkyType)
	mw.WriteField(version, strconv.Itoa(int(eventHeader.Version))) //TODO
	mw.WriteField(created, eventHeader.Created.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
	writer, err := mw.CreateFormFile(documentLevelCMSFieldName, documentLevelCMSFileName)
	if err != nil {
		return nil, "", errors.Wrap(err, "error in CreateFormFile")
	}
	_, err = writer.Write(documentLevelCMS)
	if err != nil {
		return nil, "", errors.Wrap(err, "error in write file")
	}
	boundary = mw.Boundary()
	err = mw.Close()
	if err != nil {
		return nil, "", errors.Wrap(err, "error in closing the multipart writer")
	}
	return body.Bytes(), boundary, nil
}

// ServiceLevelMIMEDeSerialize will get the documentlevelCMS out of the serviceLevelMIME
func ServiceLevelMIMEDeSerialize(serviceLevelMIME []byte, boundary string) (documentLevelCMS []byte, eventHeader event.Header, err error) {

	reader := multipart.NewReader(bytes.NewReader(serviceLevelMIME), boundary)
	form, err := reader.ReadForm(32 << 30)
	if err != nil {
		return nil, event.Header{}, errors.Wrap(err, "error in ReadForm")
	}
	values := form.Value
	for key, value := range values {
		switch key {
		case id:
			eventHeader.ID = value[0]
		case skyType:
			eventHeader.SkyType = value[0]
		case version:
			num, err := strconv.Atoi(value[0])
			if err != nil {
				hs.Log.Error(err)
			}
			eventHeader.Version = uint(num)
		case created:
			eventHeader.Created, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", value[0])
			if err != nil {
				hs.Log.Error(err)
			}

		}
	}
	fileHeaders := form.File
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader[0].Open()
		if err != nil {
			return nil, event.Header{}, errors.Wrap(err, "error in file header open")
		}
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, event.Header{}, errors.Wrap(err, "error in file readall")
		}
		documentLevelCMS = b
	}

	return documentLevelCMS, eventHeader, nil
}

func ContainsString(slice []string, str string) bool {
	for _, check := range slice {
		if check == str {
			return true
		}
	}
	return false
}
