package filegen

import (
	"flag"
	"strings"
	"log"
	"../populatefile"
	"../tmpl"
	"../runcmd"
)

func Create(kind, fileName, path string)  {
	generatedFileName := fileName+".yaml"
	
	if kind == "" {
		log.Print("Error! Expects the `kind` flag")
		return
	}

	switch kind {
	case "namespace", "ns":
		out, err := runcmd.Execute(path, generatedFileName)
		fullPath := strings.TrimSpace(string(out[:])) + "/" + generatedFileName

		if err != nil {
			log.Fatal(err)
			return
		}

		populatefile.WriteContentToFile(fullPath, tmpl.Namespace)
		break

	case "deployment", "deploy":
	case "persistentvolume", "pv":
	case "pod", "po":
	case "persistentvolumeclaim", "pvc":
	case "configmap", "cm":
	case "secret", "sc":
	case "serviceaccount", "sa":
	case "service", "svc":
	default:
		log.Print(`doee`)
		return
	}
}

func Run() {
	kind := flag.String("kind", "", "Kubernetes API Object")
	path := flag.String("path", ".", "Path to create the file [OPTIONAL]")
	fileName := flag.String("filename", "default", "Name for the generated file [OPTIONAL]")
	
	flag.Parse()

	Create(*kind, *fileName, *path)
}
