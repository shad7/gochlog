package core

import (
	"fmt"

	// log "github.com/Sirupsen/logrus"

	"github.com/shad7/gochlog/styles"
	"github.com/shad7/gochlog/types"

	// Import each style so that they initialize
	_ "github.com/shad7/gochlog/styles/standard"
)

// ListAvailableStyles provides a list of all registered styles
func ListAvailableStyles() []string {
	return styles.GetStylers()
}

// GenerateChangeLog creates/updates a change log using git commits and a specified style
func GenerateChangeLog(style string, gopts *types.GitOptions, popts *types.ParserOptions, fopts *types.FormatterOptions) error {
	styler, err := styles.GetStyler(style)
	if err != nil {
		return err
	}
	if styler == nil {
		return fmt.Errorf("unknown style provided '%s', select one of these: %s", style, ListAvailableStyles())
	}

	commits, err := fetchChanges(gopts)
	if err != nil {
		return err
	}
	fmt.Println(commits)
	return nil
}
