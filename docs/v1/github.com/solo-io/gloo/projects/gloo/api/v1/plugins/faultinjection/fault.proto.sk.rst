
===================================================
Package: `fault.plugins.gloo.solo.io`
===================================================  
TODO: this was copied form the transformation filter.
TODO: instead of manually copying, we want to do it via script, similar to the java-control-plane
TODO: to solo-kit/api/envoy




.. _fault.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/faultinjection/fault.proto:


**Types:**


- :ref:`fault.plugins.gloo.solo.io.RouteAbort`
- :ref:`fault.plugins.gloo.solo.io.RouteDelay`
- :ref:`fault.plugins.gloo.solo.io.RouteFaults`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/faultinjection/fault.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/faultinjection/fault.proto>`_





.. _fault.plugins.gloo.solo.io.RouteAbort:

RouteAbort
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "percentage": float
   "http_status": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `percentage` | `float` | Percentage of requests that should be aborted, defaulting to 0. This should be a value between 0.0 and 100.0, with up to 6 significant digits. | 
   `http_status` | `int` | This should be a standard HTTP status, i.e. 503. Defaults to 0. | 



.. _fault.plugins.gloo.solo.io.RouteDelay:

RouteDelay
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "percentage": float
   "fixed_delay": .google.protobuf.Duration

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `percentage` | `float` | Percentage of requests that should be delayed, defaulting to 0. This should be a value between 0.0 and 100.0, with up to 6 significant digits. | 
   `fixed_delay` | `.google.protobuf.Duration<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration>`_ | Fixed delay, defaulting to 0. | 



.. _fault.plugins.gloo.solo.io.RouteFaults:

RouteFaults
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "abort": .fault.plugins.gloo.solo.io.RouteAbort
   "delay": .fault.plugins.gloo.solo.io.RouteDelay

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `abort` | :ref:`fault.plugins.gloo.solo.io.RouteAbort` |  | 
   `delay` | :ref:`fault.plugins.gloo.solo.io.RouteDelay` |  | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
