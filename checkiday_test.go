// Package checkiday allows you to make use of the world's most complete holiday listing website, checkiday.com.
// There are at least 4300 unique holidays on the site that checkiday has verified for authenticity. This Go
// package is not endorsed by checkiday.com and serves as a simple wrapper on their API.
package checkiday

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func teardown() {
	server.Close()
}

func fixture(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	return ioutil.ReadAll(f)
}

func TestOn(t *testing.T) {

	t.Run("check invalid date format", func(t *testing.T) {
		setup()
		defer teardown()

		checkidayURL = server.URL

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			f, err := fixture(filepath.Join("testdata", "bad-date-error.json"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = w.Write(f)
			if err != nil {
				t.Fatal(err)
			}
		})

		_, err := On("27/03/2019")

		assert.Error(t, err)
		assert.EqualError(t, err, "checkiday error Could not parse date.")
	})

	t.Run("check valid date format", func(t *testing.T) {
		setup()
		defer teardown()

		checkidayURL = server.URL

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			f, err := fixture(filepath.Join("testdata", "good-response.json"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = w.Write(f)
			if err != nil {
				t.Fatal(err)
			}
		})

		cd, err := On("03/27/2019")
		assert.NoError(t, err)
		assert.Equal(t, "03/27/2019", cd.Date)
		assert.Equal(t, 9, len(cd.Holidays))
		assert.Equal(t, "none", cd.Error)
	})
}

func TestToday(t *testing.T) {

	today := time.Now().Local().Format("01/02/2006")

	t.Run("check valid date format", func(t *testing.T) {
		setup()
		defer teardown()

		checkidayURL = server.URL

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			f, err := fixture(filepath.Join("testdata", "good-response.json"))
			if err != nil {
				t.Fatal(err)
			}
			f = bytes.ReplaceAll(f, []byte("03/27/2019"), []byte(today))
			_, err = w.Write(f)
			if err != nil {
				t.Fatal(err)
			}
		})

		cd, err := Today()
		assert.NoError(t, err)
		assert.Equal(t, today, cd.Date)
		assert.Equal(t, int64(9), cd.Number)
		assert.Equal(t, "none", cd.Error)
	})

}
