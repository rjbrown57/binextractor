package binextractor

/*
https://github.com/google/go-containerregistry/blob/main/pkg/crane/export.go
*/

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

// getImage returns a v1.Image to be acted on
func getImage(image string) v1.Image {
	ref, err := name.ParseReference(image)
	if err != nil {
		panic(err)
	}

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		panic(err)
	}

	return img
}

func Extract(image string, sourcePath string, destinationPath string) {

	fmt.Printf("Getting image %s, extracting %s to %s\n", image, sourcePath, destinationPath)

	sourcePath = strings.TrimPrefix(sourcePath, "/")
	img := getImage(image)

	flatFS := mutate.Extract(img)
	defer flatFS.Close()

	imageTar := tar.NewReader(flatFS)

	for {
		header, err := imageTar.Next()
		if err == io.EOF {
			fmt.Printf("%s not found\n", sourcePath)
			panic(err)
		} else if err != nil {
			panic(err)
		}

		if header.Name == sourcePath {
			break
		}
	}

	f, err := os.Create(destinationPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(f, imageTar)
	if err != nil {
		panic(err)
	}

	err = os.Chmod(destinationPath, 0700)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Extracted %s to %s\n", sourcePath, destinationPath)
}
