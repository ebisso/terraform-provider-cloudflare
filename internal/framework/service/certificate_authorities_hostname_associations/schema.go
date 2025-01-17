package certificate_authorities_hostname_associations

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *CertificateAuthoritiesHostnameAssociationsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Certificate Authorities Hostname Associations resource.",
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "The UUID for a certificate that was uploaded to the mTLS Certificate Management endpoint. If no mtls_certificate_id is given, the hostnames will be associated to your active Cloudflare Managed CA.",
				Optional:    true,
			},
			"hostnames": schema.ListAttribute{
				Description: "A list of hostnames associated to the certificate.",
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}
