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
  vpc_id        = "vpc-1234567"

  build_infos = {
    subnet_id    = "subnet-1234567"
    ssh_username = "admin"
    source_ami   = "ami-1234567"
  }

  environment_infos = {
    instance_profile = "iam.ec2.demo"
    key_name         = "ghost-demo"

    subnet_ids      = ["subnet-1234567"]
    security_groups = ["sg-1234567"]
  }

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

    // You can load scripts from files
    post_deploy = "${file("post_deploy.txt")}"

    // You can also use heredocs
    pre_deploy = <<-SCRIPT
                    #!/bin/bash
                    echo "EXAMPLE_CONFIG" >> /var/www/html/wp-config.php
                    SCRIPT
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
    name        = "feature_1"
    version     = "1"
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
  }
}
