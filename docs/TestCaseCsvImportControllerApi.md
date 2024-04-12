# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Info**](TestCaseCsvImportControllerApi.md#Info) | **Post** /testcase/import/csv/{importRequestId}/info | Get testcase csv import file and return import info
[**Preview**](TestCaseCsvImportControllerApi.md#Preview) | **Post** /testcase/import/csv/{importRequestId}/preview | Preview testcase csv import
[**Submit**](TestCaseCsvImportControllerApi.md#Submit) | **Post** /testcase/import/csv/{importRequestId}/submit | Submit testcase csv import

# **Info**
> ImportRequestInfoDto Info(ctx, body, importRequestId)
Get testcase csv import file and return import info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CsvImportOptions**](CsvImportOptions.md)|  | 
  **importRequestId** | **int64**|  | 

### Return type

[**ImportRequestInfoDto**](ImportRequestInfoDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Preview**
> TestCaseOverviewDto Preview(ctx, body, importRequestId)
Preview testcase csv import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseCsvPreviewOptions**](TestCaseCsvPreviewOptions.md)|  | 
  **importRequestId** | **int64**|  | 

### Return type

[**TestCaseOverviewDto**](TestCaseOverviewDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Submit**
> Submit(ctx, body, importRequestId)
Submit testcase csv import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseCsvImportOptions**](TestCaseCsvImportOptions.md)|  | 
  **importRequestId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

