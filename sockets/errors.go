package sockets

import "errors"

const (
	pkgPrefix = "github.com/cezarmathe/daemon/socket: "
)

var (
	// ErrChanNegativeBufSize signals negative channel buffer sizes
	ErrChanNegativeBufSize = errors.New(pkgPrefix + "channel buffer sizes cannot be negative")
)
