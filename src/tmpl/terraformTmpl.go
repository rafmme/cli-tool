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


const TF_Job = `resource "kubernetes_job" "demo" {
	metadata {
	  name = "demo"
	}
	spec {
	  template {
		metadata {}
		spec {
		  container {
			name    = "pi"
			image   = "perl"
			command = ["perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"]
		  }
		  restart_policy = "Never"
		}
	  }
	  backoff_limit = 4
	}
}`

const TF_Daemonset = `resource "kubernetes_daemonset" "example" {
	metadata {
	  name      = "terraform-example"
	  namespace = "something"
	  labels = {
		test = "MyExampleApp"
	  }
	}
  
	spec {
	  selector {
		match_labels = {
		  test = "MyExampleApp"
		}
	  }
  
	  template {
		metadata {
		  labels = {
			test = "MyExampleApp"
		  }
		}
  
		spec {
		  container {
			image = "nginx:1.7.8"
			name  = "example"
  
			resources {
			  limits {
				cpu    = "0.5"
				memory = "512Mi"
			  }
			  requests {
				cpu    = "250m"
				memory = "50Mi"
			  }
			}
  
			liveness_probe {
			  http_get {
				path = "/nginx_status"
				port = 80
  
				http_header {
				  name  = "X-Custom-Header"
				  value = "Awesome"
				}
			  }
  
			  initial_delay_seconds = 3
			  period_seconds        = 3
			}
  
		  }
		}
	  }
	}
}`


const TF_Deployment = `resource "kubernetes_deployment" "example" {
	metadata {
	  name = "terraform-example"
	  labels = {
		test = "MyExampleApp"
	  }
	}
  
	spec {
	  replicas = 3
  
	  selector {
		match_labels = {
		  test = "MyExampleApp"
		}
	  }
  
	  template {
		metadata {
		  labels = {
			test = "MyExampleApp"
		  }
		}
  
		spec {
		  container {
			image = "nginx:1.7.8"
			name  = "example"
  
			resources {
			  limits {
				cpu    = "0.5"
				memory = "512Mi"
			  }
			  requests {
				cpu    = "250m"
				memory = "50Mi"
			  }
			}
  
			liveness_probe {
			  http_get {
				path = "/nginx_status"
				port = 80
  
				http_header {
				  name  = "X-Custom-Header"
				  value = "Awesome"
				}
			  }
  
			  initial_delay_seconds = 3
			  period_seconds        = 3
			}
		  }
		}
	  }
	}
}`


const TF_Endpoint = `resource "kubernetes_endpoints" "example" {
	metadata {
	  name = "terraform-example"
	}
  
	subset {
	  address {
		ip = "10.0.0.4"
	  }
  
	  address {
		ip = "10.0.0.5"
	  }
  
	  port {
		name     = "http"
		port     = 80
		protocol = "TCP"
	  }
  
	  port {
		name     = "https"
		port     = 443
		protocol = "TCP"
	  }
	}
  
	subset {
	  address {
		ip = "10.0.1.4"
	  }
  
	  address {
		ip = "10.0.1.5"
	  }
  
	  port {
		name     = "http"
		port     = 80
		protocol = "TCP"
	  }
  
	  port {
		name     = "https"
		port     = 443
		protocol = "TCP"
	  }
	}
}
  
  resource "kubernetes_service" "example" {
	metadata {
	  name = "${kubernetes_endpoints.example.metadata.0.name}"
	}
  
	spec {
	  port {
		port        = 8080
		target_port = 80
	  }
  
	  port {
		port        = 8443
		target_port = 443
	  }
	}
}`

const TF_HPA = `resource "kubernetes_horizontal_pod_autoscaler" "example" {
	metadata {
	  name = "terraform-example"
	}
	spec {
	  max_replicas = 10
	  min_replicas = 8
	  scale_target_ref {
		kind = "ReplicationController"
		name = "MyApp"
	  }
	}
}`


const TF_Pod = `resource "kubernetes_pod" "test" {
	metadata {
	  name = "terraform-example"
	}
  
	spec {
	  container {
		image = "nginx:1.7.9"
		name  = "example"
  
		env {
		  name  = "environment"
		  value = "test"
		}
  
		liveness_probe {
		  http_get {
			path = "/nginx_status"
			port = 80
  
			http_header {
			  name  = "X-Custom-Header"
			  value = "Awesome"
			}
		  }
  
		  initial_delay_seconds = 3
		  period_seconds        = 3
		}
	  }
  
	  dns_config {
		nameservers = ["1.1.1.1", "8.8.8.8", "9.9.9.9"]
		searches    = ["example.com"]
  
		option {
		  name  = "ndots"
		  value = 1
		}
  
		option {
		  name = "use-vc"
		}
	  }
  
	  dns_policy = "None"
	}
}`


