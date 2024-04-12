# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrateStats**](BlobStorageUpgradeControllerApi.md#MigrateStats) | **Get** /storage/upgrade/v2 | 
[**MigrateTestCaseAttachments**](BlobStorageUpgradeControllerApi.md#MigrateTestCaseAttachments) | **Post** /storage/upgrade/v2/test-case-attachment | 
[**MigrateTestFixtureResult**](BlobStorageUpgradeControllerApi.md#MigrateTestFixtureResult) | **Post** /storage/upgrade/v2/test-fixture-result | 
[**MigrateTestFixtureResultAttachment**](BlobStorageUpgradeControllerApi.md#MigrateTestFixtureResultAttachment) | **Post** /storage/upgrade/v2/test-fixture-result-attachment | 
[**MigrateTestResult**](BlobStorageUpgradeControllerApi.md#MigrateTestResult) | **Post** /storage/upgrade/v2/test-result | 
[**MigrateTestResultAttachment**](BlobStorageUpgradeControllerApi.md#MigrateTestResultAttachment) | **Post** /storage/upgrade/v2/test-result-attachment | 

# **MigrateStats**
> BlobStorageUpdateStats MigrateStats(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BlobStorageUpdateStats**](BlobStorageUpdateStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateTestCaseAttachments**
> MigrateTestCaseAttachments(ctx, projectId, limit, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **limit** | **int32**|  | 
 **optional** | ***BlobStorageUpgradeControllerApiMigrateTestCaseAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStorageUpgradeControllerApiMigrateTestCaseAttachmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeOnError** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateTestFixtureResult**
> MigrateTestFixtureResult(ctx, projectId, limit, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **limit** | **int32**|  | 
 **optional** | ***BlobStorageUpgradeControllerApiMigrateTestFixtureResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStorageUpgradeControllerApiMigrateTestFixtureResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeOnError** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateTestFixtureResultAttachment**
> MigrateTestFixtureResultAttachment(ctx, projectId, limit, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **limit** | **int32**|  | 
 **optional** | ***BlobStorageUpgradeControllerApiMigrateTestFixtureResultAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStorageUpgradeControllerApiMigrateTestFixtureResultAttachmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeOnError** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateTestResult**
> MigrateTestResult(ctx, projectId, limit, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **limit** | **int32**|  | 
 **optional** | ***BlobStorageUpgradeControllerApiMigrateTestResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStorageUpgradeControllerApiMigrateTestResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeOnError** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateTestResultAttachment**
> MigrateTestResultAttachment(ctx, projectId, limit, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 
  **limit** | **int32**|  | 
 **optional** | ***BlobStorageUpgradeControllerApiMigrateTestResultAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStorageUpgradeControllerApiMigrateTestResultAttachmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeOnError** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

