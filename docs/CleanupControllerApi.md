# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanupLaunch**](CleanupControllerApi.md#CleanupLaunch) | **Post** /cleanup/launch | 
[**TriggerBlobRemoveTask**](CleanupControllerApi.md#TriggerBlobRemoveTask) | **Post** /cleanup/scheduler/blob_remove_task | 
[**TriggerCleanup**](CleanupControllerApi.md#TriggerCleanup) | **Post** /cleanup/scheduler/cleaner_schema_project | 
[**TriggerGlobalCleanup**](CleanupControllerApi.md#TriggerGlobalCleanup) | **Post** /cleanup/scheduler/cleaner_schema_global | 

# **CleanupLaunch**
> CleanupLaunch(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LaunchCleanupRequest**](LaunchCleanupRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerBlobRemoveTask**
> TriggerBlobRemoveTask(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerCleanup**
> TriggerCleanup(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerGlobalCleanup**
> TriggerGlobalCleanup(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

