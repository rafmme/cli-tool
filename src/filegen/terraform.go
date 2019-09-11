package filegen

import (
	"log"
	"strings"
	"../tmpl"
)

func CreateTerraformFile(kind, fileName, path string)  {
	generatedFileName := fileName+".tf"
	
	if kind == "" {
		log.Print("Error! Expects the `kind` flag")
		return
	}

	switch strings.ToLower(kind) {
	case "namespace", "ns":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Namespace)

	case "provider", "pvd":
		HandleFileCreation(path, generatedFileName, tmpl.TF_K8S_PROVIDER)

	case "cronjob", "crj":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Cron_Job)

	case "clusterrolebinding", "crb":
		HandleFileCreation(path, generatedFileName, tmpl.TF_CRB)

	case "clusterrole", "cr":
		HandleFileCreation(path, generatedFileName, tmpl.TF_CR)

	case "job":

	case "daemonset", "dmt":

	case "deployment", "deploy", "dpl":

	case "persistentvolume", "pv":

	case "pod", "po":
	
	case "ingress", "ing":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Ingress)

	case "persistentvolumeclaim", "pvc":

	case "role":

	case "rolebinding", "rb":

	case "statefulset", "sts":

	case "configmap", "cm":

	case "secret", "sc":

	case "serviceaccount", "sa":

	case "service", "svc":

	case "policy", "npol":

	default:
		log.Printf(`Can't process k8s object type of %s`, kind)
		return
	}
}
