provider "ghost" {
  user     = "demo"
  password = "${var.password}"
  endpoint = "https://demo.ghost.morea.fr"
}

resource "ghost_app" "test" {
  name = "wordpress"
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

  modules = [{
    name       = "symfony2"
    pre_deploy = "ZXhpdCAx"
    path       = "/var/www"
    scope      = "code"
    git_repo   = "https://github.com/KnpLabs/KnpIpsum.git"
  }]

  features = [{
    version = "5.4"
    name    = "php5"
  },
  {
    version = "2.2"
    name    = "apache2"
  }]
}
