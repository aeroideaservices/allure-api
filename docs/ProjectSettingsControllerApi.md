# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLaunchCloseConfig**](ProjectSettingsControllerApi.md#GetLaunchCloseConfig) | **Get** /projectsettings/launchclose | Get launch close config
[**GetLaunchLiveDocConfig**](ProjectSettingsControllerApi.md#GetLaunchLiveDocConfig) | **Get** /projectsettings/launchlivedoc | Get launch live documentation config
[**SetLaunchCloseConfig**](ProjectSettingsControllerApi.md#SetLaunchCloseConfig) | **Patch** /projectsettings/launchclose | Save launch close config
[**SetLaunchLiveDocConfig**](ProjectSettingsControllerApi.md#SetLaunchLiveDocConfig) | **Patch** /projectsettings/launchlivedoc | Save launch live documentation config

# **GetLaunchCloseConfig**
> LaunchCloseConfigDto GetLaunchCloseConfig(ctx, projectId)
Get launch close config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**LaunchCloseConfigDto**](LaunchCloseConfigDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaunchLiveDocConfig**
> LaunchLiveDocConfigDto GetLaunchLiveDocConfig(ctx, projectId)
Get launch live documentation config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**LaunchLiveDocConfigDto**](LaunchLiveDocConfigDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLaunchCloseConfig**
> SetLaunchCloseConfig(ctx, body)
Save launch close config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LaunchCloseConfigDto**](LaunchCloseConfigDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLaunchLiveDocConfig**
> SetLaunchLiveDocConfig(ctx, body)
Save launch live documentation config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LaunchLiveDocConfigDto**](LaunchLiveDocConfigDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

