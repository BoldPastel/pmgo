/*
Copyright (c) 2022 Bold Pastel https://boldpastel.com/

*/
package cmd

import (
	"os"
	//"pmgo/pmgo"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pmgo",
	Short: "Permissionset Utilities for permission definitions mastered on predefined CSVs",
	Long: `This is a set of very basic tools to help work with permissionset xml files.
The tool primarily generates permissionsets from a CSV defintion.

In addition, the app also supports merging two permissions and generating a single one and 
supports the modified spec of generating field permissions based on object permissions.

This is to help master permission definitions on a human readable, annotateable XLX/CSV format,
allowing to consistently generate permissions based on that master reliably and with confidence.

It is expected that the fields required as available on the CSV or generated from the repo using the following
script.

This is merely a tool and should be used only to help reduce the manual effort. This does not reduce or
replace the thinking effort of a consultant.

supplied as is where is ~~~~(3:>
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pmgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
