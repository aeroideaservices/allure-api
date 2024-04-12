# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FixLinks**](IntegrationIssueControllerApi.md#FixLinks) | **Post** /integration/issue/{integrationId}/fixlinks | Fix issue links without url
[**GetFields**](IntegrationIssueControllerApi.md#GetFields) | **Get** /integration/issue/field | Get available fields for specified project integration, project key and issue type
[**GetIssues3**](IntegrationIssueControllerApi.md#GetIssues3) | **Get** /integration/issue/suggest | Get available issues for specified project integration
[**GetProjects**](IntegrationIssueControllerApi.md#GetProjects) | **Get** /integration/issue/project | Get available projects for specified project integration
[**GetTypes**](IntegrationIssueControllerApi.md#GetTypes) | **Get** /integration/issue/type | Get available issue types for specified project integration and project key

# **FixLinks**
> IntegrationLinksFixedDto FixLinks(ctx, integrationId)
Fix issue links without url

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationId** | **int64**|  | 

### Return type

[**IntegrationLinksFixedDto**](IntegrationLinksFixedDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFields**
> []Object GetFields(ctx, projectId, integrationId, projectKey, issueTypeId)
Get available fields for specified project integration, project key and issue type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
  **projectKey** | **string**|  | 
  **issueTypeId** | **string**|  | 

### Return type

**[]Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIssues3**
> []ExtIssueLink GetIssues3(ctx, projectId, integrationId, optional)
Get available issues for specified project integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
 **optional** | ***IntegrationIssueControllerApiGetIssues3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationIssueControllerApiGetIssues3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**|  | 

### Return type

[**[]ExtIssueLink**](ExtIssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> []ExtProject GetProjects(ctx, projectId, integrationId, optional)
Get available projects for specified project integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
 **optional** | ***IntegrationIssueControllerApiGetProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationIssueControllerApiGetProjectsOpts struct
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

# **GetTypes**
> []ExtIssueType GetTypes(ctx, projectId, integrationId, projectKey)
Get available issue types for specified project integration and project key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
  **projectKey** | **string**|  | 

### Return type

[**[]ExtIssueType**](ExtIssueType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

