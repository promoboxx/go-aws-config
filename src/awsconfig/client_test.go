package awsconfig

import (
	"testing"
	"time"
)

var (
	testPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7
mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBp
HssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2
XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3b
ODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy
7wIDAQAB
-----END PUBLIC KEY-----`

	// generated with openssl: req -x509 -nodes -days 9999 -newkey rsa:2048 -keyout privateKey.key -out certificate.crt
	testPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCzlf76gg7V7Xok
kv46kR1MmHDc/3MVRge9Iv3yzxgXdSQqBjDm04CFaeBb3Rg5AZ62WFbHZBPuNWt/
t0238iT2Cd0BFOPulHl1arqXiRihFYAeInLJaAgj0ILFH32xKvVKgGnFPhbzsu8Y
oFTwRbNIDrc+mpGDyHJf+cSJ9fflIuna1tjGy9BM/MoS3WX8KkqbNaHEFArTtfRY
TcrCHcTCR9Pb95o360eL+iMKZT4XBMfBxE9SUMVd/DZbpEDYUKNA6v/RyxnJCLg7
QEjWib5evWDdz1E/N0lapzHjsBmaroKQf1XK6zZJh9ir6gwrVFgST6P+SUQuKuKV
cwiCnkBXAgMBAAECggEAbwO9cJxvkV1RUUSw9gF47D2cH3GmcbMt24TDGZNd4Dpg
1b2oAzkhzNdrgz5E2BChTTWEx5OdIndRcc0dtSVyJcppHV8NnBGal7QXjs+IMyP4
ZCiFbu3pgkJGZJcX+yqEIb8KI41DYPjvBvkuKK3dqyC1tHSWmbGSyrO7BFHSIYmP
i/IEZL4AeIAcZEUhd/z53q4ScXEnb7Ero2O6+D5aIE93cuJnVv8m7dqzz/tZbZkg
MKs7AkQOYvJucwomHXVkazJRkWELw39Yr9esYYlfIY52zJtX3L++rpLxrR9Ncg8N
tGlu3Nz3aUNm9drxGY80/EaFjWLVuAcYNk3/ps7ZKQKBgQDY+Iu8z1mQ4y0gB+d3
CtPTrWeWArmb7vVCI3bkFdhfv29gdiFtE0yWJasUFu/YNrjCo2i/vBcwiO/BtC1K
HEDN5R8NhLi990dewVeMivjuxQSgtDStw9xahki4O0JWnGV0JDgzmXm+ovXCEg9F
cyysWyWNnjdmObQtj+5AVDSjawKBgQDT4+HYi3jC3qn9xEEpiAeMSGmsvxmVHJIR
AZciWW8NmmOtjIjs9xUV4Ju43Q6N5pmK1duKckUG9xpJDoKw2XK4N9fsIrsWLLdL
fbxzatO+UyxNnkhhuecaisfi80iFPf1uui72OVSTS2748CwVJAr97A6lLzCwsw0P
FUoauIQ9xQKBgQCEQEOUy+KxQPgBfS/mTNA/R4RLWM/gL3CJZuqSLoqcGij+aCMJ
xGi7YKx961k4tmo6Iba4oCKWb/GMZZHxiXUqy0z5RXwCNtbm9/ywawk/KRIgDpfJ
jwgimZV7zosqFdx1RZqIQTWHMPeR2sY6M/D4AfrK7rSf9+5Ok1vLFEidjwKBgQC+
CznZGt7pCQS2kntPYK5EZ/4/7fZoAwQPNLn1GPm93adhVRbKUqIaySViHQKcyyMT
ntQVzH+Uy7RLqjQVojJ+f7euF0htjxWnI9MOQdZAciDeTQTmgfKBn8/AAiwdNYhE
88CDHtB4e8PAisk+/ODO9hX8meK12SHxUUrxxGT3cQKBgAjn3dtMAB20nVCdEW7b
NJpH/1cn0uBvzFB0Z/8VO/Y+U2kVikZtSnm98h1WpvLN74bjMeCd6iqGz408D4Qj
IZvYSPFfvUFV7TnvnBhPztn3Pvy8jz3pwKenGsVU+Jy9fcT1tihYrNNwze21arU/
6m+Lh/4e1oTHMChLOUpx5p59
-----END PRIVATE KEY-----`

	testCertificate = `-----BEGIN CERTIFICATE-----
MIIDXjCCAkYCCQD0CAizfDHRsjANBgkqhkiG9w0BAQsFADBxMQswCQYDVQQGEwJV
UzELMAkGA1UECAwCTUExDzANBgNVBAcMBkJvc3RvbjEMMAoGA1UECgwDRk9PMQww
CgYDVQQLDANmb28xDDAKBgNVBAMMA2ZvbzEaMBgGCSqGSIb3DQEJARYLZm9vQGJh
ci5jb20wHhcNMjEwOTAzMjAyNTI2WhcNNDkwMTE4MjAyNTI2WjBxMQswCQYDVQQG
EwJVUzELMAkGA1UECAwCTUExDzANBgNVBAcMBkJvc3RvbjEMMAoGA1UECgwDRk9P
MQwwCgYDVQQLDANmb28xDDAKBgNVBAMMA2ZvbzEaMBgGCSqGSIb3DQEJARYLZm9v
QGJhci5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCzlf76gg7V
7Xokkv46kR1MmHDc/3MVRge9Iv3yzxgXdSQqBjDm04CFaeBb3Rg5AZ62WFbHZBPu
NWt/t0238iT2Cd0BFOPulHl1arqXiRihFYAeInLJaAgj0ILFH32xKvVKgGnFPhbz
su8YoFTwRbNIDrc+mpGDyHJf+cSJ9fflIuna1tjGy9BM/MoS3WX8KkqbNaHEFArT
tfRYTcrCHcTCR9Pb95o360eL+iMKZT4XBMfBxE9SUMVd/DZbpEDYUKNA6v/RyxnJ
CLg7QEjWib5evWDdz1E/N0lapzHjsBmaroKQf1XK6zZJh9ir6gwrVFgST6P+SUQu
KuKVcwiCnkBXAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAEdK5bU7vEymUkKvsT/f
VVs2AnWqo6pgb5PLS5ANcGy00u7fWampyQyhi4Wm6vmVVb9GEwXw1EP+BHOB5vVX
y8whwyFYJyAW5Eq7u3VYrx2xRb7inQnndueetBKFa7T1fH0DYplM7p2tFd7MZEp6
3O0zfn7ZS0qwq8a3aX1C6eXNLiY3X2Bt/6Kk1FWuK9XPk2E5tGa7tqhMDdANfaGY
sBJ+GH4nAXxldDkECB8BtBqzLdxClxfU6A0awVgYHKciKYBwS4vmQeh608UoYRJC
wbgU0X9cXLuh55PwXy83+b6phNgxFB6nBxCvcUVZ83bAMhKTyQtbQbdSw5wOzj9s
/x0=
-----END CERTIFICATE-----`
)

