# Knative Eventing Sources

[![GoDoc](https://godoc.org/github.com/knative/eventing-sources?status.svg)](https://godoc.org/github.com/knative/eventing-sources)
[![Go Report Card](https://goreportcard.com/badge/knative/eventing-sources)](https://goreportcard.com/report/knative/eventing-sources)

Knative Eventing Sources builds on Kubernetes and Knative to provide useful
generic strategies for developing a source to deliver messaging events to an
[Addressable](https://github.com/knative/eventing/tree/master/docs/spec/interfaces.md#addressable)
target.

The Knative Eventing Sources project provides source implementations that:

- Run your code in a container
- Integrate with GitHub
- Integrate with Kubernetes Events
- Integrate with Pub/Sub
- Expose an ingress

Also see the corresponding [knative/eventing](../eventing) repository where the
Eventing API spec is defined.

For complete documentation about Knative Eventing, see
[Knative Eventing](https://github.com/knative/docs/tree/master/eventing) or
[Knative Overview](https://github.com/knative/docs) to learn about Knative.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).
