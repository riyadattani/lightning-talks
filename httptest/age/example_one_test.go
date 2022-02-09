package age_test

import (
	"github.com/matryer/is"
	"lightning-talks/httptest/age"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAge(t *testing.T) {
	t.Run("returns riya's age", func(t *testing.T) {
		is := is.New(t)
		name := "riya"
		expectedAge := "21"

		req := httptest.NewRequest(http.MethodGet, "/"+name, nil)
		res := httptest.NewRecorder()

		age.AgeHandler(res, req)

		is.Equal(res.Body.String(), expectedAge)
		is.Equal(res.Code, http.StatusOK)
	})
}
