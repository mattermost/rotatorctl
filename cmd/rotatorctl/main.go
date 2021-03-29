package main

import (
	"encoding/json"
	"os"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func printJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}

// newRootCmd creates the root command
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rotatorctl",
		Short: "Rotatorctl is a cli-tool to rotate K8s cluster nodes",
	}
	rootCmd.AddCommand(newRotateCmd())
	return rootCmd
}
