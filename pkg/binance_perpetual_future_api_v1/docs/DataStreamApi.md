# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewListenKey**](DataStreamApi.md#CreateNewListenKey) | **Post** /fapi/v1/listenKey | Create New Listen Key
[**DeleteListenKey**](DataStreamApi.md#DeleteListenKey) | **Delete** /fapi/v1/listenKey | Delete Listen Key
[**RenewListenKey**](DataStreamApi.md#RenewListenKey) | **Put** /fapi/v1/listenKey | Renew Listen Key

# **CreateNewListenKey**
> CreateNewListenKey(ctx, contentType, xMBXAPIKEY)
Create New Listen Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListenKey**
> DeleteListenKey(ctx, contentType, xMBXAPIKEY)
Delete Listen Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewListenKey**
> RenewListenKey(ctx, contentType, xMBXAPIKEY)
Renew Listen Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