const TF_PersistentVolume = `resource "kubernetes_persistent_volume" "example" {
	metadata {
	  name = "terraform-example"
	}
	spec {
	  capacity = {
		storage = "2Gi"
	  }
	  access_modes = ["ReadWriteMany"]
	  persistent_volume_source {
		vsphere_volume {
		  volume_path = "/absolute/path"
		}
	  }
	}
}`


const TF_PVC = `resource "kubernetes_persistent_volume_claim" "example" {
	metadata {
	  name = "exampleclaimname"
	}
	spec {
	  access_modes = ["ReadWriteMany"]
	  resources {
		requests = {
		  storage = "5Gi"
		}
	  }
	  volume_name = "${kubernetes_persistent_volume.example.metadata.0.name}"
	}
}
  
  resource "kubernetes_persistent_volume" "example" {
	metadata {
	  name = "examplevolumename"
	}
	spec {
	  capacity = {
		storage = "10Gi"
	  }
	  access_modes = ["ReadWriteMany"]
	  persistent_volume_source {
		gce_persistent_disk {
		  pd_name = "test-123"
		}
	  }
	}
}`

const TF_Rsq = `resource "kubernetes_resource_quota" "example" {
	metadata {
	  name = "terraform-example"
	}
	spec {
	  hard = {
		pods = 10
	  }
	  scopes = ["BestEffort"]
	}
}`

const TF_ConfigMap = `resource "kubernetes_config_map" "example" {
	metadata {
	  name = "my-config"
	}
  
	data = {
	  api_host             = "myhost:443"
	  db_host              = "dbhost:5432"
	  "my_config_file.yml" = "${file("${path.module}/my_config_file.yml")}"
	}
  
	binary_data = {
	  "my_payload.bin" = "${filebase64("${path.module}/my_payload.bin")}"
	}
}`


const TF_Role = `resource "kubernetes_role" "example" {
	metadata {
	  name = "terraform-example"
	  labels = {
		test = "MyRole"
	  }
	}
  
	rule {
	  api_groups     = [""]
	  resources      = ["pods"]
	  resource_names = ["foo"]
	  verbs          = ["get", "list", "watch"]
	}
	rule {
	  api_groups = ["apps"]
	  resources  = ["deployments"]
	  verbs      = ["get", "list"]
	}
}`