func TestIntegration_Initialize(t *testing.T) {
	c := NewAWSLoader("dev", "test")

	err := c.Initialize()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUnit_MustGetString(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"foo": "bar"},
	}
	ret := c.MustGetString("foo")
	if ret != "bar" {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetBool(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"bool": "true", "bool2": "false"},
	}
	ret := c.MustGetBool("bool")
	if !ret {
		t.Fatalf("Unexpected value returned")
	}
	ret = c.MustGetBool("bool2")
	if ret {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetInt(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"int": "1234567890"},
	}
	ret := c.MustGetInt("int")
	if ret != 1234567890 {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetDuration(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"duration": "1m"},
	}
	ret := c.MustGetDuration("duration")
	if ret != time.Minute {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetObject(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"object": `{"foo": "bar"}`},
	}

	type testObject struct {
		Foo string `json:"foo"`
	}

	sample := new(testObject)
	c.MustGetObject("object", sample)
	if sample.Foo != "bar" {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetPublicKey(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"public_key": testPublicKey},
	}
	ret := c.MustGetPublicKey("public_key")
	if ret == nil {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetPrivateKey(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{"private_key": testPrivateKey},
	}
	ret := c.MustGetPrivateKey("private_key")
	if ret == nil {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetCertificate(t *testing.T) {
	c := &awsLoader{
		config: map[string]string{
			"certificate": testCertificate,
			"private_key": testPrivateKey,
		},
	}
	ret := c.MustGetCertificate("certificate", "private_key")
	if ret == nil {
		t.Fatalf("Unexpected value returned")
	}
}
