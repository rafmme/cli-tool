package tmpl

import "fmt"

const Namespace = `apiVersion: v1
kind: Namespace
metadata:
  annotations:
    "name": name
  name: name
  labels:
    app: name
`

const Secret = `apiVersion: v1
kind: Secret
metadata:
  name:  secretName
data:
  secretKey:  BASE64_ENCODED_VALUE
type: Opaque`


const ServiceAccount = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: name
  labels:
    app: app-name`


const Deployment = `apiVersion: extensions/v1beta1
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


const Kustomize = `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: name
commonLabels:
  app.kubernetes.io/name: name
  app.kubernetes.io/part-of: name
resources:
- resource_name
images:
- name: image/name
  newTag: 0.0.1
vars:
- fieldref:
    fieldPath: metadata.name
  name: NAME
  objref:`


const Service = `apiVersion: v1
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

const ConfigMap = `apiVersion: v1
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

const StatefulSet = `apiVersion: apps/v1beta1
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

const Pod = `apiVersion: v1
kind: Pod
metadata:
  name: name
  labels:
    app: app
spec:
  containers:
    - name: name
      image: image/name
      ports:
        - containerPort: 80
    - name: name
      image: image/name
      env:
      - name:  ENV_VAR_NAME
        value:  ENV_VAR_VALUE       
      ports:
      - containerPort:  5000
  restartPolicy: Always
  imagePullPolicy: Always`


const Ingress = `apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: name
  labels:
    app: app-label
  annotations:
    ingress.kubernetes.io/backends: ""
spec:
  rules:
    - host: host
      http:
        paths:
          - backend:
              serviceName: serviceName
              servicePort: servicePort
  tls:
  - hosts:
    - host.com
    - host-api.com
    secretName: secrets-name
`


const CronJob = `apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: name
  namespace: name
spec:
  jobTemplate:
    metadata:
    spec:
      template:
        metadata:
        spec:
          containers:
            - name: name
              image: image/name
              command: ['sh', '-c']
              args:
                - |
                  echo "hello"
              volumeMounts:
                - name: vol-mnts
                  mountPath: /vol
                  readOnly: true
          restartPolicy: OnFailure
          schedulerName: default-scheduler
          securityContext: {}
          terminationGracePeriodSeconds: 30
          volumes:
            - name: vol-mnts
              secret:
                defaultMode: 420
                secretName: vol-mnts
  schedule: '0 */23 * * *'
  successfulJobsHistoryLimit: 1
`


const Job = `apiVersion: batch/v1
kind: Job
metadata:
  name: name
spec:
  backoffLimit: 5
  activeDeadlineSeconds: 100
  template:
    metadata:
      name: name
    spec:
      containers:
      - name: name
        image: image/name
        command: ["echo",  "hello"]
      restartPolicy: Never
`

const PersistentVolume = `apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-volume
  labels:
    type: local
spec:
  storageClassName: standard
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data"
`

const PersistentVolumeClaim = `apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pv-claim
spec:
  storageClassName: standard
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
`

const Policy = `apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: name
  labels:
    app: app
spec:
  minAvailable: 1
  maxUnavailable: 2
  selector:
    matchLabels:
      app: app
`

const ClusterRole = `apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: name
  labels:
    app: app
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/proxy
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
`

const Role = `apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: name
  labels:
    app: app
rules:
- apiGroups:
  - extensions
  resources:
  - podsecuritypolicies
  resourceNames:
  - privileged
  verbs:
  - use`


const DaemonSet = `apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: name
  namespace: name
  labels:
    app: name
spec:
  selector:
    matchLabels:
      name: name
  template:
    metadata:
      labels:
        name: name
    spec:
      tolerations:
      - key: node-role.kubernetes.io/master
        effect: NoSchedule
      containers:
      - name: name
        image: gcr.io/image/name
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
        - name: varlibdockercontainers
          mountPath: /var/lib/docker/containers
          readOnly: true
      terminationGracePeriodSeconds: 30
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
`
  

func GenerateRoleBinding(roleBindingType, roleType string) string  {
  return fmt.Sprintf(`apiVersion: rbac.authorization.k8s.io/v1beta1
  kind: %s
  metadata:
    name: name
    labels:
      app: app
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: %s
    name: role-name
  subjects:
  - kind: ServiceAccount
    name: sa-name
  `, roleBindingType, roleType)
}
