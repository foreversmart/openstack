package options

import (
	"net/url"
	"strconv"
)

type ShowTokenOpts struct {
	Nocatalog    *string `json:"nocatalog"`
	AllowExpired *string `json:"allow_expired"`
}

func (opts *ShowTokenOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.Nocatalog != nil {
			options.Add("nocatalog", *opts.Nocatalog)
		}

		if opts.AllowExpired != nil {
			options.Add("allow_expired", *opts.AllowExpired)
		}
	}

	return
}

type HeadTokenOpts struct {
	AllowExpired *bool `json:"allow_expired"`
}

func (opts *HeadTokenOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.AllowExpired != nil {
			options.Add("allow_expired", strconv.FormatBool(*opts.AllowExpired))
		}
	}

	return
}
