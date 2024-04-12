# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalPermissions**](GlobalSettingsControllerApi.md#GetGlobalPermissions) | **Get** /globalsettings/globalpermissions | Returns all global permissions for user
[**GetGlobalSettings**](GlobalSettingsControllerApi.md#GetGlobalSettings) | **Get** /globalsettings | Returns global settings
[**PatchProjectCreate**](GlobalSettingsControllerApi.md#PatchProjectCreate) | **Patch** /globalsettings/projectcreate | Patch global settings

# **GetGlobalPermissions**
> GlobalPermissionsDto GetGlobalPermissions(ctx, )
Returns all global permissions for user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalPermissionsDto**](GlobalPermissionsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalSettings**
> GlobalSettingsDto GetGlobalSettings(ctx, )
Returns global settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalSettingsDto**](GlobalSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProjectCreate**
> PatchProjectCreate(ctx, body)
Patch global settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalSettingsProjectCreatePatchDto**](GlobalSettingsProjectCreatePatchDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

