# ipd2 Helm Chart

words words words words words

## Pre-requisites

For `HTTP/2` mode to work, you'll need to provide the chart a TLS certificate/key, you might want to use [cert-manager](https://github.com/jetstack/cert-manager/) to help with this.

You will then need to enable both the `https` service *AND* `https` ingress, since the secret is defined at the ingress level.
