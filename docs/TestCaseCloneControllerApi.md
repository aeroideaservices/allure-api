# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Clone**](TestCaseCloneControllerApi.md#Clone) | **Post** /testcase/{id}/clone | Clone test case

# **Clone**
> TestCaseRowDto Clone(ctx, body, id)
Clone test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseCloneRqDto**](TestCaseCloneRqDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseRowDto**](TestCaseRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

