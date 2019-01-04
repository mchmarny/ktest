# tellmeall

Simple test app for Knative deployments

> To use the `tellmeall` application you will have to deploy it to an existing [Knative](https://github.com/knative) cluster. See [Knative Installation](https://github.com/knative/docs/tree/master/install) if you don't have one.

## Demo

https://tellmeall.default.knative.tech/

## Deploy

```shell
kubectl apply -f https://raw.githubusercontent.com/mchmarny/tellmeall/master/app.yaml
```

After deployment, the `tellmeall` application will expose the following endpoints:

* `/` Landing page with links to this endpoints
* `/kn` Knative-specific data as defined in the [Runtime Contract](https://github.com/knative/serving/blob/master/docs/runtime-contract.md)
* `/env` All environment variables in a key/value format
* `/head` All request header variables in a key/value format
* `/host` Serving node info (ID, Hostname, OS, Boot-time etc.)
* `/mem` Total, used and free system memory information
* `/log` Content of specific log or log list in dir (e.g. /log?logpath=/var/log/app.log)
* `/_health` Responds with 'OK' (ala health check)



