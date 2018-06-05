provider "ghost" {
  user     = "demo"
  password = "${var.password}"
  endpoint = "https://localhost"
}

// This example shows the minimal configuration required to create an app.
// Be aware that such a minimal configuration is not recommended for real usage.
resource "ghost_app" "basic" {
  name   = "wordpress"
  env    = "dev"
  role   = "webfront"
  vpc_id = "vpc-1234567"

  build_infos = {
    subnet_id  = "subnet-1234567"
    source_ami = "ami-1234567"
  }

  environment_infos = {}

  modules = []
}
