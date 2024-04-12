# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssues**](TestResultIssueControllerApi.md#GetIssues) | **Get** /testresult/{testResultId}/issue | Find issues for test result
[**SetIssues**](TestResultIssueControllerApi.md#SetIssues) | **Post** /testresult/{testResultId}/issue | Set issues to test result

# **GetIssues**
> []IssueDto GetIssues(ctx, testResultId)
Find issues for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]IssueDto**](IssueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetIssues**
> []IssueDto SetIssues(ctx, body, testResultId)
Set issues to test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]IssueDto**](IssueDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**[]IssueDto**](IssueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

