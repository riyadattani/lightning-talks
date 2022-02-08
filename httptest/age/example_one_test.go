package age_test

import (
	"github.com/matryer/is"
	"io"
	"lightning-talks/httptest/age"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAge(t *testing.T) {
	t.Run("returns riya's age", func(t *testing.T) {
		is := is.New(t)
		name := "riya"
		expectedAge := "21"

		req := httptest.NewRequest(http.MethodGet, "/"+name, nil)
		res := httptest.NewRecorder()

		age.AgeServer(res, req)

		is.Equal(res.Body.String(), expectedAge)
		is.Equal(res.Code, http.StatusOK)

		is.True(strings.Contains(req.URL.Path, name))
	})

	t.Run("returns tamara's age on an endpoint", func(t *testing.T) {
		is := is.New(t)
		name := "tamara"
		expectedAge := "18"

		ts := httptest.NewServer(http.HandlerFunc(age.AgeServer))
		defer ts.Close()

		res, err := http.Get(ts.URL + "/" + name)
		is.NoErr(err)

		defer res.Body.Close()

		gotAge, err := io.ReadAll(res.Body)
		is.NoErr(err)

		is.Equal(string(gotAge), expectedAge)
	})
}
