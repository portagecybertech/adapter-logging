package logging

import (
	"testing"
)

func TestInit(t *testing.T) {
	log := Init()
	log.Level()
	log.Name()
	if log == nil {
		t.Errorf("Init failed")
	}
}

func TestL(t *testing.T) {
	log := L()
	if log == nil {
		t.Errorf("L failed")
	}
}

func TestNamed(t *testing.T) {
	name := "Testname"
	log := Named(name)
	if log == nil {
		t.Errorf("Named failed")
	}
	if name != log.Name() {
		t.Errorf("Name does not match")
	}
}

func TestNew(t *testing.T) {
	log := New()
	if log == nil {
		t.Errorf("New failed")
	}
}
