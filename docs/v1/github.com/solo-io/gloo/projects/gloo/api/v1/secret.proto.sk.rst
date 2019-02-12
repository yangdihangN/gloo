
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto:


**Types:**


- :ref:`gloo.solo.io.Secret` **Top-Level Resource**
- :ref:`gloo.solo.io.AwsSecret`
- :ref:`gloo.solo.io.AzureSecret`
- :ref:`gloo.solo.io.TlsSecret`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/secret.proto>`_





.. _gloo.solo.io.Secret:

Secret
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Certain plugins such as the AWS Lambda Plugin require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.

Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:

Kubernetes Secrets
Hashicorp Vault
Plaintext files (recommended only for testing)
Secrets must adhere to a structure, specified by the plugin that requires them.

Gloo's secret backend can be configured in Gloo's bootstrap options


::


   "aws": .gloo.solo.io.AwsSecret
   "azure": .gloo.solo.io.AzureSecret
   "tls": .gloo.solo.io.TlsSecret
   "extension": .gloo.solo.io.Extension
   "metadata": .core.solo.io.Metadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `aws` | :ref:`gloo.solo.io.AwsSecret` |  | 
   `azure` | :ref:`gloo.solo.io.AzureSecret` |  | 
   `tls` | :ref:`gloo.solo.io.TlsSecret` |  | 
   `extension` | :ref:`gloo.solo.io.Extension` |  | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 



.. _gloo.solo.io.AwsSecret:

AwsSecret
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "access_key": string
   "secret_key": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `access_key` | `string` |  | 
   `secret_key` | `string` |  | 



.. _gloo.solo.io.AzureSecret:

AzureSecret
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "api_keys": map<string, string>

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `api_keys` | `map<string, string>` |  | 



.. _gloo.solo.io.TlsSecret:

TlsSecret
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "cert_chain": string
   "private_key": string
   "root_ca": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `cert_chain` | `string` |  | 
   `private_key` | `string` |  | 
   `root_ca` | `string` |  | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
