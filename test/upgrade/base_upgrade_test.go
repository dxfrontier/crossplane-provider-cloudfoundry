//go:build upgrade

package upgrade

import (
	"context"
	"testing"

	"k8s.io/client-go/features"
)

func (c *CustomUpgradeTest) Build(t *testing.T) features.Feature {
	feature := features.New(c.name)
	feature = feature.Setup(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		t.Logf("Setting up resources from version %s", c.fromVersion)
		//TODO: install provider at fromVersion
		//TODO: Apply manifest at resource directories
		return ctx
	})
	// Pre-upgrade assessments
	for name, assestmentFunc := range c.preUpgradeAssessment {
		assestment := assestmentFunc //Capture for closure
		assestmentName := name 
		feature = feature.Assess(assestmentName, func(ctx context.Context, t *testing.T, cfg *envconf.Config) config.Context {
			t.Logf("Running pre-upgrade Assestment %s": assestmentName)
			return assestment(ctx, t, cfg)
		})

	}

	//Upgrade Tests
	feature = feature.Assess("Upgrade Provider", func(ctx context.Context, t *testing.T, cfg *envconf.Config) config.Context {
			t.Logf("Upgrading provider from %s to %s", c.fromVersion, c.toVersion)
			//TODO: Put in code for upgrade provider
			return ctx
	})

	// Post-upgrade assessments
	for name, assestmentFunc := range c.postUpgradeAssessment {
		assestment := assestmentFunc //Capture for closure
		assestmentName := name 
		feature = feature.Assess(assestmentName, func(ctx context.Context, t *testing.T, cfg *envconf.Config) config.Context {
			t.Logf("Running post-upgrade Assestment %s": assestmentName)
			return assestment(ctx, t, cfg)
		})

	}
	//Teardown Phase
	feature = feature.Teardown(func(ctx context.Context, t *Testing.T, cfg *envconf.Config) context.Context {
		t.Logf("Cleaning up test resources")
		//TODO: put in code to delete test resources
		return ctx
	})
	return feature

 }

 func (c *CustomUpgradeTest) Run(t *Testing.T, cfg *envconf.Config) {
	feature := c.Build(t)
	testenv.Test(t, feature.Feature())
 }


