package models

import (
	"github.com/canhlinh/go-api/src/utils"
	"github.com/gin-gonic/gin"
)

type Error struct {
	ID         string             `json:"id"`
	StatusCode int                `json:"-"`
	Message    string             `json:"message"`
	Errors     []*ErrorValidation `json:"errors"`
	RequestID  string             `json:"request_id"`
}

type ErrorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewError(id string, params map[string]interface{}, statusCode int) *Error {
	return &Error{
		ID:         id,
		StatusCode: statusCode,
		Message:    utils.T(id, params),
	}
}

func NewErrorUnexpected(err error, statusCode int) *Error {
	return NewError("unexpected.app_error", map[string]interface{}{
		"Message": err.Error(),
	}, statusCode)
}

func (err *Error) Render(c *gin.Context) {
	requestID, exist := c.Get("request_id")
	if exist {
		err.RequestID = requestID.(string)
	}
	c.JSON(err.StatusCode, err)
	c.Abort()
}
