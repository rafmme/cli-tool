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

	case "api_service", "aps":
		HandleFileCreation(path, generatedFileName, tmpl.TF_APS)

	case "limit_range", "lrg":
		HandleFileCreation(path, generatedFileName, tmpl.TF_LRange)

	case "replication_controller", "rplc":
		HandleFileCreation(path, generatedFileName, tmpl.TF_ReplC)

	case "resource_quota", "rsq":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Rsq)

	case "provider", "pvd":
		HandleFileCreation(path, generatedFileName, tmpl.TF_K8S_PROVIDER)

	case "cronjob", "crj":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Cron_Job)

	case "clusterrolebinding", "crb":
		HandleFileCreation(path, generatedFileName, tmpl.TF_CRB)

	case "clusterrole", "cr":
		HandleFileCreation(path, generatedFileName, tmpl.TF_CR)

	case "job":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Job)

	case "daemonset", "dmt":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Daemonset)

	case "deployment", "deploy", "dpl":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Deployment)

	case "endpoint", "ept":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Endpoint)

	case "hpa", "horizontal_pod_autoscaler":
		HandleFileCreation(path, generatedFileName, tmpl.TF_HPA)

	case "persistentvolume", "pv":
		HandleFileCreation(path, generatedFileName, tmpl.TF_PersistentVolume)

	case "pod", "po":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Pod)
	
	case "ingress", "ing":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Ingress)

	case "persistentvolumeclaim", "pvc":
		HandleFileCreation(path, generatedFileName, tmpl.TF_PVC)

	case "role":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Role)

	case "rolebinding", "rb":
		HandleFileCreation(path, generatedFileName, tmpl.TF_RB)

	case "statefulset", "sts":
		HandleFileCreation(path, generatedFileName, tmpl.TF_STS)

	case "configmap", "cm":
		HandleFileCreation(path, generatedFileName, tmpl.TF_ConfigMap)

	case "secret", "sc":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Secret)

	case "serviceaccount", "sa":
		HandleFileCreation(path, generatedFileName, tmpl.TF_SA)

	case "service", "svc":
		HandleFileCreation(path, generatedFileName, tmpl.TF_HPA)

	case "storage_class", "stc":
		HandleFileCreation(path, generatedFileName, tmpl.TF_STC)

	case "network_policy", "npol":
		HandleFileCreation(path, generatedFileName, tmpl.TF_Npol)

	default:
		log.Printf(`Can't process k8s object type of %s`, kind)
		return
	}
}
