// Sample test to show how to write a basic unit table test.
package ch0912

import (
	"testing"
	"net/http"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload validates the http Get function can downloaded
// content and handles different status conditions properly.
func TestDownload(t *testing.T) {
	var urls = []struct{
		url 	string
		statusCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK,
		},
		{
			"http://www.baidu.com?zz=ttt",
			http.StatusBadRequest,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to get the url.", ballotX, err)
				}
				t.Log("\t\tShould be able to get the url.", checkMark)
				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}

