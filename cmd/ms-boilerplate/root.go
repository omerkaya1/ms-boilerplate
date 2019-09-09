package main

import (
	"github.com/omerkaya1/ms-boilerplate/internal"
	"github.com/spf13/cobra"
	"log"
)

var cfgPath string

var rootCmd = &cobra.Command{
	Use:   "ms-boilerplate",
	Short: "ms-boilerplate is a simple web server that a",
	Run:   rootCommand,
}

// Execute is a method that runs the root command of the programme
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func rootCommand(cmd *cobra.Command, args []string) {
	s, err := internal.NewServer(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	s.Run()
}
