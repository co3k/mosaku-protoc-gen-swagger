# \LibraryServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBook**](LibraryServiceApi.md#CreateBook) | **Post** /v1/{parent}/books | 
[**DeleteBook**](LibraryServiceApi.md#DeleteBook) | **Delete** /v1/{name} | 
[**GetBook**](LibraryServiceApi.md#GetBook) | **Get** /v1/{name} | 
[**UpdateBook**](LibraryServiceApi.md#UpdateBook) | **Patch** /v1/{book.name} | 


# **CreateBook**
> V1Book CreateBook(ctx, parent, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**|  | 
  **body** | [**V1Book**](V1Book.md)|  | 

### Return type

[**V1Book**](v1Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBook**
> ProtobufEmpty DeleteBook(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 

### Return type

[**ProtobufEmpty**](protobufEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBook**
> V1Book GetBook(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 

### Return type

[**V1Book**](v1Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBook**
> V1Book UpdateBook(ctx, bookName, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bookName** | **string**|  | 
  **body** | [**V1Book**](V1Book.md)|  | 

### Return type

[**V1Book**](v1Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

