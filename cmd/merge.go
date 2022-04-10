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
var merge = &cobra.Command{
	Use:   "merge",
	Short: "merge two permissionsets",
	Long:  `Non intelligent merge, where the operant permissions are applied, regardless of whether the target had more liberal permissions`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return errors.New("Require Two PermissionSet files")
		}
		if strings.Split(args[0], ".")[2] != "xml" && strings.Split(args[1], ".")[2] != "xml" {
			return errors.New("File needs to be a XML permissionset")
		}
		if strings.Split(args[0], ".")[1] != "permissionset-meta" && strings.Split(args[1], ".")[1] != "permissionset-meta" {
			return errors.New("Not a genuine permissionset file")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		pmgo.WriteXML(pmgo.MergeXML(pmgo.XmlUnpack(args[0]), pmgo.XmlUnpack(args[1])))
	},
}

func init() {
	rootCmd.AddCommand(merge)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pmgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
