package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// NewCmdRoot Create Root
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adbkitting",
		Short: "Android Device Bridge kitting tools",
	}
	cobra.OnInitialize(initConfig)
	// cmd.PersistentFlags().StringVar(&cfgFile, "confg", "config.yaml", "config file (default is config.yaml)")

	cmd.AddCommand(NewCmdInstallAll())
	cmd.AddCommand(NewCmdList())
	return cmd
}

// Execute command
func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOut(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func initConfig() {
	// viper.SetConfigFile(cfgFile)
}
