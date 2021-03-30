// Code generated by "mapstructure-to-hcl2 -type Config,Blueprint,AWSTarget"; DO NOT EDIT.

package main

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatAWSTarget is an auto-generated flat version of AWSTarget.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAWSTarget struct {
	Architecture      *string  `mapstructure:"architecture,required" cty:"architecture" hcl:"architecture"`
	ShareWithAccounts []string `mapstructure:"share_with_accounts" cty:"share_with_accounts" hcl:"share_with_accounts"`
}

// FlatMapstructure returns a new FlatAWSTarget.
// FlatAWSTarget is an auto-generated flat version of AWSTarget.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AWSTarget) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatAWSTarget)
}

// HCL2Spec returns the hcl spec of a AWSTarget.
// This spec is used by HCL to read the fields of AWSTarget.
// The decoded values from this spec will then be applied to a FlatAWSTarget.
func (*FlatAWSTarget) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"architecture":        &hcldec.AttrSpec{Name: "architecture", Type: cty.String, Required: false},
		"share_with_accounts": &hcldec.AttrSpec{Name: "share_with_accounts", Type: cty.List(cty.String), Required: false},
	}
	return s
}

// FlatBlueprint is an auto-generated flat version of Blueprint.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatBlueprint struct {
	Distribution *string  `mapstructure:"distribution,required" cty:"distribution" hcl:"distribution"`
	Packages     []string `mapstructure:"packages" cty:"packages" hcl:"packages"`
}

// FlatMapstructure returns a new FlatBlueprint.
// FlatBlueprint is an auto-generated flat version of Blueprint.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Blueprint) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatBlueprint)
}

// HCL2Spec returns the hcl spec of a Blueprint.
// This spec is used by HCL to read the fields of Blueprint.
// The decoded values from this spec will then be applied to a FlatBlueprint.
func (*FlatBlueprint) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"distribution": &hcldec.AttrSpec{Name: "distribution", Type: cty.String, Required: false},
		"packages":     &hcldec.AttrSpec{Name: "packages", Type: cty.List(cty.String), Required: false},
	}
	return s
}

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	OfflineToken        *string           `mapstructure:"offline_token" cty:"offline_token" hcl:"offline_token"`
	Blueprint           *FlatBlueprint    `mapstructure:"blueprint" cty:"blueprint" hcl:"blueprint"`
	AWSTargets          []FlatAWSTarget   `mapstructure:"aws_targets" cty:"aws_targets" hcl:"aws_targets"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"offline_token":              &hcldec.AttrSpec{Name: "offline_token", Type: cty.String, Required: false},
		"blueprint":                  &hcldec.BlockSpec{TypeName: "blueprint", Nested: hcldec.ObjectSpec((*FlatBlueprint)(nil).HCL2Spec())},
		"aws_targets":                &hcldec.BlockListSpec{TypeName: "aws_targets", Nested: hcldec.ObjectSpec((*FlatAWSTarget)(nil).HCL2Spec())},
	}
	return s
}