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

const TF_Ingress = `resource "kubernetes_ingress" "example_ingress" {
	metadata {
	  name = "example-ingress"
	}
  
	spec {
	  backend {
		service_name = "MyApp1"
		service_port = 8080
	  }
  
	  rule {
		http {
		  path {
			backend {
			  service_name = "MyApp1"
			  service_port = 8080
			}
  
			path = "/app1/*"
		  }
  
		  path {
			backend {
			  service_name = "MyApp2"
			  service_port = 8080
			}
  
			path = "/app2/*"
		  }
		}
	  }
  
	  tls {
		secret_name = "tls-secret"
	  }
	}
  }`
