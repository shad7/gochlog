package core

import (
	"fmt"

	// log "github.com/Sirupsen/logrus"

	"github.com/davecgh/go-spew/spew"
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
	fmt.Println(spew.Sdump(commits))

	mpopts, err := types.ApplyDefaultParserOptions(popts, styler.GetParserOptions())
	if err != nil {
		return err
	}
	fmt.Println(spew.Sdump(mpopts))

	data, err := parse(mpopts, commits)
	if err != nil {
		return err
	}

	mfopts, err := types.ApplyDefaultFormatterOptions(fopts, styler.GetFormatterOptions())
	if err != nil {
		return err
	}
	fmt.Println(spew.Sdump(mfopts))

	data, err = format(mfopts, data)
	if err != nil {
		return err
	}

	err = write(data)
	if err != nil {
		return err
	}

	return nil
}
