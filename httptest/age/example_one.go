package age

import (
	"fmt"
	"net/http"
	"strings"
)

func AgeServer(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	_, _ = fmt.Fprint(w, GetAge(name))
}

func GetAge(name string) string {
	if name == "riya" {
		return "21"
	}

	if name == "tamara" {
		return "18"
	}

	return ""
}
