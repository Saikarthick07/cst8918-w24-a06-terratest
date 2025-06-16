package test

import (
	"testing"
	
	"strings"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureLinuxVMCreation(t *testing.T) {
	t.Parallel()

	subscriptionID := "b8f4d954-1b62-49d7-800a-57d856397796" // Replace this
	labelPrefix := "saikarthick"

	terraformOptions := &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"labelPrefix": labelPrefix,
		},

		EnvVars: map[string]string{
			"ARM_SUBSCRIPTION_ID": subscriptionID,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Example test: check the VM name output
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	assert.True(t, strings.HasPrefix(vmName, labelPrefix))
}
