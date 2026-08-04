package quorum

import "errors"

var ErrNoQuorum = errors.New("quorum was not reached")