
===================================================
Package: `envoy.config.filter.http.aws.v2`
===================================================

.. _envoy.config.filter.http.aws.v2.github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/filter.proto:


**Types:**


- :ref:`message.envoy.config.filter.http.aws.v2.LambdaPerRoute`
- :ref:`message.envoy.config.filter.http.aws.v2.LambdaProtocolExtension`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/filter.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/pkg/plugins/aws/filter.proto>`_




.. _message.envoy.config.filter.http.aws.v2.LambdaPerRoute:

LambdaPerRoute
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
AWS Lambda contains the configuration necessary to perform transform regular http calls to
AWS Lambda invocations.


::


   "name": string
   "qualifier": string
   "async": bool



.. _field.envoy.config.filter.http.aws.v2.LambdaPerRoute.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: The name of the function 



.. _field.envoy.config.filter.http.aws.v2.LambdaPerRoute.qualifier:

qualifier
++++++++++++++++++++++++++

Type: `string` 

Description: The qualifier of the function (defaults to $LATEST if not specified) 



.. _field.envoy.config.filter.http.aws.v2.LambdaPerRoute.async:

async
++++++++++++++++++++++++++

Type: `bool` 

Description: Invocation type - async or regular. 






.. _message.envoy.config.filter.http.aws.v2.LambdaProtocolExtension:

LambdaProtocolExtension
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "host": string
   "region": string
   "access_key": string
   "secret_key": string



.. _field.envoy.config.filter.http.aws.v2.LambdaProtocolExtension.host:

host
++++++++++++++++++++++++++

Type: `string` 

Description: The host header for AWS this cluster 



.. _field.envoy.config.filter.http.aws.v2.LambdaProtocolExtension.region:

region
++++++++++++++++++++++++++

Type: `string` 

Description: The region for this cluster 



.. _field.envoy.config.filter.http.aws.v2.LambdaProtocolExtension.access_key:

access_key
++++++++++++++++++++++++++

Type: `string` 

Description: The access_key for AWS this cluster 



.. _field.envoy.config.filter.http.aws.v2.LambdaProtocolExtension.secret_key:

secret_key
++++++++++++++++++++++++++

Type: `string` 

Description: The secret_key for AWS this cluster 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
