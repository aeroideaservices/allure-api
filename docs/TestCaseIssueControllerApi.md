# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssues1**](TestCaseIssueControllerApi.md#GetIssues1) | **Get** /testcase/{testCaseId}/issue | Find issues for test case
[**SetIssues2**](TestCaseIssueControllerApi.md#SetIssues2) | **Post** /testcase/{testCaseId}/issue | Set issues to test case

# **GetIssues1**
> []IssueDto GetIssues1(ctx, testCaseId)
Find issues for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]IssueDto**](IssueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetIssues2**
> []IssueDto SetIssues2(ctx, body, testCaseId)
Set issues to test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]IssueDto**](IssueDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]IssueDto**](IssueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

