package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayWithSimplePods(t *testing.T) {
	var tr Traverse
	var tree, ruletree map[string]interface{}
	assert.Nil(t, json.Unmarshal([]byte(simplePods), &tree))
	assert.Nil(t, json.Unmarshal([]byte(simpleRule), &ruletree))
	assert.Len(t, tree, 2)
	tr.traverse("", tree, ruletree)
	assert.Nil(t, tr.lastError)
}

func TestPlayWithFullPods(t *testing.T) {
	var tr Traverse
	var tree, ruletree map[string]interface{}
	assert.Nil(t, json.Unmarshal([]byte(fullPods), &tree))
	assert.Nil(t, json.Unmarshal([]byte(simpleRule), &ruletree))
	assert.Len(t, tree, 4)
	tr.traverse("", tree, ruletree)
	assert.Nil(t, tr.lastError)
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

var simpleRule = `{
	"apiVersion": {
		"equals": "v1",
		"matches": "v"
	},
	"items": {
		"length": "1"
	}
}
`
var simplePods = `{
	"apiVersion": "v1",
	"items":[{
		"a":"1"
		}]
}`

var fullPods = `{
	"apiVersion": "v1",
	"items": [
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"creationTimestamp": "2018-02-24T14:10:10Z",
							"generateName": "heapster-",
							"labels": {
									"addonmanager.kubernetes.io/mode": "Reconcile",
									"k8s-app": "heapster"
							},
							"name": "heapster-7cltx",
							"namespace": "kube-system",
							"ownerReferences": [
									{
											"apiVersion": "v1",
											"blockOwnerDeletion": true,
											"controller": true,
											"kind": "ReplicationController",
											"name": "heapster",
											"uid": "6c9de11a-196c-11e8-bb9b-54ee75a8e55c"
									}
							],
							"resourceVersion": "225",
							"selfLink": "/api/v1/namespaces/kube-system/pods/heapster-7cltx",
							"uid": "6c9e1d3d-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"command": [
													"/heapster",
													"--source=kubernetes.summary_api:''",
													"--sink=influxdb:http://monitoring-influxdb:8086",
													"--metric_resolution=60s"
											],
											"image": "k8s.gcr.io/heapster-amd64:v1.5.0",
											"imagePullPolicy": "IfNotPresent",
											"name": "heapster",
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/etc/ssl/certs",
															"name": "ssl-certs",
															"readOnly": true
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"volumes": [
									{
											"hostPath": {
													"path": "/etc/ssl/certs",
													"type": ""
											},
											"name": "ssl-certs"
									},
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:15Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://487f678687b7b41b25ee3acc412576128e77ed167396458ebc833f355753c944",
											"image": "k8s.gcr.io/heapster-amd64:v1.5.0",
											"imageID": "docker-pullable://k8s.gcr.io/heapster-amd64@sha256:da3288b0fe2312c621c2a6d08f24ccc56183156ec70767987501287db4927b9d",
											"lastState": {},
											"name": "heapster",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:14Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "172.17.0.5",
							"qosClass": "BestEffort",
							"startTime": "2018-02-24T14:10:13Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"creationTimestamp": "2018-02-24T14:10:10Z",
							"generateName": "influxdb-grafana-",
							"labels": {
									"addonmanager.kubernetes.io/mode": "Reconcile",
									"k8s-app": "influx-grafana"
							},
							"name": "influxdb-grafana-9k8qj",
							"namespace": "kube-system",
							"ownerReferences": [
									{
											"apiVersion": "v1",
											"blockOwnerDeletion": true,
											"controller": true,
											"kind": "ReplicationController",
											"name": "influxdb-grafana",
											"uid": "6cb2c25c-196c-11e8-bb9b-54ee75a8e55c"
									}
							],
							"resourceVersion": "232",
							"selfLink": "/api/v1/namespaces/kube-system/pods/influxdb-grafana-9k8qj",
							"uid": "6cb2fe82-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"image": "k8s.gcr.io/heapster-influxdb-amd64:v1.3.3",
											"imagePullPolicy": "IfNotPresent",
											"name": "influxdb",
											"ports": [
													{
															"containerPort": 8083,
															"name": "http",
															"protocol": "TCP"
													},
													{
															"containerPort": 8086,
															"name": "api",
															"protocol": "TCP"
													}
											],
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/data",
															"name": "influxdb-storage"
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									},
									{
											"env": [
													{
															"name": "INFLUXDB_SERVICE_URL",
															"value": "http://localhost:8086"
													},
													{
															"name": "GF_AUTH_BASIC_ENABLED",
															"value": "false"
													},
													{
															"name": "GF_AUTH_ANONYMOUS_ENABLED",
															"value": "true"
													},
													{
															"name": "GF_AUTH_ANONYMOUS_ORG_ROLE",
															"value": "Admin"
													},
													{
															"name": "GF_SERVER_ROOT_URL",
															"value": "/"
													}
											],
											"image": "k8s.gcr.io/heapster-grafana-amd64:v4.4.3",
											"imagePullPolicy": "IfNotPresent",
											"name": "grafana",
											"ports": [
													{
															"containerPort": 3000,
															"name": "ui",
															"protocol": "TCP"
													}
											],
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/var",
															"name": "grafana-storage"
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"volumes": [
									{
											"emptyDir": {},
											"name": "influxdb-storage"
									},
									{
											"emptyDir": {},
											"name": "grafana-storage"
									},
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:15Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://ea4430368c8c2bb0720ef2c9caf3c3c888c228a190c6dd0e072542cb9fcc9bc3",
											"image": "k8s.gcr.io/heapster-grafana-amd64:v4.4.3",
											"imageID": "docker-pullable://k8s.gcr.io/heapster-grafana-amd64@sha256:4a472eb4df03f4f557d80e7c6b903d9c8fe31493108b99fbd6da6540b5448d70",
											"lastState": {},
											"name": "grafana",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:15Z"
													}
											}
									},
									{
											"containerID": "docker://0b89d1c5911adf65a90093b834f6f028785ac5f8442b2625c83b5dacb8319314",
											"image": "k8s.gcr.io/heapster-influxdb-amd64:v1.3.3",
											"imageID": "docker-pullable://k8s.gcr.io/heapster-influxdb-amd64@sha256:f433e331c1865ad87bc5387589965528b78cd6b1b2f61697e589584d690c1edd",
											"lastState": {},
											"name": "influxdb",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:15Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "172.17.0.6",
							"qosClass": "BestEffort",
							"startTime": "2018-02-24T14:10:13Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"annotations": {
									"kubernetes.io/config.hash": "c4c3188325a93a2d7fb1714e1abf1259",
									"kubernetes.io/config.mirror": "c4c3188325a93a2d7fb1714e1abf1259",
									"kubernetes.io/config.seen": "2018-02-24T15:10:01.998278415+01:00",
									"kubernetes.io/config.source": "file"
							},
							"creationTimestamp": "2018-02-24T14:10:07Z",
							"labels": {
									"component": "kube-addon-manager",
									"kubernetes.io/minikube-addons": "addon-manager",
									"version": "v6.5"
							},
							"name": "kube-addon-manager-sirius",
							"namespace": "kube-system",
							"resourceVersion": "85",
							"selfLink": "/api/v1/namespaces/kube-system/pods/kube-addon-manager-sirius",
							"uid": "6a9cada1-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"env": [
													{
															"name": "KUBECONFIG",
															"value": "/var/lib/localkube/kubeconfig"
													}
											],
											"image": "gcr.io/google-containers/kube-addon-manager:v6.5",
											"imagePullPolicy": "IfNotPresent",
											"name": "kube-addon-manager",
											"resources": {
													"requests": {
															"cpu": "5m",
															"memory": "50Mi"
													}
											},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/etc/kubernetes/",
															"name": "addons",
															"readOnly": true
													},
													{
															"mountPath": "/var/lib/localkube/",
															"name": "kubeconfig",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"hostNetwork": true,
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"terminationGracePeriodSeconds": 30,
							"tolerations": [
									{
											"effect": "NoExecute",
											"operator": "Exists"
									}
							],
							"volumes": [
									{
											"hostPath": {
													"path": "/etc/kubernetes/",
													"type": ""
											},
											"name": "addons"
									},
									{
											"hostPath": {
													"path": "/var/lib/localkube/",
													"type": ""
											},
											"name": "kubeconfig"
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:07Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:08Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:07Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://2bd1342b5ffc0653923584a2ac733dbde60a4dbc4d800a26ee1ede8ba1ed2136",
											"image": "gcr.io/google-containers/kube-addon-manager:v6.5",
											"imageID": "docker://sha256:d166ffa9201aa156eb76d3a221c3fdab07bb1a0b6407548b1b1f03dc111c0e39",
											"lastState": {
													"terminated": {
															"containerID": "docker://b929a4e9382243f4abe448303a81d69bd5c5b983bc7e689f67ec0dd8636c6502",
															"exitCode": 137,
															"finishedAt": "2018-02-22T13:18:23Z",
															"reason": "Error",
															"startedAt": "2018-02-22T13:17:40Z"
													}
											},
											"name": "kube-addon-manager",
											"ready": true,
											"restartCount": 5,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:08Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "192.168.2.15",
							"qosClass": "Burstable",
							"startTime": "2018-02-24T14:10:07Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"annotations": {
									"scheduler.alpha.kubernetes.io/critical-pod": ""
							},
							"creationTimestamp": "2018-02-24T14:10:10Z",
							"generateName": "kube-dns-54cccfbdf8-",
							"labels": {
									"k8s-app": "kube-dns",
									"pod-template-hash": "1077796894"
							},
							"name": "kube-dns-54cccfbdf8-nrzmp",
							"namespace": "kube-system",
							"ownerReferences": [
									{
											"apiVersion": "extensions/v1beta1",
											"blockOwnerDeletion": true,
											"controller": true,
											"kind": "ReplicaSet",
											"name": "kube-dns-54cccfbdf8",
											"uid": "6cc9a050-196c-11e8-bb9b-54ee75a8e55c"
									}
							],
							"resourceVersion": "10043",
							"selfLink": "/api/v1/namespaces/kube-system/pods/kube-dns-54cccfbdf8-nrzmp",
							"uid": "6ccb3a27-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"args": [
													"--domain=cluster.local.",
													"--dns-port=10053",
													"--config-map=kube-dns",
													"--v=2"
											],
											"env": [
													{
															"name": "PROMETHEUS_PORT",
															"value": "10055"
													}
											],
											"image": "k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.5",
											"imagePullPolicy": "IfNotPresent",
											"livenessProbe": {
													"failureThreshold": 5,
													"httpGet": {
															"path": "/healthcheck/kubedns",
															"port": 10054,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 60,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 5
											},
											"name": "kubedns",
											"ports": [
													{
															"containerPort": 10053,
															"name": "dns-local",
															"protocol": "UDP"
													},
													{
															"containerPort": 10053,
															"name": "dns-tcp-local",
															"protocol": "TCP"
													},
													{
															"containerPort": 10055,
															"name": "metrics",
															"protocol": "TCP"
													}
											],
											"readinessProbe": {
													"failureThreshold": 3,
													"httpGet": {
															"path": "/readiness",
															"port": 8081,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 3,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 5
											},
											"resources": {
													"limits": {
															"memory": "170Mi"
													},
													"requests": {
															"cpu": "100m",
															"memory": "70Mi"
													}
											},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/kube-dns-config",
															"name": "kube-dns-config"
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									},
									{
											"args": [
													"-v=2",
													"-logtostderr",
													"-configDir=/etc/k8s/dns/dnsmasq-nanny",
													"-restartDnsmasq=true",
													"--",
													"-k",
													"--cache-size=1000",
													"--log-facility=-",
													"--server=/cluster.local/127.0.0.1#10053",
													"--server=/in-addr.arpa/127.0.0.1#10053",
													"--server=/ip6.arpa/127.0.0.1#10053"
											],
											"image": "k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.5",
											"imagePullPolicy": "IfNotPresent",
											"livenessProbe": {
													"failureThreshold": 5,
													"httpGet": {
															"path": "/healthcheck/dnsmasq",
															"port": 10054,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 60,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 5
											},
											"name": "dnsmasq",
											"ports": [
													{
															"containerPort": 53,
															"name": "dns",
															"protocol": "UDP"
													},
													{
															"containerPort": 53,
															"name": "dns-tcp",
															"protocol": "TCP"
													}
											],
											"resources": {
													"requests": {
															"cpu": "150m",
															"memory": "20Mi"
													}
											},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/etc/k8s/dns/dnsmasq-nanny",
															"name": "kube-dns-config"
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									},
									{
											"args": [
													"--v=2",
													"--logtostderr",
													"--probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.cluster.local.,5,A",
													"--probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.cluster.local.,5,A"
											],
											"image": "k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.5",
											"imagePullPolicy": "IfNotPresent",
											"livenessProbe": {
													"failureThreshold": 5,
													"httpGet": {
															"path": "/metrics",
															"port": 10054,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 60,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 5
											},
											"name": "sidecar",
											"ports": [
													{
															"containerPort": 10054,
															"name": "metrics",
															"protocol": "TCP"
													}
											],
											"resources": {
													"requests": {
															"cpu": "10m",
															"memory": "20Mi"
													}
											},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "Default",
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"tolerations": [
									{
											"key": "CriticalAddonsOnly",
											"operator": "Exists"
									}
							],
							"volumes": [
									{
											"configMap": {
													"defaultMode": 420,
													"name": "kube-dns",
													"optional": true
											},
											"name": "kube-dns-config"
									},
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:20Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://01f436aa233a503bb05f200c41c6a7a157cfc61d191685817e379b75807d9b50",
											"image": "k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.5",
											"imageID": "docker://sha256:459944ce8cc4f08ebade5c05bb884e4da053d73e61ec6afe82a0b1687317254c",
											"lastState": {
													"terminated": {
															"containerID": "docker://cfdaf7a3ce2c284275d607484f408c786ea0a2f6a8dce50f54b26412cf3ec262",
															"exitCode": 137,
															"finishedAt": "2018-02-24T18:21:40Z",
															"reason": "Error",
															"startedAt": "2018-02-24T18:11:41Z"
													}
											},
											"name": "dnsmasq",
											"ready": true,
											"restartCount": 25,
											"state": {
													"running": {
															"startedAt": "2018-02-24T18:21:40Z"
													}
											}
									},
									{
											"containerID": "docker://c049bd90d3a758cd771103d2c2b1c8033dfe2a0f5f49a19a58d1c7d13166a74c",
											"image": "k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.5",
											"imageID": "docker-pullable://k8s.gcr.io/k8s-dns-kube-dns-amd64@sha256:1a3fc069de481ae690188f6f1ba4664b5cc7760af37120f70c86505c79eea61d",
											"lastState": {},
											"name": "kubedns",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:15Z"
													}
											}
									},
									{
											"containerID": "docker://286bef7ed81d9d38964642342a09cf478c307918d1c0f366795aace2469f6a02",
											"image": "k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.5",
											"imageID": "docker://sha256:fed89e8b4248a788655d528d96fe644aff012879c782784cd486ff6894ef89f6",
											"lastState": {},
											"name": "sidecar",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:16Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "172.17.0.7",
							"qosClass": "Burstable",
							"startTime": "2018-02-24T14:10:13Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"creationTimestamp": "2018-02-24T14:10:10Z",
							"generateName": "kubernetes-dashboard-77d8b98585-",
							"labels": {
									"addonmanager.kubernetes.io/mode": "Reconcile",
									"app": "kubernetes-dashboard",
									"pod-template-hash": "3384654141",
									"version": "v1.8.1"
							},
							"name": "kubernetes-dashboard-77d8b98585-k8j6z",
							"namespace": "kube-system",
							"ownerReferences": [
									{
											"apiVersion": "extensions/v1beta1",
											"blockOwnerDeletion": true,
											"controller": true,
											"kind": "ReplicaSet",
											"name": "kubernetes-dashboard-77d8b98585",
											"uid": "6c772d69-196c-11e8-bb9b-54ee75a8e55c"
									}
							],
							"resourceVersion": "219",
							"selfLink": "/api/v1/namespaces/kube-system/pods/kubernetes-dashboard-77d8b98585-k8j6z",
							"uid": "6c787399-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"image": "k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.1",
											"imagePullPolicy": "IfNotPresent",
											"livenessProbe": {
													"failureThreshold": 3,
													"httpGet": {
															"path": "/",
															"port": 9090,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 30,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 30
											},
											"name": "kubernetes-dashboard",
											"ports": [
													{
															"containerPort": 9090,
															"protocol": "TCP"
													}
											],
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"volumes": [
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:15Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://776bad242d8c06eaa3707a02dfd14efb357962c84638f1334afc3c085755f0b9",
											"image": "k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.1",
											"imageID": "docker-pullable://k8s.gcr.io/kubernetes-dashboard-amd64@sha256:3861695e962972965a4c611bcabc2032f885d8cbdb0bccc9bf513ef16335fe33",
											"lastState": {},
											"name": "kubernetes-dashboard",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:14Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "172.17.0.3",
							"qosClass": "BestEffort",
							"startTime": "2018-02-24T14:10:13Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"creationTimestamp": "2018-02-24T14:10:10Z",
							"generateName": "kubernetes-dashboard-",
							"labels": {
									"addonmanager.kubernetes.io/mode": "Reconcile",
									"app": "kubernetes-dashboard",
									"version": "v1.8.0"
							},
							"name": "kubernetes-dashboard-xk57z",
							"namespace": "kube-system",
							"ownerReferences": [
									{
											"apiVersion": "v1",
											"blockOwnerDeletion": true,
											"controller": true,
											"kind": "ReplicationController",
											"name": "kubernetes-dashboard",
											"uid": "6c7eea82-196c-11e8-bb9b-54ee75a8e55c"
									}
							],
							"resourceVersion": "216",
							"selfLink": "/api/v1/namespaces/kube-system/pods/kubernetes-dashboard-xk57z",
							"uid": "6c8215ba-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"image": "gcr.io/google_containers/kubernetes-dashboard-amd64:v1.8.0",
											"imagePullPolicy": "IfNotPresent",
											"livenessProbe": {
													"failureThreshold": 3,
													"httpGet": {
															"path": "/",
															"port": 9090,
															"scheme": "HTTP"
													},
													"initialDelaySeconds": 30,
													"periodSeconds": 10,
													"successThreshold": 1,
													"timeoutSeconds": 30
											},
											"name": "kubernetes-dashboard",
											"ports": [
													{
															"containerPort": 9090,
															"protocol": "TCP"
													}
											],
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"volumes": [
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:15Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:13Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://e1b58a2f535ac6cb38d4da6703b87d9e9d965d22143f70f6a500612c7f51a633",
											"image": "gcr.io/google_containers/kubernetes-dashboard-amd64:v1.8.0",
											"imageID": "docker-pullable://gcr.io/google_containers/kubernetes-dashboard-amd64@sha256:71a0de5c6a21cb0c2fbcad71a4fef47acd3e61cd78109822d35e1742f9d8140d",
											"lastState": {},
											"name": "kubernetes-dashboard",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:14Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "172.17.0.4",
							"qosClass": "BestEffort",
							"startTime": "2018-02-24T14:10:13Z"
					}
			},
			{
					"apiVersion": "v1",
					"kind": "Pod",
					"metadata": {
							"creationTimestamp": "2018-02-24T14:10:09Z",
							"labels": {
									"addonmanager.kubernetes.io/mode": "EnsureExists",
									"integration-test": "storage-provisioner"
							},
							"name": "storage-provisioner",
							"namespace": "kube-system",
							"resourceVersion": "208",
							"selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
							"uid": "6c184dc7-196c-11e8-bb9b-54ee75a8e55c"
					},
					"spec": {
							"containers": [
									{
											"command": [
													"/storage-provisioner"
											],
											"image": "gcr.io/k8s-minikube/storage-provisioner:v1.8.1",
											"imagePullPolicy": "IfNotPresent",
											"name": "storage-provisioner",
											"resources": {},
											"terminationMessagePath": "/dev/termination-log",
											"terminationMessagePolicy": "File",
											"volumeMounts": [
													{
															"mountPath": "/tmp",
															"name": "tmp"
													},
													{
															"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
															"name": "default-token-wdxzb",
															"readOnly": true
													}
											]
									}
							],
							"dnsPolicy": "ClusterFirst",
							"hostNetwork": true,
							"nodeName": "sirius",
							"restartPolicy": "Always",
							"schedulerName": "default-scheduler",
							"securityContext": {},
							"serviceAccount": "default",
							"serviceAccountName": "default",
							"terminationGracePeriodSeconds": 30,
							"volumes": [
									{
											"hostPath": {
													"path": "/tmp",
													"type": "Directory"
											},
											"name": "tmp"
									},
									{
											"name": "default-token-wdxzb",
											"secret": {
													"defaultMode": 420,
													"secretName": "default-token-wdxzb"
											}
									}
							]
					},
					"status": {
							"conditions": [
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:12Z",
											"status": "True",
											"type": "Initialized"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:14Z",
											"status": "True",
											"type": "Ready"
									},
									{
											"lastProbeTime": null,
											"lastTransitionTime": "2018-02-24T14:10:12Z",
											"status": "True",
											"type": "PodScheduled"
									}
							],
							"containerStatuses": [
									{
											"containerID": "docker://cf9b1865b48df28afb20a1d5200ae882b397d8042635c065c459af20c0e7cde3",
											"image": "gcr.io/k8s-minikube/storage-provisioner:v1.8.0",
											"imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:088daa9fcbccf04c3f415d77d5a6360d2803922190b675cb7fc88a9d2d91985a",
											"lastState": {},
											"name": "storage-provisioner",
											"ready": true,
											"restartCount": 0,
											"state": {
													"running": {
															"startedAt": "2018-02-24T14:10:13Z"
													}
											}
									}
							],
							"hostIP": "192.168.2.15",
							"phase": "Running",
							"podIP": "192.168.2.15",
							"qosClass": "BestEffort",
							"startTime": "2018-02-24T14:10:12Z"
					}
			}
	],
	"kind": "List",
	"metadata": {
			"resourceVersion": "",
			"selfLink": ""
	}
}

`
