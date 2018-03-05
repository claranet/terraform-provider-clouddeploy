package ghost

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestGhostAppBasic(t *testing.T) {
	resourceName := "ghost_app.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testGhostAppConfig(resourceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "wordpress"),
					resource.TestCheckResourceAttrSet("env", "dv"),
				),
			},
		},
	})
}

func testGhostAppConfig(name string) string {
	return fmt.Sprintf(`
			resource "ghost_app" "test" {
				name = "%s"
			  env  = "dev"
			  role = "webfront"

			  region        = "eu-west-1"
			  instance_type = "t2.micro"
			  vpc_id        = "vpc-3f1eb65a"

			  log_notifications = [
			    "ghost-devops@domain.com",
			  ]

			  build_infos = {
			    subnet_id    = "subnet-a7e849fe"
			    ssh_username = "admin"
			    source_ami   = "ami-03ce4474"
			  }

			  environment_infos = {
			    instance_profile  = "iam.ec2.demo"
			    key_name          = "ghost-demo"
			    root_block_device = {}
			    optional_volumes  = []
			    subnet_ids        = ["subnet-a7e849fe"]
			    security_groups   = ["sg-6814f60c"]
			  }

			  autoscale = {
			    name = ""
			  }

			  module = {
			    name       = "symfony2"
			    pre_deploy = "ZXhpdCAx"
			    path       = "/var/www"
			    scope      = "code"
			    git_repo   = "https://github.com/KnpLabs/KnpIpsum.git"
			  }

			  feature = {
			    version = "5.4"
			    name    = "php5"
			  }

			  feature = {
			    version = "2.2"
			    name    = "apache2"
			  }
			}
			`, name)
}
