package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"

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

		myFigure := figure.NewFigure(outPutString, font, true)
		myFigure.Print()
	},
}

func init() {
	AsciingCmd.PersistentFlags().StringVar(&font, "font", "", "font name")
	err := AsciingCmd.MarkFlagRequired("font")
	if err != nil {
		if font != "" {
			AsciingCmd.Run = func(cmd *cobra.Command, args []string) {
				err := VaridateFont(font) // ここでフォントの存在を確認します
				if err != nil {
					os.Exit(1)
				}
				figure.NewFigure(args[0], font, true).Print()
			}
		}
	}

	rootCmd.AddCommand(AsciingCmd)
}

func VaridateFont(font string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(font) {
		return fmt.Errorf("Invalid font name")
	}
	fontpath := "fonts/" + font + ".flf"

	if _, err := os.Stat(fontpath); os.IsNotExist(err) {
		log.Println(err)                    // 詳細なエラーメッセージをログに記録します
		return fmt.Errorf("Font not found") // ユーザーには一般的なエラーメッセージを表示します
	}
	return nil
}
