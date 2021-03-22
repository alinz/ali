package session_test

import (
	"testing"

	"github.com/trisolaria/ali/pkg/crypt"
	"github.com/trisolaria/ali/pkg/session"
)

func TestUserSession(t *testing.T) {
	testCases := []struct {
		username      string
		password      string
		authenticator crypt.Authenticator
		expected      bool
	}{
		{
			username: "user1",
			password: "pass1",
			authenticator: crypt.AuthenticatorFunc(func(u, p string) bool {
				return u == "user1" && p == "pass1"
			}),
			expected: true,
		},
		{
			username: "user1",
			password: "pass1",
			authenticator: crypt.AuthenticatorFunc(func(u, p string) bool {
				return u == "user2" && p == "pass2"
			}),
			expected: false,
		},
		{
			username:      "user1",
			password:      "pass1",
			authenticator: nil,
			expected:      false,
		},
	}

	for _, testCase := range testCases {
		userSession := session.NewUserSession(testCase.authenticator)
		result := userSession.Authenticate(testCase.username, testCase.password)
		if result != testCase.expected {
			t.Errorf("expected %t, but got %t for %s:%s credential", testCase.expected, result, testCase.username, testCase.password)
		}
	}
}
