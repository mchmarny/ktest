# tellmeall

Simple test app for Knative deployments

## Use

To use the `tellmeall` application you will have to deploy it to an existing [Knative](https://github.com/knative) cluster. See [Knative Installation](https://github.com/knative/docs/tree/master/install) if you don't have one.

### Deploy

```shell
kubectl apply -f https://raw.githubusercontent.com/mchmarny/tellmeall/master/app.yaml
```

### Endpoints

The `tellmeall` application exposes a few endpoints

* `/` responds with simple `OK`
* `/env` responds with all environment variables in a key/value format
* `/head` responds with all request header variables in a key/value format
* `/mem` responds with total, used and free system memory information
* `/host` responds with container info (ID, Hostname, OS, Boot-time etc.)
* `/log` responds with content of specific log (e.g. /log?logpath=/var/log/app.log)
* `/kn` responds with Knative-specific data as defined in the [Runtime Contract](https://github.com/knative/serving/blob/master/docs/runtime-contract.md)
* `/help` responds with this list of endpoints as URLs



