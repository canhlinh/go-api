package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	ContentJSON = "application/json"
)

type Render interface {
	Write() error
}

type JsonRender struct {
	writer http.ResponseWriter
	status int
	value  interface{}
}

func NewJsonRender(w http.ResponseWriter, status int, value interface{}) Render {
	return JsonRender{
		writer: w,
		status: status,
		value:  value,
	}
}

func (r JsonRender) Write() error {
	r.writer.Header().Set("Content-Type", ContentJSON)
	r.writer.WriteHeader(r.status)
	data, err := json.Marshal(r.value)
	if err != nil {
		return err
	}
	_, err = r.writer.Write(data)
	return err
}

type FileStream struct {
	writer      http.ResponseWriter
	reader      io.ReadCloser
	contentType string
	fileName    string
}

func NewFileStream(dst http.ResponseWriter, src io.ReadCloser, contentType string) Render {
	return FileStream{
		writer:      dst,
		reader:      src,
		contentType: contentType,
	}
}

func (r FileStream) Write() error {
	r.writer.Header().Set("Content-Disposition", "inline")
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.Header().Set("X-Frame-Options", "DENY")
	r.writer.Header().Set("Content-Security-Policy", "Frame-ancestors 'none'")
	r.writer.WriteHeader(http.StatusOK)
	_, err := io.Copy(r.writer, r.reader)
	return err
}
