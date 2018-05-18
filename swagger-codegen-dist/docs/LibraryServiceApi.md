# \LibraryServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBook**](LibraryServiceApi.md#CreateBook) | **Post** /v1/{parent}/books | 
[**DeleteBook**](LibraryServiceApi.md#DeleteBook) | **Delete** /v1/{name} | 
[**GetBook**](LibraryServiceApi.md#GetBook) | **Get** /v1/{name} | 
[**ListBooks**](LibraryServiceApi.md#ListBooks) | **Get** /v1/{parent}/books | ListBooks は本を列挙します
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

[OAuth2](../README.md#OAuth2)

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

[OAuth2](../README.md#OAuth2)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBooks**
> V1ListBooksResponse ListBooks(ctx, parent, optional)
ListBooks は本を列挙します

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parent** | **string**|  | 
 **pageSize** | **int32**| The maximum number of items to return. | 
 **pageToken** | **string**| The next_page_token value returned from a previous List request, if any. | 

### Return type

[**V1ListBooksResponse**](v1ListBooksResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

