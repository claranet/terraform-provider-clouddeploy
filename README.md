Terraform Provider that manages Ghost apps
==========================================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Bulding The Provider
--------------------
Clone repository to: `$GOPATH/src/cloud-deploy.io/terraform-provider-ghost`

Using Go:
```sh
$ go get -d cloud-deploy.io/terraform-provider-ghost
```

Using git:
```sh
$ mkdir -p $GOPATH/src/cloud-deploy.io; cd $GOPATH/src/cloud-deploy.io
$ git clone git@bitbucket.org:morea/terraform-provider-ghost.git
```

Enter the repository and build the provider
```sh
$ cd $GOPATH/src/cloud-deploy.io/terraform-provider-ghost
make
```

Using the provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

An example is available in the examples directory.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

In order to test the provider, you can simply run `make test`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources. For the tests to pass, you'll need to setup a ghost instance (can be local), and set the following environment variables:

```sh
$ export GHOST_USER=myuser
$ export GHOST_PASSWORD=mypwd
$ export GHOST_ENDPOINT=http://localhost

$ make testacc
```

TF example:
-----------
```
provider "ghost" {
  user     = "admin"
  password = "mypass"
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
    root_block_device = {
      name = "testblockdevice"
      size = 20
    }
    optional_volumes  = [{
      device_name = "/dev/xvdd"
      volume_type = "gp2"
      volume_size = 20
    }]
    subnet_ids        = ["subnet-a7e849fe"]
    security_groups   = ["sg-6814f60c", "sg-2414f60c"]
    instance_tags			= [{
      tag_name  = "Name"
      tag_value = "wordpress"
    },
    {
      tag_name  = "Type"
      tag_value = "front"
    }]
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
  }]

  features = [{
    version = "5.4"
    name    = "php5"
  },
  {
    version = "2.2"
    name    = "apache2"
  }]

  lifecycle_hooks = {
    pre_buildimage  = "#!/usr/bin/env bash"
    post_buildimage = "#!/usr/bin/env bash"
  }

  environment_variables = [{
    key   = "myvar"
    value = "myvalue"
  }]
}
```
