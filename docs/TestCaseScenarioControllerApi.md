# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create14**](TestCaseScenarioControllerApi.md#Create14) | **Post** /testcase/step | Create scenario step
[**DeleteById1**](TestCaseScenarioControllerApi.md#DeleteById1) | **Delete** /testcase/step/{id} | Delete a specified scenario step
[**DeleteNewScenario**](TestCaseScenarioControllerApi.md#DeleteNewScenario) | **Delete** /testcase/{id}/step | Delete scenario for test case
[**DeleteScenario**](TestCaseScenarioControllerApi.md#DeleteScenario) | **Delete** /testcase/{id}/scenario | Delete scenario for test case
[**GetNormalizedScenario**](TestCaseScenarioControllerApi.md#GetNormalizedScenario) | **Get** /testcase/{id}/step | Get scenario for test case
[**GetScenario**](TestCaseScenarioControllerApi.md#GetScenario) | **Get** /testcase/{id}/scenario | Find scenario for test case
[**GetScenarioFromLastRun**](TestCaseScenarioControllerApi.md#GetScenarioFromLastRun) | **Get** /testcase/{id}/scenariofromrun | Find scenario for test case from last run
[**MigrateScenario**](TestCaseScenarioControllerApi.md#MigrateScenario) | **Post** /testcase/{id}/migrate | Migrate scenario for test case
[**Move**](TestCaseScenarioControllerApi.md#Move) | **Post** /testcase/step/{id}/move | Move scenario step
[**Move1**](TestCaseScenarioControllerApi.md#Move1) | **Post** /testcase/step/{id}/copy | Copy scenario step
[**PatchById**](TestCaseScenarioControllerApi.md#PatchById) | **Patch** /testcase/step/{id} | Patch a specified scenario step
[**UpdateScenario**](TestCaseScenarioControllerApi.md#UpdateScenario) | **Post** /testcase/{id}/scenario | Update scenario for test case

# **Create14**
> ScenarioStepCreatedResponseDto Create14(ctx, body, optional)
Create scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepCreateDto**](ScenarioStepCreateDto.md)|  | 
 **optional** | ***TestCaseScenarioControllerApiCreate14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestCaseScenarioControllerApiCreate14Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beforeId** | **optional.**|  | 
 **afterId** | **optional.**|  | 

### Return type

[**ScenarioStepCreatedResponseDto**](ScenarioStepCreatedResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteById1**
> NormalizedScenarioDto DeleteById1(ctx, id)
Delete a specified scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNewScenario**
> DeleteNewScenario(ctx, id)
Delete scenario for test case

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

# **DeleteScenario**
> DeleteScenario(ctx, id)
Delete scenario for test case

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

# **GetNormalizedScenario**
> NormalizedScenarioDto GetNormalizedScenario(ctx, id)
Get scenario for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScenario**
> TestCaseScenarioDto GetScenario(ctx, id)
Find scenario for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestCaseScenarioDto**](TestCaseScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScenarioFromLastRun**
> TestCaseScenarioDto GetScenarioFromLastRun(ctx, id)
Find scenario for test case from last run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestCaseScenarioDto**](TestCaseScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateScenario**
> MigrateScenario(ctx, id)
Migrate scenario for test case

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

# **Move**
> NormalizedScenarioDto Move(ctx, body, id)
Move scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepMoveDto**](ScenarioStepMoveDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Move1**
> NormalizedScenarioDto Move1(ctx, body, id)
Copy scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepCopyDto**](ScenarioStepCopyDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchById**
> NormalizedScenarioDto PatchById(ctx, body, id)
Patch a specified scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepPatchDto**](ScenarioStepPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateScenario**
> TestCaseScenarioDto UpdateScenario(ctx, body, id)
Update scenario for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseScenarioDto**](TestCaseScenarioDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseScenarioDto**](TestCaseScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

