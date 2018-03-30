package ghost

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGhostAppImportBasic(t *testing.T) {
	envName := fmt.Sprintf("import_ghost_app_acc_env_basic_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGhostAppConfig(envName),
			},

			resource.TestStep{
				ResourceName:      "ghost_app.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
