# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search2**](LaunchSearchControllerApi.md#Search2) | **Get** /launch/__search | Find all launches by given AQL
[**ValidateQuery2**](LaunchSearchControllerApi.md#ValidateQuery2) | **Get** /launch/query/validate | Find all launches by given AQL

# **Search2**
> PageLaunchDto Search2(ctx, projectId, rql, optional)
Find all launches by given AQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **rql** | **string**|  | 
 **optional** | ***LaunchSearchControllerApiSearch2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchSearchControllerApiSearch2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;created_date,DESC&quot;]]

### Return type

[**PageLaunchDto**](PageLaunchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateQuery2**
> AqlValidateResponseDto ValidateQuery2(ctx, projectId, rql)
Find all launches by given AQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **rql** | **string**|  | 

### Return type

[**AqlValidateResponseDto**](AqlValidateResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

