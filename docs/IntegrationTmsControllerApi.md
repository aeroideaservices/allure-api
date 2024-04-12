# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindProjects**](IntegrationTmsControllerApi.md#FindProjects) | **Get** /integration/tms/projects | Get available projects for tms

# **FindProjects**
> []ExtProject FindProjects(ctx, projectId, integrationId, optional)
Get available projects for tms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
 **optional** | ***IntegrationTmsControllerApiFindProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationTmsControllerApiFindProjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**|  | 

### Return type

[**[]ExtProject**](ExtProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

