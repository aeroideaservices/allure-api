# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create39**](EnvVarSchemaControllerApi.md#Create39) | **Post** /evschema | Create a new env var schema
[**Delete35**](EnvVarSchemaControllerApi.md#Delete35) | **Delete** /evschema/{id} | Delete env var schema by id
[**FindAll37**](EnvVarSchemaControllerApi.md#FindAll37) | **Get** /evschema | Find all env var schemas for given project
[**FindOne31**](EnvVarSchemaControllerApi.md#FindOne31) | **Get** /evschema/{id} | Find env var schema by id
[**Patch34**](EnvVarSchemaControllerApi.md#Patch34) | **Patch** /evschema/{id} | Patch env var schema

# **Create39**
> EnvVarSchemaDto Create39(ctx, body)
Create a new env var schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarSchemaCreateDto**](EnvVarSchemaCreateDto.md)|  | 

### Return type

[**EnvVarSchemaDto**](EnvVarSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete35**
> Delete35(ctx, id)
Delete env var schema by id

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

# **FindAll37**
> PageEnvVarSchemaDto FindAll37(ctx, projectId, optional)
Find all env var schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***EnvVarSchemaControllerApiFindAll37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvVarSchemaControllerApiFindAll37Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageEnvVarSchemaDto**](PageEnvVarSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne31**
> EnvVarSchemaDto FindOne31(ctx, id)
Find env var schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**EnvVarSchemaDto**](EnvVarSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch34**
> EnvVarSchemaDto Patch34(ctx, body, id)
Patch env var schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarSchemaPatchDto**](EnvVarSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**EnvVarSchemaDto**](EnvVarSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

