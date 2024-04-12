# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create20**](SharedStepAttachmentControllerApi.md#Create20) | **Post** /sharedstep/attachment | Upload new shared step attachments
[**Delete18**](SharedStepAttachmentControllerApi.md#Delete18) | **Delete** /sharedstep/attachment/{id} | Delete shared step attachment
[**FindAll17**](SharedStepAttachmentControllerApi.md#FindAll17) | **Get** /sharedstep/attachment | Find attachments for shared step
[**Patch17**](SharedStepAttachmentControllerApi.md#Patch17) | **Patch** /sharedstep/attachment/{id} | Patch shared step attachment
[**ReadContent3**](SharedStepAttachmentControllerApi.md#ReadContent3) | **Get** /sharedstep/attachment/{id}/content | Get attachment content by id
[**UpdateContent3**](SharedStepAttachmentControllerApi.md#UpdateContent3) | **Put** /sharedstep/attachment/{id}/content | Update shared step attachment content

# **Create20**
> []SharedStepAttachmentRowDto Create20(ctx, sharedStepId, optional)
Upload new shared step attachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharedStepId** | **int64**|  | 
 **optional** | ***SharedStepAttachmentControllerApiCreate20Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepAttachmentControllerApiCreate20Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**[]SharedStepAttachmentRowDto**](SharedStepAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete18**
> Delete18(ctx, id)
Delete shared step attachment

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

# **FindAll17**
> PageSharedStepAttachmentRowDto FindAll17(ctx, sharedStepId, optional)
Find attachments for shared step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharedStepId** | **int64**|  | 
 **optional** | ***SharedStepAttachmentControllerApiFindAll17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepAttachmentControllerApiFindAll17Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageSharedStepAttachmentRowDto**](PageSharedStepAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch17**
> SharedStepAttachmentRowDto Patch17(ctx, body, id)
Patch shared step attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedStepAttachmentPatchDto**](SharedStepAttachmentPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**SharedStepAttachmentRowDto**](SharedStepAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContent3**
> StreamingResponseBody ReadContent3(ctx, id, optional)
Get attachment content by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***SharedStepAttachmentControllerApiReadContent3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepAttachmentControllerApiReadContent3Opts struct
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

# **UpdateContent3**
> SharedStepAttachmentRowDto UpdateContent3(ctx, id, optional)
Update shared step attachment content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***SharedStepAttachmentControllerApiUpdateContent3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepAttachmentControllerApiUpdateContent3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**SharedStepAttachmentRowDto**](SharedStepAttachmentRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

