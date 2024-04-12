# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Assign2**](TestPlanControllerApi.md#Assign2) | **Post** /testplan/{id}/assign | Assign test plan test cases to user
[**Create7**](TestPlanControllerApi.md#Create7) | **Post** /testplan | Create a new test plan
[**Delete7**](TestPlanControllerApi.md#Delete7) | **Delete** /testplan/{id} | Delete test plan by given id
[**FindAllByProject**](TestPlanControllerApi.md#FindAllByProject) | **Get** /testplan | Find all test plans for given project
[**FindOne5**](TestPlanControllerApi.md#FindOne5) | **Get** /testplan/{id} | Find test plan by id
[**GetDiff**](TestPlanControllerApi.md#GetDiff) | **Get** /testplan/{id}/diff | Get test plan test cases changes
[**GetGroups1**](TestPlanControllerApi.md#GetGroups1) | **Get** /testplan/{id}/tree/group | Find tree groups for node
[**GetJobs**](TestPlanControllerApi.md#GetJobs) | **Get** /testplan/{id}/job | Get test plan jobs statistic
[**GetLeafs1**](TestPlanControllerApi.md#GetLeafs1) | **Get** /testplan/{id}/tree/leaf | Find tree leafs for node
[**GetMembers1**](TestPlanControllerApi.md#GetMembers1) | **Get** /testplan/{id}/member | Get test plan members statistic
[**Patch6**](TestPlanControllerApi.md#Patch6) | **Patch** /testplan/{id} | Patch test plan
[**ResetJobs**](TestPlanControllerApi.md#ResetJobs) | **Post** /testplan/{id}/resetjob | Reset test plan
[**Run**](TestPlanControllerApi.md#Run) | **Post** /testplan/{id}/run | Run test plan by given id
[**SetJobParameters**](TestPlanControllerApi.md#SetJobParameters) | **Post** /testplan/{id}/jobparameter | Configure test plan job parameters
[**Suggest2**](TestPlanControllerApi.md#Suggest2) | **Get** /testplan/suggest | Suggest for test plans
[**Sync**](TestPlanControllerApi.md#Sync) | **Post** /testplan/{id}/sync | Sync test plan

# **Assign2**
> Assign2(ctx, body, id)
Assign test plan test cases to user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestPlanAssignDto**](TestPlanAssignDto.md)|  | 
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create7**
> TestPlanDto Create7(ctx, body)
Create a new test plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestPlanCreateDto**](TestPlanCreateDto.md)|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete7**
> Delete7(ctx, id)
Delete test plan by given id

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

# **FindAllByProject**
> PageTestPlanDto FindAllByProject(ctx, projectId, optional)
Find all test plans for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestPlanControllerApiFindAllByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestPlanControllerApiFindAllByProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestPlanDto**](PageTestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne5**
> TestPlanDto FindOne5(ctx, id)
Find test plan by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiff**
> TestPlanDiffDto GetDiff(ctx, id)
Get test plan test cases changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestPlanDiffDto**](TestPlanDiffDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups1**
> PageTestCaseTreeGroupDto GetGroups1(ctx, id, optional)
Find tree groups for node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestPlanControllerApiGetGroups1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestPlanControllerApiGetGroups1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **treeId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **username** | **optional.String**|  | 
 **jobId** | **optional.Int64**|  | 
 **manual** | **optional.Bool**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestCaseTreeGroupDto**](PageTestCaseTreeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobs**
> []TestPlanJobStatDto GetJobs(ctx, id)
Get test plan jobs statistic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**[]TestPlanJobStatDto**](TestPlanJobStatDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeafs1**
> PageTestCaseTreeLeafDto GetLeafs1(ctx, id, optional)
Find tree leafs for node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestPlanControllerApiGetLeafs1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestPlanControllerApiGetLeafs1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **treeId** | **optional.Int64**|  | 
 **path** | [**optional.Interface of []int64**](int64.md)|  | [default to []]
 **username** | **optional.String**|  | 
 **jobId** | **optional.Int64**|  | 
 **manual** | **optional.Bool**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestCaseTreeLeafDto**](PageTestCaseTreeLeafDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMembers1**
> []TestPlanMemberStatDto GetMembers1(ctx, id)
Get test plan members statistic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**[]TestPlanMemberStatDto**](TestPlanMemberStatDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch6**
> TestPlanDto Patch6(ctx, body, id)
Patch test plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestPlanPatchDto**](TestPlanPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetJobs**
> TestPlanDto ResetJobs(ctx, id)
Reset test plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Run**
> LaunchDto Run(ctx, body, id)
Run test plan by given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestPlanRunRequestDto**](TestPlanRunRequestDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**LaunchDto**](LaunchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetJobParameters**
> SetJobParameters(ctx, body, id)
Configure test plan job parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestPlanJobParametersDto**](TestPlanJobParametersDto.md)|  | 
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest2**
> PageTestPlanRowDto Suggest2(ctx, optional)
Suggest for test plans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TestPlanControllerApiSuggest2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestPlanControllerApiSuggest2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **projectId** | **optional.Int64**|  | 
 **id** | [**optional.Interface of []int64**](int64.md)|  | 
 **ignoreId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestPlanRowDto**](PageTestPlanRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Sync**
> TestPlanDto Sync(ctx, id)
Sync test plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

