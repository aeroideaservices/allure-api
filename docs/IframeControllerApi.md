# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateTestCases**](IframeControllerApi.md#BulkCreateTestCases) | **Post** /iframe/testcases | Bulk create a new test cases from iframe
[**CreateTestCase**](IframeControllerApi.md#CreateTestCase) | **Post** /iframe/testcase | Create a new test case from iframe
[**GetLaunches**](IframeControllerApi.md#GetLaunches) | **Get** /iframe/launch | 
[**GetStatistic1**](IframeControllerApi.md#GetStatistic1) | **Get** /iframe/launch/{launchId}/statistic | Get launch statistic
[**GetTestCases**](IframeControllerApi.md#GetTestCases) | **Get** /iframe/testcase | Get pageble list of testcases
[**GetTestResults**](IframeControllerApi.md#GetTestResults) | **Get** /iframe/testresult | Get test results for launch
[**LinkTestCaseWithIssue**](IframeControllerApi.md#LinkTestCaseWithIssue) | **Post** /iframe/testcases/linkissue | Link test cases with issue from iframe

# **BulkCreateTestCases**
> []TestCaseDto BulkCreateTestCases(ctx, body)
Bulk create a new test cases from iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IframeBulkCreateTestCaseDto**](IframeBulkCreateTestCaseDto.md)|  | 

### Return type

[**[]TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTestCase**
> TestCaseDto CreateTestCase(ctx, body)
Create a new test case from iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IframeCreateTestCaseDto**](IframeCreateTestCaseDto.md)|  | 

### Return type

[**TestCaseDto**](TestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaunches**
> PageLaunchPreviewDto GetLaunches(ctx, issueKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueKey** | **string**|  | 
 **optional** | ***IframeControllerApiGetLaunchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IframeControllerApiGetLaunchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **optional.Int64**|  | 
 **issueTrackerId** | **optional.Int64**|  | 
 **closed** | **optional.Bool**|  | 
 **withJob** | **optional.Bool**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;created_date,DESC&quot;]]

### Return type

[**PageLaunchPreviewDto**](PageLaunchPreviewDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatistic1**
> []TestStatusCount GetStatistic1(ctx, launchId)
Get launch statistic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 

### Return type

[**[]TestStatusCount**](TestStatusCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestCases**
> PageTestCaseDto GetTestCases(ctx, issueKey, optional)
Get pageble list of testcases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueKey** | **string**|  | 
 **optional** | ***IframeControllerApiGetTestCasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IframeControllerApiGetTestCasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **optional.Int64**|  | 
 **issueTrackerId** | **optional.Int64**|  | 
 **search** | **optional.String**|  | 
 **status** | [**optional.Interface of []int64**](int64.md)|  | 
 **automated** | **optional.Bool**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestCaseDto**](PageTestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestResults**
> PageTestResultDto GetTestResults(ctx, launchId, optional)
Get test results for launch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***IframeControllerApiGetTestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IframeControllerApiGetTestResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []TestStatus**](TestStatus.md)|  | 
 **manual** | **optional.Bool**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestResultDto**](PageTestResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkTestCaseWithIssue**
> LinkTestCaseWithIssue(ctx, body)
Link test cases with issue from iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IframeTestCaseWithIssueDto**](IframeTestCaseWithIssueDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

