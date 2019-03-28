// Package checkiday allows you to make use of the world's most complete holiday listing website, checkiday.com.
// There are at least 4300 unique holidays on the site that checkiday has verified for authenticity. This Go
// package is not endorsed by checkiday.com and serves as a simple wrapper on their API.
package checkiday

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOn(t *testing.T) {
	Checkiday, err := On("27/03/2019")
	assert.Error(t, err)
	assert.EqualError(t, err, "checkiday error Could not parse date.")

	Checkiday, err = On("03/27/2019")
	assert.NoError(t, err)
	assert.Equal(t, Checkiday.Date, "03/27/2019")
	assert.Equal(t, len(Checkiday.Holidays), 8)
	assert.Equal(t, Checkiday.Error, "none")
}

func TestToday(t *testing.T) {
	today := time.Now().Local().Format("01/02/2006")
	CheckidayOn, err := On(today)
	assert.NoError(t, err)

	CheckidayToday, err := Today()
	assert.NoError(t, err)

	assert.Equal(t, CheckidayToday.Date, today)
	assert.Equal(t, CheckidayOn.Number, CheckidayToday.Number)
	assert.Equal(t, CheckidayToday.Error, "none")
}
