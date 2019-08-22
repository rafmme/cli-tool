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
