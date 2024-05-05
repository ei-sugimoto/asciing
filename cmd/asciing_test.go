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

func setArgs(arg string) {
	os.Args = append(originalArgs, "asciing", arg)
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
	setArgs("hello")
	defer resetArgs()
	myFigure := figure.NewFigure("hello", "", true)
	AsciingCmd := cmd.AsciingCmd
	err := cmd.AsciingCmd.Execute()
    if err != nil {
        log.Fatalf("Failed to execute AsciingCmd: %v", err)
    }
	got := PickStdout(t, func() { AsciingCmd.Execute() }) + "\n"
	want := "asciing called\n" + myFigure.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}