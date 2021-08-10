/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"github.com/zdz1715/cron-log/pkg"
)

type options struct {
	User  string `json:"user"`
	Shell string `valid:"in(sh|bash)" json:"shell"`
}

var opt options

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cron-log [flags] args",
	Short: "Format and record Linux Cron output ",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := govalidator.ValidateStruct(opt)
		fmt.Printf("%+v\n", opt)
		if err != nil {
			return err
		}
		pkg.Collect(opt.Shell, opt.User, "echo sss")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&opt.User, "user", "u", "root", "executing user")
	rootCmd.Flags().StringVarP(&opt.Shell, "shell", "s", "sh", "sh:'/bin/sh' | bash:'/bin/bash'")
}