package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestMachine(t *testing.T) {
	values := url.Values{}

	values.Add("userName", "root1")
	resp, err := http.PostForm("http://127.0.0.1:8111/machine/check/user", values)
	if err != nil {
		t.Fatal(err)
	}
	by, err := httputil.DumpResponse(resp, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(by))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatal(resp.Status)
	}
}
