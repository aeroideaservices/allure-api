# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup1**](TestCaseTreeControllerApi.md#AddGroup1) | **Post** /testcasetree/group | Add a new group (AQL)
[**AddLeaf**](TestCaseTreeControllerApi.md#AddLeaf) | **Post** /testcasetree/leaf | Add a new group
[**CountLeaves2**](TestCaseTreeControllerApi.md#CountLeaves2) | **Get** /testcasetree/countleaves | Count all tree leaves for given path and filter
[**GetGroups21**](TestCaseTreeControllerApi.md#GetGroups21) | **Get** /testcasetree/group | Find tree groups for node (AQL)
[**GetJobsInfo**](TestCaseTreeControllerApi.md#GetJobsInfo) | **Post** /testcasetree/job | Get information about jobs that will be used to run selected test cases
[**GetLeaves1**](TestCaseTreeControllerApi.md#GetLeaves1) | **Get** /testcasetree/leaf | Find tree leaves for node (AQL)
[**GetRunStats**](TestCaseTreeControllerApi.md#GetRunStats) | **Post** /testcasetree/runstats | Get run information
[**RenameGroup1**](TestCaseTreeControllerApi.md#RenameGroup1) | **Post** /testcasetree/group/rename | Rename tree group (AQL)
[**RenameLeaf**](TestCaseTreeControllerApi.md#RenameLeaf) | **Post** /testcasetree/leaf/rename | Rename tree leaf
[**Suggest4**](TestCaseTreeControllerApi.md#Suggest4) | **Get** /testcasetree/suggest | Tree groups suggest

# **AddGroup1**
> TestCaseTreeGroupDto AddGroup1(ctx, body, projectId, optional)
Add a new group (AQL)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeGroupAddDto**](TestCaseTreeGroupAddDto.md)|  | 
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiAddGroup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiAddGroup1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterId** | **optional.**|  | 
 **search** | **optional.**|  | 
 **treeId** | **optional.**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **baseRql** | **optional.**|  | 

### Return type

[**TestCaseTreeGroupDto**](TestCaseTreeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddLeaf**
> TestCaseTreeLeafDto AddLeaf(ctx, body, projectId, optional)
Add a new group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeLeafAddDto**](TestCaseTreeLeafAddDto.md)|  | 
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiAddLeafOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiAddLeafOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **treeId** | **optional.**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]

### Return type

[**TestCaseTreeLeafDto**](TestCaseTreeLeafDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountLeaves2**
> TestCaseTreeFilterCountDto CountLeaves2(ctx, projectId, optional)
Count all tree leaves for given path and filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiCountLeaves2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiCountLeaves2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **filterId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]

### Return type

[**TestCaseTreeFilterCountDto**](TestCaseTreeFilterCountDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups21**
> PageTestCaseTreeGroupDto GetGroups21(ctx, projectId, optional)
Find tree groups for node (AQL)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiGetGroups21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiGetGroups21Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **filterId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 100]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]
 **baseRql** | **optional.String**|  | 

### Return type

[**PageTestCaseTreeGroupDto**](PageTestCaseTreeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobsInfo**
> []JobTestCasesStatDto GetJobsInfo(ctx, body)
Get information about jobs that will be used to run selected test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeRunStatRequestDto**](TestCaseTreeRunStatRequestDto.md)|  | 

### Return type

[**[]JobTestCasesStatDto**](JobTestCasesStatDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeaves1**
> PageTestCaseTreeLeafDto GetLeaves1(ctx, projectId, optional)
Find tree leaves for node (AQL)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiGetLeaves1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiGetLeaves1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **filterId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 100]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]
 **baseRql** | **optional.String**|  | 

### Return type

[**PageTestCaseTreeLeafDto**](PageTestCaseTreeLeafDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunStats**
> TestCaseRunByStats GetRunStats(ctx, body)
Get run information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeRunStatRequestDto**](TestCaseTreeRunStatRequestDto.md)|  | 

### Return type

[**TestCaseRunByStats**](TestCaseRunByStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameGroup1**
> TestCaseTreeGroupDto RenameGroup1(ctx, body, projectId, optional)
Rename tree group (AQL)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeGroupRenameDto**](TestCaseTreeGroupRenameDto.md)|  | 
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiRenameGroup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiRenameGroup1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterId** | **optional.**|  | 
 **search** | **optional.**|  | 
 **treeId** | **optional.**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **baseRql** | **optional.**|  | 

### Return type

[**TestCaseTreeGroupDto**](TestCaseTreeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameLeaf**
> TestCaseTreeLeafDto RenameLeaf(ctx, body, projectId, leafId)
Rename tree leaf

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseTreeLeafRenameDto**](TestCaseTreeLeafRenameDto.md)|  | 
  **projectId** | **int64**|  | 
  **leafId** | **int64**|  | 

### Return type

[**TestCaseTreeLeafDto**](TestCaseTreeLeafDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest4**
> PageIdAndNameOnlyDto Suggest4(ctx, projectId, optional)
Tree groups suggest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseTreeControllerApiSuggest4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseTreeControllerApiSuggest4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **treeId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **id** | [**optional.Interface of []int64**](int64.md)|  | 
 **ignoreId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageIdAndNameOnlyDto**](PageIdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

