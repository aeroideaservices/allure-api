# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobRunByUid1**](UploadControllerApi.md#GetJobRunByUid1) | **Get** /upload/jobrun | Get information about job run by id
[**SessionJobRun1**](UploadControllerApi.md#SessionJobRun1) | **Post** /upload/session | Creates test session for manual upload
[**Start**](UploadControllerApi.md#Start) | **Post** /upload/run | Notifies about external job run start
[**Start1**](UploadControllerApi.md#Start1) | **Post** /upload/start | Notifies about external job run start
[**Stop**](UploadControllerApi.md#Stop) | **Post** /upload/stop | Notifies about external job run stop
[**UploadResults**](UploadControllerApi.md#UploadResults) | **Post** /upload | Upload test results
[**UploadResultsArchives**](UploadControllerApi.md#UploadResultsArchives) | **Post** /upload/archive | Upload archives with test results
[**UploadResultsFiles**](UploadControllerApi.md#UploadResultsFiles) | **Post** /upload/file | Upload files with test results

# **GetJobRunByUid1**
> ExternalRunResponseDto GetJobRunByUid1(ctx, projectId, jobUid, jobRunUid, jobRunId)
Get information about job run by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **jobUid** | **string**|  | 
  **jobRunUid** | **string**|  | 
  **jobRunId** | **int64**|  | 

### Return type

[**ExternalRunResponseDto**](ExternalRunResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionJobRun1**
> TestSessionResponseDto SessionJobRun1(ctx, body)
Creates test session for manual upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ManualSessionRequestDto**](ManualSessionRequestDto.md)|  | 

### Return type

[**TestSessionResponseDto**](TestSessionResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Start**
> ExternalRunResponseDto Start(ctx, body)
Notifies about external job run start

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalRunStartRequestDto**](ExternalRunStartRequestDto.md)|  | 

### Return type

[**ExternalRunResponseDto**](ExternalRunResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Start1**
> ExternalRunResponseDto Start1(ctx, body)
Notifies about external job run start

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalRunStartRequestDto**](ExternalRunStartRequestDto.md)|  | 

### Return type

[**ExternalRunResponseDto**](ExternalRunResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stop**
> ExternalRunResponseDto Stop(ctx, body)
Notifies about external job run stop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalRunStopRequestDto**](ExternalRunStopRequestDto.md)|  | 

### Return type

[**ExternalRunResponseDto**](ExternalRunResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadResults**
> UploadResults(ctx, id, optional)
Upload test results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***UploadControllerApiUploadResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadControllerApiUploadResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 
 **archive** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadResultsArchives**
> UploadResultsArchives(ctx, id, optional)
Upload archives with test results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***UploadControllerApiUploadResultsArchivesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadControllerApiUploadResultsArchivesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadResultsFiles**
> UploadResultsFiles(ctx, id, optional)
Upload files with test results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***UploadControllerApiUploadResultsFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadControllerApiUploadResultsFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of []*os.File**](*os.File.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

