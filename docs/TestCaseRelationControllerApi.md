# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelations**](TestCaseRelationControllerApi.md#GetRelations) | **Get** /testcase/{testCaseId}/relation | Find relations for test case
[**SetRelations**](TestCaseRelationControllerApi.md#SetRelations) | **Post** /testcase/{testCaseId}/relation | Set relations for test case

# **GetRelations**
> []TestCaseRelationDto GetRelations(ctx, testCaseId)
Find relations for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestCaseRelationDto**](TestCaseRelationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRelations**
> []TestCaseRelationDto SetRelations(ctx, body, testCaseId)
Set relations for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TestCaseRelationDto**](TestCaseRelationDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestCaseRelationDto**](TestCaseRelationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

