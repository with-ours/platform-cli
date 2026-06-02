// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestLocationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "list",
		)
	})
}

func TestLocationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "create",
			"--additional-address", "[{latitude: 0, longitude: 0, city: city, country: country, line1: line1, line2: line2, name: name, phoneNumber: phoneNumber, state: state, websiteLinkText: websiteLinkText, websiteUrl: websiteUrl, zip: zip}]",
			"--center", "{}",
			"--city", "city",
			"--country", "country",
			"--custom-domain", "customDomain",
			"--latitude", "0",
			"--line1", "line1",
			"--line2", "line2",
			"--longitude", "0",
			"--map-name", "mapName",
			"--name", "name",
			"--phone-number", "phoneNumber",
			"--state", "state",
			"--website-link-text", "websiteLinkText",
			"--website-url", "websiteUrl",
			"--zip", "zip",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(locationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "create",
			"--additional-address.latitude", "0",
			"--additional-address.longitude", "0",
			"--additional-address.city", "city",
			"--additional-address.country", "country",
			"--additional-address.line1", "line1",
			"--additional-address.line2", "line2",
			"--additional-address.name", "name",
			"--additional-address.phone-number", "phoneNumber",
			"--additional-address.state", "state",
			"--additional-address.website-link-text", "websiteLinkText",
			"--additional-address.website-url", "websiteUrl",
			"--additional-address.zip", "zip",
			"--center", "{}",
			"--city", "city",
			"--country", "country",
			"--custom-domain", "customDomain",
			"--latitude", "0",
			"--line1", "line1",
			"--line2", "line2",
			"--longitude", "0",
			"--map-name", "mapName",
			"--name", "name",
			"--phone-number", "phoneNumber",
			"--state", "state",
			"--website-link-text", "websiteLinkText",
			"--website-url", "websiteUrl",
			"--zip", "zip",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"additionalAddresses:\n" +
			"  - latitude: 0\n" +
			"    longitude: 0\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    name: name\n" +
			"    phoneNumber: phoneNumber\n" +
			"    state: state\n" +
			"    websiteLinkText: websiteLinkText\n" +
			"    websiteUrl: websiteUrl\n" +
			"    zip: zip\n" +
			"center: {}\n" +
			"city: city\n" +
			"country: country\n" +
			"customDomain: customDomain\n" +
			"latitude: 0\n" +
			"line1: line1\n" +
			"line2: line2\n" +
			"longitude: 0\n" +
			"mapName: mapName\n" +
			"name: name\n" +
			"phoneNumber: phoneNumber\n" +
			"state: state\n" +
			"websiteLinkText: websiteLinkText\n" +
			"websiteUrl: websiteUrl\n" +
			"zip: zip\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"locations", "create",
		)
	})
}

func TestLocationsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "update",
			"--id", "id",
			"--additional-address", "[{latitude: 0, longitude: 0, city: city, country: country, line1: line1, line2: line2, name: name, phoneNumber: phoneNumber, state: state, websiteLinkText: websiteLinkText, websiteUrl: websiteUrl, zip: zip}]",
			"--center", "{}",
			"--city", "city",
			"--country", "country",
			"--custom-domain", "customDomain",
			"--latitude", "0",
			"--line1", "line1",
			"--line2", "line2",
			"--longitude", "0",
			"--map-name", "mapName",
			"--name", "name",
			"--phone-number", "phoneNumber",
			"--state", "state",
			"--website-link-text", "websiteLinkText",
			"--website-url", "websiteUrl",
			"--zip", "zip",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(locationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "update",
			"--id", "id",
			"--additional-address.latitude", "0",
			"--additional-address.longitude", "0",
			"--additional-address.city", "city",
			"--additional-address.country", "country",
			"--additional-address.line1", "line1",
			"--additional-address.line2", "line2",
			"--additional-address.name", "name",
			"--additional-address.phone-number", "phoneNumber",
			"--additional-address.state", "state",
			"--additional-address.website-link-text", "websiteLinkText",
			"--additional-address.website-url", "websiteUrl",
			"--additional-address.zip", "zip",
			"--center", "{}",
			"--city", "city",
			"--country", "country",
			"--custom-domain", "customDomain",
			"--latitude", "0",
			"--line1", "line1",
			"--line2", "line2",
			"--longitude", "0",
			"--map-name", "mapName",
			"--name", "name",
			"--phone-number", "phoneNumber",
			"--state", "state",
			"--website-link-text", "websiteLinkText",
			"--website-url", "websiteUrl",
			"--zip", "zip",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"additionalAddresses:\n" +
			"  - latitude: 0\n" +
			"    longitude: 0\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    line1: line1\n" +
			"    line2: line2\n" +
			"    name: name\n" +
			"    phoneNumber: phoneNumber\n" +
			"    state: state\n" +
			"    websiteLinkText: websiteLinkText\n" +
			"    websiteUrl: websiteUrl\n" +
			"    zip: zip\n" +
			"center: {}\n" +
			"city: city\n" +
			"country: country\n" +
			"customDomain: customDomain\n" +
			"latitude: 0\n" +
			"line1: line1\n" +
			"line2: line2\n" +
			"longitude: 0\n" +
			"mapName: mapName\n" +
			"name: name\n" +
			"phoneNumber: phoneNumber\n" +
			"state: state\n" +
			"websiteLinkText: websiteLinkText\n" +
			"websiteUrl: websiteUrl\n" +
			"zip: zip\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"locations", "update",
			"--id", "id",
		)
	})
}

func TestLocationsEmbedCode(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"locations", "embed-code",
			"--id", "id",
			"--color", "#007EA8",
			"--color-scheme", "light",
			"--include-address-box=true",
			"--include-controls", "yes",
			"--include-seo-schema=true",
			"--map-style", "Standard",
			"--theme", "default",
			"--zoom", "1",
		)
	})
}
