# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroups**](TestResultTreeControllerApi.md#GetGroups) | **Get** /testresulttree/group | Find tree groups for node
[**GetLeafs**](TestResultTreeControllerApi.md#GetLeafs) | **Get** /testresulttree/leaf | Find tree leafs for node

# **GetGroups**
> PageTestResultTreeGroupDto GetGroups(ctx, launchId, optional)
Find tree groups for node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***TestResultTreeControllerApiGetGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultTreeControllerApiGetGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestResultTreeGroupDto**](PageTestResultTreeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeafs**
> PageTestResultTreeLeafDto GetLeafs(ctx, launchId, optional)
Find tree leafs for node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***TestResultTreeControllerApiGetLeafsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultTreeControllerApiGetLeafsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestResultTreeLeafDto**](PageTestResultTreeLeafDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

