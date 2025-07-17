/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "h2go",
	Short: "HTMX golang component library",
	Long: `h2go is a CLI tool for scaffolding Go + HTMX UI components, inspired by shadcn/ui.  
Quickly add prebuilt, TailwindCSS-styled components to your Go applications.

Usage:
  h2go [command]

Available Commands:
  init        Initialize your project with TailwindCSS config and component directories
  add         Add a UI component to your project (button, modal, form, etc.)
  list        List available UI components you can add
  update      Update your components to the latest version
  help        Help about any command

Flags:
  -h, --help      help for h2go
  -v, --version   version for h2go`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.h2go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
