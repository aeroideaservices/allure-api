# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCfv**](TestResultCustomFieldControllerApi.md#GetCfv) | **Get** /testresult/{testResultId}/cfv | Find custom field values for test result
[**SetIssues1**](TestResultCustomFieldControllerApi.md#SetIssues1) | **Post** /testresult/{testResultId}/cfv | Set custom field values to test result

# **GetCfv**
> []CustomFieldValueDto GetCfv(ctx, testResultId)
Find custom field values for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]CustomFieldValueDto**](CustomFieldValueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetIssues1**
> []CustomFieldValueDto SetIssues1(ctx, body, testResultId)
Set custom field values to test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CustomFieldValueDto**](CustomFieldValueDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**[]CustomFieldValueDto**](CustomFieldValueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

