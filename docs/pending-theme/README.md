This shortcode will be merged with https://github.com/solo-io/hugo-theme-soloio prior to merge

### protobuf shortcode

- usage:
  - choose a protobuf name
    - solo-kit generates the `data/ProtoMap.yaml` file which maps the protobuf name to the docs url.
  - shortcode generates a link to the docs for that proto
    -  Additional derivative links may be added later, such as yaml samples.

```
{{% protobuf name="cors.plugins.gloo.solo.io" %}}
```
