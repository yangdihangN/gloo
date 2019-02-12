
===================================================
Package: `ingress.solo.io`
===================================================

.. _ingress.solo.io.github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto:


**Types:**


- :ref:`message.ingress.solo.io.Ingress` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto <https://github.com/solo-io/gloo/blob/master/projects/ingress/api/v1/ingress.proto>`_




.. _message.ingress.solo.io.Ingress:

Ingress
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

A simple wrapper for a Kubernetes Ingress Object.


::


   "kube_ingress_spec": .google.protobuf.Any
   "kube_ingress_status": .google.protobuf.Any
   "metadata": .core.solo.io.Metadata



.. _field.ingress.solo.io.Ingress.kube_ingress_spec:

kube_ingress_spec
++++++++++++++++++++++++++

Type: `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ 

Description: a raw byte representation of the kubernetes ingress this resource wraps 



.. _field.ingress.solo.io.Ingress.kube_ingress_status:

kube_ingress_status
++++++++++++++++++++++++++

Type: `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ 

Description: a raw byte representation of the ingress status of the kubernetes ingress object 



.. _field.ingress.solo.io.Ingress.metadata:

metadata
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Metadata` 

Description: Metadata contains the object metadata for this resource 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
