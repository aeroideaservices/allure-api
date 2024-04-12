# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Generate**](TestCaseExampleControllerApi.md#Generate) | **Post** /testcase/example/nwise | 
[**GetExamples**](TestCaseExampleControllerApi.md#GetExamples) | **Get** /testcase/{testCaseId}/example | 
[**SetExamples**](TestCaseExampleControllerApi.md#SetExamples) | **Post** /testcase/{testCaseId}/example | 

# **Generate**
> []TestCaseExampleDto Generate(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TestCaseParameterValues**](TestCaseParameterValues.md)|  | 
 **optional** | ***TestCaseExampleControllerApiGenerateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseExampleControllerApiGenerateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **optional.**|  | [default to 1]

### Return type

[**[]TestCaseExampleDto**](TestCaseExampleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExamples**
> PageTestCaseExampleDto GetExamples(ctx, testCaseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 
 **optional** | ***TestCaseExampleControllerApiGetExamplesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseExampleControllerApiGetExamplesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;id,ASC&quot;]]

### Return type

[**PageTestCaseExampleDto**](PageTestCaseExampleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetExamples**
> []TestCaseExampleDto SetExamples(ctx, body, testCaseId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[][]ParameterValueDto**](array.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]TestCaseExampleDto**](TestCaseExampleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

