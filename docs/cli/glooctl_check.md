---
title: "glooctl check"
weight: 5
---
## glooctl check

Checks for any configuration errors.

### Synopsis

Checks for any configuration errors.

```
glooctl check [flags]
```

### Options

```
  -A, --all-namespaces       check for resources in all namespaces
  -f, --file string          file to be read or written to
  -h, --help                 help for check
      --name string          name of the resource to read or write
  -n, --namespace string     namespace for reading or writing resources (default "gloo-system")
      --namespaces strings   check for resources in namespace list
```

### Options inherited from parent commands

```
  -i, --interactive   use interactive mode
```

### SEE ALSO

* [glooctl](../glooctl)	 - CLI for Gloo
* [glooctl check upstream](../glooctl_check_upstream)	 - Check upstreams for any configuration errors.

