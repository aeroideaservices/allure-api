# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindById1**](JobRunTestPlanControllerApi.md#FindById1) | **Get** /jobrun/{id}/plan | Find test plan for execution by external id

# **FindById1**
> []TestCaseInfo FindById1(ctx, id, optional)
Find test plan for execution by external id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***JobRunTestPlanControllerApiFindById1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobRunTestPlanControllerApiFindById1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expected** | **optional.Bool**|  | [default to false]

### Return type

[**[]TestCaseInfo**](TestCaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

