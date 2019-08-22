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

	switch strings.ToLower(kind) {
	case "namespace", "ns":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Namespace)

	case "deployment", "deploy", "dpl":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Deployment)

	case "persistentvolume", "pv":
	case "pod", "po":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Pod)
		
	case "persistentvolumeclaim", "pvc":
	case "statefulset", "sts":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.StatefulSet)

	case "configmap", "cm":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.ConfigMap)

	case "secret", "sc":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Secret)

	case "serviceaccount", "sa":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.ServiceAccount)

	case "service", "svc":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Service)

	case "kustomize", "kz":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Kustomize)

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
