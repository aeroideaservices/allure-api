# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create13**](TestCaseControllerApi.md#Create13) | **Post** /testcase | Create a new test case
[**Delete13**](TestCaseControllerApi.md#Delete13) | **Delete** /testcase/{id} | Delete test case by id
[**DetachAutomation**](TestCaseControllerApi.md#DetachAutomation) | **Post** /testcase/{id}/detachautomation | Detach automation from test case
[**FindAll11**](TestCaseControllerApi.md#FindAll11) | **Get** /testcase | Find all test cases for specified project
[**FindAllDeleted**](TestCaseControllerApi.md#FindAllDeleted) | **Get** /testcase/deleted | Find all deleted test cases for given project
[**FindAllMuted**](TestCaseControllerApi.md#FindAllMuted) | **Get** /testcase/muted | Find all muted test cases for given project
[**FindHistory1**](TestCaseControllerApi.md#FindHistory1) | **Get** /testcase/{id}/history | Find run history for test case
[**FindHistory2**](TestCaseControllerApi.md#FindHistory2) | **Get** /testcase/history | Find run history for test case
[**FindOne11**](TestCaseControllerApi.md#FindOne11) | **Get** /testcase/{id} | Find test case by id
[**FindWorkflow**](TestCaseControllerApi.md#FindWorkflow) | **Get** /testcase/{id}/workflow | Find workflow for test case
[**Patch12**](TestCaseControllerApi.md#Patch12) | **Patch** /testcase/{id} | 
[**Restore**](TestCaseControllerApi.md#Restore) | **Post** /testcase/{id}/restore | Restore test case by id
[**Suggest5**](TestCaseControllerApi.md#Suggest5) | **Get** /testcase/suggest | Find suggest for test case

# **Create13**
> TestCaseDto Create13(ctx, body)
Create a new test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseCreateDto**](TestCaseCreateDto.md)|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete13**
> Delete13(ctx, id, optional)
Delete test case by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestCaseControllerApiDelete13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiDelete13Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachAutomation**
> TestCaseDto DetachAutomation(ctx, body, id)
Detach automation from test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseDetachAutomationRqDto**](TestCaseDetachAutomationRqDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll11**
> PageTestCaseRowDto FindAll11(ctx, projectId, optional)
Find all test cases for specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseControllerApiFindAll11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiFindAll11Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestCaseRowDto**](PageTestCaseRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllDeleted**
> PageTestCaseRowDto FindAllDeleted(ctx, projectId, optional)
Find all deleted test cases for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseControllerApiFindAllDeletedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiFindAllDeletedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestCaseRowDto**](PageTestCaseRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllMuted**
> PageTestCaseRowDto FindAllMuted(ctx, projectId, optional)
Find all muted test cases for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***TestCaseControllerApiFindAllMutedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiFindAllMutedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestCaseRowDto**](PageTestCaseRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindHistory1**
> PageTestResultHistoryDto FindHistory1(ctx, id, optional)
Find run history for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestCaseControllerApiFindHistory1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiFindHistory1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestResultHistoryDto**](PageTestResultHistoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindHistory2**
> PageTestResultHistoryDto FindHistory2(ctx, testCaseId, optional)
Find run history for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 
 **optional** | ***TestCaseControllerApiFindHistory2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiFindHistory2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **optional.Int64**|  | 
 **launchId** | **optional.Int64**|  | 
 **testResultId** | **optional.Int64**|  | 
 **search** | **optional.String**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestResultHistoryDto**](PageTestResultHistoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne11**
> TestCaseDto FindOne11(ctx, id)
Find test case by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindWorkflow**
> WorkflowDto FindWorkflow(ctx, id)
Find workflow for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**WorkflowDto**](WorkflowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch12**
> TestCaseDto Patch12(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCasePatchDto**](TestCasePatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Restore**
> TestCaseDto Restore(ctx, id)
Restore test case by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest5**
> PageIdAndNameOnlyDto Suggest5(ctx, optional)
Find suggest for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TestCaseControllerApiSuggest5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseControllerApiSuggest5Opts struct
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

[**PageIdAndNameOnlyDto**](PageIdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

