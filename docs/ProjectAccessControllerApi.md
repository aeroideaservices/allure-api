# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectCollaborators**](ProjectAccessControllerApi.md#AddProjectCollaborators) | **Post** /project/access/{projectId}/collaborator | Add collaborators to project
[**AddProjectGroups**](ProjectAccessControllerApi.md#AddProjectGroups) | **Post** /project/access/{projectId}/group | Add groups to project
[**DeleteProjectGroups**](ProjectAccessControllerApi.md#DeleteProjectGroups) | **Delete** /project/access/{projectId}/group | Delete groups from project
[**DeleteUsers**](ProjectAccessControllerApi.md#DeleteUsers) | **Delete** /project/access/{projectId}/collaborator | Delete collaborators from project
[**GetProjectAccessGroups**](ProjectAccessControllerApi.md#GetProjectAccessGroups) | **Get** /project/access/{projectId}/group | Get project access groups
[**GetProjectCollaborators**](ProjectAccessControllerApi.md#GetProjectCollaborators) | **Get** /project/access/{projectId}/collaborator | Get project collaborators

# **AddProjectCollaborators**
> []ProjectCollaboratorAccessDto AddProjectCollaborators(ctx, body, projectId)
Add collaborators to project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectAccessCollaboratorAddDto**](ProjectAccessCollaboratorAddDto.md)|  | 
  **projectId** | **int64**|  | 

### Return type

[**[]ProjectCollaboratorAccessDto**](ProjectCollaboratorAccessDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddProjectGroups**
> []ProjectGroupAccessDto AddProjectGroups(ctx, body, projectId)
Add groups to project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectAccessGroupAddDto**](ProjectAccessGroupAddDto.md)|  | 
  **projectId** | **int64**|  | 

### Return type

[**[]ProjectGroupAccessDto**](ProjectGroupAccessDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectGroups**
> DeleteProjectGroups(ctx, projectId, groupId)
Delete groups from project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **groupId** | [**[]int64**](int64.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsers**
> DeleteUsers(ctx, projectId, username)
Delete collaborators from project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **username** | [**[]string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectAccessGroups**
> PageProjectGroupAccessDto GetProjectAccessGroups(ctx, projectId, optional)
Get project access groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***ProjectAccessControllerApiGetProjectAccessGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectAccessControllerApiGetProjectAccessGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **permissionsSetId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageProjectGroupAccessDto**](PageProjectGroupAccessDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectCollaborators**
> PageProjectCollaboratorAccessDto GetProjectCollaborators(ctx, projectId, optional)
Get project collaborators

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
 **optional** | ***ProjectAccessControllerApiGetProjectCollaboratorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectAccessControllerApiGetProjectCollaboratorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **permissionsSetId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;username,ASC&quot;]]

### Return type

[**PageProjectCollaboratorAccessDto**](PageProjectCollaboratorAccessDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

