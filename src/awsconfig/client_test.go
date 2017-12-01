package awsconfig

import "testing"
import "time"

func TestUnit_Initialize(t *testing.T) {
	c := NewAWSLoader("dev", "test")

	err := c.Initialize()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUnit_MustGetString(t *testing.T) {
	c := NewAWSLoader("dev", "test")
	err := c.Initialize()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	ret := c.MustGetString("foo")
	if ret != "bar" {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetBool(t *testing.T) {
	c := NewAWSLoader("dev", "test")
	err := c.Initialize()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
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
	c := NewAWSLoader("dev", "test")
	err := c.Initialize()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	ret := c.MustGetInt("int")
	if ret != 1234567890 {
		t.Fatalf("Unexpected value returned")
	}
}

func TestUnit_MustGetDuration(t *testing.T) {
	c := NewAWSLoader("dev", "test")
	err := c.Initialize()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	ret := c.MustGetDuration("duration")
	if ret != time.Minute {
		t.Fatalf("Unexpected value returned")
	}
}
