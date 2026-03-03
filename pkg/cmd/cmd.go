// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
	"github.com/with-ours/platform-cli/internal/autocomplete"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "oursprivacy",
		Usage:   "CLI for the ours-privacy-platform API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("OURS_PRIVACY_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "destinations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&destinationsCreate,
					&destinationsRetrieve,
					&destinationsUpdate,
					&destinationsList,
					&destinationsDelete,
				},
			},
			{
				Name:     "sources",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sourcesCreate,
					&sourcesRetrieve,
					&sourcesUpdate,
					&sourcesList,
					&sourcesDelete,
				},
			},
			{
				Name:     "allowed-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&allowedEventsCreate,
					&allowedEventsRetrieve,
					&allowedEventsList,
					&allowedEventsDelete,
				},
			},
			{
				Name:     "consent-settings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&consentSettingsCreate,
					&consentSettingsRetrieve,
					&consentSettingsUpdate,
					&consentSettingsList,
					&consentSettingsDelete,
				},
			},
			{
				Name:     "global-dispatch-centers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&globalDispatchCentersCreate,
					&globalDispatchCentersRetrieve,
					&globalDispatchCentersUpdate,
					&globalDispatchCentersList,
					&globalDispatchCentersDelete,
				},
			},
			{
				Name:     "replay-settings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&replaySettingsCreate,
					&replaySettingsRetrieve,
					&replaySettingsUpdate,
					&replaySettingsList,
					&replaySettingsDelete,
				},
			},
			{
				Name:     "versions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&versionsCreate,
					&versionsRetrieve,
					&versionsUpdate,
					&versionsList,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "oursprivacy @manpages [-o oursprivacy.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "oursprivacy.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "oursprivacy.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
