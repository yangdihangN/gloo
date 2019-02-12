
===================================================
Package: `core.solo.io`
===================================================

.. _core.solo.io.github.com/solo-io/solo-kit/api/v1/solo-kit.proto:


**Types:**


- :ref:`message.core.solo.io.Resource`
  



**Source File:** `github.com/solo-io/solo-kit/api/v1/solo-kit.proto <https://github.com/solo-io/solo-kit/blob/master/api/v1/solo-kit.proto>`_




.. _message.core.solo.io.Resource:

Resource
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "short_name": string
   "plural_name": string
   "cluster_scoped": bool
   "skip_docs_gen": bool



.. _field.core.solo.io.Resource.short_name:

short_name
++++++++++++++++++++++++++

Type: `string` 

Description: becomes the kubernetes short name for the generated crd 



.. _field.core.solo.io.Resource.plural_name:

plural_name
++++++++++++++++++++++++++

Type: `string` 

Description: becomes the kubernetes plural name for the generated crd 



.. _field.core.solo.io.Resource.cluster_scoped:

cluster_scoped
++++++++++++++++++++++++++

Type: `bool` 

Description: the resource lives at the cluster level, namespace is ignored by the server 



.. _field.core.solo.io.Resource.skip_docs_gen:

skip_docs_gen
++++++++++++++++++++++++++

Type: `bool` 

Description: indicates whether documentation generation has to be skipped for the given resource, defaults to false 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
