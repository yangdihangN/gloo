
===================================================
Package: `core.solo.io`
===================================================

.. _core.solo.io.github.com/solo-io/solo-kit/api/v1/solo-kit.proto:


**Types:**


- :ref:`core.solo.io.Resource`
  



**Source File:** `github.com/solo-io/solo-kit/api/v1/solo-kit.proto <https://github.com/solo-io/solo-kit/blob/master/api/v1/solo-kit.proto>`_





.. _core.solo.io.Resource:

Resource
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "short_name": string
   "plural_name": string
   "cluster_scoped": bool
   "skip_docs_gen": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `short_name` | `string` | becomes the kubernetes short name for the generated crd | 
   `plural_name` | `string` | becomes the kubernetes plural name for the generated crd | 
   `cluster_scoped` | `bool` | the resource lives at the cluster level, namespace is ignored by the server | 
   `skip_docs_gen` | `bool` | indicates whether documentation generation has to be skipped for the given resource, defaults to false | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
