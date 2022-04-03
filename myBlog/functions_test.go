package myBlog

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApplication_Indexhandle(t *testing.T) {
	r := chi.NewRouter()
	Register(r)
	testServer := httptest.NewServer(r)

	req, err := http.NewRequest(http.MethodGet, testServer.URL+"/", nil)
	if err != nil {
		t.Errorf("%v", err)
	}
	//resp := httptest.NewRecorder()
	//indexhandler(resp, req)

	resp, err := http.DefaultClient.Do(req)
	log.Println(resp)
}

func TestApplication_PostContenthandler(t *testing.T) {
	r := chi.NewRouter()
	Register(r)
	testServer := httptest.NewServer(r)
	form := url.Values{}
	form.Add("title", "Active")
	form.Add("content", "Active CSS")

	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/addpost", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	req.Form = form
	resp := httptest.NewRecorder()
	postContenthandler(resp, req)
	//log.Println(resp.Body)

	if resp.Code != http.StatusFound {
		t.Errorf("expected %v but got %v", http.StatusFound, resp.Code)
	}

}
