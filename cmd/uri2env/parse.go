package uri2env

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/u-clarkdeveloper/uri2env/pkg/uri2env"
)

var parseCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p"},
	Short:   "Parse a URI",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := uri2env.Parse(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
			os.Exit(1)
		}
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
