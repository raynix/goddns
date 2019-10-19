# goddns
Use public IP service to update DNS records so this is kind of a DDNS self service.

For now the DNS is assumed hosted in CloudFlare.

## Deploy to Kubernetes cluster
Assuming the `kubectl` setup is done already
```
$ kubectl apply -k k8s
$ echo -n "<cloudflare api key>" > /tmp/CF_API_KEY
$ echo -n "<CloudFlare email>" > /tmp/CF_API_EMAIL
$ echo -n "my.com,you.com,other.com" > /tmp/GODDNS_DOMAINS
$ kubectl create secret generic goddns  -n goddns --from-file=/tmp/CF_API_KEY --from-file=/tmp/CF_API_EMAIL --from-file=/tmp/GODDNS_DOMAINS
```


## Run in container
```
$ docker run --rm -e CF_API_KEY=xxx -e CF_API_EMAIL=xxx -e GODDNS_DOMAINS=xxx raynix/goddns:v1.0.0
```
