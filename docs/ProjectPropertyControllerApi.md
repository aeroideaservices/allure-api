# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create23**](ProjectPropertyControllerApi.md#Create23) | **Post** /projectproperty | Create a new project property
[**Delete21**](ProjectPropertyControllerApi.md#Delete21) | **Delete** /projectproperty/{id} | Delete project by id
[**FindAll20**](ProjectPropertyControllerApi.md#FindAll20) | **Get** /projectproperty | Find all project properties
[**FindOne18**](ProjectPropertyControllerApi.md#FindOne18) | **Get** /projectproperty/{id} | Find project property by id
[**Patch20**](ProjectPropertyControllerApi.md#Patch20) | **Patch** /projectproperty/{id} | Patch project property

# **Create23**
> ProjectPropertyDto Create23(ctx, body)
Create a new project property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectPropertyCreateDto**](ProjectPropertyCreateDto.md)|  | 

### Return type

[**ProjectPropertyDto**](ProjectPropertyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete21**
> Delete21(ctx, id)
Delete project by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll20**
> []ProjectPropertyDto FindAll20(ctx, projectId)
Find all project properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**[]ProjectPropertyDto**](ProjectPropertyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne18**
> ProjectPropertyDto FindOne18(ctx, id)
Find project property by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**ProjectPropertyDto**](ProjectPropertyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch20**
> ProjectPropertyDto Patch20(ctx, body, id)
Patch project property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectPropertyPatchDto**](ProjectPropertyPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**ProjectPropertyDto**](ProjectPropertyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

