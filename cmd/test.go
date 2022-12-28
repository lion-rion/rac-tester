/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

/*グローバル変数*/
// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//get flag value
		sim, _ := cmd.Flags().GetBool("sim")
		printAa("test.txt")
		prompt := promptui.Select{
			Label:     "where test?",
			Items:     []string{"kick", "wheel", "spinner"},
			CursorPos: 0,
		}
		idx, _, err := prompt.Run() //入力を受け取る
		if err != nil {
			fmt.Println(err)
			return
		}
		switch idx {
		case 0:
			KickTest(sim)
		case 1:
			WheelTest(sim)
		case 2:
			SpinTest(sim)
		default:
		}

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().Bool("sim", false, "Use simulator")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
