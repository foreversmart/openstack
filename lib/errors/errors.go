package errors

import "errors"

var (
	ErrNotFound                   = errors.New("Not found!")
	ErrInvalidParams              = errors.New("Invalid params")
	ErrInvalidPassword            = errors.New("You must provide username, password and tenant id.")
	ErrInvalidCredential          = errors.New("You must provide either username/password or token values.")
	ErrInvalidToken               = errors.New("You must provide token.")
	ErrTokenID                    = errors.New("Invalid response token ID")
	ErrInvalidAuth                = errors.New("Unknown authentication.")
	ErrInvalidTenant              = errors.New("You must provide tenant id or tenant name.")
	ErrNotImplemented             = errors.New("Not implemented!")
	ErrInvalidRegion              = errors.New("Invalid region.")
	ErrInvalidProvider            = errors.New("Invalid platform provider.")
	ErrInvalidFloatingipRateLimit = errors.New("The rate limit of floating ip must between 1 and 1000.")
	ErrEmptySubnetValue           = errors.New("Cannot find subnet of the network")
	ErrEmptyAddress               = errors.New("Cannot find the address of the vm")
)
