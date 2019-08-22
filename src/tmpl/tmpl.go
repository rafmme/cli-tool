package tmpl

var Namespace string = `apiVersion: v1
kind: Namespace
metadata:
  annotations:
    "name": name
  name: name
  labels:
    app: name
`

var Secret string = `apiVersion: v1
kind: Secret
metadata:
  name:  secretName
data:
  secretKey:  BASE64_ENCODED_VALUE
type: Opaque`


var ServiceAccount string = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: name
  labels:
    app: app-name`


var Deployment string = `apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name:  name
  labels:
    name:  name
spec:
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        name:  name
    spec:
      containers:
      - image:  docker/name
        name:  name
        resources:
          requests:
            cpu: "20m"
            memory: "55M"
        livenessProbe:
          httpGet:
            path: /_status/healthz
            port: 5000
          initialDelaySeconds: 90
          timeoutSeconds: 10
        readinessProbe:
          httpGet:
            path: /_status/healthz
            port: 5000
          initialDelaySeconds: 30
          timeoutSeconds: 10
        env:
        - name:  ENVVARNAME
          value:  ENVVARVALUE       
        ports:
        - containerPort:  8000
          name:  my-name
        volumeMounts:
        - mountPath: /data
          name: data
      volumes:
        - name: data
          emptyDir: {}
      restartPolicy: Always
      imagePullPolicy: Always`


var Kustomize string = `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
- resource_name

namespace: name`


var Service string = `apiVersion: v1
kind: Service
metadata:
  name:  Service Name
spec:
  selector:
    app:  Selector Label
  type:  LoadBalancer | ClusterIP | NodePort
  ports:
  - name:  name-of-the-port
    port:  80
    targetPort:  8080
`

var ConfigMap = `apiVersion: v1
kind: ConfigMap
metadata:
  name: app-configmap
  labels:
    app: name
data:
  my-key: my-value
 my-properties-file.yaml: |
    # yaml file
    key: value
`

var StatefulSet string = `apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: app-name
  labels:
    app: app
spec:
  serviceName: "serviceName"
  selector:
    matchLabels:
      app: app
  replicas: 1
  template:
    metadata:
      labels:
        app: app
    spec:
      containers:
      - name: name
        image: image/name
        ports:
        - containerPort: 80
          name: name
        volumeMounts:
        - name: vol-mnt
          mountPath: /usr/
  volumeClaimTemplates:
  - metadata:
      name: vol-mnt
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
      storageClassName: disk
`
