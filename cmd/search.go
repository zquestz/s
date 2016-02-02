package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zquestz/s/providers"
	"github.com/zquestz/s/server"
)

const (
	appName = "s"
	version = "0.2.1"
)

// Flag variables
var displayVersion bool
var verbose bool
var provider string
var listProviders bool
var binary string
var serverMode bool
var port int
var certpem string
var keypem string

// SearchCmd is the main command for Cobra.
var SearchCmd = &cobra.Command{
	Use:   "s <query>",
	Short: "Web search from the terminal",
	Long:  `Web search from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Error] %s\n", err)
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
	SearchCmd.PersistentFlags().StringVarP(
		&binary, "binary", "b", "", "binary to launch search uri")
	SearchCmd.PersistentFlags().BoolVarP(
		&serverMode, "server", "s", false, "launch web server")
	SearchCmd.PersistentFlags().IntVarP(
		&port, "port", "", 8080, "server port")
	SearchCmd.PersistentFlags().StringVarP(
		&certpem, "cert", "c", "", "location of cert.pem for TLS")
	SearchCmd.PersistentFlags().StringVarP(
		&keypem, "key", "k", "", "location of key.pem for TLS")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if displayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if listProviders {
		fmt.Printf(providers.DisplayProviders())
		return nil
	}

	if serverMode {
		err := server.Run(port, certpem, keypem, provider)
		if err != nil {
			return err
		}

		return nil
	}

	query := strings.Join(args, " ")

	st, err := os.Stdin.Stat()
	if err != nil {
		// os.Stdin.Stat() can be unavailable on Windows.
		if runtime.GOOS != "windows" {
			return fmt.Errorf("Failed to stat Stdin: %s", err)
		}
	} else {
		if st.Mode()&os.ModeNamedPipe != 0 {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("Failed to read from Stdin: %s", err)
			}

			query = strings.TrimSpace(fmt.Sprintf("%s %s", query, bytes))
		}
	}

	if query != "" {
		err := providers.Search(binary, provider, query, verbose)
		if err != nil {
			return err
		}
	} else {
		// Don't return an error, help screen is more appropriate.
		cmd.Help()
	}

	return nil
}
