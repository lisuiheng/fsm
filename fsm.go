package fsm

import "github.com/dgrijalva/jwt-go"

const AUTH_KEY = "AUTHORIZATION"
const ROLE_KEY = "ROLE"
const NO_AUTH = "NO_AUTH"

// Error is a domain error encountered while processing chronograf requests
type Error string

func (e Error) Error() string {
	return string(e)
}

// General errors.
const (
	ErrUserNotFound = Error("user not found")
)

// Logger represents an abstracted structured logging implementation. It
// provides methods to trigger log messages at various alert levels and a
// WithField method to set keys for a structured log message.
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Error(...interface{})
	Errorf(template string, args ...interface{})
}

type Sensitivity string

const SENSITIVE_DATA = Sensitivity("[SENSITIVE DATA]")

func (s Sensitivity) MarshalJSON() ([]byte, error) {
	return []byte(`"[SENSITIVE DATA]"`), nil
}

type User struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Password Sensitivity `json:"Password"`
	Role     int         `json:"role"`
}

type AuthClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	UseAuth  bool   `json:"useAuth"`
	jwt.StandardClaims
}
