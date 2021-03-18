package main

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rotatorctl",
	Short: "Rotatorctl is a cli-tool to rotate K8s cluster nodes",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("dry-run", false, "When set to true, only print the API request without sending it.")
}

func printJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}
