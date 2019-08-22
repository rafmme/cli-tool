package filegen

import (
	"flag"
	"strings"
	"log"
)

func Run() {
	kind := flag.String("kind", "", "Kubernetes API Resource Object type **REQUIRED")
	path := flag.String("path", ".", "Destination Path to create the generated file [OPTIONAL: Uses current working directory if no path was specified]")
	fileName := flag.String("filename", "default", "Name for the generated file [OPTIONAL]")
	extension := flag.String(
		"ext", "yaml",
		"Specify file type to generate (`tf` or `yaml`); yaml is the default [OPTIONAL]",
	)

	flag.Parse()
	var file string = *fileName

	if strings.ToLower(*extension) != "tf" && strings.ToLower(*extension) != "yaml" {
		log.Printf(".%s file extension is not supported. Try tf or yaml", *extension)
		return
	}

	if *fileName == "default" {
		file = *kind
	}
	
	if strings.ToLower(*kind) == "kz" || strings.ToLower(*kind) == "kustomize" {
		file = "kustomization"
	}

	if strings.ToLower(*extension) == "tf" && (strings.ToLower(*kind) != "kz" && strings.ToLower(*kind) != "kustomize") {
		CreateTerraformFile(*kind, file, *path)
		return
	}

	CreateYAML(*kind, file, *path)
}
