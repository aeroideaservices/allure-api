# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create15**](TestCaseAttachmentControllerApi.md#Create15) | **Post** /testcase/attachment | Upload new test case attachments
[**Delete14**](TestCaseAttachmentControllerApi.md#Delete14) | **Delete** /testcase/attachment/{id} | Delete test case attachment
[**FindAll13**](TestCaseAttachmentControllerApi.md#FindAll13) | **Get** /testcase/attachment | Find attachments for test case
[**Patch13**](TestCaseAttachmentControllerApi.md#Patch13) | **Patch** /testcase/attachment/{id} | Patch test case attachment
[**ReadContent2**](TestCaseAttachmentControllerApi.md#ReadContent2) | **Get** /testcase/attachment/{id}/content | Get attachment content by id
[**UpdateContent2**](TestCaseAttachmentControllerApi.md#UpdateContent2) | **Put** /testcase/attachment/{id}/content | Update test case attachment content

# **Create15**
> []TestCaseAttachmentRowDto Create15(ctx, testCaseId, optional)
Upload new test case attachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 
 **optional** | ***TestCaseAttachmentControllerApiCreate15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseAttachmentControllerApiCreate15Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**[]TestCaseAttachmentRowDto**](TestCaseAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete14**
> Delete14(ctx, id)
Delete test case attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll13**
> PageTestCaseAttachmentRowDto FindAll13(ctx, testCaseId, optional)
Find attachments for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 
 **optional** | ***TestCaseAttachmentControllerApiFindAll13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseAttachmentControllerApiFindAll13Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestCaseAttachmentRowDto**](PageTestCaseAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch13**
> TestCaseAttachmentRowDto Patch13(ctx, body, id)
Patch test case attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseAttachmentPatchDto**](TestCaseAttachmentPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseAttachmentRowDto**](TestCaseAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContent2**
> StreamingResponseBody ReadContent2(ctx, id, optional)
Get attachment content by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestCaseAttachmentControllerApiReadContent2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseAttachmentControllerApiReadContent2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inline** | **optional.Bool**|  | [default to false]

### Return type

[**StreamingResponseBody**](StreamingResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContent2**
> TestCaseAttachmentRowDto UpdateContent2(ctx, id, optional)
Update test case attachment content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestCaseAttachmentControllerApiUpdateContent2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseAttachmentControllerApiUpdateContent2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**TestCaseAttachmentRowDto**](TestCaseAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

