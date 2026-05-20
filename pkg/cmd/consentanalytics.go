// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/with-ours/platform-cli/internal/apiquery"
	"github.com/with-ours/platform-cli/internal/requestflag"
	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/option"
)

var consentAnalyticsList = cli.Command{
	Name:    "list",
	Usage:   "Account-wide blocking stats from the Global Consent Center for the window: how\nmany dispatches were attempted, how many were blocked, and a per-category\nbreakdown of the blocks (with `percentageBlocked` = share of `totalDispatches`).\nNot scoped to a single consent settings record — this aggregates across every\ndestination in the account. The endpoint is identified by query params rather\nthan a path id because the report is account-scoped; this is a documented\nderived-read exception. Reuses the API-key scope `consentSettings:list` because\nthe report crosses every consent setting. Requires scope: consentSettings:list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
	},
	Action:          handleConsentAnalyticsList,
	HideHelpCommand: true,
}

func handleConsentAnalyticsList(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := oursprivacy.ConsentAnalyticsListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentAnalytics.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "consent-analytics list",
		Transform:      transform,
	})
}
