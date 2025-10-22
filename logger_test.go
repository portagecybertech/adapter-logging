package logging

import (
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	log := Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	LOG_ENV := os.Getenv("LOG_ENV")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	if log == nil {
		t.Errorf("Init failed")
	}

	if len(LOG_LEVEL) > 0 {
		if strings.ToUpper(LOG_LEVEL) != log.Level().CapitalString() {
			t.Errorf("Log level mismatch")
		}
	} else {
		if LOG_ENV == "dev" && log.Level().CapitalString() != "INFO" {
			t.Errorf("Log level default setting for dev is incorrect")
		}
		if LOG_ENV != "dev" && log.Level().CapitalString() != "ERROR" {
			t.Errorf("Log level default setting for non-dev environment is incorrect")
		}
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	LOG_ENV := os.Getenv("LOG_ENV")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	if log == nil {
		t.Errorf("New failed")
	}

	if len(LOG_LEVEL) > 0 {
		if strings.ToUpper(LOG_LEVEL) != log.Level().CapitalString() {
			t.Errorf("Log level mismatch")
		}
	} else {
		if LOG_ENV == "dev" && log.Level().CapitalString() != "INFO" {
			t.Errorf("Log level default setting for dev is incorrect")
		}
		if LOG_ENV != "dev" && log.Level().CapitalString() != "ERROR" {
			t.Errorf("Log level default setting for non-dev environment is incorrect")
		}
	}
}
