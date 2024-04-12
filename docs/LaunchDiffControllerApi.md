# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FailedToPassedDiff**](LaunchDiffControllerApi.md#FailedToPassedDiff) | **Get** /launch/diff/fixed | Find fixed
[**GetNew**](LaunchDiffControllerApi.md#GetNew) | **Get** /launch/diff/new | New tests
[**GetStatusMatrix**](LaunchDiffControllerApi.md#GetStatusMatrix) | **Get** /launch/diff/matrix | Get status matrix for given launches with overlay parameter
[**Missed**](LaunchDiffControllerApi.md#Missed) | **Get** /launch/diff/missed | Missed tests
[**PassedToFailedDiff**](LaunchDiffControllerApi.md#PassedToFailedDiff) | **Get** /launch/diff/failed | Find failed
[**StatusChanged**](LaunchDiffControllerApi.md#StatusChanged) | **Get** /launch/diff/status-changed | Find status changed difference

# **FailedToPassedDiff**
> []LaunchDiffStatusChangeDto FailedToPassedDiff(ctx, from, to, optional)
Find fixed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **int64**|  | 
  **to** | **int64**|  | 
 **optional** | ***LaunchDiffControllerApiFailedToPassedDiffOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiFailedToPassedDiffOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**|  | 

### Return type

[**[]LaunchDiffStatusChangeDto**](LaunchDiffStatusChangeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNew**
> []LaunchDiffTestResultDto GetNew(ctx, from, to, optional)
New tests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **int64**|  | 
  **to** | **int64**|  | 
 **optional** | ***LaunchDiffControllerApiGetNewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiGetNewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**|  | 

### Return type

[**[]LaunchDiffTestResultDto**](LaunchDiffTestResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatusMatrix**
> PageLaunchDiffRow GetStatusMatrix(ctx, launchIds, optional)
Get status matrix for given launches with overlay parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchIds** | [**[]int64**](int64.md)|  | 
 **optional** | ***LaunchDiffControllerApiGetStatusMatrixOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiGetStatusMatrixOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | [**optional.Interface of LaunchDiffMode**](.md)|  | 
 **statusChange** | **optional.Bool**|  | [default to false]
 **search** | **optional.String**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,DESC&quot;]]

### Return type

[**PageLaunchDiffRow**](PageLaunchDiffRow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Missed**
> []LaunchDiffTestResultDto Missed(ctx, from, to, optional)
Missed tests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **int64**|  | 
  **to** | **int64**|  | 
 **optional** | ***LaunchDiffControllerApiMissedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiMissedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**|  | 

### Return type

[**[]LaunchDiffTestResultDto**](LaunchDiffTestResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PassedToFailedDiff**
> []LaunchDiffStatusChangeDto PassedToFailedDiff(ctx, from, to, optional)
Find failed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **int64**|  | 
  **to** | **int64**|  | 
 **optional** | ***LaunchDiffControllerApiPassedToFailedDiffOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiPassedToFailedDiffOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**|  | 

### Return type

[**[]LaunchDiffStatusChangeDto**](LaunchDiffStatusChangeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusChanged**
> []LaunchDiffStatusChangeDto StatusChanged(ctx, from, to, optional)
Find status changed difference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **int64**|  | 
  **to** | **int64**|  | 
 **optional** | ***LaunchDiffControllerApiStatusChangedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaunchDiffControllerApiStatusChangedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **optional.String**|  | 

### Return type

[**[]LaunchDiffStatusChangeDto**](LaunchDiffStatusChangeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

