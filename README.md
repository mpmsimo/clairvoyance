# clairvoyance
Drift detection and reporting for Terraform.

Repository was bootstrapped with [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang).

## Getting started
This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

### Project setup
This project requires Go to be installed. 

On OS X with Homebrew you can just run `brew install go`.


## Development
### Build and Run
Run the binary after it's been packaged:
```console
$ make
$ ./bin/clairvoyance
```

### Testing
``make test``

## Additional Information
Repository was initially bootstrapped with [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang).

### Notable packages
Packages can be downloaded from public GitHub repositories, like so:
`https://github.com/$USER/$REPO`

Modules that are intended to be used are documented below.
```
[tfvar](https://github.com/shihanng/tfvar) - programatic definition and generation of variables based on user input
[hclwrite](https://github.com/hashicorp/hcl/tree/v2.0.0/hclwrite) - write HCL on the fly
[terraform-exec](https://github.com/kmoe/terraform-exec) - so we can init/plan/apply via the Terraform CLI programmatically.
[terrafmt](https://github.com/terrycain/terrafmt) - format the HCL output, if live update is used
```
