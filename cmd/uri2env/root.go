package uri2env

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uri2env",
	Short: "uri2env is a tool to convert uri to env and back",
	Long: `uri2env is a tool to convert uri to env and back.
	For example:
	> uri2env -p "MY_PREFIX" "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1"
	MY_PREFIX_USER=userA
	MY_PREFIX_PASSWORD=passwordA
	MY_PREFIX_HOST=localhost
	MY_PREFIX_PORT=8080
	MY_PREFIX_PATH=/path/stuff/here
	MY_PREFIX_QUERY=yes
	MY_PREFIX_TESTING=1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uri2env")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
