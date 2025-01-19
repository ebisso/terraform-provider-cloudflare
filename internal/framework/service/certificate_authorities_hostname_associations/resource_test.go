package certificate_authorities_hostname_associations_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareCertificateAuthoritiesHostnameAssociations_BYO_CA(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_certificate_authorities_hostname_associations.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname1 := rnd + "." + zoneName
	hostname2 := rnd + "." + zoneName

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Account(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificateAuthoritiesHostnameAssociationsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCertificateAuthoritiesHostnameAssociationsConfigBYOCA(rnd, accountID, zoneID, []string{hostname1}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "mtls_certificate_id"),
					//resource.TestCheckResourceAttr(name, "hostnames.0", hostname1),
					//resource.TestCheckResourceAttr(name, "hostnames.1", hostname2),
				),
			},
			{
				Config: testCertificateAuthoritiesHostnameAssociationsConfigBYOCA(rnd, accountID, zoneID, []string{hostname1, hostname2}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "mtls_certificate_id"),
					//resource.TestCheckResourceAttr(name, "hostnames.0", hostname1),
					//resource.TestCheckResourceAttr(name, "hostnames.1", hostname2),
				),
			},
		},
	})
}

func TestAccCloudflareCertificateAuthoritiesHostnameAssociations_DefaultCA(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_certificate_authorities_hostname_associations.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname1 := rnd + "." + zoneName
	hostname2 := rnd + "." + zoneName

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Account(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificateAuthoritiesHostnameAssociationsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCertificateAuthoritiesHostnameAssociationsConfigDefaultCA(rnd, zoneID, []string{hostname1}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "hostnames.0", hostname1),
				),
			},
			{
				Config: testCertificateAuthoritiesHostnameAssociationsConfigDefaultCA(rnd, zoneID, []string{hostname1, hostname2}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "hostnames.0", hostname1),
					resource.TestCheckResourceAttr(name, "hostnames.1", hostname2),
				),
			},
		},
	})
}

func testAccCheckCloudflareCertificateAuthoritiesHostnameAssociationsDestroy(s *terraform.State) error {
	client, err := acctest.SharedV1Client()
	if err != nil {
		return fmt.Errorf("Failed to create Cloudflare client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_certificate_authorities_hostname_associations" {
			continue
		}

		if rs.Primary.Attributes[consts.ZoneIDSchemaKey] != "" {
			hostnames, _ := client.ListCertificateAuthoritiesHostnameAssociations(context.Background(), cfv1.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), cfv1.ListCertificateAuthoritiesHostnameAssociationsParams{
				MTLSCertificateID: rs.Primary.Attributes["mtls_certificate_id"],
			})
			if len(hostnames) != 0 {
				return fmt.Errorf("CertificateAuthoritiesHostnameAssociations still exists")
			}
		}

		if rs.Primary.Attributes[consts.AccountIDSchemaKey] != "" {
			hostnames, _ := client.ListCertificateAuthoritiesHostnameAssociations(context.Background(), cfv1.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), cfv1.ListCertificateAuthoritiesHostnameAssociationsParams{})
			if len(hostnames) != 0 {
				return fmt.Errorf("CertificateAuthoritiesHostnameAssociations still exists")
			}
		}
	}

	return nil
}

func testCertificateAuthoritiesHostnameAssociationsConfigBYOCA(rnd string, accountID string, zoneID string, hostnames []string) string {
	return fmt.Sprintf(`
resource "cloudflare_mtls_certificate" "%[1]s" {
	account_id   = "%[2]s"
	name         = ""
	certificates = "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoCgAwIBAgIUKTOAZNjcXVZRj4oQt0SHsl1c1vMwDQYJKoZIhvcNAQELBQAwUTELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDVNhbiBGcmFuY2lzY28xEzARBgNVBAcMCkNhbGlmb3JuaWExFTATBgNVBAoMDEV4YW1wbGUgSW5jLjAgFw0yMjExMjIxNjU5NDdaGA8yMTIyMTAyOTE2NTk0N1owUTELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDVNhbiBGcmFuY2lzY28xEzARBgNVBAcMCkNhbGlmb3JuaWExFTATBgNVBAoMDEV4YW1wbGUgSW5jLjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMRcORwgJFTdcG/2GKI+cFYiOBNDKjCZUXEOvXWY42BkH9wxiMT869CO+enA1w5pIrXow6kCM1sQspHHaVmJUlotEMJxyoLFfA/8Kt1EKFyobOjuZs2SwyVyJ2sStvQuUQEosULZCNGZEqoH5g6zhMPxaxm7ZLrrsDZ9maNGVqo7EWLWHrZ57Q/5MtTrbxQL+eXjUmJ9K3kS+3uEwMdqR6Z3BluU1ivanpPc1CN2GNhdO0/hSY4YkGEnuLsqJyDd3cIiB1MxuCBJ4ZaqOd2viV1WcP3oU3dxVPm4MWyfYIldMWB14FahScxLhWdRnM9YZ/i9IFcLypXsuz7DjrJPtPUCAwEAAaNmMGQwHQYDVR0OBBYEFP5JzLUawNF+c3AXsYTEWHh7z2czMB8GA1UdIwQYMBaAFP5JzLUawNF+c3AXsYTEWHh7z2czMA4GA1UdDwEB/wQEAwIBBjASBgNVHRMBAf8ECDAGAQH/AgEBMA0GCSqGSIb3DQEBCwUAA4IBAQBc+Be7NDhpE09y7hLPZGRPl1cSKBw4RI0XIv6rlbSTFs5EebpTGjhx/whNxwEZhB9HZ7111Oa1YlT8xkI9DshB78mjAHCKBAJ76moK8tkG0aqdYpJ4ZcJTVBB7l98Rvgc7zfTii7WemTy72deBbSeiEtXavm4EF0mWjHhQ5Nxpnp00Bqn5g1x8CyTDypgmugnep+xG+iFzNmTdsz7WI9T/7kDMXqB7M/FPWBORyS98OJqNDswCLF8bIZYwUBEe+bRHFomoShMzaC3tvim7WCb16noDkSTMlfKO4pnvKhpcVdSgwcruATV7y+W+Lvmz2OT/Gui4JhqeoTewsxndhDDE\n-----END CERTIFICATE-----"
	private_key  = ""
	ca           = true
}
resource "cloudflare_certificate_authorities_hostname_associations" "%[1]s" {
	zone_id             = "%[3]s"
    mtls_certificate_id = cloudflare_mtls_certificate.%[1]s.id
	hostnames           = ["%[4]s"]
}
`, rnd, accountID, zoneID, strings.Join(hostnames, `","`))
}

func testCertificateAuthoritiesHostnameAssociationsConfigDefaultCA(rnd string, zoneID string, hostnames []string) string {
	return fmt.Sprintf(`
resource "cloudflare_certificate_authorities_hostname_associations" "%[1]s" {
	zone_id   = "%[2]s"
	hostnames = ["%[3]s"]
}
`, rnd, zoneID, strings.Join(hostnames, `","`))
}
