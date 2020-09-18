# ccv2ctl - SAP Commerce Cloud Portal CLI

[![REUSE status](https://api.reuse.software/badge/github.com/SAP/commerce-ccv2ctl)](https://api.reuse.software/info/github.com/SAP/commerce-ccv2ctl)

This tool allows you to:

- get data about builds and create (=start) new builds:

      ccv2ctl get build 20180930.2
      # data about a specific build
      
      ccv2ctl get build --all
      # data of all builds
      
      ccv2ctl create build --name some-name --branch production
      # create a new build
      
      ccv2ctl logs 20180930.2
      # dump build log to stdout
      
- check the data about running deployments

      ccv2ctl get deployment d1
      
- trigger a deployment

      ccv2ctl create deployment --build 20180930.3 --environment d1

- get data about previous deployments

      ccv2ctl get deploymenthistory d1

- get data about initial passwords

      ccv2ctl get initialpassword d1

- get customer properties

      ccv2ctl get customerproperties d1 --aspect hcs_common

- set customer properties

      ccv2ctl set customerproperties d1 --aspect hcs_common --propertyfile hcs_common.properties
      # set properties from file hcs_common.properties

      echo accstore.environment=dev | ccv2ctl set customerproperties d1 --aspect hcs_common
      # set properties from stdin

**Every command has a detailed help available, make sure to check it! (use `ccv2ctl help` as an entry point)**

Since Go compiles executables as statically linked binaries, you can easily distribute this tool to you CI/CD servers (Jenkins, for example) and use it to automate builds and deployments to the Commerce Cloud.

<!-- TOC depthFrom:2 depthTo:2 -->

- [Installation](#installation)
- [How it works](#how-it-works)
- [Support](#support)
- [Contributing](#contributing)

<!-- /TOC -->

## Installation

Authentication is done via the client certificate for a S-User, the [SAP Passport](https://support.sap.com/en/my-support/single-sign-on-passports.html)

The S-User needs to be configured as `CUSTOMER_DEVELOPER` or `CUSTOMER_SYS_ADMIN` in the CCV2 Cloud Portal.

**Make sure to *disable* two-factor authentication when you add the S-User to your subscription!**\
If you enable 2FA _even once_, you _cannot_ get rid of it without a support ticket.

You need to export the certificate and the key into two PEM-encoded files so they can be used for `ccv2ctl` (you will be prompted for the keystore password):

    openssl pkcs12 -in /path/to/store.pfx -nokeys -nodes | openssl x509 -out certfile.pem
    openssl pkcs12 -in /path/to/store.pfx -nocerts -nodes |  openssl rsa -out keyfile.pem

1. Clone the repo and build the binary
   ```
    go get github.com/sap-commerce-tools/ccv2ctl
    go build github.com/sap-commerce-tools/ccv2ctl
   ```
1. put the binary somewhere on your `$PATH`, or just run `go install github.com/sap-commerce-tools/ccv2ctl`
1. create the config file `.ccv2ctl.yaml` in your [home directory](https://en.wikipedia.org/wiki/Home_directory) with following content:

        certfile: /path/to/certfile.pem
        # Path to PEM-encoded SAP Passport client certificate
    
        keyfile: /path/to/keyfile.pem
        # Path to PEM-encoded key of SAP Passport client certificate
    
        # subscription: c0deba5ec0deba5ec0deba5ec0deba5e
        # (optional) Default subscription-ID to use for all commands. You can find the ID in the URL of the cloud portal.
        # https://portal.commerce.ondemand.com/subscription/c0deba5ec0deba5ec0deba5ec0deba5e/...
        #                                                   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    
        # cookiejar: /path/to/jar
        # (optional) path to HTTP cookie jar to avoid DDOSing the portal.
        # Default value: $HOME/.ccv2jar
        
   I recommend configuring the default `subscription`, provided you only work on a single project.
   
1. Enjoy

> You can provide all configuration options also on the commend line or as environment variables (prefixed with `CCV2_`) \
> Check the output of `ccv2ctl help` for further details.

## How it works

It simulates the single-sign-on flow of the frontend using a SAP Passport client certificate.

Once we have an authenticated session, we can call the same REST APIs that the Cloud Portal uses.

## Support

If you encounter any bugs or have an idea for a new feature, please submit a new [issue] in this repository.

[issue]: https://github.com/SAP/commerce-ccv2ctl/issues

## Contributing

Any contributions are welcome!

Please:
1. [Fork] the repository
1. Implement and
1. **Test** your changes
1. Format the code (`gofmt -l -s -w .`)
1. Open a new [pull request][pr]

[Fork]: https://help.github.com/articles/fork-a-repo
[pr]: https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request

### Developer Certificate of Origin (DCO)

Due to legal reasons, contributors will be asked to accept a DCO before they submit the first pull request to this projects, this happens in an automated fashion during the submission process. SAP uses [the standard DCO text of the Linux Foundation](https://developercertificate.org/).
