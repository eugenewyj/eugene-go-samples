// Sample test to show how to mock an HTTP GET call internally.
// Differs slightly from the book to show more.
package ch0913

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u1713"
const ballotX = "\u2717"

// mockServer returns a pointer to a server to handle the get call.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>Going Go Programming</title>
		<description>Golang : https://github.com/goinggo</description>
		<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
			<title>Object Oriented Programming Mechanics</title>
			<description>Go is an object oriented language.</description>
			<link>http://www.goinggo.net/2015/03/object-oriented</link>
		</item>
	</channel>
</rss>`

// mockServer return a pointer to a server to handle the get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Gvien the need to test downloading content.")
	{
		t.Logf("\tWhen check \"%s\" for status code \"%d\"", server.URL, statusCode)
		{

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatal("\t\tShould be able to unmarshal the response.", ballotX, err)
			}
			t.Log("\t\tSould be able to unmarshal the response.", checkMark)

			if len(d.Channel.Items) == 1 {
				t.Log("\t\tShould have \"1\" item in the feed.")
			} else {
				t.Errorf("\t\tShould have \"1\" item in the feed.", ballotX, len(d.Channel.Items))
			}
		}
	}
}
