provider "ghost" {
  user     = "demo"
  password = "${var.password}"
  endpoint = "https://localhost"
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
    instance_profile = "iam.ec2.demo"
    key_name         = "ghost-demo"

    root_block_device = {
      name = "testblockdevice"
      size = 20
    }

    optional_volumes = [{
      device_name = "/dev/xvdd"
      volume_type = "gp2"
      volume_size = 20
    }]

    subnet_ids      = ["subnet-a7e849fe"]
    security_groups = ["sg-6814f60c", "sg-2414f60c"]

    instance_tags = [{
      tag_name  = "Name"
      tag_value = "wordpress"
    },
      {
        tag_name  = "Type"
        tag_value = "front"
      },
    ]
  }

  autoscale = {
    name = "autoscale"
    min  = 1
    max  = 3
  }

  lifecycle_hooks = {
    pre_buildimage  = "#!/usr/bin/env bash"
    post_buildimage = "#!/usr/bin/env bash"
  }

  environment_variables = [{
    key   = "myvar"
    value = "myvalue"
  }]

  // Several modules and/or lists of modules can be merged together.
  modules = "${concat(list(local.custom_module_1, local.custom_module_2), local.basic_modules)}"

  features = ["${local.custom_feature}"]
}

// Defining modules in locals allows to reuse them into different apps without having
// to rewrite them.
locals {
  custom_module_1 = {
    name     = "module_1"
    path     = "/var/www"
    scope    = "code"
    git_repo = "https://github.com/KnpLabs/KnpIpsum.git"
  }

  custom_module_2 = {
    name     = "module_2"
    path     = "/var/www"
    scope    = "code"
    git_repo = "https://github.com/KnpLabs/KnpIpsum.git"
  }

  // It's also possible to define lists
  basic_modules = [
    {
      name     = "module_3"
      path     = "/var/w"
      scope    = "code"
      git_repo = "https://github.com/KnpLabs/KnpIpsum.git"
    },
    {
      name     = "module_4"
      path     = "/var/www"
      scope    = "code"
      git_repo = "https://github.com/KnpLabs/KnpIpsum.git"
    },
  ]
}

locals {
  custom_feature = {
    name    = "feature_1"
    version = "1"
  }
}
