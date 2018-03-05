package main

import (
	"bytes"
	"fmt"
	"github.com/craigivy/dalog"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/skyrepublic/sky/pkg/attachment"
	"io/ioutil"

	"mime/multipart"
	"net/http"
	"strings"
	"github.com/satori/go.uuid"
)

type server struct {
	Log dalog.Log
}

var hs = server{dalog.NoContext()}

func main() {
	fmt.Println("listening on 8084 connect with /upload")
	router := mux.NewRouter()
	router.HandleFunc("/upload", handler).Methods("POST")
	err := http.ListenAndServe(":8084", router)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	//get the multipart file from request
	multiPartReader, err := r.MultipartReader()
	if err != nil {
		hs.Log.Error(err)
	}
	form, err := multiPartReader.ReadForm(32 << 30)
	if err != nil {
		hs.Log.Error(err)
	}
	var attachs []attachment.Attachment
	fileHeaders := form.File
	for _, val := range fileHeaders {
		attach := attachment.Attachment{}
		for i, fileHeader := range val {
			file, err := fileHeader.Open()
			if err != nil {
				hs.Log.Error(err)
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				hs.Log.Error(err)
			}
			fn := strings.SplitN(fileHeader.Filename, ".", 2)
			attach.ID = uuid.NewV4().String()
			attach.FileType = fn[1]
			attach.Size = fileHeader.Size
			attach.Payload = b
			attach.Filename = fileHeader.Filename

			fmt.Println(i,"attachements comming from client:", string(attach.Payload))
			fmt.Println("attachements comming from client:size", attach.Size)
			fmt.Println("attachements comming from client:filename", attach.Filename)
			fmt.Println("attachements comming from client:filetype", attach.FileType)
		}
		attachs = append(attachs, attach)

	}

	fmt.Println("======================================creating a multipart file======================================")

	//==========================================================================
	//create a multipart file
	//get the four attachements for the given backend id

	type values struct {
		boundary string
		body []byte
		contentType string
	}

	mws := []values{}
	data:=[][]byte{}

	for i, attach := range attachs {

		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		//mw.CreatePart()
		//multipart.FileHeader{}
        mw.CreateFormField("here--------------")

		mw.WriteField("id", attach.ID)
		mw.WriteField("id", attach.ID)
		mw.WriteField("type", attach.FileType)
		mw.WriteField("last_updated", attach.LastUpdated.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
		mw.WriteField("created", attach.Created.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
		mw.WriteField("encoding", attach.Encoding)
		mw.WriteField("size", string(attach.Size))
		mw.WriteField("filename", attach.Filename)
		writer, err := mw.CreateFormFile(attach.Filename, attach.Filename)
		if err != nil {
			hs.Log.Error(errors.Wrap(err, "error in CreateFormFile"))
		}
		_, err = writer.Write(attachs[i].Payload)
		if err != nil {
			hs.Log.Error(errors.Wrap(err, "error in write file"))
		}
		err = mw.Close()
		if err != nil {
			hs.Log.Error(err)
		}
		mws=append(mws,values{contentType:mw.FormDataContentType(),boundary:mw.Boundary(),body:body.Bytes()})
		data=append(data,body.Bytes())

	}
	fmt.Println("attachements after mutlipart in bytes:%v",data)
	fmt.Println("======================================getting attachements from a multipart file======================================")

	//==========================================================================
	//get the attachements from multipart files

	//ç
	//	reader := multipart.NewReader(bytes.NewReader(mw.body), mw.boundary)
//
	//	for {
	//		p, err := reader.NextPart()
	//		if err == io.EOF {
	//			hs.Log.Debug("end of file in nextpart")
	//			break
	//		}
	//		if err != nil {
	//			fmt.Println(errors.Wrap(err, "error in next part"))
	//			return
	//		}
	//		formName := p.FormName()
	//		slurp, err := ioutil.ReadAll(p)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Printf("formname:=%s and formdata:=%s \n", formName,string(slurp))
//
	//	}
	//}

	reader := multipart.NewReader(bytes.NewReader(mws[0].body), mws[0].boundary)
		form1, err := reader.ReadForm(32 << 30)
		if err != nil {
			hs.Log.Error(errors.Wrap(err, "error in ReadForm"))
			return
		}
		fmt.Println("==============================", form1.Value)
        //values:=form1.Value
		fileHeaders = form1.File
		for _, val := range fileHeaders {
			attach := attachment.Attachment{}
			fileHeader := val[0]
			file, err := fileHeader.Open()
			if err != nil {
				hs.Log.Error(err)
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				hs.Log.Error(err)
			}
			//hs.Log.Debug(string(b), "--------------------")
			attach.Payload = b
			attachs = append(attachs, attach)
		}
		fmt.Println("finally",string(attachs[0].Payload))

		//fmt.Println("finally:====================82983912839812==================", string(attachs[0].Payload))

		//============================================================================

	}

