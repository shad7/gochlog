package commands

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	// The fork has multiple enhancements being used by Docker but not merged
	// into the base project as of yet. Once these are included then the switch
	// back can be made. This project also leverages tagged versions/releases
	// which make stability a strong option.
	// "github.com/spf13/cobra"
	"github.com/dnephin/cobra"
	"github.com/spf13/viper"

	"github.com/shad7/gochlog/core"
	"github.com/shad7/gochlog/types"
)

// RootCmd is the base command all other commands with extend from
var RootCmd = &cobra.Command{
	Use:   "gochlog",
	Short: "Generate Change Log",
	Long:  "Generate Change Log",
	RunE: func(cmd *cobra.Command, args []string) error {
		if verFlag {
			fmt.Print(core.GetVersionDisplay())
			return nil
		}

		err := core.GenerateChangeLog(style, gitOpts, parseOpts, formatOpts)
		if err != nil {
			return err
		}

		return nil
	},
}

var cfgFile string
var verFlag bool
var style string

var gitOpts *types.GitOptions
var parseOpts *types.ParserOptions
var formatOpts *types.FormatterOptions

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolVarP(&verFlag, "version", "v", false, "Print version information and quit")

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .gochlog.yaml)")
	RootCmd.PersistentFlags().StringVar(&style, "style", "standard", "select one of the available styles to use for formatting")

	gitOpts = &types.GitOptions{}
	RootCmd.PersistentFlags().StringVar(&gitOpts.TagPattern, "pattern", "[0-9.]{1,}", "Regular expresion for tags")

	parseOpts = &types.ParserOptions{}

	formatOpts = &types.FormatterOptions{}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gochlog") // name of the file w/o extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Debug("no config file found")
	} else {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}
}
