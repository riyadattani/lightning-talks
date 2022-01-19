package main

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAgify(t *testing.T) {
	t.Run("returns an age given a name", func(t *testing.T) {
		is := is.New(t)
		req, _ := http.NewRequest(http.MethodGet, "/riya", nil)
		res := httptest.NewRecorder()

		AgeServer(res, req)

		got := res.Body.String()
		want := "21"

		is.Equal(got, want)
	})
}
