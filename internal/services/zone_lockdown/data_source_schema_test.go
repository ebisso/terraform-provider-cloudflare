// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_lockdown"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneLockdownDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_lockdown.ZoneLockdownDataSourceModel)(nil)
	schema := zone_lockdown.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
