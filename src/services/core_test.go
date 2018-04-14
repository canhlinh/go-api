package services

import (
	"testing"

	"github.com/canhlinh/go-api/src/stores"
	"github.com/canhlinh/go-api/src/utils"
	"github.com/stretchr/testify/assert"
	"goji.io"
)

func TestMain(m *testing.M) {
	utils.Init("../i18n")
	m.Run()
}

func TestNewSrv(t *testing.T) {
	srv := NewServer(goji.NewMux(), stores.NewStore())
	assert.NotNil(t, srv)
	assert.NotNil(t, srv.Router, "Router should not be nil")
	assert.NotNil(t, srv.Store, "Store should not be nil")
}
