# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFixtures**](TestResultFixtureControllerApi.md#GetFixtures) | **Get** /testresult/{testResultId}/fixture | Find fixtures for test result
[**GetFixturesAttachments**](TestResultFixtureControllerApi.md#GetFixturesAttachments) | **Get** /testresult/{testResultId}/fixture/attachment | Find fixtures attachments for test result

# **GetFixtures**
> []TestFixtureResultDto GetFixtures(ctx, testResultId)
Find fixtures for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]TestFixtureResultDto**](TestFixtureResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFixturesAttachments**
> PageTestFixtureResultAttachmentRowDto GetFixturesAttachments(ctx, testResultId, optional)
Find fixtures attachments for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
 **optional** | ***TestResultFixtureControllerApiGetFixturesAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultFixtureControllerApiGetFixturesAttachmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestFixtureResultAttachmentRowDto**](PageTestFixtureResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

