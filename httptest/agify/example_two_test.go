package agify_test

import (
	"encoding/json"
	"github.com/matryer/is"
	"lightning-talks/httptest/agify"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAgify_GetAge(t *testing.T) {
	t.Run("get the age of a person using agify", func(t *testing.T) {
		is := is.New(t)

		name := "riya"
		expectedAge := 21

		testAgifyData := agify.APIResponse{
			Name:  name,
			Age:   expectedAge,
			Count: 1234,
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := json.NewEncoder(w).Encode(testAgifyData)
			is.NoErr(err)
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		agifyClient := agify.NewAgify(ts.URL)

		gotAge, err := agifyClient.GetAge(name)
		is.NoErr(err)

		is.Equal(gotAge, expectedAge)
	})
}
