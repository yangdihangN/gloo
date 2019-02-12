
===================================================
Package: `ingress.solo.io`
===================================================

.. _ingress.solo.io.github.com/solo-io/gloo/projects/ingress/api/v1/service.proto:


**Types:**


- :ref:`message.ingress.solo.io.KubeService` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/ingress/api/v1/service.proto <https://github.com/solo-io/gloo/blob/master/projects/ingress/api/v1/service.proto>`_




.. _message.ingress.solo.io.KubeService:

KubeService
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

A simple wrapper for a Kubernetes Service Object.


::


   "kube_service_spec": .google.protobuf.Any
   "kube_service_status": .google.protobuf.Any
   "metadata": .core.solo.io.Metadata



.. _field.ingress.solo.io.KubeService.kube_service_spec:

kube_service_spec
++++++++++++++++++++++++++

Type: `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ 

Description: a raw byte representation of the kubernetes service this resource wraps 



.. _field.ingress.solo.io.KubeService.kube_service_status:

kube_service_status
++++++++++++++++++++++++++

Type: `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ 

Description: a raw byte representation of the service status of the kubernetes service object 



.. _field.ingress.solo.io.KubeService.metadata:

metadata
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Metadata` 

Description: Metadata contains the object metadata for this resource 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
