# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create48**](CustomFieldSchemaControllerApi.md#Create48) | **Post** /cfschema | Create a new custom field schema
[**Delete44**](CustomFieldSchemaControllerApi.md#Delete44) | **Delete** /cfschema/{id} | Delete custom field schema by id
[**FindAll43**](CustomFieldSchemaControllerApi.md#FindAll43) | **Get** /cfschema | Find all custom field schemas for given project
[**FindOne38**](CustomFieldSchemaControllerApi.md#FindOne38) | **Get** /cfschema/{id} | Find custom field schema by id
[**Patch43**](CustomFieldSchemaControllerApi.md#Patch43) | **Patch** /cfschema/{id} | Patch custom field schema

# **Create48**
> CustomFieldSchemaDto Create48(ctx, body)
Create a new custom field schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomFieldSchemaCreateDto**](CustomFieldSchemaCreateDto.md)|  | 

### Return type

[**CustomFieldSchemaDto**](CustomFieldSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete44**
> Delete44(ctx, id)
Delete custom field schema by id

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

# **FindAll43**
> PageCustomFieldSchemaDto FindAll43(ctx, projectId, optional)
Find all custom field schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***CustomFieldSchemaControllerApiFindAll43Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomFieldSchemaControllerApiFindAll43Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;id,ASC&quot;]]

### Return type

[**PageCustomFieldSchemaDto**](PageCustomFieldSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne38**
> CustomFieldSchemaDto FindOne38(ctx, id)
Find custom field schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**CustomFieldSchemaDto**](CustomFieldSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch43**
> CustomFieldSchemaDto Patch43(ctx, body, id)
Patch custom field schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomFieldSchemaPatchDto**](CustomFieldSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**CustomFieldSchemaDto**](CustomFieldSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

