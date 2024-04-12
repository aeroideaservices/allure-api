# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create32**](IssueSchemaControllerApi.md#Create32) | **Post** /issueschema | Create a new issue schema
[**Delete29**](IssueSchemaControllerApi.md#Delete29) | **Delete** /issueschema/{id} | Delete an issue schema by id
[**FindAll29**](IssueSchemaControllerApi.md#FindAll29) | **Get** /issueschema | Find all issue schemas for given project
[**FindOne25**](IssueSchemaControllerApi.md#FindOne25) | **Get** /issueschema/{id} | Find an issue schema by id
[**Patch27**](IssueSchemaControllerApi.md#Patch27) | **Patch** /issueschema/{id} | Patch an issue schema

# **Create32**
> IssueSchemaDto Create32(ctx, body)
Create a new issue schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueSchemaCreateDto**](IssueSchemaCreateDto.md)|  | 

### Return type

[**IssueSchemaDto**](IssueSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete29**
> Delete29(ctx, id)
Delete an issue schema by id

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

# **FindAll29**
> PageIssueSchemaDto FindAll29(ctx, projectId, optional)
Find all issue schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***IssueSchemaControllerApiFindAll29Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueSchemaControllerApiFindAll29Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;key,ASC&quot;]]

### Return type

[**PageIssueSchemaDto**](PageIssueSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne25**
> IssueSchemaDto FindOne25(ctx, id)
Find an issue schema by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**IssueSchemaDto**](IssueSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch27**
> IssueSchemaDto Patch27(ctx, body, id)
Patch an issue schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueSchemaPatchDto**](IssueSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**IssueSchemaDto**](IssueSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

