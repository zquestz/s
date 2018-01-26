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
	appName         = "s"
	version         = "0.5.12"
	defaultPort     = 8080
	defaultProvider = "google"
)

// Stores configuration data.
var config Config

// SearchCmd is the main command for Cobra.
var SearchCmd = &cobra.Command{
	Use:   "s <query>",
	Short: "Web search from the terminal",
	Long:  `Web search from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			bail(err)
		}
	},
}

func init() {
	err := config.Load()
	if err != nil {
		bail(fmt.Errorf("Failed to load configuration: %s", err))
	}

	loadCustomProviders()

	prepareFlags()
}

func bail(err error) {
	fmt.Fprintf(os.Stderr, "[Error] %s\n", err)
	os.Exit(1)
}

func loadCustomProviders() {
	for _, p := range config.CustomProviders {
		err := p.Valid()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Warn] Invalid provider %q: %s\n", p.Name, err)
			continue
		}

		providers.AddProvider(p.Name, p)
	}
}

func prepareFlags() {
	if config.Provider == "" {
		config.Provider = defaultProvider
	}

	if config.Port == 0 {
		config.Port = defaultPort
	}

	SearchCmd.PersistentFlags().BoolVarP(
		&config.DisplayVersion, "version", "", false, "display version")
	SearchCmd.PersistentFlags().BoolVarP(
		&config.Verbose, "verbose", "v", config.Verbose, "verbose mode")
	SearchCmd.PersistentFlags().BoolVarP(
		&config.Output, "output", "o", config.Output, "output only mode")
	SearchCmd.PersistentFlags().StringVarP(
		&config.Provider, "provider", "p", config.Provider, "search provider")
	SearchCmd.PersistentFlags().StringVarP(
		&config.Tag, "tag", "t", config.Tag, "search tag")
	SearchCmd.PersistentFlags().BoolVarP(
		&config.ListProviders, "list-providers", "l", false, "list supported providers")
	SearchCmd.PersistentFlags().BoolVarP(
		&config.ListTags, "list-tags", "", false, "list available tags")
	SearchCmd.PersistentFlags().StringVarP(
		&config.Binary, "binary", "b", config.Binary, "binary to launch search URI")
	SearchCmd.PersistentFlags().BoolVarP(
		&config.ServerMode, "server", "s", false, "launch web server")
	SearchCmd.PersistentFlags().IntVarP(
		&config.Port, "port", "", config.Port, "server port")
	SearchCmd.PersistentFlags().StringVarP(
		&config.Cert, "cert", "c", config.Cert, "path to cert.pem for TLS")
	SearchCmd.PersistentFlags().StringVarP(
		&config.Key, "key", "k", config.Key, "path to key.pem for TLS")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if config.DisplayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	providers.SetBlacklist(config.Blacklist)
	providers.SetWhitelist(config.Whitelist)

	if config.ListProviders {
		fmt.Printf(providers.DisplayProviders(config.Verbose))
		return nil
	}

	if config.ListTags {
		fmt.Printf(providers.DisplayTags(config.Verbose))
		return nil
	}

	if config.ServerMode {
		err := server.Run(config.Port, config.Cert, config.Key, config.Provider, config.Verbose)
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
		err := providers.Search(
			config.Binary,
			config.Provider,
			config.Tag,
			query,
			cmd.Flags().Changed("provider"),
			config.Output,
			config.Verbose,
		)
		if err != nil {
			return err
		}
	} else {
		// Don't return an error, help screen is more appropriate.
		help := cmd.HelpFunc()
		help(cmd, args)
	}

	return nil
}
