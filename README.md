# go-conception

> Container + Inception = Conception
>
> I'm not good with names... and the verb 'conceive' fits nicely

_Work in progress. Use at your own discretion._

A Go package for running one-off containerized applications through Docker or Kubernetes API (coming soon).


## Rationale

This package was created with the intent of easing the programmatic execution of commands in containerized applications, and subsequently, capturing the resulting output. Think `docker run --rm <image> <command>`, but without depending on having Docker CLI available on the executing machine.

But why?

It may for example, be useful for launching containerized applications within a containerized application. In the aforementioned Docker example, we simply need to bind-mount `/var/run/docker.sock` to the parent containerized application, and this will allow it to spawn containers on the same level as the parent. If the child containerized application is using this package, neither the child nor the parent image will require Docker to be installed.

Nevertheless, this package should work fine should you decide to go down the _Docker-in-Docker (DinD)_ route instead, and you can benefit from having the package manage the full lifecycle of a one-off containerized application. That is, it will create, attach, start, and clean up a container with a given image name, command, stdout writer, and stderr writer.

The plan is to introduce an _adapter_ for Kubernetes in the future that will allow a container in a Pod to spawn a one-off job or Pod in a similar fashion.
