# Terraform provider that manages Ghost apps #

TF example:

```
provider "ghost" {
  user     = "admin"
  password = "mypass"
  endpoint = "https://demo.ghost.morea.fr"
}

resource "ghost_app" "wordpress" {
  name = "wordpress"
  env  = "prod"
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
```

# Building and testing in a Dockerized environment #

The dockerized environment can be invoked through docker-compose. It mounts an external volume named _terraform-provider-ghost-root_user_ in which you can personalize the root session.

Hint: in the _terraform-provider-ghost-root_user_ volume, you can import your personal SSH keys and add a .gitconfig file with a content like this:    
```
[url "git@bitbucket.org:"]
  insteadOf = https://bitbucket.org/
```
To allow go to load git dependencies from private repositories using your SSH key.

To use the environment, just run the command:   
`docker-compose run --rm [test|build]-env` in the root directory of this project.

# Running terraform in a Dockerized environment #

To run the provider after building, just run the command:
`docker-compose run --rm terraform [plan|<terraform command>]` in the root directory of this project

the _sources/build_ directory is mounted directly in the container and should be linked by a _.terraformrc_ file in the _terraform-provider-ghost-root_user_ volume, example:
```
providers {
  ghost = "/terraform-providers/ghost/terraform-provider-ghost"
}
```
