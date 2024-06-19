package util

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func ReplaceIfInt64Diff() planmodifier.Int64 {
	return int64planmodifier.RequiresReplaceIf(func(ctx context.Context, request planmodifier.Int64Request, response *int64planmodifier.RequiresReplaceIfFuncResponse) {
		response.RequiresReplace = !request.PlanValue.Equal(request.StateValue)
	},
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.",
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.")
}

func ReplaceIfStringDiff() planmodifier.String {
	return stringplanmodifier.RequiresReplaceIf(func(ctx context.Context, request planmodifier.StringRequest, response *stringplanmodifier.RequiresReplaceIfFuncResponse) {
		response.RequiresReplace = !request.PlanValue.Equal(request.StateValue)
	},
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.",
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.")
}
