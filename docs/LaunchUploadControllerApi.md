# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](LaunchUploadControllerApi.md#Upload) | **Post** /launch/{launchId}/upload | Manually upload launch results
[**Upload1**](LaunchUploadControllerApi.md#Upload1) | **Post** /launch/upload | Create launch from uploaded results
[**UploadArchives**](LaunchUploadControllerApi.md#UploadArchives) | **Post** /launch/{launchId}/upload/archive | Manually upload launch results
[**UploadFiles**](LaunchUploadControllerApi.md#UploadFiles) | **Post** /launch/{launchId}/upload/file | Manually upload launch results

# **Upload**
> LaunchUploadResponseDto Upload(ctx, launchId, optional)
Manually upload launch results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***LaunchUploadControllerApiUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchUploadControllerApiUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of LaunchExistingUploadDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 
 **archive** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**LaunchUploadResponseDto**](LaunchUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Upload1**
> LaunchUploadResponseDto Upload1(ctx, optional)
Create launch from uploaded results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaunchUploadControllerApiUpload1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchUploadControllerApiUpload1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of LaunchCreateAndUploadDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 
 **archive** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**LaunchUploadResponseDto**](LaunchUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadArchives**
> LaunchUploadResponseDto UploadArchives(ctx, launchId, optional)
Manually upload launch results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***LaunchUploadControllerApiUploadArchivesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchUploadControllerApiUploadArchivesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of LaunchExistingUploadDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**LaunchUploadResponseDto**](LaunchUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFiles**
> LaunchUploadResponseDto UploadFiles(ctx, launchId, optional)
Manually upload launch results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***LaunchUploadControllerApiUploadFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchUploadControllerApiUploadFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of LaunchExistingUploadDto**](.md)|  | 
 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

[**LaunchUploadResponseDto**](LaunchUploadResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

