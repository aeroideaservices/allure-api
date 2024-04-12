# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvVarValues**](TestResultEnvVarControllerApi.md#GetEnvVarValues) | **Get** /testresult/{testResultId}/evv | Find environment variables for test result

# **GetEnvVarValues**
> []EnvVarValueDto GetEnvVarValues(ctx, testResultId)
Find environment variables for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]EnvVarValueDto**](EnvVarValueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

