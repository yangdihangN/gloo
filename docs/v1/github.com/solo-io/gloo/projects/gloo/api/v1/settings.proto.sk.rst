
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto:


**Types:**


- :ref:`gloo.solo.io.Settings` **Top-Level Resource**
- :ref:`gloo.solo.io.Settings.KubernetesCrds`
- :ref:`gloo.solo.io.Settings.KubernetesSecrets`
- :ref:`gloo.solo.io.Settings.VaultSecrets`
- :ref:`gloo.solo.io.Settings.KubernetesConfigmaps`
- :ref:`gloo.solo.io.Settings.Directory`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/settings.proto>`_





.. _gloo.solo.io.Settings:

Settings
~~~~~~~~~~~~~~~~~~~~~~~~~~

 



::


   "discovery_namespace": string
   "watch_namespaces": []string
   "kubernetes_config_source": .gloo.solo.io.Settings.KubernetesCrds
   "directory_config_source": .gloo.solo.io.Settings.Directory
   "kubernetes_secret_source": .gloo.solo.io.Settings.KubernetesSecrets
   "vault_secret_source": .gloo.solo.io.Settings.VaultSecrets
   "directory_secret_source": .gloo.solo.io.Settings.Directory
   "kubernetes_artifact_source": .gloo.solo.io.Settings.KubernetesConfigmaps
   "directory_artifact_source": .gloo.solo.io.Settings.Directory
   "bind_addr": string
   "refresh_rate": .google.protobuf.Duration
   "dev_mode": bool
   "extensions": .gloo.solo.io.Extensions
   "metadata": .core.solo.io.Metadata
   "status": .core.solo.io.Status

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `discovery_namespace` | `string` | namespace to write discovered data | 
   `watch_namespaces` | `[]string` | namespaces to watch for user config as well as services TODO(ilackarms): split out watch_namespaces and service_discovery_namespaces... | 
   `kubernetes_config_source` | :ref:`gloo.solo.io.Settings.KubernetesCrds` |  | 
   `directory_config_source` | :ref:`gloo.solo.io.Settings.Directory` |  | 
   `kubernetes_secret_source` | :ref:`gloo.solo.io.Settings.KubernetesSecrets` |  | 
   `vault_secret_source` | :ref:`gloo.solo.io.Settings.VaultSecrets` |  | 
   `directory_secret_source` | :ref:`gloo.solo.io.Settings.Directory` |  | 
   `kubernetes_artifact_source` | :ref:`gloo.solo.io.Settings.KubernetesConfigmaps` |  | 
   `directory_artifact_source` | :ref:`gloo.solo.io.Settings.Directory` |  | 
   `bind_addr` | `string` | where the gloo xds server should bind (should not need configuration by user) | 
   `refresh_rate` | `.google.protobuf.Duration<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration>`_ | how frequently to resync watches, etc | 
   `dev_mode` | `bool` | enable serving debug data on port 9090 | 
   `extensions` | :ref:`gloo.solo.io.Extensions` | Settings for extensions | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 
   `status` | :ref:`core.solo.io.Status` | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation | 



.. _gloo.solo.io.Settings.KubernetesCrds:

KubernetesCrds
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
ilackarms(todo: make sure these are configurable)


::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _gloo.solo.io.Settings.KubernetesSecrets:

KubernetesSecrets
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _gloo.solo.io.Settings.VaultSecrets:

VaultSecrets
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _gloo.solo.io.Settings.KubernetesConfigmaps:

KubernetesConfigmaps
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _gloo.solo.io.Settings.Directory:

Directory
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "directory": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `directory` | `string` |  | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
