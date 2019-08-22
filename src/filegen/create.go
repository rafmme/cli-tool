package filegen

import (
	"flag"
	"strings"
)

func Run() {
	kind := flag.String("kind", "", "Kubernetes API Object")
	path := flag.String("path", ".", "Path to create the file [OPTIONAL]")
	fileName := flag.String("filename", "default", "Name for the generated file [OPTIONAL]")
	extension := flag.String(
		"ext", "yaml",
		"Specify file type to generate (`tf` or `yaml`); yaml is the default [OPTIONAL]",
	)

	flag.Parse()
	var file string = *fileName

	if *fileName == "default" {
		file = *kind
	}
	
	if strings.ToLower(*kind) == "kz" || strings.ToLower(*kind) == "kustomize" {
		file = "kustomization"
	}

	if strings.ToLower(*extension) == "tf" && (strings.ToLower(*kind) != "kz" && strings.ToLower(*kind) != "kustomize") {
		CreateTerraformFile()
		return
	}

	CreateYAML(*kind, file, *path)
}
