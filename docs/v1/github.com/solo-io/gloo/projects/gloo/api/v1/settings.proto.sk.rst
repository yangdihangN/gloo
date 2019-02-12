
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto:


**Types:**


- :ref:`message.gloo.solo.io.Settings` **Top-Level Resource**
- :ref:`message.gloo.solo.io.Settings.KubernetesCrds`
- :ref:`message.gloo.solo.io.Settings.KubernetesSecrets`
- :ref:`message.gloo.solo.io.Settings.VaultSecrets`
- :ref:`message.gloo.solo.io.Settings.KubernetesConfigmaps`
- :ref:`message.gloo.solo.io.Settings.Directory`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/settings.proto>`_




.. _message.gloo.solo.io.Settings:

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



.. _field.gloo.solo.io.Settings.discovery_namespace:

discovery_namespace
++++++++++++++++++++++++++

Type: `string` 

Description: namespace to write discovered data 



.. _field.gloo.solo.io.Settings.watch_namespaces:

watch_namespaces
++++++++++++++++++++++++++

Type: `[]string` 

Description: namespaces to watch for user config as well as services TODO(ilackarms): split out watch_namespaces and service_discovery_namespaces... 



.. _field.gloo.solo.io.Settings.kubernetes_config_source:

kubernetes_config_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.KubernetesCrds` 

Description:  



.. _field.gloo.solo.io.Settings.directory_config_source:

directory_config_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.Directory` 

Description:  



.. _field.gloo.solo.io.Settings.kubernetes_secret_source:

kubernetes_secret_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.KubernetesSecrets` 

Description:  



.. _field.gloo.solo.io.Settings.vault_secret_source:

vault_secret_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.VaultSecrets` 

Description:  



.. _field.gloo.solo.io.Settings.directory_secret_source:

directory_secret_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.Directory` 

Description:  



.. _field.gloo.solo.io.Settings.kubernetes_artifact_source:

kubernetes_artifact_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.KubernetesConfigmaps` 

Description:  



.. _field.gloo.solo.io.Settings.directory_artifact_source:

directory_artifact_source
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Settings.Directory` 

Description:  



.. _field.gloo.solo.io.Settings.bind_addr:

bind_addr
++++++++++++++++++++++++++

Type: `string` 

Description: where the gloo xds server should bind (should not need configuration by user) 



.. _field.gloo.solo.io.Settings.refresh_rate:

refresh_rate
++++++++++++++++++++++++++

Type: `.google.protobuf.Duration<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration>`_ 

Description: how frequently to resync watches, etc 



.. _field.gloo.solo.io.Settings.dev_mode:

dev_mode
++++++++++++++++++++++++++

Type: `bool` 

Description: enable serving debug data on port 9090 



.. _field.gloo.solo.io.Settings.extensions:

extensions
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Extensions` 

Description: Settings for extensions 



.. _field.gloo.solo.io.Settings.metadata:

metadata
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Metadata` 

Description: Metadata contains the object metadata for this resource 



.. _field.gloo.solo.io.Settings.status:

status
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Status` 

Description: Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation 






.. _message.gloo.solo.io.Settings.KubernetesCrds:

KubernetesCrds
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
ilackarms(todo: make sure these are configurable)


::








.. _message.gloo.solo.io.Settings.KubernetesSecrets:

KubernetesSecrets
~~~~~~~~~~~~~~~~~~~~~~~~~~



::








.. _message.gloo.solo.io.Settings.VaultSecrets:

VaultSecrets
~~~~~~~~~~~~~~~~~~~~~~~~~~



::








.. _message.gloo.solo.io.Settings.KubernetesConfigmaps:

KubernetesConfigmaps
~~~~~~~~~~~~~~~~~~~~~~~~~~



::








.. _message.gloo.solo.io.Settings.Directory:

Directory
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "directory": string



.. _field.gloo.solo.io.Settings.Directory.directory:

directory
++++++++++++++++++++++++++

Type: `string` 

Description:  







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
