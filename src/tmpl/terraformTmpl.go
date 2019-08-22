package tmpl


const TF_Namespace = `resource "kubernetes_namespace" "example" {
	metadata {
	  annotations = {
		name = "example-annotation"
	  }
  
	  labels = {
		mylabel = "label-value"
	  }
  
	  name = "terraform-example-namespace"
	}
  }`

const TF_K8S_PROVIDER = `provider "kubernetes" {
	host = "https://104.196.242.174"
  
	client_certificate     = file("~/.kube/client-cert.pem")
	client_key             = file("~/.kube/client-key.pem")
	cluster_ca_certificate = file("~/.kube/cluster-ca-cert.pem")

	username = "ClusterMaster"
	password = "MindTheGap"
  }`
