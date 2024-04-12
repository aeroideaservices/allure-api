# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create10**](TestKeySchemaControllerApi.md#Create10) | **Post** /testkeyschema | Create a new test key schema
[**Delete10**](TestKeySchemaControllerApi.md#Delete10) | **Delete** /testkeyschema/{id} | Delete a test key schema by id
[**FindAll8**](TestKeySchemaControllerApi.md#FindAll8) | **Get** /testkeyschema | Find all test key schemas for given project
[**FindOne8**](TestKeySchemaControllerApi.md#FindOne8) | **Get** /testkeyschema/{id} | Find a test key schema by id
[**Patch9**](TestKeySchemaControllerApi.md#Patch9) | **Patch** /testkeyschema/{id} | Patch a test key schema

# **Create10**
> TestKeySchemaDto Create10(ctx, body)
Create a new test key schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestKeySchemaCreateDto**](TestKeySchemaCreateDto.md)|  | 

### Return type

[**TestKeySchemaDto**](TestKeySchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete10**
> Delete10(ctx, id)
Delete a test key schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll8**
> PageTestKeySchemaDto FindAll8(ctx, projectId, optional)
Find all test key schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestKeySchemaControllerApiFindAll8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestKeySchemaControllerApiFindAll8Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;key,ASC&quot;]]

### Return type

[**PageTestKeySchemaDto**](PageTestKeySchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne8**
> TestKeySchemaDto FindOne8(ctx, id)
Find a test key schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestKeySchemaDto**](TestKeySchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch9**
> TestKeySchemaDto Patch9(ctx, body, id)
Patch a test key schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestKeySchemaPatchDto**](TestKeySchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestKeySchemaDto**](TestKeySchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

