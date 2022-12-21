package main

import (
	"fmt"
	"os"
	"time"

	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	ld "github.com/launchdarkly/go-server-sdk/v6"
)

// Set sdkKey to your LaunchDarkly SDK key.
const sdkKey = ""

// Set featureFlagKey to the feature flag key you want to evaluate.
const featureFlagKey = "limsRoute"

var ldConfig ld.Config

func showMessage(s string) { fmt.Printf("*** %s\n\n", s) }

func main() {
	if sdkKey == "" {
		showMessage("Please edit main.go to set sdkKey to your LaunchDarkly SDK key first")
		os.Exit(1)
	}
	//ldConfig.Offline = true
	ldClient, _ := ld.MakeCustomClient(sdkKey, ldConfig, 5*time.Second)
	if ldClient.Initialized() {
		showMessage("SDK successfully initialized!")
	} else {
		showMessage("SDK failed to initialize")
		os.Exit(1)
	}

	// Set up the user properties. This user should appear on your LaunchDarkly users dashboard
	// soon after you run the demo.

	// context-based request (new in v6)
	context := ldcontext.New("whatev")
	// user-based anonymous request - doesnt make a user in LD upon request. without Anonymous(true) it does do this, which isnt appropriate for requests from a backend service
	//user := ldcontext.New("whatev").Name("ensemble").Anonymous(true).Build()

	flagValue, err := ldClient.BoolVariation(featureFlagKey, context, false)
	if err != nil {
		showMessage("error: " + err.Error())
	}

	showMessage(fmt.Sprintf("Feature flag '%s' is %t for this user", featureFlagKey, flagValue))

	// Here we ensure that the SDK shuts down cleanly and has a chance to deliver analytics
	// events to LaunchDarkly before the program exits. If analytics events are not delivered,
	// the user properties and flag usage statistics will not appear on your dashboard. In a
	// normal long-running application, the SDK would continue running and events would be
	// delivered automatically in the background.
	ldClient.Close()
}