const TF_RB = `resource "kubernetes_role_binding" "example" {
	metadata {
	  name      = "terraform-example"
	  namespace = "default"
	}
	role_ref {
	  api_group = "rbac.authorization.k8s.io"
	  kind      = "Role"
	  name      = "admin"
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


const TF_STS = `resource "kubernetes_stateful_set" "prometheus" {
	metadata {
	  annotations = {
		SomeAnnotation = "foobar"
	  }
  
	  labels = {
		k8s-app                           = "prometheus"
		"kubernetes.io/cluster-service"   = "true"
		"addonmanager.kubernetes.io/mode" = "Reconcile"
		version                           = "v2.2.1"
	  }
  
	  name = "prometheus"
	}
  
	spec {
	  pod_management_policy  = "Parallel"
	  replicas               = 1
	  revision_history_limit = 5
  
	  selector {
		match_labels = {
		  k8s-app = "prometheus"
		}
	  }
  
	  service_name = "prometheus"
  
	  template {
		metadata {
		  labels = {
			k8s-app = "prometheus"
		  }
  
		  annotations = {}
		}
  
		spec {
		  service_account_name = "prometheus"
  
		  init_container {
			name              = "init-chown-data"
			image             = "busybox:latest"
			image_pull_policy = "IfNotPresent"
			command           = ["chown", "-R", "65534:65534", "/data"]
  
			volume_mount {
			  name       = "prometheus-data"
			  mount_path = "/data"
			  sub_path   = ""
			}
		  }
  
		  container {
			name              = "prometheus-server-configmap-reload"
			image             = "jimmidyson/configmap-reload:v0.1"
			image_pull_policy = "IfNotPresent"
  
			args = [
			  "--volume-dir=/etc/config",
			  "--webhook-url=http://localhost:9090/-/reload",
			]
  
			volume_mount {
			  name       = "config-volume"
			  mount_path = "/etc/config"
			  read_only  = true
			}
  
			resources {
			  limits {
				cpu    = "10m"
				memory = "10Mi"
			  }
  
			  requests {
				cpu    = "10m"
				memory = "10Mi"
			  }
			}
		  }
  
		  container {
			name              = "prometheus-server"
			image             = "prom/prometheus:v2.2.1"
			image_pull_policy = "IfNotPresent"
  
			args = [
			  "--config.file=/etc/config/prometheus.yml",
			  "--storage.tsdb.path=/data",
			  "--web.console.libraries=/etc/prometheus/console_libraries",
			  "--web.console.templates=/etc/prometheus/consoles",
			  "--web.enable-lifecycle",
			]
  
			port {
			  container_port = 9090
			}
  
			resources {
			  limits {
				cpu    = "200m"
				memory = "1000Mi"
			  }
  
			  requests {
				cpu    = "200m"
				memory = "1000Mi"
			  }
			}
  
			volume_mount {
			  name       = "config-volume"
			  mount_path = "/etc/config"
			}
  
			volume_mount {
			  name       = "prometheus-data"
			  mount_path = "/data"
			  sub_path   = ""
			}
  
			readiness_probe {
			  http_get {
				path = "/-/ready"
				port = 9090
			  }
  
			  initial_delay_seconds = 30
			  timeout_seconds       = 30
			}
  
			liveness_probe {
			  http_get {
				path   = "/-/healthy"
				port   = 9090
				scheme = "HTTPS"
			  }
  
			  initial_delay_seconds = 30
			  timeout_seconds       = 30
			}
		  }
  
		  termination_grace_period_seconds = 300
  
		  volume {
			name = "config-volume"
  
			config_map {
			  name = "prometheus-config"
			}
		  }
		}
	  }
  
	  update_strategy {
		type = "RollingUpdate"
  
		rolling_update {
		  partition = 1
		}
	  }
  
	  volume_claim_template {
		metadata {
		  name = "prometheus-data"
		}
  
		spec {
		  access_modes       = ["ReadWriteOnce"]
		  storage_class_name = "standard"
  
		  resources {
			requests = {
			  storage = "16Gi"
			}
		  }
		}
	  }
	}
}`


const TF_Secret = `resource "kubernetes_secret" "example" {
	metadata {
	  name = "basic-auth"
	}
  
	data = {
	  username = "admin"
	  password = "P4ssw0rd"
	}
  
	type = "kubernetes.io/basic-auth"
}`


const TF_SA = `resource "kubernetes_service_account" "example" {
	metadata {
	  name = "terraform-example"
	}
	secret {
	  name = "${kubernetes_secret.example.metadata.0.name}"
	}
  }
  
  resource "kubernetes_secret" "example" {
	metadata {
	  name = "terraform-example"
	}
}`


const TF_STC = `resource "kubernetes_storage_class" "example" {
	metadata {
	  name = "terraform-example"
	}
	storage_provisioner = "kubernetes.io/gce-pd"
	reclaim_policy      = "Retain"
	parameters = {
	  type = "pd-standard"
	}
}`


const TF_Npol = `resource "kubernetes_network_policy" "example" {
	metadata {
	  name      = "terraform-example-network-policy"
	  namespace = "default"
	}
  
	spec {
	  pod_selector {
		match_expressions {
		  key      = "name"
		  operator = "In"
		  values   = ["webfront", "api"]
		}
	  }
  
	  ingress {
		ports {
		  port     = "http"
		  protocol = "TCP"
		}
		ports {
		  port     = "8125"
		  protocol = "UDP"
		}
  
		from {
		  namespace_selector {
			match_labels = {
			  name = "default"
			}
		  }
		}
  
		from {
		  ip_block {
			cidr = "10.0.0.0/8"
			except = [
			  "10.0.0.0/24",
			  "10.0.1.0/24",
			]
		  }
		}
	  }
  
	  egress {} # single empty rule to allow all egress traffic
  
	  policy_types = ["Ingress", "Egress"]
	}
}`


const TF_APS = `resource "kubernetes_api_service" "example" {
	metadata {
	  name = "terraform-example"
	}
	spec {
	  selector {
		app = "${kubernetes_pod.example.metadata.0.labels.app}"
	  }
	  session_affinity = "ClientIP"
	  port {
		port = 8080
		target_port = 80
	  }
  
	  type = "LoadBalancer"
	}
}`

const TF_LRange = `resource "kubernetes_limit_range" "example" {
	metadata {
	  name = "terraform-example"
	}
	spec {
	  limit {
		type = "Pod"
		max = {
		  cpu    = "200m"
		  memory = "1024M"
		}
	  }
	  limit {
		type = "PersistentVolumeClaim"
		min = {
		  storage = "24M"
		}
	  }
	  limit {
		type = "Container"
		default = {
		  cpu    = "50m"
		  memory = "24M"
		}
	  }
	}
}`


const TF_ReplC = `resource "kubernetes_replication_controller" "example" {
	metadata {
	  name = "terraform-example"
	  labels = {
		test = "MyExampleApp"
	  }
	}
  
	spec {
	  selector = {
		test = "MyExampleApp"
	  }
	  template {
		metadata {
		  labels = {
			test = "MyExampleApp"
		  }
		  annotations = {
			"key1" = "value1"
		  }
		}
  
		spec {
		  container {
			image = "nginx:1.7.8"
			name  = "example"
  
			liveness_probe {
			  http_get {
				path = "/nginx_status"
				port = 8080
  
				http_header {
				  name  = "X-Custom-Header"
				  value = "Awesome"
				}
			  }
  
			  initial_delay_seconds = 3
			  period_seconds        = 3
			}
  
			resources {
			  limits {
				cpu    = "0.5"
				memory = "512Mi"
			  }
			  requests {
				cpu    = "250m"
				memory = "50Mi"
			  }
			}
		  }
		}
	  }
	}
}`
