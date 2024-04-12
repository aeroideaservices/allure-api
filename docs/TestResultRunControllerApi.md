# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Assign**](TestResultRunControllerApi.md#Assign) | **Post** /testresult/{id}/assign | Assign test result
[**Resolve**](TestResultRunControllerApi.md#Resolve) | **Post** /testresult/{id}/resolve | Resolve test result

# **Assign**
> TestResultRowDto Assign(ctx, body, id)
Assign test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssignRequestDto**](AssignRequestDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestResultRowDto**](TestResultRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Resolve**
> TestResultRowDto Resolve(ctx, body, id)
Resolve test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResolveRequestDto**](ResolveRequestDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestResultRowDto**](TestResultRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

