# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create6**](TestResultAttachmentControllerApi.md#Create6) | **Post** /testresult/attachment | 
[**Delete6**](TestResultAttachmentControllerApi.md#Delete6) | **Delete** /testresult/attachment/{id} | 
[**FindAll5**](TestResultAttachmentControllerApi.md#FindAll5) | **Get** /testresult/attachment | 
[**Patch5**](TestResultAttachmentControllerApi.md#Patch5) | **Patch** /testresult/attachment/{id} | 
[**ReadContent**](TestResultAttachmentControllerApi.md#ReadContent) | **Get** /testresult/attachment/{id}/content | 
[**UpdateContent**](TestResultAttachmentControllerApi.md#UpdateContent) | **Put** /testresult/attachment/{id}/content | 

# **Create6**
> []TestResultAttachmentRowDto Create6(ctx, testResultId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
 **optional** | ***TestResultAttachmentControllerApiCreate6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultAttachmentControllerApiCreate6Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**[]TestResultAttachmentRowDto**](TestResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete6**
> Delete6(ctx, id)


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

# **FindAll5**
> PageTestResultAttachmentRowDto FindAll5(ctx, testResultId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
 **optional** | ***TestResultAttachmentControllerApiFindAll5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultAttachmentControllerApiFindAll5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageTestResultAttachmentRowDto**](PageTestResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch5**
> TestResultAttachmentRowDto Patch5(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultAttachmentPatchDto**](TestResultAttachmentPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestResultAttachmentRowDto**](TestResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContent**
> StreamingResponseBody ReadContent(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestResultAttachmentControllerApiReadContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultAttachmentControllerApiReadContentOpts struct
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

# **UpdateContent**
> TestResultAttachmentRowDto UpdateContent(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestResultAttachmentControllerApiUpdateContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultAttachmentControllerApiUpdateContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**TestResultAttachmentRowDto**](TestResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

