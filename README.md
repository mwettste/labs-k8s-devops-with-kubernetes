# labs-k8s-devops-with-kubernetes

A learning lab where I learn about Kubernetes. Guided by the online course at [University of Helsinki](https://devopswithkubernetes.com/part-0).

I decided to do the course in minikube rather than k3d. To also get some practice with `go`, I decided to use that rather than `C#`.

## Learnings
* `fmt.Scanln()` should not be used in a go program that should keep running, since it will automatically complete when started in a container without `-it` flag
    * use a channel to react to SIGINT instead
* first containerized go app ran without issue, but the second one said "./app not found" even though the executable was obviously there
    * this error can pop up if cgo is used to build the application. Depending on the packages used, glibc can be dynamically linked which is e.g. not available on Alpine Linux (which uses musl libc)
* when containerizing a `go` application therefore consider setting `ENV CGO_ENABLED=0`, this enables static compilation and can be important when e.g. running the application on an Alpine Linux
* i could not get ingress to work with minikube despite activating the addons ingress and ingress and ingress-dns (followed [this guide](https://minikube.sigs.k8s.io/docs/handbook/addons/ingress-dns/))