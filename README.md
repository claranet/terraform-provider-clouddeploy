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

Examples are available in the examples directory:
- basic_main: shows how to define a simple ghost application
- shared_modules_features_main: shows how modules and features can be shared across ghost_app resources using `locals`

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
