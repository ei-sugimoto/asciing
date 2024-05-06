package cmd_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/common-nighthawk/go-figure"
	"github.com/ei-sugimoto/asciing/cmd"
)

var originalArgs = os.Args

func setArgs(arg ...string) {
	os.Args = append(os.Args[:1], arg...)
}

func resetArgs() {
	os.Args = originalArgs
}

func PickStdout(t *testing.T, fnc func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("fail pipe: %v", err)
	}
	os.Stdout = w
	fnc()
	w.Close()
	var buffer bytes.Buffer
	if n, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v - number: %v", err, n)
	}
	s := buffer.String()
	return s[:len(s)-1]
}

func TestAsciingCmd(t *testing.T) {
	setArgs("create", "hello", "world")
	defer resetArgs()
	myFigure := figure.NewFigure("hello world", "", true)
	AsciingCmd := cmd.AsciingCmd

	var err error
	got := PickStdout(t, func() { err = AsciingCmd.Execute() }) + "\n"
	if err != nil {
		log.Fatalf("Failed to execute AsciingCmd: %v", err)
	}
	want := "asciing called\n" + myFigure.String()
	if got != want {
		t.Errorf("got: \n%q \n want:%q\n", got, want)
	}
}

func TestAsciingCmdWithFlagFont(t *testing.T) {
	setArgs("create", "hello", "world", "--font", "isometric1")
	defer resetArgs()
	myFigure := figure.NewFigure("hello world", "isometric1", true)
	AsciingCmd := cmd.AsciingCmd
	var err error

	got := PickStdout(t, func() { err = AsciingCmd.Execute() }) + "\n"
	if err != nil {
		log.Fatalf("Failed to execute AsciingCmd: %v", err)
	}

	want := "asciing called\n" + myFigure.String()

	if got != want {
		t.Errorf("got: %q\n want:%q\n", got, want)
	}
}

func TestFontVaridate(t *testing.T) {
	err := cmd.ValidateFont("notFoundFont")

	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
