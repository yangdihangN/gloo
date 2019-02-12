
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto:


**Types:**


- :ref:`gloo.solo.io.Upstream` **Top-Level Resource**
- :ref:`gloo.solo.io.DiscoveryMetadata`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/upstream.proto>`_





.. _gloo.solo.io.Upstream:

Upstream
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
[clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster.html?highlight=cluster) in Envoy terminology.
Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
Each upstream type is handled by a corresponding Gloo plugin.


::


   "upstream_spec": .gloo.solo.io.UpstreamSpec
   "status": .core.solo.io.Status
   "metadata": .core.solo.io.Metadata
   "discovery_metadata": .gloo.solo.io.DiscoveryMetadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `upstream_spec` | :ref:`gloo.solo.io.UpstreamSpec` | Type-specific configuration. Examples include static, kubernetes, and aws. The type-specific config for the upstream is called a spec. | 
   `status` | :ref:`core.solo.io.Status` | Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 
   `discovery_metadata` | :ref:`gloo.solo.io.DiscoveryMetadata` | Upstreams and their configuration can be automatically by Gloo Discovery if this upstream is created or modified by Discovery, metadata about the operation will be placed here. | 



.. _gloo.solo.io.DiscoveryMetadata:

DiscoveryMetadata
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
created by discovery services


::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |






.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
