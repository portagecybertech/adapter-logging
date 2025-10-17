package logging

import (
	"testing"
)

func TestInit(t *testing.T) {
	log := Init()
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

func TestNamedName(t *testing.T) {
	name := "Gladys"
	log := Named(name)
	if log == nil {
		t.Errorf("Named failed")
	}
}

func TestNew(t *testing.T) {
	log := New()
	if log == nil {
		t.Errorf("New failed")
	}
}
