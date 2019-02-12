
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto:


**Types:**


- :ref:`gloo.solo.io.ListenerPlugins`
- :ref:`gloo.solo.io.VirtualHostPlugins`
- :ref:`gloo.solo.io.RoutePlugins`
- :ref:`gloo.solo.io.DestinationSpec`
- :ref:`gloo.solo.io.UpstreamSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins.proto>`_





.. _gloo.solo.io.ListenerPlugins:

ListenerPlugins
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Plugin-specific configuration that lives on listeners
Each ListenerPlugin object contains configuration for a specific plugin
Note to developers: new Listener Plugins must be added to this struct
to be usable by Gloo.


::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _gloo.solo.io.VirtualHostPlugins:

VirtualHostPlugins
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Plugin-specific configuration that lives on virtual hosts
Each VirtualHostPlugin object contains configuration for a specific plugin
Note to developers: new Virtual Host Plugins must be added to this struct
to be usable by Gloo.


::


   "extensions": .gloo.solo.io.Extensions

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `extensions` | :ref:`gloo.solo.io.Extensions` |  | 



.. _gloo.solo.io.RoutePlugins:

RoutePlugins
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Plugin-specific configuration that lives on routes
Each RoutePlugin object contains configuration for a specific plugin
Note to developers: new Route Plugins must be added to this struct
to be usable by Gloo.


::


   "transformations": .transformation.plugins.gloo.solo.io.RouteTransformations
   "faults": .fault.plugins.gloo.solo.io.RouteFaults
   "prefix_rewrite": .transformation.plugins.gloo.solo.io.PrefixRewrite
   "timeout": .google.protobuf.Duration
   "retries": .retries.plugins.gloo.solo.io.RetryPolicy
   "extensions": .gloo.solo.io.Extensions

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `transformations` | :ref:`transformation.plugins.gloo.solo.io.RouteTransformations` |  | 
   `faults` | :ref:`fault.plugins.gloo.solo.io.RouteFaults` |  | 
   `prefix_rewrite` | :ref:`transformation.plugins.gloo.solo.io.PrefixRewrite` |  | 
   `timeout` | `.google.protobuf.Duration<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration>`_ |  | 
   `retries` | :ref:`retries.plugins.gloo.solo.io.RetryPolicy` |  | 
   `extensions` | :ref:`gloo.solo.io.Extensions` |  | 



.. _gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Configuration for Destinations that are tied to the UpstreamSpec or ServiceSpec on that destination


::


   "aws": .aws.plugins.gloo.solo.io.DestinationSpec
   "azure": .azure.plugins.gloo.solo.io.DestinationSpec
   "rest": .rest.plugins.gloo.solo.io.DestinationSpec
   "grpc": .grpc.plugins.gloo.solo.io.DestinationSpec

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `aws` | :ref:`aws.plugins.gloo.solo.io.DestinationSpec` |  | 
   `azure` | :ref:`azure.plugins.gloo.solo.io.DestinationSpec` |  | 
   `rest` | :ref:`rest.plugins.gloo.solo.io.DestinationSpec` |  | 
   `grpc` | :ref:`grpc.plugins.gloo.solo.io.DestinationSpec` |  | 



.. _gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
Each upstream type is handled by a corresponding Gloo plugin.


::


   "kube": .kubernetes.plugins.gloo.solo.io.UpstreamSpec
   "static": .static.plugins.gloo.solo.io.UpstreamSpec
   "aws": .aws.plugins.gloo.solo.io.UpstreamSpec
   "azure": .azure.plugins.gloo.solo.io.UpstreamSpec
   "consul": .consul.plugins.gloo.solo.io.UpstreamSpec

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `kube` | :ref:`kubernetes.plugins.gloo.solo.io.UpstreamSpec` |  | 
   `static` | :ref:`static.plugins.gloo.solo.io.UpstreamSpec` |  | 
   `aws` | :ref:`aws.plugins.gloo.solo.io.UpstreamSpec` |  | 
   `azure` | :ref:`azure.plugins.gloo.solo.io.UpstreamSpec` |  | 
   `consul` | :ref:`consul.plugins.gloo.solo.io.UpstreamSpec` |  | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
