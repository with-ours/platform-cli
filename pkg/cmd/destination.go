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

var destinationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new destination. Requires scope: destination:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS", "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook", "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol", "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer", "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest", "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination", "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo", "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel", "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari", "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk", "ZoomInfo".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleDestinationsCreate,
	HideHelpCommand: true,
}

var destinationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single destination by ID. Requires scope: destination:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDestinationsRetrieve,
	HideHelpCommand: true,
}

var destinationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a destination. Requires scope: destination:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "Disabled", "Enabled".`,
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[*string]{
			Name:     "facebook-conversion-api-key",
			BodyPath: "facebookConversionAPIKey",
		},
		&requestflag.Flag[*string]{
			Name:     "facebook-pixel-id",
			BodyPath: "facebookPixelId",
		},
		&requestflag.Flag[*string]{
			Name:     "g4-analytics-api-key",
			BodyPath: "g4AnalyticsApiKey",
		},
		&requestflag.Flag[*string]{
			Name:     "g4-analytics-measurement-id",
			BodyPath: "g4AnalyticsMeasurementId",
		},
		&requestflag.Flag[*bool]{
			Name:     "g4-analytics-track-on-page",
			BodyPath: "g4AnalyticsTrackOnPage",
		},
		&requestflag.Flag[*string]{
			Name:     "hashing-salt",
			BodyPath: "hashingSalt",
		},
		&requestflag.Flag[*string]{
			Name:     "http-destination-url",
			BodyPath: "httpDestinationUrl",
		},
		&requestflag.Flag[any]{
			Name:     "limited-to-source-id",
			BodyPath: "limitedToSourceIds",
		},
		&requestflag.Flag[*string]{
			Name:     "manager-google-customer-id",
			BodyPath: "managerGoogleCustomerId",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "project-api-key",
			BodyPath: "projectAPIKey",
		},
		&requestflag.Flag[*string]{
			Name:     "project-token",
			BodyPath: "projectToken",
		},
		&requestflag.Flag[*string]{
			Name:     "selected-account-id",
			BodyPath: "selectedAccountId",
		},
		&requestflag.Flag[any]{
			Name:     "settings",
			BodyPath: "settings",
		},
	},
	Action:          handleDestinationsUpdate,
	HideHelpCommand: true,
}

var destinationsList = cli.Command{
	Name:            "list",
	Usage:           "List all destinations. Requires scope: destination:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDestinationsList,
	HideHelpCommand: true,
}

var destinationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a destination. Requires scope: destination:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDestinationsDelete,
	HideHelpCommand: true,
}

func handleDestinationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomwithoursplatformsdkgo.DestinationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Destinations.New(ctx, params, options...)
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
		Title:          "destinations create",
		Transform:      transform,
	})
}

func handleDestinationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Destinations.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "destinations retrieve",
		Transform:      transform,
	})
}

func handleDestinationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomwithoursplatformsdkgo.DestinationUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Destinations.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
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
		Title:          "destinations update",
		Transform:      transform,
	})
}

func handleDestinationsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Destinations.List(ctx, options...)
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
		Title:          "destinations list",
		Transform:      transform,
	})
}

func handleDestinationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Destinations.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "destinations delete",
		Transform:      transform,
	})
}
