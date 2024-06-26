# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTags**](TestCaseTagControllerApi.md#GetTags) | **Get** /testcase/{testCaseId}/tag | Find tags for test case
[**SetTags**](TestCaseTagControllerApi.md#SetTags) | **Post** /testcase/{testCaseId}/tag | Set test tags for test case

# **GetTags**
> []TestTagDto GetTags(ctx, testCaseId)
Find tags for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestTagDto**](TestTagDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTags**
> []TestTagDto SetTags(ctx, body, testCaseId)
Set test tags for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TestTagDto**](TestTagDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestTagDto**](TestTagDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

