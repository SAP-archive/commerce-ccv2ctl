
**Disclaimer**: *This tool is not officially supported by SAP. Use at your own risk*

# ccv2ctl - SAP Commerce Cloud Portal CLI

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/W7W7VS24)

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

       ccv2ctl create deployment --build 20180930.3 --environment d1 --mode none

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

## Installation

Authentication is done via the client certificate for a S-User, the [SAP Passport](https://support.sap.com/en/my-support/single-sign-on-passports.html)

The S-User needs to be configured as `CUSTOMER_DEVELOPER` or `CUSTOMER_SYS_ADMIN` in the CCV2 Cloud Portal.

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
        
   I recommend configuring the default `subscription`
   
1. Enjoy

> You can provide all configuration options also on the commend line or as environment variables (prefixed with `CCV2_`) \
> Check the output of `ccv2ctl help` for further details.

## How it works

It simulates the single-sign-on flow of the frontend using a SAP Passport client certificate.

Once we have an authenticated session, we can call the same REST APIs that the Cloud Portal uses.

## Contribution Guide

Here is a good explaination on how to fork and contribute to a  Go project without changing import paths in your fork:

<http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path>



