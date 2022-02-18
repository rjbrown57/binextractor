package cmd

import (
	binextractor "github.com/rjbrown57/binextractor/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string
var imageName string
var sourcePath string
var destinationPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "binextractor",
	Short: "Extract binaries from Container images",
	Run: func(cmd *cobra.Command, args []string) {
		binextractor.Extract(imageName, sourcePath, destinationPath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.binextractor.yaml)")
	rootCmd.PersistentFlags().StringVarP(&imageName, "image", "i", "", "image to extract binary from")
	rootCmd.PersistentFlags().StringVarP(&sourcePath, "sourcePath", "s", "", "Path in binary to extract from")
	rootCmd.PersistentFlags().StringVarP(&destinationPath, "destinationPath", "d", "", "Destination to write extracted binary to")

	rootCmd.MarkPersistentFlagRequired("image")
	rootCmd.MarkPersistentFlagRequired("sourcePath")
	rootCmd.MarkPersistentFlagRequired("destinationPath")

}
