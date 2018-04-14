package services

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/canhlinh/go-api/src/models"
)

const (
	ContentJSON = "application/json; charset=utf-8"
)

type RenderFunc func(httpCode int, value interface{}) error
type HandlerFunc func(c *Context) RenderFunc

type Context struct {
	Srv          *Srv
	Request      *http.Request
	Writer       http.ResponseWriter
	RequestID    string
	UserID       *string
	ResponseCode int
	ResponseData interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request, srv *Srv) *Context {
	return &Context{
		RequestID: models.NewID(),
		Request:   r,
		Writer:    w,
		Srv:       srv,
	}
}

func (c *Context) JSON(statusCode int, value interface{}) RenderFunc {
	c.ResponseCode = statusCode
	c.ResponseData = value
	return c.RenderJSON
}

func (c *Context) Error(apperr *models.Error) RenderFunc {
	apperr.RequestID = c.RequestID
	c.ResponseCode = apperr.StatusCode
	c.ResponseData = apperr
	return c.RenderJSON
}

func (c *Context) Log(level logrus.Level, message string) {
	logEntry := logrus.WithFields(logrus.Fields{
		"request_id": c.RequestID,
	})
	logEntry.WriterLevel(level).Write([]byte(message))
}

func (c *Context) RenderJSON(httpCode int, v interface{}) error {
	c.Writer.WriteHeader(httpCode)
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.Writer.Header().Add("Content-Type", ContentJSON)
	c.Writer.Write(data)
	return nil
}
