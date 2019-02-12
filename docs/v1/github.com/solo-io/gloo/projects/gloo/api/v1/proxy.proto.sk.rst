
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto:


**Types:**


- :ref:`message.gloo.solo.io.Proxy` **Top-Level Resource**
- :ref:`message.gloo.solo.io.Listener`
- :ref:`message.gloo.solo.io.HttpListener`
- :ref:`message.gloo.solo.io.VirtualHost`
- :ref:`message.gloo.solo.io.Route`
- :ref:`message.gloo.solo.io.Matcher`
- :ref:`message.gloo.solo.io.HeaderMatcher`
- :ref:`message.gloo.solo.io.QueryParameterMatcher`
- :ref:`message.gloo.solo.io.RouteAction`
- :ref:`message.gloo.solo.io.Destination`
- :ref:`message.gloo.solo.io.MultiDestination`
- :ref:`message.gloo.solo.io.WeightedDestination`
- :ref:`message.gloo.solo.io.RedirectAction`
- [RedirectResponseCode](#RedirectResponseCode)
- :ref:`message.gloo.solo.io.DirectResponseAction`
- :ref:`message.gloo.solo.io.SslConfig`
- :ref:`message.gloo.solo.io.SSLFiles`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/proxy.proto>`_




.. _message.gloo.solo.io.Proxy:

Proxy
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

A Proxy is a container for the entire set of configuration that will to be applied to one or more Proxy instances.
Proxies can be understood as a set of listeners, represents a different bind address/port where the proxy will listen
for connections. Each listener has its own set of configuration.

If any of the sub-resources within a listener is declared invalid (e.g. due to invalid user configuration), the
proxy will be marked invalid by Gloo.

Proxy instances that register with Gloo are assigned the proxy configuration corresponding with
a proxy-specific identifier.
In the case of Envoy, proxy instances are identified by their Node ID. Node IDs must match a existing Proxy
Node ID can be specified in Envoy with the `--service-node` flag, or in the Envoy instance's bootstrap config.


::


   "listeners": []gloo.solo.io.Listener
   "status": .core.solo.io.Status
   "metadata": .core.solo.io.Metadata



.. _field.gloo.solo.io.Proxy.listeners:

listeners
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Listener` 

Description: Define here each listener the proxy should create. Listeners define the a set of behaviors for a single bind address/port where the proxy will listen If no listeners are specified, the instances configured with the proxy resource will not accept connections. 



.. _field.gloo.solo.io.Proxy.status:

status
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Status` 

Description: Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation 



.. _field.gloo.solo.io.Proxy.metadata:

metadata
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Metadata` 

Description: Metadata contains the object metadata for this resource 






.. _message.gloo.solo.io.Listener:

Listener
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Listeners define the address:port where the proxy will listen for incoming connections
A Listener accepts connections (currently only HTTP is supported) and apply user-defined behavior for those connections,
e.g. performing SSL termination, HTTP retries, and rate limiting.


::


   "name": string
   "bind_address": string
   "bind_port": int
   "http_listener": .gloo.solo.io.HttpListener
   "ssl_configuations": []gloo.solo.io.SslConfig



.. _field.gloo.solo.io.Listener.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: the name of the listener. names must be unique for each listener within a proxy 



.. _field.gloo.solo.io.Listener.bind_address:

bind_address
++++++++++++++++++++++++++

Type: `string` 

Description: the bind address for the listener. both ipv4 and ipv6 formats are supported 



.. _field.gloo.solo.io.Listener.bind_port:

bind_port
++++++++++++++++++++++++++

Type: `int` 

Description: the port to bind on ports numbers must be unique for listeners within a proxy 



.. _field.gloo.solo.io.Listener.http_listener:

http_listener
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.HttpListener` 

Description: The HTTP Listener is currently the only supported listener type. It contains configuration options for GLoo's HTTP-level features including request-based routing 



.. _field.gloo.solo.io.Listener.ssl_configuations:

ssl_configuations
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.SslConfig` 

Description: SSL Config is optional for the listener. If provided, the listener will serve TLS for connections on this port Multiple SslConfigs are supported for the pupose of SNI. Be aware that the SNI domain provided in the SSL Config must match a domain in virtual host TODO(ilackarms): ensure that ssl configs without a matching virtual host are errored 






.. _message.gloo.solo.io.HttpListener:

HttpListener
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Use this listener to configure proxy behavior for any HTTP-level features including defining routes (via virtualservices).
HttpListeners also contain plugin configuration that applies globally across all virtaul hosts on the listener.
Some plugins can be configured to work both on the listener and virtual host level (such as the rate limit plugin)


::


   "virtual_hosts": []gloo.solo.io.VirtualHost
   "listener_plugins": .gloo.solo.io.ListenerPlugins



.. _field.gloo.solo.io.HttpListener.virtual_hosts:

virtual_hosts
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.VirtualHost` 

Description: the set of virtual hosts that will be accessible by clients connecting to this listener. at least one virtual host must be specified for this listener to be active (else connections will be refused) the set of domains for each virtual host must be unique, or the config will be considered invalid 



.. _field.gloo.solo.io.HttpListener.listener_plugins:

listener_plugins
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.ListenerPlugins` 

Description: Plugins contains top-level plugin configuration to be applied to a listener Listener config is applied to all HTTP traffic that connects to this listener. Some configuration here can be overridden in Virtual Host Plugin configuration or Route Plugin configuration Plugins should be specified here in the form of `"plugin_name": {..//plugin_config...}` to allow specifying multiple plugins. 






.. _message.gloo.solo.io.VirtualHost:

VirtualHost
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Virtual Hosts group an ordered list of routes under one or more domains.
Each Virtual Host has a logical name, which must be unique for the listener.
An HTTP request is first matched to a virtual host based on its host header, then to a route within the virtual host.
If a request is not matched to any virtual host or a route therein, the target proxy will reply with a 404.


::


   "name": string
   "domains": []string
   "routes": []gloo.solo.io.Route
   "virtual_host_plugins": .gloo.solo.io.VirtualHostPlugins



.. _field.gloo.solo.io.VirtualHost.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: the logical name of the virtual host. names must be unique for each virtual host within a listener 



.. _field.gloo.solo.io.VirtualHost.domains:

domains
++++++++++++++++++++++++++

Type: `[]string` 

Description: The list of domains (i.e.: matching the `Host` header of a request) that belong to this virtual host. Note that the wildcard will not match the empty string. e.g. “*-bar.foo.com” will match “baz-bar.foo.com” but not “-bar.foo.com”. Additionally, a special entry “*” is allowed which will match any host/authority header. Only a single virtual host in the entire route configuration can match on “*”. A domain must be unique across all virtual hosts or the config will be invalidated by Gloo Domains on virtual hosts obey the same rules as [Envoy Virtual Hosts](https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto) 



.. _field.gloo.solo.io.VirtualHost.routes:

routes
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Route` 

Description: The list of HTTP routes define routing actions to be taken for incoming HTTP requests whose host header matches this virtual host. If the request matches more than one route in the list, the first route matched will be selected. If the list of routes is empty, the virtual host will be ignored by Gloo. 



.. _field.gloo.solo.io.VirtualHost.virtual_host_plugins:

virtual_host_plugins
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.VirtualHostPlugins` 

Description: Plugins contains top-level plugin configuration to be applied to a listener Listener config is applied to all HTTP traffic that connects to this listener. Some configuration here can be overridden in Virtual Host Plugin configuration or Route Plugin configuration Plugins should be specified here in the form of `"plugin_name": {..//plugin_config...}` to allow specifying multiple plugins. 






.. _message.gloo.solo.io.Route:

Route
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
*
Routes declare the entrypoints on virtual hosts and the action to take for matched requests.


::


   "matcher": .gloo.solo.io.Matcher
   "route_action": .gloo.solo.io.RouteAction
   "redirect_action": .gloo.solo.io.RedirectAction
   "direct_response_action": .gloo.solo.io.DirectResponseAction
   "route_plugins": .gloo.solo.io.RoutePlugins



.. _field.gloo.solo.io.Route.matcher:

matcher
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Matcher` 

Description: The matcher contains parameters for matching requests (i.e.: based on HTTP path, headers, etc.) 



.. _field.gloo.solo.io.Route.route_action:

route_action
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.RouteAction` 

Description: This action is the primary action to be selected for most routes. The RouteAction tells the proxy to route requests to an upstream. 



.. _field.gloo.solo.io.Route.redirect_action:

redirect_action
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.RedirectAction` 

Description: Redirect actions tell the proxy to return a redirect response to the downstream client 



.. _field.gloo.solo.io.Route.direct_response_action:

direct_response_action
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.DirectResponseAction` 

Description: Return an arbitrary HTTP response directly, without proxying. 



.. _field.gloo.solo.io.Route.route_plugins:

route_plugins
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.RoutePlugins` 

Description: Route Plugins extend the behavior of routes. Route plugins include configuration such as retries, rate limiting, and request/resonse transformation. Plugins should be specified here in the form of `"plugin_name": {..//plugin_config...}` to allow specifying multiple plugins. 






.. _message.gloo.solo.io.Matcher:

Matcher
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Parameters for matching routes to requests received by a Gloo-managed proxy


::


   "prefix": string
   "exact": string
   "regex": string
   "headers": []gloo.solo.io.HeaderMatcher
   "query_parameters": []gloo.solo.io.QueryParameterMatcher
   "methods": []string



.. _field.gloo.solo.io.Matcher.prefix:

prefix
++++++++++++++++++++++++++

Type: `string` 

Description: If specified, the route is a prefix rule meaning that the prefix must match the beginning of the *:path* header. 



.. _field.gloo.solo.io.Matcher.exact:

exact
++++++++++++++++++++++++++

Type: `string` 

Description: If specified, the route is an exact path rule meaning that the path must exactly match the *:path* header once the query string is removed. 



.. _field.gloo.solo.io.Matcher.regex:

regex
++++++++++++++++++++++++++

Type: `string` 

Description: If specified, the route is a regular expression rule meaning that the regex must match the *:path* header once the query string is removed. The entire path (without the query string) must match the regex. The rule will not match if only a subsequence of the *:path* header matches the regex. The regex grammar is defined `here <http://en.cppreference.com/w/cpp/regex/ecmascript>`_. Examples: * The regex */b[io]t* matches the path */bit* * The regex */b[io]t* matches the path */bot* * The regex */b[io]t* does not match the path */bite* * The regex */b[io]t* does not match the path */bit/bot* 



.. _field.gloo.solo.io.Matcher.headers:

headers
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.HeaderMatcher` 

Description: Specifies a set of headers that the route should match on. The router will check the request’s headers against all the specified headers in the route config. A match will happen if all the headers in the route are present in the request with the same values (or based on presence if the value field is not in the config). 



.. _field.gloo.solo.io.Matcher.query_parameters:

query_parameters
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.QueryParameterMatcher` 

Description: Specifies a set of URL query parameters on which the route should match. The router will check the query string from the *path* header against all the specified query parameters. If the number of specified query parameters is nonzero, they all must match the *path* header's query string for a match to occur. 



.. _field.gloo.solo.io.Matcher.methods:

methods
++++++++++++++++++++++++++

Type: `[]string` 

Description: HTTP Method/Verb(s) to match on. If none specified, the matcher will ignore the HTTP Method 






.. _message.gloo.solo.io.HeaderMatcher:

HeaderMatcher
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Internally, Gloo always uses the HTTP/2 *:authority* header to represent the HTTP/1 *Host*
  header. Thus, if attempting to match on *Host*, match on *:authority* instead.

  In the absence of any header match specifier, match will default to `present_match`
  i.e, a request that has the `name` header will match, regardless of the header's
  value.


::


   "name": string
   "value": string
   "regex": bool



.. _field.gloo.solo.io.HeaderMatcher.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the name of the header in the request. 



.. _field.gloo.solo.io.HeaderMatcher.value:

value
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the value of the header. If the value is absent a request that has the name header will match, regardless of the header’s value. 



.. _field.gloo.solo.io.HeaderMatcher.regex:

regex
++++++++++++++++++++++++++

Type: `bool` 

Description: Specifies whether the header value should be treated as regex or not. 






.. _message.gloo.solo.io.QueryParameterMatcher:

QueryParameterMatcher
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Query parameter matching treats the query string of a request's :path header
as an ampersand-separated list of keys and/or key=value elements.


::


   "name": string
   "value": string
   "regex": bool



.. _field.gloo.solo.io.QueryParameterMatcher.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the name of a key that must be present in the requested *path*'s query string. 



.. _field.gloo.solo.io.QueryParameterMatcher.value:

value
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the value of the key. If the value is absent, a request that contains the key in its query string will match, whether the key appears with a value (e.g., "?debug=true") or not (e.g., "?debug") 



.. _field.gloo.solo.io.QueryParameterMatcher.regex:

regex
++++++++++++++++++++++++++

Type: `bool` 

Description: Specifies whether the query parameter value is a regular expression. Defaults to false. The entire query parameter value (i.e., the part to the right of the equals sign in "key=value") must match the regex. E.g., the regex "\d+$" will match "123" but not "a123" or "123a". 






.. _message.gloo.solo.io.RouteAction:

RouteAction
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
RouteActions are used to route matched requests to upstreams.


::


   "single": .gloo.solo.io.Destination
   "multi": .gloo.solo.io.MultiDestination



.. _field.gloo.solo.io.RouteAction.single:

single
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Destination` 

Description: Use SingleDestination to route to a single upstream 



.. _field.gloo.solo.io.RouteAction.multi:

multi
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.MultiDestination` 

Description: Use MultiDestination to load balance requests between multiple upstreams (by weight) 






.. _message.gloo.solo.io.Destination:

Destination
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Destinations define routable destinations for proxied requests


::


   "upstream": .core.solo.io.ResourceRef
   "destination_spec": .gloo.solo.io.DestinationSpec



.. _field.gloo.solo.io.Destination.upstream:

upstream
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.ResourceRef` 

Description: The upstream to route requests to 



.. _field.gloo.solo.io.Destination.destination_spec:

destination_spec
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.DestinationSpec` 

Description: Some upstreams utilize plugins which require or permit additional configuration on routes targeting them. gRPC upstreams, for example, allow specifying REST-style parameters for JSON-to-gRPC transcoding in the destination config. If the destination config is required for the upstream and not provided by the user, Gloo will invalidate the destination and its parent resources. 






.. _message.gloo.solo.io.MultiDestination:

MultiDestination
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
MultiDestination is a container for a set of weighted destinations. Gloo will load balance traffic for a single
route across multiple destinations according to their specified weights.


::


   "destinations": []gloo.solo.io.WeightedDestination



.. _field.gloo.solo.io.MultiDestination.destinations:

destinations
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.WeightedDestination` 

Description: This list must contain at least one destination or the listener housing this route will be invalid, causing Gloo to error the parent proxy resource. 






.. _message.gloo.solo.io.WeightedDestination:

WeightedDestination
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
WeightedDestination attaches a weight to a single destination.


::


   "destination": .gloo.solo.io.Destination
   "weight": int



.. _field.gloo.solo.io.WeightedDestination.destination:

destination
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.Destination` 

Description:  



.. _field.gloo.solo.io.WeightedDestination.weight:

weight
++++++++++++++++++++++++++

Type: `int` 

Description: Weight must be greater than zero Routing to each destination will be balanced by the ratio of the destination's weight to the total weight on a route 






.. _message.gloo.solo.io.RedirectAction:

RedirectAction
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
TODO(ilackarms): evaluate how much to differentiate (or if even to include) RedirectAction
Notice: RedirectAction is copioed directly from https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto


::


   "host_redirect": string
   "path_redirect": string
   "prefix_rewrite": string
   "response_code": .gloo.solo.io.RedirectAction.RedirectResponseCode
   "https_redirect": bool
   "strip_query": bool



.. _field.gloo.solo.io.RedirectAction.host_redirect:

host_redirect
++++++++++++++++++++++++++

Type: `string` 

Description: The host portion of the URL will be swapped with this value. 



.. _field.gloo.solo.io.RedirectAction.path_redirect:

path_redirect
++++++++++++++++++++++++++

Type: `string` 

Description: The path portion of the URL will be swapped with this value. 



.. _field.gloo.solo.io.RedirectAction.prefix_rewrite:

prefix_rewrite
++++++++++++++++++++++++++

Type: `string` 

Description: Indicates that during redirection, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request. Pay attention to the use of trailing slashes as mentioned in `RouteAction`'s `prefix_rewrite`. 



.. _field.gloo.solo.io.RedirectAction.response_code:

response_code
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.RedirectAction.RedirectResponseCode` 

Description: The HTTP status code to use in the redirect response. The default response code is MOVED_PERMANENTLY (301). 



.. _field.gloo.solo.io.RedirectAction.https_redirect:

https_redirect
++++++++++++++++++++++++++

Type: `bool` 

Description: The scheme portion of the URL will be swapped with "https". 



.. _field.gloo.solo.io.RedirectAction.strip_query:

strip_query
++++++++++++++++++++++++++

Type: `bool` 

Description: Indicates that during redirection, the query portion of the URL will be removed. Default value is false. 






---
### <a name="RedirectResponseCode">RedirectResponseCode</a>



.. csv-table:: Enum Reference
   :header: "Name", "Description"
   :delim: |


   `MOVED_PERMANENTLY` | Moved Permanently HTTP Status Code - 301.

   `FOUND` | Found HTTP Status Code - 302.

   `SEE_OTHER` | See Other HTTP Status Code - 303.

   `TEMPORARY_REDIRECT` | Temporary Redirect HTTP Status Code - 307.

   `PERMANENT_REDIRECT` | Permanent Redirect HTTP Status Code - 308.




.. _message.gloo.solo.io.DirectResponseAction:

DirectResponseAction
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
TODO(ilackarms): evaluate how much to differentiate (or if even to include) DirectResponseAction
DirectResponseAction is copied directly from https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto


::


   "status": int
   "body": string



.. _field.gloo.solo.io.DirectResponseAction.status:

status
++++++++++++++++++++++++++

Type: `int` 

Description: Specifies the HTTP response status to be returned. 



.. _field.gloo.solo.io.DirectResponseAction.body:

body
++++++++++++++++++++++++++

Type: `string` 

Description: Specifies the content of the response body. If this setting is omitted, no body is included in the generated response. Note: Headers can be specified using the Header Modification plugin in the enclosing Route, Virtual Host, or Listener. 






.. _message.gloo.solo.io.SslConfig:

SslConfig
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
SslConfig contains the options necessary to configure a virtual host or listener to use TLS


::


   "secret_ref": .core.solo.io.ResourceRef
   "ssl_files": .gloo.solo.io.SSLFiles
   "sni_domains": []string



.. _field.gloo.solo.io.SslConfig.secret_ref:

secret_ref
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.ResourceRef` 

Description: * SecretRef contains the secret ref to a gloo secret containing the following structure: { "tls.crt": <ca chain data...>, "tls.key": <private key data...> } 



.. _field.gloo.solo.io.SslConfig.ssl_files:

ssl_files
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.SSLFiles` 

Description: SSLFiles reference paths to certificates which are local to the proxy 



.. _field.gloo.solo.io.SslConfig.sni_domains:

sni_domains
++++++++++++++++++++++++++

Type: `[]string` 

Description: optional. the SNI domains that should be considered for TLS connections 






.. _message.gloo.solo.io.SSLFiles:

SSLFiles
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
SSLFiles reference paths to certificates which can be read by the proxy off of its local filesystem


::


   "tls_cert": string
   "tls_key": string
   "root_ca": string



.. _field.gloo.solo.io.SSLFiles.tls_cert:

tls_cert
++++++++++++++++++++++++++

Type: `string` 

Description:  



.. _field.gloo.solo.io.SSLFiles.tls_key:

tls_key
++++++++++++++++++++++++++

Type: `string` 

Description:  



.. _field.gloo.solo.io.SSLFiles.root_ca:

root_ca
++++++++++++++++++++++++++

Type: `string` 

Description: for client cert validation. optional 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
