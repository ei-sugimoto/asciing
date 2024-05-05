package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var AsciingCmd = &cobra.Command{
	Use:   "asciing",
	Short: "asciing is a tool for converting string to ASCII art.",
	Long: "asciing takes a string as its argument. It converts the string into ASCII art and outputs it.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asciing called")
		outPutString := ""
		for i := 0; i < len(args); i++ {
    		outPutString += args[i]
			outPutString += " "
		}
		myFigure := figure.NewFigure(outPutString, "", true)
		myFigure.Print()
	},
}


func init() {
	rootCmd.AddCommand(AsciingCmd)
}