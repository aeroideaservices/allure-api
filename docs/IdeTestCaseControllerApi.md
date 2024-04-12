# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScenario1**](IdeTestCaseControllerApi.md#GetScenario1) | **Get** /ide/testcase/{id}/scenario | Get scenario for test case

# **GetScenario1**
> IdeScenarioDto GetScenario1(ctx, id)
Get scenario for test case

To use with IDE plugins only

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**IdeScenarioDto**](IdeScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

