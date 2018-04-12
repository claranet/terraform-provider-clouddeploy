provider "ghost" {
  user     = "demo"
  password = "${var.password}"
  endpoint = "https://localhost"
}

resource "ghost_app" "basic_import" {
  // You can ignore parameters using lifecycle.ignore_changes meta-parameter
  lifecycle {
    ignore_changes = ["log_notifications", "safe_deployment"]
  }

  name = "wordpress"
  env  = "dev"
  role = "webfront"

  region        = "eu-west-1"
  instance_type = "t2.micro"
  vpc_id        = "vpc-1234567"

  build_infos = {
    subnet_id    = "subnet-1234567"
    ssh_username = "admin"
    source_ami   = "ami-1234567"
  }

  environment_infos = {
    instance_profile  = "iam.ec2.demo"
    key_name          = "ghost-demo"
    public_ip_address = true

    root_block_device = {
      name = "testblockdevice"
      size = 20
    }

    optional_volumes = [
      {
        device_name = "/dev/xvdd"
        volume_type = "gp2"
        volume_size = 20
        iops        = 0
      },
    ]

    subnet_ids      = ["subnet-1234567"]
    security_groups = ["sg-1234567", "sg-1234567"]

    instance_tags = [
      {
        tag_name  = "Name"
        tag_value = "wordpress"
      },
    ]
  }

  autoscale = {
    name           = "autoscale"
    min            = 1
    max            = 3
    enable_metrics = true
  }

  modules = [
    {
      name             = "wordpress"
      path             = "/var/www"
      scope            = "code"
      git_repo         = "https://github.com/KnpLabs/KnpIpsum.git"
      uid              = 0
      gid              = 0
      build_pack       = ""
      pre_deploy       = ""
      post_deploy      = ""
      after_all_deploy = ""
    },
  ]

  features = [
    {
      name        = "php5"
      version     = "5.4"
      provisioner = "salt"
    },
    {
      name        = "package"
      provisioner = "ansible"

      parameters = <<-JSON
                      {
                        "package_name" : [
                          "nano",
                          "cowsay",
                          "ffmpeg",
                          "curl"
                        ]
                      }
                      JSON
    },
  ]

  lifecycle_hooks = {
    pre_buildimage = "echo PRE_BUILD >> /var/www/html/wp-config.php"

    post_buildimage = <<-SCRIPT
                         #!/bin/bash
                         echo "EXAMPLE_CONFIG" >> /var/www/html/wp-config.php
                         SCRIPT

    pre_bootstrap  = ""
    post_bootstrap = ""
  }

  environment_variables = [
    {
      key   = "myvar"
      value = "myvalue"
    },
  ]
}
