# BinExtractor

Extract a binary from a container image. For example `binextractor -i "bitnami/kubectl:1.20.12" -s "opt/bitnami/kubectl/bin/kubectl" -d "kubectl"`

utilizes https://github.com/google/go-containerregistry

## Flags 
* -i --image, Image name
* -s --sourcepath, path within image to extract
* -d --destpath, destination path of image

## Fixes
* handle multiplatform images
* Verify + sanitize output destination pre image pull
* Wildcards for the search binary
* Optional configureable perms
