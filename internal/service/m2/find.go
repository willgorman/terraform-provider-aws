// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package m2

import (
	"context"

	m2 "github.com/aws/aws-sdk-go-v2/service/m2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindEnvByID(ctx context.Context, conn *m2.Client, id string) (*awstypes.EnvironmentSummary, error) {
	in := &m2.GetEnvironmentInput{
		EnvironmentId: aws.String(id),
	}
	out, err := conn.GetEnvironment(ctx, in)

	if errs.IsA[*awstypes.ResourceNotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: in,
		}
	}

	if err != nil {
		return nil, err
	}

	if out == nil || out.EnvironmentId == nil {
		return nil, tfresource.NewEmptyResultError(in)
	}

	return &awstypes.EnvironmentSummary{}, nil
}
