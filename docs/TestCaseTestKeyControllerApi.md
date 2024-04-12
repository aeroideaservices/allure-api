# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKeys1**](TestCaseTestKeyControllerApi.md#GetKeys1) | **Get** /testcase/{testCaseId}/testkey | Find test keys for test case
[**SetKeys1**](TestCaseTestKeyControllerApi.md#SetKeys1) | **Post** /testcase/{testCaseId}/testkey | Set test keys to test case

# **GetKeys1**
> []TestKeyDto GetKeys1(ctx, testCaseId)
Find test keys for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestKeyDto**](TestKeyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetKeys1**
> []TestKeyDto SetKeys1(ctx, body, testCaseId)
Set test keys to test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TestKeyDto**](TestKeyDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestKeyDto**](TestKeyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

