package filegen

import (
	"strings"
	"log"
	"../tmpl"
)

func CreateYAML(kind, fileName, path string)  {
	generatedFileName := fileName+".yaml"
	
	if kind == "" {
		log.Print("Error! Expects the `kind` flag")
		return
	}

	switch strings.ToLower(kind) {
	case "namespace", "ns":
		HandleFileCreation(path, generatedFileName, tmpl.Namespace)

	case "cronjob", "crj":
		HandleFileCreation(path, generatedFileName, tmpl.CronJob)

	case "clusterrolebinding", "crb":
		HandleFileCreation(
			path, generatedFileName,
			tmpl.GenerateRoleBinding("ClusterRoleBinding", "ClusterRole"),
		)

	case "clusterrole", "cr":
		HandleFileCreation(path, generatedFileName, tmpl.ClusterRole)

	case "job":
		HandleFileCreation(path, generatedFileName, tmpl.Job)

	case "daemonset", "dmt":
		HandleFileCreation(path, generatedFileName, tmpl.DaemonSet)

	case "deployment", "deploy", "dpl":
		HandleFileCreation(path, generatedFileName, tmpl.Deployment)

	case "persistentvolume", "pv":
		HandleFileCreation(path, generatedFileName, tmpl.PersistentVolume)

	case "pod", "po":
		HandleFileCreation(path, generatedFileName, tmpl.Pod)
	
	case "ingress", "ing":
		HandleFileCreation(path, generatedFileName, tmpl.Ingress)

	case "persistentvolumeclaim", "pvc":
		HandleFileCreation(path, generatedFileName, tmpl.PersistentVolumeClaim)

	case "role":
		HandleFileCreation(path, generatedFileName, tmpl.Role)

	case "rolebinding", "rb":
		HandleFileCreation(
			path, generatedFileName,
			tmpl.GenerateRoleBinding("RoleBinding", "Role"),
		)

	case "statefulset", "sts":
		HandleFileCreation(path, generatedFileName, tmpl.StatefulSet)

	case "configmap", "cm":
		HandleFileCreation(path, generatedFileName, tmpl.ConfigMap)

	case "secret", "sc":
		HandleFileCreation(path, generatedFileName, tmpl.Secret)

	case "serviceaccount", "sa":
		HandleFileCreation(path, generatedFileName, tmpl.ServiceAccount)

	case "service", "svc":
		HandleFileCreation(path, generatedFileName, tmpl.Service)

	case "kustomize", "kz":
		HandleFileCreation(path, generatedFileName, tmpl.Kustomize)

	case "policy", "pol":
		HandleFileCreation(path, generatedFileName, tmpl.Policy)

	default:
		log.Printf(`Can't process k8s object type of %s`, kind)
		return
	}
}
