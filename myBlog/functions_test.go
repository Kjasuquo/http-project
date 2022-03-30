package myBlog

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestRegister(t *testing.T) {
	r:= chi.NewRouter()
	Register(r)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil{
		return
	}
	resp:= httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var tests = []struct{
		
	}
}

func TestError(t *testing.T) {
	//router := chi.NewRouter()
	//
	//req, err := http.NewRequest(http.MethodGet, "/", nil)
	//resp := httptest.NewRecorder()
	//router.ServeHTTP(resp, req)
	var tests = []struct {
		input string
		want  string
	}{
		{"No Error", "No Error"},
	}

	for _, test := range tests {
		got := Error(test.input)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
