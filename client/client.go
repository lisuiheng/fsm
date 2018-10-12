package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

type uploadResp struct {
	Fid      string
	FileName string
	FileUrl  string
	Size     int64
	Error    string
}

type Filer struct {
	nrUrl string
	orUrl string
	trUrl string
}

type Rep int

const (
	NO_REP  Rep = 0
	ONE_REP Rep = 1
	TWO_REP Rep = 2
)

func (f *Filer) Upload(pathname string, mimeType string, file io.Reader, rep Rep) (r *uploadResp, err error) {
	formData, contentType, err := makeFormData(pathname, mimeType, file)
	if err != nil {
		return
	}

	if !strings.HasPrefix(pathname, "/") {
		pathname = "/" + pathname
	}

	url := ""
	switch rep {
	case NO_REP:
		url = f.nrUrl
	case ONE_REP:
		url = f.orUrl
	case TWO_REP:
		url = f.trUrl
	default:
		err = errors.New(fmt.Sprintf("invalid Rep %d", rep))
		return
	}
	resp, err := http.Post(url+pathname, contentType, formData)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	upload := new(uploadResp)
	if err = decodeJson(resp.Body, upload); err != nil {
		return
	}

	if upload.Error != "" {
		err = errors.New(upload.Error)
		return
	}

	r = upload

	return
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func createFormFile(writer *multipart.Writer, fieldname, filename, mime string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	if len(mime) == 0 {
		mime = "application/octet-stream"
	}
	h.Set("Content-Type", mime)
	return writer.CreatePart(h)
}

func makeFormData(filename, mimeType string, content io.Reader) (formData io.Reader, contentType string, err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, err := createFormFile(writer, "file", filename, mimeType)
	//log.Println(filename, mimeType)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(part, content)
	if err != nil {
		log.Println(err)
		return
	}

	formData = buf
	contentType = writer.FormDataContentType()
	//log.Println(contentType)
	writer.Close()

	return
}

func decodeJson(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
