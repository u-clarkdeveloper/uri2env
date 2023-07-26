package uri2env

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/u-clarkdeveloper/uri2env/pkg/uri2env"
)

var version string = "0.0.1"
var prefix string = ""
var verbose bool
var rootCmd = &cobra.Command{
	Use:     "uri2env [OPTIONS] [COMMANDS] [URI]",
	Version: version,
	Args:    cobra.ExactArgs(1),
	Short:   "uri2env is a tool to convert uri to env and back",
	Long: `uri2env is a tool to convert uri to env and back.

For example:

> uri2env -v -p "MY_PREFIX" "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#some_fragment"

Input: [http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment]
Prefix: MY_PREFIX
MY_PREFIX_SCHEME=http
MY_PREFIX_HOST=localhost
MY_PREFIX_PORT=8080
MY_PREFIX_USER=userA
MY_PREFIX_PASSWORD=passwordA
MY_PREFIX_PATH=/path/stuff/here
MY_PREFIX_QUERY=yes
MY_PREFIX_TESTING=1
MY_PREFIX_FRAGMENT=my_fragment

	`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("Input: %s\n", args)
			fmt.Printf("Prefix: %s\n", prefix)
		}
		output, err := uri2env.ConvertUri(args[0], prefix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
			os.Exit(1)
		}
		fmt.Println(output)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Prefix for environment variables")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output: include input and prefix in output")
}
