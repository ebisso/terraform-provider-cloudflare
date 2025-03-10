package workers_for_platforms_dispatch_namespace_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_workers_for_platforms_dispatch_namespace", &resource.Sweeper{
		Name: "cloudflare_workers_for_platforms_dispatch_namespace",
		F: func(region string) error {
			client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			resp, err := client.ListWorkersForPlatformsDispatchNamespaces(ctx, cfv1.AccountIdentifier(accountID))
			if err != nil {
				return err
			}

			for _, namespace := range resp.Result {
				err := client.DeleteWorkersForPlatformsDispatchNamespace(ctx, cfv1.AccountIdentifier(accountID), namespace.NamespaceName)
				if err != nil {
					return err
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareWorkersForPlatforms_NamespaceManagement(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersForPlatformsNamespaceManagement(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// {
			// 	ResourceName:        resourceName,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func testAccCheckCloudflareWorkersForPlatformsNamespaceManagement(rnd, accountID string) string {
	return acctest.LoadTestCase("workersforplatformsnamespacemanagement.tf", rnd, accountID)
}

func testAccCheckCloudflareWorkerScriptDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_script" {
			continue
		}

		client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if err != nil {
			return fmt.Errorf("error establishing client: %w", err)
		}
		r, _ := client.GetWorkerWithDispatchNamespace(
			context.Background(),
			cfv1.AccountIdentifier(accountID),
			rs.Primary.Attributes["name"],
			rs.Primary.Attributes["dispatch_namespace"],
		)

		if r.Script != "" {
			return fmt.Errorf("namespaced worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
