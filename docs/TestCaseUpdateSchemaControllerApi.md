# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create12**](TestCaseUpdateSchemaControllerApi.md#Create12) | **Post** /testcaseupdateschema | Create a new test case update schema
[**Delete12**](TestCaseUpdateSchemaControllerApi.md#Delete12) | **Delete** /testcaseupdateschema/{id} | Delete test case update schema by id
[**FindAll10**](TestCaseUpdateSchemaControllerApi.md#FindAll10) | **Get** /testcaseupdateschema | Find all test case update schemas for given project
[**FindOne10**](TestCaseUpdateSchemaControllerApi.md#FindOne10) | **Get** /testcaseupdateschema/{id} | Find a test case update schemas by id
[**Patch11**](TestCaseUpdateSchemaControllerApi.md#Patch11) | **Patch** /testcaseupdateschema/{id} | Patch test case update schema

# **Create12**
> TestCaseUpdateSchemaDto Create12(ctx, body)
Create a new test case update schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseUpdateSchemaCreateDto**](TestCaseUpdateSchemaCreateDto.md)|  | 

### Return type

[**TestCaseUpdateSchemaDto**](TestCaseUpdateSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete12**
> Delete12(ctx, id)
Delete test case update schema by id

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

# **FindAll10**
> []TestCaseUpdateSchemaDto FindAll10(ctx, projectId)
Find all test case update schemas for given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**|  | 

### Return type

[**[]TestCaseUpdateSchemaDto**](TestCaseUpdateSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne10**
> TestCaseUpdateSchemaDto FindOne10(ctx, id)
Find a test case update schemas by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**TestCaseUpdateSchemaDto**](TestCaseUpdateSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch11**
> TestCaseUpdateSchemaDto Patch11(ctx, body, id)
Patch test case update schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseUpdateSchemaPatchDto**](TestCaseUpdateSchemaPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**TestCaseUpdateSchemaDto**](TestCaseUpdateSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

