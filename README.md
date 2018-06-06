# terraform-provider-cloud-deploy
[![API Reference](http://img.shields.io/badge/api-reference-blue.svg)](https://docs.cloud-deploy.io/rst/api.html) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/claranet/terraform-provider-cloud-deploy/blob/master/LICENSE)

Terraform Provider that manages Cloud Deploy (Ghost) applications.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
- [Cloud Deploy](https://docs.cloud-deploy.io/) 18.05

Bulding The Provider
--------------------
Clone repository to: `$GOPATH/src/cloud-deploy.io/terraform-provider-cloud-deploy`

Using Go:
```sh
$ go get -d cloud-deploy.io/terraform-provider-cloud-deploy
```

Using git:
```sh
$ mkdir -p $GOPATH/src/cloud-deploy.io; cd $GOPATH/src/cloud-deploy.io
$ git clone git@github.com:claranet/terraform-provider-cloud-deploy.git
```

Enter the repository and build the provider
```sh
$ cd $GOPATH/src/cloud-deploy.io/terraform-provider-cloud-deploy
make
```

Using the provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

Examples are available in the examples directory:

- `full_app_model`: exposes all the configuration parameters available to create your cloud deploy app.
- `minimal_app_model`: shows the minimal configuration required to create a cloud deploy app.
- `basic_import`: shows how to ignore parameters during imports.
- `shared_modules_features`: shows how modules and features can be shared across ghost\_app resources using `locals`. It also shows how to write or import scripts.

Create a new Ghost App
---------------------------
First make sure the provider is installed as described above.

Create your app configuration using the examples available.

Run
```sh
$ terraform init # or tfwrapper init
$ terraform apply # or tfwrapper apply
```

Import an existing Ghost App
---------------------------
First make sure the provider is installed as described above.

Create your app configuration using the import examples available.

Run terraform import ghost_app.your_app app_id. Example:
```sh
$ terraform import ghost_app.basic_import 5accabf63d7eba00014e5679 # or tfwrapper import
```

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
