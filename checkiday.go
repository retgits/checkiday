// Package checkiday allows you to make use of the world's most complete holiday listing website, checkiday.com.
// There are at least 4300 unique holidays on the site that checkiday has verified for authenticity. This Go
// package is not endorsed by checkiday.com and serves as a simple wrapper on their API.
package checkiday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Checkiday is the result coming from checkiday.com
type Checkiday struct {
	Error      string    `json:"error"`
	Date       string    `json:"date"`
	Holidays   []Holiday `json:"holidays"`
	Number     int64     `json:"number"`
	LastUpdate int64     `json:"lastUpdate"`
}

// Holiday is a struct to capture the details of every holiday sent back by checkiday
type Holiday struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const (
	checkidayURL = "https://www.checkiday.com/api/3/?d=%s"
)

// Today calls the checkiday.com API and returns all the holidays that are today
// (based on the locale of the machine) or an error.
func Today() (Checkiday, error) {
	return On(time.Now().Local().Format("01/02/2006"))
}

// On calls the checkiday.com API and returns all the holidays for the mentioned
// date or an error. The date must be in form of mm/dd/yyyy.
func On(date string) (Checkiday, error) {
	days := Checkiday{}

	req, err := http.NewRequest("GET", fmt.Sprintf(checkidayURL, date), nil)
	if err != nil {
		return days, fmt.Errorf("error creating HTTP request %s", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return days, fmt.Errorf("error executing HTTP request %s", err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return days, fmt.Errorf("error reading HTTP response %s", err.Error())
	}

	days, err = unmarshalCheckiday(body)
	if err != nil {
		return days, fmt.Errorf("error creating checkiday structure %s", err.Error())
	}

	if days.Error != "none" {
		return days, fmt.Errorf("checkiday error %s", days.Error)
	}

	return days, nil
}

func unmarshalCheckiday(data []byte) (Checkiday, error) {
	var r Checkiday
	err := json.Unmarshal(data, &r)
	return r, err
}
