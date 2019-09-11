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


const TF_Cron_Job = `resource "kubernetes_cron_job" "demo" {
	metadata {
	  name = "demo"
	}
	spec {
	  concurrency_policy            = "Replace"
	  failed_jobs_history_limit     = 5
	  schedule                      = "1 0 * * *"
	  starting_deadline_seconds     = 10
	  successful_jobs_history_limit = 10
	  suspend                       = true
	  job_template {
		metadata {}
		spec {
		  backoff_limit = 2
		  template {
			metadata {}
			spec {
			  container {
				name    = "hello"
				image   = "busybox"
				command = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]
			  }
			  restart_policy = "OnFailure"
			}
		  }
		}
	  }
	}
}`


const TF_CR = `resource "kubernetes_cluster_role" "example" {
	metadata {
	  name = "terraform-example"
	}
  
	rule {
	  api_groups = [""]
	  resources  = ["namespaces", "pods"]
	  verbs      = ["get", "list", "watch"]
	}
}`

const TF_CRB = `resource "kubernetes_cluster_role_binding" "example" {
	metadata {
	  name = "terraform-example"
	}
	role_ref {
	  api_group = "rbac.authorization.k8s.io"
	  kind      = "ClusterRole"
	  name      = "cluster-admin"
	}
	subject {
	  kind      = "User"
	  name      = "admin"
	  api_group = "rbac.authorization.k8s.io"
	}
	subject {
	  kind      = "ServiceAccount"
	  name      = "default"
	  namespace = "kube-system"
	}
	subject {
	  kind      = "Group"
	  name      = "system:masters"
	  api_group = "rbac.authorization.k8s.io"
	}
}`
