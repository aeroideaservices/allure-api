# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create35**](IntegrationWebhookControllerApi.md#Create35) | **Post** /integration/webhook | Create a new webhook config
[**DeleteById4**](IntegrationWebhookControllerApi.md#DeleteById4) | **Delete** /integration/webhook/{id} | Delete webhook config
[**FindAll31**](IntegrationWebhookControllerApi.md#FindAll31) | **Get** /integration/webhook | Find all webhook configs for integration
[**FindAllLogs**](IntegrationWebhookControllerApi.md#FindAllLogs) | **Get** /integration/webhook/log | Find all webhook logs by integration or webhook
[**Patch30**](IntegrationWebhookControllerApi.md#Patch30) | **Patch** /integration/webhook/{id} | Patch webhook config

# **Create35**
> IntegrationWebhookTokenDto Create35(ctx, body)
Create a new webhook config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IntegrationWebhookCreateDto**](IntegrationWebhookCreateDto.md)|  | 

### Return type

[**IntegrationWebhookTokenDto**](IntegrationWebhookTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteById4**
> DeleteById4(ctx, id)
Delete webhook config

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

# **FindAll31**
> PageIntegrationWebhookDto FindAll31(ctx, integrationId, optional)
Find all webhook configs for integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationId** | **int64**|  | 
 **optional** | ***IntegrationWebhookControllerApiFindAll31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationWebhookControllerApiFindAll31Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageIntegrationWebhookDto**](PageIntegrationWebhookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllLogs**
> PageIntegrationWebhookLogDto FindAllLogs(ctx, optional)
Find all webhook logs by integration or webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegrationWebhookControllerApiFindAllLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationWebhookControllerApiFindAllLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **optional.Int64**|  | 
 **webhookId** | **optional.Int64**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageIntegrationWebhookLogDto**](PageIntegrationWebhookLogDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch30**
> IntegrationWebhookDto Patch30(ctx, body, id)
Patch webhook config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IntegrationWebhookPatchDto**](IntegrationWebhookPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**IntegrationWebhookDto**](IntegrationWebhookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

