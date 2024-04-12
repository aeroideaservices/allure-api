# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create8**](TestLayerSchemaControllerApi.md#Create8) | **Post** /testlayerschema | Create a new test layer schema
[**Delete8**](TestLayerSchemaControllerApi.md#Delete8) | **Delete** /testlayerschema/{id} | Delete test layer schema by id
[**FindAll6**](TestLayerSchemaControllerApi.md#FindAll6) | **Get** /testlayerschema | Find all test layer schemas for given project
[**FindOne6**](TestLayerSchemaControllerApi.md#FindOne6) | **Get** /testlayerschema/{id} | Find test layer schema by id
[**Patch7**](TestLayerSchemaControllerApi.md#Patch7) | **Patch** /testlayerschema/{id} | Patch test layer schema

# **Create8**
> TestLayerSchemaDto Create8(ctx, body)
Create a new test layer schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestLayerSchemaCreateDto**](TestLayerSchemaCreateDto.md)|  | 

### Return type

[**TestLayerSchemaDto**](TestLayerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete8**
> Delete8(ctx, id)
Delete test layer schema by id

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

# **FindAll6**
> PageTestLayerSchemaDto FindAll6(ctx, projectId, optional)
Find all test layer schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestLayerSchemaControllerApiFindAll6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestLayerSchemaControllerApiFindAll6Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;id,ASC&quot;]]

### Return type

[**PageTestLayerSchemaDto**](PageTestLayerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne6**
> TestLayerSchemaDto FindOne6(ctx, id)
Find test layer schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestLayerSchemaDto**](TestLayerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch7**
> TestLayerSchemaDto Patch7(ctx, body, id)
Patch test layer schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestLayerSchemaPatchDto**](TestLayerSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestLayerSchemaDto**](TestLayerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

