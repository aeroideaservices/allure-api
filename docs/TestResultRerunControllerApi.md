# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Retry**](TestResultRerunControllerApi.md#Retry) | **Post** /testresult/{testResultId}/rerun | Schedule manual rerun for test case
[**Retry1**](TestResultRerunControllerApi.md#Retry1) | **Post** /testresult/{testResultId}/retry | Schedule manual rerun for test case

# **Retry**
> IdAndNameOnlyDto Retry(ctx, body, testResultId)
Schedule manual rerun for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultRerunDto**](TestResultRerunDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**IdAndNameOnlyDto**](IdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Retry1**
> IdAndNameOnlyDto Retry1(ctx, body, testResultId)
Schedule manual rerun for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultRerunDto**](TestResultRerunDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**IdAndNameOnlyDto**](IdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

