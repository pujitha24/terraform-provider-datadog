package fwprovider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestApmRetentionFilterUpdateStateTraceRatePrecision(t *testing.T) {
	r := &ApmRetentionFilterResource{}
	state := &ApmRetentionFilterModel{
		Rate:      types.StringValue("0.1"),
		TraceRate: types.StringValue("0.01"),
	}
	r.updateState(context.Background(), state, "id", "foo", 0.1, true, 0.01, "service:foo", true, "spans-sampling-processor")
	if got := state.Rate.ValueString(); got != "0.1" {
		t.Errorf("rate = %q, want 0.1", got)
	}
	if got := state.TraceRate.ValueString(); got != "0.01" {
		t.Errorf("trace_rate = %q, want 0.01", got)
	}
}
