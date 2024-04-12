# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create46**](CleanerSchemaControllerApi.md#Create46) | **Post** /cleanerschema | Create a new cleaner schema
[**Delete42**](CleanerSchemaControllerApi.md#Delete42) | **Delete** /cleanerschema/{id} | Delete cleaner schema by id
[**FindAll41**](CleanerSchemaControllerApi.md#FindAll41) | **Get** /cleanerschema | Find all cleaner schemas for given project
[**FindAllGlobalSchemas**](CleanerSchemaControllerApi.md#FindAllGlobalSchemas) | **Get** /cleanerschema/global | Find all global cleaner schemas
[**FindOne36**](CleanerSchemaControllerApi.md#FindOne36) | **Get** /cleanerschema/{id} | Find cleaner schema by id
[**Patch41**](CleanerSchemaControllerApi.md#Patch41) | **Patch** /cleanerschema/{id} | Patch cleaner schema

# **Create46**
> CleanerSchemaDto Create46(ctx, body)
Create a new cleaner schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CleanerSchemaCreateDto**](CleanerSchemaCreateDto.md)|  | 

### Return type

[**CleanerSchemaDto**](CleanerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete42**
> Delete42(ctx, id)
Delete cleaner schema by id

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

# **FindAll41**
> PageCleanerSchemaDto FindAll41(ctx, projectId, optional)
Find all cleaner schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***CleanerSchemaControllerApiFindAll41Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CleanerSchemaControllerApiFindAll41Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;id,ASC&quot;]]

### Return type

[**PageCleanerSchemaDto**](PageCleanerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllGlobalSchemas**
> PageCleanerSchemaDto FindAllGlobalSchemas(ctx, optional)
Find all global cleaner schemas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CleanerSchemaControllerApiFindAllGlobalSchemasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CleanerSchemaControllerApiFindAllGlobalSchemasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;id,ASC&quot;]]

### Return type

[**PageCleanerSchemaDto**](PageCleanerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne36**
> CleanerSchemaDto FindOne36(ctx, id)
Find cleaner schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**CleanerSchemaDto**](CleanerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch41**
> CleanerSchemaDto Patch41(ctx, body, id)
Patch cleaner schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CleanerSchemaPatchDto**](CleanerSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**CleanerSchemaDto**](CleanerSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

