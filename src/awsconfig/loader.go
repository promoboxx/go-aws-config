package awsconfig

import (
	"crypto/rsa"
	"crypto/tls"
	"time"
)

type Loader interface {
	Import(data []byte) error
	Initialize() error
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error

	// Must functions will panic if they can't do what is requested.
	// They are maingly meant for use with configs that are required for an app to start up
	MustGetString(key string) string
	MustGetBool(key string) bool
	MustGetInt(key string) int
	MustGetDuration(key string) time.Duration
	MustGetObject(key string, obj interface{})
	MustGetPublicKey(key string) *rsa.PublicKey
	MustGetPrivateKey(key string) *rsa.PrivateKey
	MustGetCertificate(certKey string, privKeyKey string) *tls.Certificate
	MustGetEnv(key string) string
}
