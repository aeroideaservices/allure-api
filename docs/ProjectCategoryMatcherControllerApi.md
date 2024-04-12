# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](ProjectCategoryMatcherControllerApi.md#Add) | **Post** /project/{projectId}/categorymatcher | 
[**FindAll22**](ProjectCategoryMatcherControllerApi.md#FindAll22) | **Get** /project/{projectId}/categorymatcher | 
[**Remove**](ProjectCategoryMatcherControllerApi.md#Remove) | **Post** /project/{projectId}/categorymatcher/remove | 

# **Add**
> CategoryMatcherDto Add(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectCategoryMatcherAddDto**](ProjectCategoryMatcherAddDto.md)|  | 
  **projectId** | **int64**|  | 

### Return type

[**CategoryMatcherDto**](CategoryMatcherDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll22**
> PageCategoryMatcherDto FindAll22(ctx, projectId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***ProjectCategoryMatcherControllerApiFindAll22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCategoryMatcherControllerApiFindAll22Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCategoryMatcherDto**](PageCategoryMatcherDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Remove**
> Remove(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectCategoryMatcherRemoveDto**](ProjectCategoryMatcherRemoveDto.md)|  | 
  **projectId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

