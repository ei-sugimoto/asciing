package asciing

import (
	"os"

	"github.com/ei-sugimoto/asciing/cmd"
)
func RunAsciing(args []string) error {
	origArgs := os.Args
	os.Args = append([]string{"asciing"}, args...)
	err := cmd.AsciingCmd.Execute()
	os.Args = origArgs
	return err
}
