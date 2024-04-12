# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallSelect**](IntegrationSelectControllerApi.md#CallSelect) | **Get** /integration/select | 

# **CallSelect**
> []ExtSelectValue CallSelect(ctx, projectId, integrationId, fieldName, params, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **integrationId** | **int64**|  | 
  **fieldName** | **string**|  | 
  **params** | [**map[string]string**](string.md)|  | 
 **optional** | ***IntegrationSelectControllerApiCallSelectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationSelectControllerApiCallSelectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **query** | **optional.String**|  | 

### Return type

[**[]ExtSelectValue**](ExtSelectValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

