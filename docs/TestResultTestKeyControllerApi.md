# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKeys**](TestResultTestKeyControllerApi.md#GetKeys) | **Get** /testresult/{testResultId}/testkey | Find test keys for test result
[**SetKeys**](TestResultTestKeyControllerApi.md#SetKeys) | **Post** /testresult/{testResultId}/testkey | Set test keys to test result

# **GetKeys**
> []TestKeyDto GetKeys(ctx, testResultId)
Find test keys for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]TestKeyDto**](TestKeyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetKeys**
> []TestKeyDto SetKeys(ctx, body, testResultId)
Set test keys to test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TestKeyDto**](TestKeyDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**[]TestKeyDto**](TestKeyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

