# Image Builder Plugin for Packer

A plugin for packer that provides a builder using [Image
Builder](https://osbuild.org) to build images.

Note that Image Builder does not spin up instances of the images it creates.
Thus, packer's provisioners do not work with this builder. To run a provisioner,
run an additional builder with this builder's output.

## Running from source

To run this plugin from source, first build it:

    go generate ./...
    go build .

Then, run it with

    PACKER_PLUGIN_PATH=$PWD PKR_VAR_offline_token=$TOKEN packer build example.pkr.hcl

Where `$TOKEN` is an offline token for Red Hat's API, generated as described in
[Red Hat's documentation][redhat-api], and `example.pkr.hcl` contains something
like this:

```hcl
variable "offline_token" {
    type = string
}

source "image-builder" "example" {
    offline_token = var.offline_token

    blueprint {
      distribution = "rhel-8"
      packages = [ "postgresql" ]
    }

    aws_targets {
      architecture = "x86_64"
      share_with_accounts = [ "YOUR_AWS_ACCOUNT_ID" ]
    }
}

build {
    sources = [ "source.image-builder.example" ]
}
```

[redhat-api]: https://access.redhat.com/articles/3626371
