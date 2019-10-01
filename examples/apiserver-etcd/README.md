# Running kube-apiserver and etcd locally

- Download both etcd and kube-apiserver:
```
os=$(go env GOOS)
arch=$(go env GOARCH)
curl -sL https://go.kubebuilder.io/dl/2.0.1/${os}/${arch} | tar -xz -C /tmp/
```

- Commands are run from /tmp (will be omitted)

- Run:
```
ETCDCTL_API=3 ./etcd & ./kube-apiserver --cert-dir . --etcd-servers=http://127.0.0.1:2379 
```

- Create configuration for kubectl 
```
kubectl config --kubeconfig=config-demo set-cluster scratch --server=https://127.0.0.1 --insecure-skip-tls-verify
```
Your configuration is stored in config-demo file.

- Execute a kubectl command:
```kubectl --kubeconfig=config-demo cluster-info```

- Create default service account:
```
kubectl --kubeconfig=config-demo create -f serviceaccount.yaml
```


# Etcdctl

- If you want to query etcd, download the etcdctl CLI:
 
https://github.com/etcd-io/etcd/releases

- List all keys:
```ETCDCTL_API=3 etcdctl --endpoints=localhost:2379 get / --prefix --keys-only```


