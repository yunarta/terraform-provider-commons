package util

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceIfInt64Diff(t *testing.T) {
	ctx := context.Background()
	request := planmodifier.Int64Request{
		StateValue: types.Int64Value(100),
		PlanValue:  types.Int64Value(200),
	}

	response := new(planmodifier.Int64Response)
	ReplaceIfInt64Diff().PlanModifyInt64(ctx, request, response)
	assert.Equal(t, true, response.RequiresReplace, "ReplaceIfInt64Diff() failed")
}

func TestReplaceIfStringDiff(t *testing.T) {
	ctx := context.Background()
	request := planmodifier.StringRequest{
		StateValue: types.StringValue("100"),
		PlanValue:  types.StringValue("200"),
	}

	response := new(planmodifier.StringResponse)
	ReplaceIfStringDiff().PlanModifyString(ctx, request, response)
	assert.Equal(t, true, response.RequiresReplace, "ReplaceIfInt64Diff() failed")
}
