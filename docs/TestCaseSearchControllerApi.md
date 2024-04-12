# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search1**](TestCaseSearchControllerApi.md#Search1) | **Get** /testcase/__search | Find all test cases by given AQL
[**ValidateQuery1**](TestCaseSearchControllerApi.md#ValidateQuery1) | **Get** /testcase/query/validate | Find all test cases by given AQL

# **Search1**
> PageTestCaseDto Search1(ctx, projectId, rql, optional)
Find all test cases by given AQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **rql** | **string**|  | 
 **optional** | ***TestCaseSearchControllerApiSearch1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseSearchControllerApiSearch1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **optional.Bool**|  | [default to false]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestCaseDto**](PageTestCaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateQuery1**
> AqlValidateResponseDto ValidateQuery1(ctx, projectId, rql, optional)
Find all test cases by given AQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **rql** | **string**|  | 
 **optional** | ***TestCaseSearchControllerApiValidateQuery1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseSearchControllerApiValidateQuery1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **optional.Bool**|  | [default to false]

### Return type

[**AqlValidateResponseDto**](AqlValidateResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

