# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindById**](JobRunControllerApi.md#FindById) | **Get** /jobrun/{id} | Get job run by id
[**Rerun1**](JobRunControllerApi.md#Rerun1) | **Post** /jobrun/{id}/rerun | Rerun job
[**Upload2**](JobRunControllerApi.md#Upload2) | **Post** /jobrun/{id}/upload | Manually upload job run results
[**UploadArchives1**](JobRunControllerApi.md#UploadArchives1) | **Post** /jobrun/{id}/upload/archive | Manually upload job run results
[**UploadFiles1**](JobRunControllerApi.md#UploadFiles1) | **Post** /jobrun/{id}/upload/file | Manually upload job run results

# **FindById**
> JobRunDto FindById(ctx, id)
Get job run by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**JobRunDto**](JobRunDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Rerun1**
> Rerun1(ctx, body, id)
Rerun job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JobRerunRequestDto**](JobRerunRequestDto.md)|  | 
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Upload2**
> FileUploadResponseDto Upload2(ctx, id, optional)
Manually upload job run results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***JobRunControllerApiUpload2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobRunControllerApiUpload2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of JobRunUploadInfoDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 
 **archive** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**FileUploadResponseDto**](FileUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadArchives1**
> FileUploadResponseDto UploadArchives1(ctx, id, optional)
Manually upload job run results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***JobRunControllerApiUploadArchives1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobRunControllerApiUploadArchives1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of JobRunUploadInfoDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**FileUploadResponseDto**](FileUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFiles1**
> FileUploadResponseDto UploadFiles1(ctx, id, optional)
Manually upload job run results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***JobRunControllerApiUploadFiles1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobRunControllerApiUploadFiles1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of JobRunUploadInfoDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**FileUploadResponseDto**](FileUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

