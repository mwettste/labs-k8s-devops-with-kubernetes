# labs-k8s-devops-with-kubernetes

A learning lab where I learn about Kubernetes. Guided by the online course at [University of Helsinki](https://devopswithkubernetes.com/part-0).

I decided to do the course in minikube rather than k3d. To also get some practice with `go`, I decided to use that rather than `C#`.

## Learnings
* `fmt.Scanln()` should not be used in a go program that should keep running, since it will automatically complete when started in a container without `-it` flag
    * use a channel to react to SIGINT instead