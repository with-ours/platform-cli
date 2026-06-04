// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestExperimentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
			"--search", "pricing hero",
			"--status", "completed",
			"--type", "ab",
		)
	})
}

func TestExperimentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "create",
			"--experiment-settings-id", "settings_01HZX9BB73EY2Q37VGK5A0VW7A",
			"--name", "Homepage Hero Headline Test",
			"--control-weight", "34",
			"--description", "description",
			"--include-query-string=true",
			"--key", "homepage-hero-headline-test",
			"--metrics", "{primary: {eventName: demo_requested, funnelId: funnelId}, secondary: [{eventName: demo_requested, funnelId: funnelId}]}",
			"--targeting-rules", "{urlPatterns: [/pricing*, /enterprise], audienceId: audienceId, queryParams: [{key: utm_campaign, operator: contains, value: spring-launch}], visitorProperties: {}, visitorStatus: visitorStatus}",
			"--traffic-allocation", "100",
			"--type", "ab",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "create",
			"--experiment-settings-id", "settings_01HZX9BB73EY2Q37VGK5A0VW7A",
			"--name", "Homepage Hero Headline Test",
			"--control-weight", "34",
			"--description", "description",
			"--include-query-string=true",
			"--key", "homepage-hero-headline-test",
			"--metrics.primary", "{eventName: demo_requested, funnelId: funnelId}",
			"--metrics.secondary", "[{eventName: demo_requested, funnelId: funnelId}]",
			"--targeting-rules.url-patterns", "[/pricing*, /enterprise]",
			"--targeting-rules.audience-id", "audienceId",
			"--targeting-rules.query-params", "[{key: utm_campaign, operator: contains, value: spring-launch}]",
			"--targeting-rules.visitor-properties", "{}",
			"--targeting-rules.visitor-status", "visitorStatus",
			"--traffic-allocation", "100",
			"--type", "ab",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"experimentSettingsId: settings_01HZX9BB73EY2Q37VGK5A0VW7A\n" +
			"name: Homepage Hero Headline Test\n" +
			"controlWeight: 34\n" +
			"description: description\n" +
			"includeQueryString: true\n" +
			"key: homepage-hero-headline-test\n" +
			"metrics:\n" +
			"  primary:\n" +
			"    eventName: demo_requested\n" +
			"    funnelId: funnelId\n" +
			"  secondary:\n" +
			"    - eventName: demo_requested\n" +
			"      funnelId: funnelId\n" +
			"targetingRules:\n" +
			"  urlPatterns:\n" +
			"    - /pricing*\n" +
			"    - /enterprise\n" +
			"  audienceId: audienceId\n" +
			"  queryParams:\n" +
			"    - key: utm_campaign\n" +
			"      operator: contains\n" +
			"      value: spring-launch\n" +
			"  visitorProperties: {}\n" +
			"  visitorStatus: visitorStatus\n" +
			"trafficAllocation: 100\n" +
			"type: ab\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "create",
		)
	})
}

func TestExperimentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "retrieve",
			"--id", "id",
		)
	})
}

func TestExperimentsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
			"--description", "description",
			"--include-query-string=true",
			"--key", "key",
			"--metrics", "{primary: {eventName: demo_requested, funnelId: funnelId}, secondary: [{eventName: demo_requested, funnelId: funnelId}]}",
			"--name", "name",
			"--targeting-rules", "{urlPatterns: [/pricing*, /enterprise], audienceId: audienceId, queryParams: [{key: utm_campaign, operator: contains, value: spring-launch}], visitorProperties: {}, visitorStatus: visitorStatus}",
			"--traffic-allocation", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
			"--description", "description",
			"--include-query-string=true",
			"--key", "key",
			"--metrics.primary", "{eventName: demo_requested, funnelId: funnelId}",
			"--metrics.secondary", "[{eventName: demo_requested, funnelId: funnelId}]",
			"--name", "name",
			"--targeting-rules.url-patterns", "[/pricing*, /enterprise]",
			"--targeting-rules.audience-id", "audienceId",
			"--targeting-rules.query-params", "[{key: utm_campaign, operator: contains, value: spring-launch}]",
			"--targeting-rules.visitor-properties", "{}",
			"--targeting-rules.visitor-status", "visitorStatus",
			"--traffic-allocation", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"includeQueryString: true\n" +
			"key: key\n" +
			"metrics:\n" +
			"  primary:\n" +
			"    eventName: demo_requested\n" +
			"    funnelId: funnelId\n" +
			"  secondary:\n" +
			"    - eventName: demo_requested\n" +
			"      funnelId: funnelId\n" +
			"name: name\n" +
			"targetingRules:\n" +
			"  urlPatterns:\n" +
			"    - /pricing*\n" +
			"    - /enterprise\n" +
			"  audienceId: audienceId\n" +
			"  queryParams:\n" +
			"    - key: utm_campaign\n" +
			"      operator: contains\n" +
			"      value: spring-launch\n" +
			"  visitorProperties: {}\n" +
			"  visitorStatus: visitorStatus\n" +
			"trafficAllocation: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
		)
	})
}

func TestExperimentsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "delete",
			"--id", "id",
		)
	})
}

func TestExperimentsStart(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "start",
			"--id", "id",
			"--publish-after-start=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("publishAfterStart: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "start",
			"--id", "id",
		)
	})
}

func TestExperimentsStop(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "stop",
			"--id", "id",
			"--winner-variant-id", "var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("winnerVariantId: var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "stop",
			"--id", "id",
		)
	})
}

func TestExperimentsPause(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "pause",
			"--id", "id",
			"--publish-after-pause=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("publishAfterPause: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "pause",
			"--id", "id",
		)
	})
}

func TestExperimentsResume(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "resume",
			"--id", "id",
			"--publish-after-resume=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("publishAfterResume: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "resume",
			"--id", "id",
		)
	})
}

func TestExperimentsResults(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "results",
			"--id", "id",
			"--event-name", "demo_requested",
		)
	})
}

func TestExperimentsResultsTimeSeries(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "results-time-series",
			"--id", "id",
			"--end-date", "2026-04-30",
			"--event-name", "demo_requested",
			"--start-date", "2026-04-01",
		)
	})
}

func TestExperimentsSessionReplays(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "session-replays",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "25",
			"--variant-id", "var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8",
		)
	})
}
