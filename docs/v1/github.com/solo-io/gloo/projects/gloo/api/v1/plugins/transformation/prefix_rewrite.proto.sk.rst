
===================================================
Package: `transformation.plugins.gloo.solo.io`
===================================================  
TODO: this was copied form the transformation filter.
TODO: instead of manually copying, we want to do it via script, similar to the java-control-plane
TODO: to solo-kit/api/envoy




.. _transformation.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/prefix_rewrite.proto:


**Types:**


- :ref:`transformation.plugins.gloo.solo.io.PrefixRewrite`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/prefix_rewrite.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/transformation/prefix_rewrite.proto>`_





.. _transformation.plugins.gloo.solo.io.PrefixRewrite:

PrefixRewrite
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
if set, prefix_rewrite will be used to rewrite the matched HTTP Path prefix on requests to this value.


::


   "prefix_rewrite": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `prefix_rewrite` | `string` | Set to an empty string to remove the matched HTTP Path prefix | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
