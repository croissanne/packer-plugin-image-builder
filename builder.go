//go:generate go run github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc mapstructure-to-hcl2 -type Config,Blueprint,AWSTarget

package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"

	"github.com/larskarlitski/packer-plugin-image-builder/imagebuilder"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	OfflineToken string      `mapstructure:"offline_token"`
	Blueprint    Blueprint   `mapstructure:"blueprint"`
	AWSTargets   []AWSTarget `mapstructure:"aws_targets"`
}

type Blueprint struct {
	Distribution string   `mapstructure:"distribution,required"`
	Packages     []string `mapstructure:"packages"`
}

type AWSTarget struct {
	Architecture      string   `mapstructure:"architecture,required"`
	ShareWithAccounts []string `mapstructure:"share_with_accounts"`
}

type Builder struct {
	config Config
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	return b.config.FlatMapstructure().HCL2Spec()
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {
	err := config.Decode(&b.config, &config.DecodeOpts{
		PluginType:  "image-builder",
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	client, err := imagebuilder.NewClient(ctx, b.config.OfflineToken)
	if err != nil {
		return nil, err
	}

	cr := imagebuilder.ComposeRequest{
		Distribution: imagebuilder.Distributions(b.config.Blueprint.Distribution),
		Customizations: &imagebuilder.Customizations{
			Packages: &b.config.Blueprint.Packages,
		},
	}
	for _, t := range b.config.AWSTargets {
		ir := imagebuilder.ImageRequest{
			Architecture: t.Architecture,
			ImageType:    "ami",
			UploadRequest: imagebuilder.UploadRequest{
				Type: imagebuilder.UploadTypesAws,
				Options: imagebuilder.AWSUploadRequestOptions{
					ShareWithAccounts: t.ShareWithAccounts,
				},
			},
		}
		cr.ImageRequests = append(cr.ImageRequests, ir)
	}

	composeId, err := client.Compose(&cr)
	if err != nil {
		return nil, err
	}

	ui.Say(fmt.Sprintf("Started compose %v. Waiting for it to finish...", composeId))

	for {
		select {
		case <-time.After(15 * time.Second):
			status, err := client.GetComposeStatus(composeId)
			if err != nil {
				return nil, fmt.Errorf("error getting compose status: %v", err)
			}
			switch status.Status {
			case "failure":
				return nil, errors.New("image build failed")
			case "success":
				return artifact{composeId}, nil
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
