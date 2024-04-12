# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCfv1**](TestCaseCustomFieldControllerApi.md#GetCfv1) | **Get** /testcase/{testCaseId}/cfv | Find custom field values for test case
[**SetCfv**](TestCaseCustomFieldControllerApi.md#SetCfv) | **Post** /testcase/{testCaseId}/cfv | Set custom field values for test case

# **GetCfv1**
> []CustomFieldValueDto GetCfv1(ctx, testCaseId)
Find custom field values for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]CustomFieldValueDto**](CustomFieldValueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetCfv**
> []CustomFieldValueDto SetCfv(ctx, body, testCaseId)
Set custom field values for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CustomFieldValueDto**](CustomFieldValueDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]CustomFieldValueDto**](CustomFieldValueDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

