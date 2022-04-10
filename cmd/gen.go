/*
Copyright (c) 2022 Bold Pastel https://boldpastel.com/

*/
package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"pmgo/pm"
	//"fmt"
)

// rootCmd represents the base command when called without any subcommands
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate permissions from the specified CSV file",
	Long: `Sample CSV saved here: https://docs.google.com/spreadsheets/d/1rr5cAQlfA2-yMgET_D9CchV7SEg8LR5PJs0UpARS01U/edit?usp=sharing

	This script will assign the permissions specified. Any blank field permissions would be ignored and not added to the permissionset.
	
	Script to be run the object folder of the local repo for reference:
	grep "<type>" */fields/* | awk -F '/' '{printf("%s.%s\n",$1, $3);}' | awk -F '.field' '{split($2,a,">");split(a[2],b,"<");printf("%s\n", $1);}'>>../../../../fields.csv

	All metadata specified here is supported: https://developer.salesforce.com/docs/atlas.en-us.196.0.api_meta.meta/api_meta/meta_permissionset.htm
	
	Also note the script is not smart enough to handle required fields and would have to be corrected post generation if required`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("requires a file")
		}
		if strings.Split(args[0], ".")[1] != "csv" {
			return errors.New("File needs to be a CSV")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		for _, x := range pmgo.OpenCSV(args[0]) {
			pmgo.WriteXML(x)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pmgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
