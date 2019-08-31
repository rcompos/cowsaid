Cowsaid
-------

Cowsaid is a Cowsay-As-A-Service applicaion for [Kubernetes][k8s-home]
clusters, providing accurate fortune telling.

TL;DR;
------

```console
$ helm install -n cowsaid --namespace cowsaid ./cowsaid-helm
```

Introduction
------------

This chart bootstraps a [Cowsaid][cowsaid-home] installation on
a [Kubernetes][k8s-home] cluster using the [Helm][helm-home] package manager.
This chart provides an simple service deployment for validation.

Prerequisites
-------------

-  Kubernetes 1.10+

Installing the Chart
--------------------

The chart can be installed as follows:

```console
$ helm install --name cowsaid stable/cowsaid
```

The command deploys Cowsaid on the Kubernetes cluster.

Uninstalling the Chart
----------------------

To uninstall/delete the `cowsaid` deployment:

```console
$ helm delete cowsaid
```

The command removes all the Kubernetes components associated with the
chart, but will not remove the release metadata from `helm` — this will prevent
you, for example, if you later try to create a release also named `cowsaid`). To
fully delete the release and release history, simply [include the `--purge`
flag][helm-usage]:

```console
$ helm delete --purge cowsaid
```


[helm-home]: https://helm.sh
[helm-usage]: https://docs.helm.sh/using_helm/
[k8s-home]: https://kubernetes.io
[cowsaid-home]: https://github.com/rcompos/cowsaid
