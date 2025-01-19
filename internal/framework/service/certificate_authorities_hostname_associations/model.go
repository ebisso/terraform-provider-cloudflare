package certificate_authorities_hostname_associations

import "github.com/hashicorp/terraform-plugin-framework/types"

type CertificateAuthoritiesHostnameAssociationsModel struct {
	ZoneID            types.String `tfsdk:"zone_id"`
	MTLSCertificateID types.String `tfsdk:"mtls_certificate_id"`
	Hostnames         types.List   `tfsdk:"hostnames"`
}
