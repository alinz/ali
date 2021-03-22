package session

import (
	"github.com/trisolaria/ali/pkg/crypt"
)

// UserSession is a new type for authenticating user using
// given authenticator from crypt.Authenticator interface
type UserSession struct {
	authenticator crypt.Authenticator
}

// Make sure UserSession is always satisfying crypt.Authenticator interface
// at compile time
var _ crypt.Authenticator = (*UserSession)(nil)

// Authenticate uses the given authenticator at initialization of UserSession
// to Authenticate the user
func (s *UserSession) Authenticate(username, password string) bool {
	if s.authenticator == nil {
		return false
	}

	return s.authenticator.Authenticate(username, password)
}

// NewUserSession creates a new UserSession based on given authenticator
// NOTE: if a nil authenticator pass to this function, all future Authenticate calls
// will be result of false value
func NewUserSession(authenticator crypt.Authenticator) *UserSession {
	return &UserSession{
		authenticator: authenticator,
	}
}
