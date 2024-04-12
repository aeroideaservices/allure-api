# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create11**](TestFixtureResultAttachmentControllerApi.md#Create11) | **Post** /testfixtureresult/attachment | 
[**Delete11**](TestFixtureResultAttachmentControllerApi.md#Delete11) | **Delete** /testfixtureresult/attachment/{id} | 
[**FindAll9**](TestFixtureResultAttachmentControllerApi.md#FindAll9) | **Get** /testfixtureresult/attachment | 
[**Patch10**](TestFixtureResultAttachmentControllerApi.md#Patch10) | **Patch** /testfixtureresult/attachment/{id} | 
[**ReadContent1**](TestFixtureResultAttachmentControllerApi.md#ReadContent1) | **Get** /testfixtureresult/attachment/{id}/content | 
[**UpdateContent1**](TestFixtureResultAttachmentControllerApi.md#UpdateContent1) | **Put** /testfixtureresult/attachment/{id}/content | 

# **Create11**
> []TestFixtureResultAttachmentRowDto Create11(ctx, tfrId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tfrId** | **int64**|  | 
 **optional** | ***TestFixtureResultAttachmentControllerApiCreate11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestFixtureResultAttachmentControllerApiCreate11Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**[]TestFixtureResultAttachmentRowDto**](TestFixtureResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete11**
> Delete11(ctx, id)


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

# **FindAll9**
> PageTestFixtureResultAttachmentRowDto FindAll9(ctx, tfrId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tfrId** | **int64**|  | 
 **optional** | ***TestFixtureResultAttachmentControllerApiFindAll9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestFixtureResultAttachmentControllerApiFindAll9Opts struct
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

# **Patch10**
> TestFixtureResultAttachmentRowDto Patch10(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestFixtureResultAttachmentPatchDto**](TestFixtureResultAttachmentPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestFixtureResultAttachmentRowDto**](TestFixtureResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContent1**
> StreamingResponseBody ReadContent1(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestFixtureResultAttachmentControllerApiReadContent1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestFixtureResultAttachmentControllerApiReadContent1Opts struct
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

# **UpdateContent1**
> TestFixtureResultAttachmentRowDto UpdateContent1(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestFixtureResultAttachmentControllerApiUpdateContent1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestFixtureResultAttachmentControllerApiUpdateContent1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**TestFixtureResultAttachmentRowDto**](TestFixtureResultAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

