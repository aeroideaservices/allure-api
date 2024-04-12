# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionsOnProject**](PermissionControllerApi.md#GetPermissionsOnProject) | **Get** /permission | Get user permissions for project

# **GetPermissionsOnProject**
> []PermissionDto GetPermissionsOnProject(ctx, projectId)
Get user permissions for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**[]PermissionDto**](PermissionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

