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

  modules = [{
    name       = "wordpress"
    pre_deploy = ""
    path       = "/var/www"
    scope      = "code"
    git_repo   = "https://github.com/KnpLabs/KnpIpsum.git"
  },
    {
      name        = "wordpress2"
      pre_deploy  = "ZXhpdCAx"
      post_deploy = "ZXhpdCAx"
      path        = "/var/www"
      scope       = "code"
      git_repo    = "https://github.com/KnpLabs/KnpIpsum.git"
    },
  ]

  features = [{
    version = "5.4"
    name    = "php5"
  },
    {
      version = "2.2"
      name    = "apache2"
    },
  ]

  lifecycle_hooks = {
    pre_buildimage  = "#!/usr/bin/env bash"
    post_buildimage = "#!/usr/bin/env bash"
  }

  environment_variables = [{
    key   = "myvar"
    value = "myvalue"
  }]
}
