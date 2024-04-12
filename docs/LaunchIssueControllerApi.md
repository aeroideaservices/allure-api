# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Export**](LaunchIssueControllerApi.md#Export) | **Post** /launch/{launchId}/issue/export | Export launch data to issue issueTracker
[**GetIssues2**](LaunchIssueControllerApi.md#GetIssues2) | **Get** /launch/issue | Get all issues used in launches

# **Export**
> Export(ctx, body, launchId)
Export launch data to issue issueTracker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]IssueDto**](IssueDto.md)|  | 
  **launchId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIssues2**
> []IssueDto GetIssues2(ctx, projectId)
Get all issues used in launches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**[]IssueDto**](IssueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

