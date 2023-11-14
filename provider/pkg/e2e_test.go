// Copyright 2023, Pulumi Corporation.  All rights reserved.

// Disable running if we've specifically selected unit tests to run as this is an integration test.
// This is easier than having to remember to explicitly tag every unit test with `//go:build unit || all`.
//go:build !unit

package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest"
)

func TestApi(t *testing.T) {
	e2eTest("api").Run(t)
}

func e2eTest(dir string, opts ...providertest.Option) *providertest.ProviderTest {
	cwd, err := os.Getwd()
	if err != nil {
		// If we can't get the current working directory, we'll fail hard.
		panic(err)
	}

	opts = append(opts,
		providertest.WithProviderName("aws-apigateway"),
		providertest.WithBaselineVersion("1.0.1"),
		providertest.WithSkippedUpgradeTestMode(
			providertest.UpgradeTestMode_Quick,
			"Quick mode is only supported for providers written in Go at the moment"),
	)

	return providertest.NewProviderTest(
		filepath.Join(cwd, "test-programs", dir),
		opts...,
	)
}