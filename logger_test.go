package logging

import (
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	var LOG_ENV string
	var LOG_LEVEL string

	err := godotenv.Load()
	if err != nil {
		LOG_ENV = "prod"
		LOG_LEVEL = "error"
	} else {
		LOG_ENV = os.Getenv("LOG_ENV")
		LOG_LEVEL = os.Getenv("LOG_LEVEL")
	}

	log := Init()

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

	//second test
	log2 := L()

	if log2 != log {
		t.Errorf("Did not return singleton")
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
	name2 := "Testname2"
	log := Named(name)
	if log == nil {
		t.Errorf("Named failed")
	}
	if name != log.Name() {
		t.Errorf("Name does not match")
	}

	//new name, second test
	log2 := Named(name2)

	if log == log2 {
		t.Errorf("Failed to return new named log")
	}
}

func TestNew(t *testing.T) {
	var LOG_ENV string
	var LOG_LEVEL string

	err := godotenv.Overload()
	if err != nil {
		LOG_ENV = "prod"
		LOG_LEVEL = "error"
	} else {
		LOG_ENV = os.Getenv("LOG_ENV")
		LOG_LEVEL = os.Getenv("LOG_LEVEL")
	}
	log := New()

	if log == nil {
		t.Errorf("New failed")
	}

	if len(LOG_LEVEL) > 0 {
		if strings.ToUpper(LOG_LEVEL) != log.Level().CapitalString() {
			t.Errorf("New log level mismatch")
		}
	} else {
		if LOG_ENV == "dev" && log.Level().CapitalString() != "INFO" {
			t.Errorf("New log level default setting for dev is incorrect")
		}
		if LOG_ENV != "dev" && log.Level().CapitalString() != "ERROR" {
			t.Errorf("New log level default setting for non-dev environment is incorrect")
		}
	}

	//new env, second test
	env, env_err := godotenv.Unmarshal("LOG_ENV=prod\nLOG_LEVEL=panic\nLOG_FORMAT=console")
	if env_err == nil {
		godotenv.Write(env, ".env")
	}

	err2 := godotenv.Overload()
	if err2 != nil {
		LOG_ENV = "prod"
		LOG_LEVEL = "error"
	} else {
		LOG_ENV = os.Getenv("LOG_ENV")
		LOG_LEVEL = os.Getenv("LOG_LEVEL")
	}
	log2 := New()

	if log2 == nil {
		t.Errorf("New failed")
	}

	if log == log2 {
		t.Errorf("Failed to return new log")
	}
}
