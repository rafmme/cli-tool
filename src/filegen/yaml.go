package filegen

import (
	"strings"
	"log"
	"../populatefile"
	"../tmpl"
	"../runcmd"
)

func CreateYAML(kind, fileName, path string)  {
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

	case "cronjob", "crj":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.CronJob)

	case "clusterrolebinding", "crb":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.GenerateRoleBinding("ClusterRoleBinding", "ClusterRole"))

	case "clusterrole", "cr":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.ClusterRole)

	case "job":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Job)

	case "daemonset", "dmt":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.DaemonSet)

	case "deployment", "deploy", "dpl":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Deployment)

	case "persistentvolume", "pv":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.PersistentVolume)

	case "pod", "po":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Pod)
	
	case "ingress", "ing":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Ingress)

	case "persistentvolumeclaim", "pvc":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.PersistentVolumeClaim)

	case "role":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Role)

	case "rolebinding", "rb":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.GenerateRoleBinding("RoleBinding", "Role"))

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

	case "policy", "pol":
		fullPath, err := runcmd.Execute(path, generatedFileName)
		populatefile.IsError(err)
		populatefile.WriteContentToFile(*fullPath, tmpl.Policy)

	default:
		log.Printf(`Can't process k8s object type of %s`, kind)
		return
	}
}
