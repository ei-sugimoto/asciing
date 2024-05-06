package cmd

import (
	"fmt"
	"os"


	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var font string

var AsciingCmd = &cobra.Command{
	Use:   "create",
	Short: "asciing is a tool for converting string to ASCII art.",
	Long:  "asciing takes a string as its argument. It converts the string into ASCII art and outputs it.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asciing called")
		outPutString := ""
		for i := 0; i < len(args); i++ {
			outPutString += args[i]
			outPutString += " "
		}
		var err error

		if font, err = cmd.Flags().GetString("font"); err != nil {
			fmt.Println("Error: ", err)
		}


		if font != "" {
			err := ValidateFont(font)
			if err != nil {
				fmt.Printf("Error: Font %s not found\n", font)
				os.Exit(1)
			}
		}
		myFigure := figure.NewFigure(outPutString, font, true)
		myFigure.Print()
	},
}

func init() {
	AsciingCmd.PersistentFlags().StringVar(&font, "font", "", "font name")
	rootCmd.AddCommand(AsciingCmd)
}

func ValidateFont(font string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Font %s not found", font)
		}
	}()
	figure.NewFigure("test", font, true)
	err = nil
	return err
}
