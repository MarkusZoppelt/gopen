package cmd

import (
	"log"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"

	apps "github.com/MarkusZoppelt/gopen/internal/applications"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopen [file]",
	Short: "Open a file with a specific application",
	Long: `Open a file with a specific application.
After running the command, a fuzzy search will be started to select the application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		applications := apps.GetInstalledApplications()

		// start fuzzy search
		idx, err := fuzzyfinder.Find(
			applications,
			func(i int) string {
				return applications[i]
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		// open the selected application
		err = apps.OpenWithApplication(applications[idx], args, verbose)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// read all installed applications
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
