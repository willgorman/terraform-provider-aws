// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package configservice

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	configservice_sdkv2 "github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newRetentionConfigurationResource,
			Name:    "Retention Configuration",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAggregateAuthorization,
			TypeName: "aws_config_aggregate_authorization",
			Name:     "Aggregate Authorization",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceConfigRule,
			TypeName: "aws_config_config_rule",
			Name:     "Config Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceConfigurationAggregator,
			TypeName: "aws_config_configuration_aggregator",
			Name:     "Configuration Aggregator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceConfigurationRecorder,
			TypeName: "aws_config_configuration_recorder",
			Name:     "Configuration Recorder",
		},
		{
			Factory:  resourceConfigurationRecorderStatus,
			TypeName: "aws_config_configuration_recorder_status",
			Name:     "Configuration Recorder Status",
		},
		{
			Factory:  resourceConformancePack,
			TypeName: "aws_config_conformance_pack",
			Name:     "Conformance Pack",
		},
		{
			Factory:  resourceDeliveryChannel,
			TypeName: "aws_config_delivery_channel",
			Name:     "Delivery Channel",
		},
		{
			Factory:  resourceOrganizationConformancePack,
			TypeName: "aws_config_organization_conformance_pack",
			Name:     "Organization Conformance Pack",
		},
		{
			Factory:  resourceOrganizationCustomPolicyRule,
			TypeName: "aws_config_organization_custom_policy_rule",
			Name:     "Organization Custom Policy Rule",
		},
		{
			Factory:  resourceOrganizationCustomRule,
			TypeName: "aws_config_organization_custom_rule",
			Name:     "Organization Custom Rule",
		},
		{
			Factory:  resourceOrganizationManagedRule,
			TypeName: "aws_config_organization_managed_rule",
			Name:     "Organization Managed Rule",
		},
		{
			Factory:  resourceRemediationConfiguration,
			TypeName: "aws_config_remediation_configuration",
			Name:     "Remediation Configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ConfigService
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*configservice_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return configservice_sdkv2.NewFromConfig(cfg, func(o *configservice_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
