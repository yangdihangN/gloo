
===================================================
Package: `grpc.plugins.gloo.solo.io`
===================================================

.. _grpc.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc.proto:


**Types:**


- :ref:`message.grpc.plugins.gloo.solo.io.ServiceSpec`
- :ref:`message.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService`
- :ref:`message.grpc.plugins.gloo.solo.io.DestinationSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/grpc/grpc.proto>`_




.. _message.grpc.plugins.gloo.solo.io.ServiceSpec:

ServiceSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Service spec describing GRPC upstreams. This will usually be filled
automatically via function discovery (if the upstream supports reflection).
If your upstream service is a GRPC service, use this service spec (an empty
spec is fine), to make sure that traffic to it is routed with http2.


::


   "descriptors": bytes
   "grpc_services": []grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService



.. _field.grpc.plugins.gloo.solo.io.ServiceSpec.descriptors:

descriptors
++++++++++++++++++++++++++

Type: `bytes` 

Description: Descriptors that contain information of the services listed below. this is a serialized google.protobuf.FileDescriptorSet 



.. _field.grpc.plugins.gloo.solo.io.ServiceSpec.grpc_services:

grpc_services
++++++++++++++++++++++++++

Type: :ref:`message.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService` 

Description: List of services used by this upstream. For a grpc upstream where you don't need to use Gloo's function routing, this can be an empty list. These services must be present in the descriptors. 






.. _message.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService:

GrpcService
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Describes a grpc service


::


   "package_name": string
   "service_name": string
   "function_names": []string



.. _field.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService.package_name:

package_name
++++++++++++++++++++++++++

Type: `string` 

Description: The package of this service. 



.. _field.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService.service_name:

service_name
++++++++++++++++++++++++++

Type: `string` 

Description: The service name of this service. 



.. _field.grpc.plugins.gloo.solo.io.ServiceSpec.GrpcService.function_names:

function_names
++++++++++++++++++++++++++

Type: `[]string` 

Description: The functions available in this service. 






.. _message.grpc.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
This is only for upstream with Grpc service spec.


::


   "package": string
   "service": string
   "function": string
   "parameters": .transformation.plugins.gloo.solo.io.Parameters



.. _field.grpc.plugins.gloo.solo.io.DestinationSpec.package:

package
++++++++++++++++++++++++++

Type: `string` 

Description: The proto package of the function. 



.. _field.grpc.plugins.gloo.solo.io.DestinationSpec.service:

service
++++++++++++++++++++++++++

Type: `string` 

Description: The name of the service of the function. 



.. _field.grpc.plugins.gloo.solo.io.DestinationSpec.function:

function
++++++++++++++++++++++++++

Type: `string` 

Description: The name of the function. 



.. _field.grpc.plugins.gloo.solo.io.DestinationSpec.parameters:

parameters
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.Parameters` 

Description: Parameters describe how to extract the function parameters from the request. 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
