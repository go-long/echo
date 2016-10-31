package middleware

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/go-long/echo"
	"github.com/go-long/echo/test"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	e.SetLogOutput(buf)
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := Recover()(echo.HandlerFunc(func(c echo.Context) error {
		panic("test")
	}))
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Status())
	assert.Contains(t, buf.String(), "PANIC RECOVER")
}
