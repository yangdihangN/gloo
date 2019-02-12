
===================================================
Package: `grpc.plugins.gloo.solo.io`
===================================================

.. _grpc.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc.proto:


**Types:**


- :ref:`grpc.plugins.gloo.solo.io.ServiceSpec`
- :ref:`grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService`
- :ref:`grpc.plugins.gloo.solo.io.DestinationSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/grpc/grpc.proto>`_





.. _grpc.plugins.gloo.solo.io.ServiceSpec:

ServiceSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Service spec describing GRPC upstreams. This will usually be filled
automatically via function discovery (if the upstream supports reflection).
If your upstream service is a GRPC service, use this service spec (an empty
spec is fine), to make sure that traffic to it is routed with http2.


::


   "descriptors": bytes
   "grpc_services": []grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `descriptors` | `bytes` | Descriptors that contain information of the services listed below. this is a serialized google.protobuf.FileDescriptorSet | 
   `grpc_services` | :ref:`[]grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService` | List of services used by this upstream. For a grpc upstream where you don't need to use Gloo's function routing, this can be an empty list. These services must be present in the descriptors. | 



.. _grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService:

GrpcService
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Describes a grpc service


::


   "package_name": string
   "service_name": string
   "function_names": []string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `package_name` | `string` | The package of this service. | 
   `service_name` | `string` | The service name of this service. | 
   `function_names` | `[]string` | The functions available in this service. | 



.. _grpc.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
This is only for upstream with Grpc service spec.


::


   "package": string
   "service": string
   "function": string
   "parameters": .transformation.plugins.gloo.solo.io.Parameters

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `package` | `string` | The proto package of the function. | 
   `service` | `string` | The name of the service of the function. | 
   `function` | `string` | The name of the function. | 
   `parameters` | :ref:`transformation.plugins.gloo.solo.io.Parameters` | Parameters describe how to extract the function parameters from the request. | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
