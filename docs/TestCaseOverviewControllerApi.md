# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOverview**](TestCaseOverviewControllerApi.md#GetOverview) | **Get** /testcase/{testCaseId}/overview | Get test case overview

# **GetOverview**
> TestCaseOverviewDto GetOverview(ctx, testCaseId)
Get test case overview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**TestCaseOverviewDto**](TestCaseOverviewDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

