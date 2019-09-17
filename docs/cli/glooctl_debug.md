---
title: "glooctl debug"
weight: 5
---
## glooctl debug

Debug a Gloo resource (requires Gloo running on Kubernetes)

### Synopsis

Debug a Gloo resource (requires Gloo running on Kubernetes)

```
glooctl debug [flags]
```

### Options

```
      --errors-only        filter for error logs only
  -f, --file string        file to be read or written to
  -h, --help               help for debug
  -n, --namespace string   namespace for reading or writing resources (default "gloo-system")
      --zip                save logs to a tar file (specify location with -f)
```

### Options inherited from parent commands

```
  -i, --interactive         use interactive mode
      --kubeconfig string   kubeconfig to use, if not standard one
```

### SEE ALSO

* [glooctl](../glooctl)	 - CLI for Gloo

