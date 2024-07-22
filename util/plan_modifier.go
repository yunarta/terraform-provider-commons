package util

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func ReplaceIfInt64Diff() planmodifier.Int64 {
	return int64planmodifier.RequiresReplaceIf(func(ctx context.Context, request planmodifier.Int64Request, response *int64planmodifier.RequiresReplaceIfFuncResponse) {
		response.RequiresReplace = !request.PlanValue.Equal(request.StateValue)
	},
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.",
		"If the value of this attribute changes, Terraform will destroy and recreate the resource.")
}

type replaceIfStringDiff struct {
}

func (r replaceIfStringDiff) Description(ctx context.Context) string {
	return "If the value of this attribute changes, Terraform will destroy and recreate the resource."
}

func (r replaceIfStringDiff) MarkdownDescription(ctx context.Context) string {
	return "If the value of this attribute changes, Terraform will destroy and recreate the resource."
}

func (r replaceIfStringDiff) PlanModifyString(ctx context.Context, request planmodifier.StringRequest, response *planmodifier.StringResponse) {
	if request.PlanValue.IsUnknown() || request.PlanValue.IsNull() {
		response.RequiresReplace = false
	} else {
		response.RequiresReplace = !request.PlanValue.Equal(request.StateValue)
	}
}

var _ planmodifier.String = &replaceIfStringDiff{}

func ReplaceIfStringDiff() planmodifier.String {
	return &replaceIfStringDiff{}
}
