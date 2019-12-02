# \V2Api

All URIs are relative to *https://api.ote-godaddy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsForwardsDelete**](V2Api.md#DomainsForwardsDelete) | **Delete** /v2/customers/{customerId}/domains/forwards/{fqdn} | Submit a forwarding cancellation request for the given fqdn
[**DomainsForwardsGet**](V2Api.md#DomainsForwardsGet) | **Get** /v2/customers/{customerId}/domains/forwards/{fqdn} | Retrieve the forwarding information for the given fqdn
[**DomainsForwardsPut**](V2Api.md#DomainsForwardsPut) | **Put** /v2/customers/{customerId}/domains/forwards/{fqdn} | Modify the forwarding information for the given fqdn


# **DomainsForwardsDelete**
> DomainsForwardsDelete(ctx, customerId, fqdn)
Submit a forwarding cancellation request for the given fqdn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsForwardsGet**
> []DomainForwarding DomainsForwardsGet(ctx, customerId, fqdn, optional)
Retrieve the forwarding information for the given fqdn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be retrieved. | 
 **optional** | ***DomainsForwardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsForwardsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSubs** | **optional.Bool**| Optionally include all sub domains if the fqdn specified is a domain and not a sub domain. | 

### Return type

[**[]DomainForwarding**](DomainForwarding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsForwardsPut**
> DomainsForwardsPut(ctx, customerId, fqdn, body)
Modify the forwarding information for the given fqdn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be modified. | 
  **body** | [**DomainForwardingCreate**](DomainForwardingCreate.md)| Domain forwarding rule to create or replace on the fqdn | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

