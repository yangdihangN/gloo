
===================================================
Package: `envoy.config.filter.http.aws.v2`
===================================================

.. _envoy.config.filter.http.aws.v2.github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/filter.proto:


**Types:**


- :ref:`envoy.config.filter.http.aws.v2.LambdaPerRoute`
- :ref:`envoy.config.filter.http.aws.v2.LambdaProtocolExtension`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/filter.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/pkg/plugins/aws/filter.proto>`_





.. _envoy.config.filter.http.aws.v2.LambdaPerRoute:

LambdaPerRoute
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
AWS Lambda contains the configuration necessary to perform transform regular http calls to
AWS Lambda invocations.


::


   "name": string
   "qualifier": string
   "async": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `name` | `string` | The name of the function | 
   `qualifier` | `string` | The qualifier of the function (defaults to $LATEST if not specified) | 
   `async` | `bool` | Invocation type - async or regular. | 



.. _envoy.config.filter.http.aws.v2.LambdaProtocolExtension:

LambdaProtocolExtension
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "host": string
   "region": string
   "access_key": string
   "secret_key": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `host` | `string` | The host header for AWS this cluster | 
   `region` | `string` | The region for this cluster | 
   `access_key` | `string` | The access_key for AWS this cluster | 
   `secret_key` | `string` | The secret_key for AWS this cluster | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
