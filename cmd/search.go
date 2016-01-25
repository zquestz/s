package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zquestz/s/providers"
)

const (
	appName = "s"
	version = "0.1.7"
)

// Flag variables
var displayVersion bool
var verbose bool
var provider string
var listProviders bool

// Main command for Cobra.
var SearchCmd = &cobra.Command{
	Use:   "s <query>",
	Short: "Web search from the terminal",
	Long:  `Web search from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(args)
		if err != nil {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	prepareFlags()
}

func prepareFlags() {
	SearchCmd.PersistentFlags().BoolVarP(
		&displayVersion, "version", "", false, "display version")
	SearchCmd.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false, "display url when opening")
	SearchCmd.PersistentFlags().StringVarP(
		&provider, "provider", "p", "google", "set search provider")
	SearchCmd.PersistentFlags().BoolVarP(
		&listProviders, "list-providers", "l", false, "list supported providers")
}

// Where all the work happens.
func performCommand(args []string) error {
	if displayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if listProviders {
		fmt.Printf(providers.DisplayProviders())
		return nil
	}

	query := strings.Join(args, " ")

	if query != "" {
		providers.Search(provider, query, verbose)
		return nil
	} else {
		// We don't display this, as the help screen is more useful.
		return fmt.Errorf("[Error] query is required.")
	}
}
