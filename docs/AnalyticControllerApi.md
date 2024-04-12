# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationChart**](AnalyticControllerApi.md#AutomationChart) | **Get** /analytic/{id}/automation_chart | 
[**GroupByAutomation**](AnalyticControllerApi.md#GroupByAutomation) | **Get** /analytic/{id}/group_by_automation | 
[**GroupByStatus**](AnalyticControllerApi.md#GroupByStatus) | **Get** /analytic/{id}/group_by_status | 
[**LastResults**](AnalyticControllerApi.md#LastResults) | **Get** /analytic/{id}/tc_last_result | 
[**LaunchDurationHistogram**](AnalyticControllerApi.md#LaunchDurationHistogram) | **Get** /analytic/{id}/launch_duration_histogram | 
[**MuteTrend**](AnalyticControllerApi.md#MuteTrend) | **Get** /analytic/{id}/mute_trend | 
[**StatisticTrend**](AnalyticControllerApi.md#StatisticTrend) | **Get** /analytic/{id}/statistic_trend | 
[**SuccessRate**](AnalyticControllerApi.md#SuccessRate) | **Get** /analytic/{id}/tc_success_rate | 

# **AutomationChart**
> []AnalyticAutomationTrendDto AutomationChart(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiAutomationChartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiAutomationChartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 
 **launchRql** | **optional.String**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | [**optional.Interface of AnalyticInterval**](.md)|  | 

### Return type

[**[]AnalyticAutomationTrendDto**](AnalyticAutomationTrendDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupByAutomation**
> []AnalyticTcAutomationCountDto GroupByAutomation(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiGroupByAutomationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiGroupByAutomationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 

### Return type

[**[]AnalyticTcAutomationCountDto**](AnalyticTcAutomationCountDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupByStatus**
> []AnalyticTcStatusCountDto GroupByStatus(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiGroupByStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiGroupByStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 

### Return type

[**[]AnalyticTcStatusCountDto**](AnalyticTcStatusCountDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LastResults**
> []TestCaseLastResultDto LastResults(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**[]TestCaseLastResultDto**](TestCaseLastResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LaunchDurationHistogram**
> []AnalyticLaunchDurationHistogramDto LaunchDurationHistogram(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiLaunchDurationHistogramOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiLaunchDurationHistogramOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 
 **launchRql** | **optional.String**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **buckets** | **optional.Int32**|  | [default to 10]

### Return type

[**[]AnalyticLaunchDurationHistogramDto**](AnalyticLaunchDurationHistogramDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteTrend**
> []AnalyticMuteTrendDto MuteTrend(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiMuteTrendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiMuteTrendOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | [**optional.Interface of AnalyticInterval**](.md)|  | 

### Return type

[**[]AnalyticMuteTrendDto**](AnalyticMuteTrendDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticTrend**
> []AnalyticTrByStatusTrendDto StatisticTrend(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiStatisticTrendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiStatisticTrendOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 
 **launchRql** | **optional.String**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | [**optional.Interface of AnalyticInterval**](.md)|  | 

### Return type

[**[]AnalyticTrByStatusTrendDto**](AnalyticTrByStatusTrendDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuccessRate**
> []AnalyticDto SuccessRate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***AnalyticControllerApiSuccessRateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticControllerApiSuccessRateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tcRql** | **optional.String**|  | 
 **launchRql** | **optional.String**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | [**optional.Interface of AnalyticInterval**](.md)|  | 

### Return type

[**[]AnalyticDto**](AnalyticDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

