
===================================================
Package: `retries.plugins.gloo.solo.io`
===================================================  
TODO: this was copied form the transformation filter.
TODO: instead of manually copying, we want to do it via script, similar to the java-control-plane
TODO: to solo-kit/api/envoy




.. _retries.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/retries/retries.proto:


**Types:**


- :ref:`message.retries.plugins.gloo.solo.io.RetryPolicy`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/retries/retries.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/retries/retries.proto>`_




.. _message.retries.plugins.gloo.solo.io.RetryPolicy:

RetryPolicy
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Retry Policy applied to a route


::


   "retry_on": string
   "num_retries": int
   "per_try_timeout": .google.protobuf.Duration



.. _field.retries.plugins.gloo.solo.io.RetryPolicy.retry_on:

retry_on
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the conditions under which retry takes place. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on) 



.. _field.retries.plugins.gloo.solo.io.RetryPolicy.num_retries:

num_retries
++++++++++++++++++++++++++

Type: `int` 

Description: Specifies the allowed number of retries. This parameter is optional and defaults to 1. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on) 



.. _field.retries.plugins.gloo.solo.io.RetryPolicy.per_try_timeout:

per_try_timeout
++++++++++++++++++++++++++

Type: `.google.protobuf.Duration<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration>`_ 

Description: Specifies a non-zero upstream timeout per retry attempt. This parameter is optional. 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
