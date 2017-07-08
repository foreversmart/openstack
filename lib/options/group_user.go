package options

import (
	"net/url"
)

/*
Filter results based on which user passwords have expired.
The query should include an operator and a timestamp with a colon (:) separating the two, for example:

password_expires_at={operator}:{timestamp}
Valid operators are: lt, lte, gt, gte, eq, and neq
lt: expiration time lower than the timestamp
lte: expiration time lower than or equal to the timestamp
gt: expiration time higher than the timestamp
gte: expiration time higher than or equal to the timestamp
eq: expiration time equal to the timestamp
neq: expiration time not equal to the timestamp
Valid timestamps are of the form: YYYY-MM-DDTHH:mm:ssZ.
For example:

/v3/users?password_expires_at=lt:2016-12-08T22:02:00Z
The example would return a list of users whose password expired before the timestamp (2016-12-08T22:02:00Z).
*/
type ListGroupUserOpts struct {
	GroupID           string  `json:"group_id"`
	PasswordExpiresAt *string `json:"password_expires_at"`
}

func (opts ListGroupUserOpts) IsValid() bool {
	if opts.GroupID == "" {
		return false
	}

	return true
}

func (opts ListGroupUserOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts.PasswordExpiresAt != nil {
		options.Add("password_expires_at", *opts.PasswordExpiresAt)
	}

	return
}

type BaseGroupUserOpts struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

func (opts *BaseGroupUserOpts) IsValid() bool {
	if opts.UserID == "" || opts.GroupID == "" {
		return false
	}

	return true
}
