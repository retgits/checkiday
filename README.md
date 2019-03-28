# Checkiday

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retgits/checkiday)

[Checkiday.com](https://www.checkiday.com/) is the world's most complete holiday listing website. There are at least 4300 unique holidays on the site that checkiday has verified for authenticity. This Go package, not endorsed by checkiday.com, serves as a simple wrapper on their API so that you can use their awesome data in your awesome app.

## Usage

Get today's holidays

```go
import "github.com/retgits/checkiday"

days, err := checkiday.Today()
if err != nil {
    fmt.Printf("Oh noes, an error occured: %s", err.Error())
}

for idx := range days.Holidays {
    fmt.Printf("Today is %s\n", days.Holidays[idx].Name)
}
```

Get the holidays of a specific date (must be formatted as `mm/dd/yyyy`)

```go
import "github.com/retgits/checkiday"

days, err := checkiday.On("07/30/2018")
if err != nil {
    fmt.Printf("Oh noes, an error occured: %s", err.Error())
}

for idx := range days.Holidays {
    fmt.Printf("On 07/30/2018 we celebrated %s\n", days.Holidays[idx].Name)
}
```

July 30th is a special day, since we celebrate **National Cheesecake Day**

## License

See the [LICENSE](./LICENSE) file in the repository

## Acknowledgements

A most sincere thanks to the team of [Checkiday](https://www.checkiday.com/), for building such an awesome service that I enjoy every day!

_This package is not endorsed by Checkiday.com_