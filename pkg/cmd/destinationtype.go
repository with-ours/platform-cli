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

var destinationTypesList = cli.Command{
	Name:            "list",
	Usage:           "Lists every destination type the platform supports, with its human-readable\nlabel, capability flags (oauth, listsAccounts, supportsRenamedEvents), and the\nsettings descriptor used to configure a destination of that type.\nAccount-agnostic — the response is the same for every API key. Requires scope:\ndestination:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDestinationTypesList,
	HideHelpCommand: true,
}

var destinationTypesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch the descriptor for a single destination type by its identifier (e.g.\n`Klaviyo`, `Facebook`, `GoogleDataManagerEventIngest`). Returns 404 if the\nidentifier is unknown. Requires scope: destination:list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     `Allowed values: "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS", "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Everflow", "Facebook", "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol", "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer", "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest", "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination", "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo", "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel", "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari", "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk", "ZoomInfo".`,
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDestinationTypesRetrieve,
	HideHelpCommand: true,
}

func handleDestinationTypesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DestinationTypes.List(ctx, options...)
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
		Title:          "destination-types list",
		Transform:      transform,
	})
}

func handleDestinationTypesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DestinationTypes.Get(ctx, oursprivacy.DestinationTypeGetParamsID(cmd.Value("id").(string)), options...)
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
		Title:          "destination-types retrieve",
		Transform:      transform,
	})
}
