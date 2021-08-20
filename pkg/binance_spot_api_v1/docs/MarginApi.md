# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMarginAssetsMARKETDATA**](MarginApi.md#GetAllMarginAssetsMARKETDATA) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**GetAllMarginPairsMARKETDATA**](MarginApi.md#GetAllMarginPairsMARKETDATA) | **Get** /sapi/v1/margin/allPairs | Get All Margin Pairs (MARKET_DATA)
[**GetBNBBurnStatusUSERDATA**](MarginApi.md#GetBNBBurnStatusUSERDATA) | **Get** /sapi/v1/bnbBurn | Get BNB Burn Status (USER_DATA)
[**GetForceLiquidationRecordUSERDATA**](MarginApi.md#GetForceLiquidationRecordUSERDATA) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**GetInterestHistoryUSERDATA**](MarginApi.md#GetInterestHistoryUSERDATA) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**GetTransferHistoryUSERDATA**](MarginApi.md#GetTransferHistoryUSERDATA) | **Get** /sapi/v1/margin/transfer | Get Transfer History (USER_DATA)
[**MarginAccountBorrowMARGIN**](MarginApi.md#MarginAccountBorrowMARGIN) | **Post** /sapi/v1/margin/loan | Margin Account Borrow (MARGIN)
[**MarginAccountCancelOrderTRADE**](MarginApi.md#MarginAccountCancelOrderTRADE) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginAccountCancelallOpenOrdersonaSymbol**](MarginApi.md#MarginAccountCancelallOpenOrdersonaSymbol) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol
[**MarginAccountNewOrderTRADE**](MarginApi.md#MarginAccountNewOrderTRADE) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginAccountRepayMARGIN**](MarginApi.md#MarginAccountRepayMARGIN) | **Post** /sapi/v1/margin/repay | Margin Account Repay (MARGIN)
[**MarginAccountTransferMARGIN**](MarginApi.md#MarginAccountTransferMARGIN) | **Post** /sapi/v1/margin/transfer | Margin Account Transfer (MARGIN)
[**QueryLoanRecordUSERDATA**](MarginApi.md#QueryLoanRecordUSERDATA) | **Get** /sapi/v1/margin/loan | Query Loan Record (USER_DATA)
[**QueryMarginAccountDetailsUSERDATA**](MarginApi.md#QueryMarginAccountDetailsUSERDATA) | **Get** /sapi/v1/margin/account | Query Margin Account Details (USER_DATA)
[**QueryMarginAccountsAllOrderUSERDATA**](MarginApi.md#QueryMarginAccountsAllOrderUSERDATA) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#x27;s All Order (USER_DATA)
[**QueryMarginAccountsOpenOrderUSERDATA**](MarginApi.md#QueryMarginAccountsOpenOrderUSERDATA) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#x27;s Open Order (USER_DATA)
[**QueryMarginAccountsOrderUSERDATA**](MarginApi.md#QueryMarginAccountsOrderUSERDATA) | **Get** /sapi/v1/margin/order | Query Margin Account&#x27;s Order (USER_DATA)
[**QueryMarginAccountsTradeListUSERDATA**](MarginApi.md#QueryMarginAccountsTradeListUSERDATA) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#x27;s Trade List (USER_DATA)
[**QueryMarginAssetMARKETDATA**](MarginApi.md#QueryMarginAssetMARKETDATA) | **Get** /sapi/v1/margin/asset | Query Margin Asset (MARKET_DATA)
[**QueryMarginInterestRateHistoryUSERDATA**](MarginApi.md#QueryMarginInterestRateHistoryUSERDATA) | **Get** /sapi/v1/margin/interestRateHistory | Query Margin Interest Rate History (USER_DATA)
[**QueryMarginPairMARKETDATA**](MarginApi.md#QueryMarginPairMARKETDATA) | **Get** /sapi/v1/margin/pair | Query Margin Pair (MARKET_DATA)
[**QueryMarginPriceIndexMARKETDATA**](MarginApi.md#QueryMarginPriceIndexMARKETDATA) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**QueryMaxBorrowUSERDATA**](MarginApi.md#QueryMaxBorrowUSERDATA) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**QueryMaxTransferOutAmountUSERDATA**](MarginApi.md#QueryMaxTransferOutAmountUSERDATA) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**QueryRepayRecordUSERDATA**](MarginApi.md#QueryRepayRecordUSERDATA) | **Get** /sapi/v1/margin/repay | Query Repay Record (USER_DATA)
[**ToggleBNBBurnOnSpotTradeAndMarginInterestUSERDATA**](MarginApi.md#ToggleBNBBurnOnSpotTradeAndMarginInterestUSERDATA) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)

# **GetAllMarginAssetsMARKETDATA**
> GetAllMarginAssetsMARKETDATA(ctx, contentType, xMBXAPIKEY)
Get All Margin Assets (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMarginPairsMARKETDATA**
> GetAllMarginPairsMARKETDATA(ctx, contentType, xMBXAPIKEY)
Get All Margin Pairs (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBNBBurnStatusUSERDATA**
> GetBNBBurnStatusUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get BNB Burn Status (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForceLiquidationRecordUSERDATA**
> GetForceLiquidationRecordUSERDATA(ctx, startTime, endTime, isolatedSymbol, current, size, timestamp, signature, contentType, xMBXAPIKEY)
Get Force Liquidation Record (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startTime** | **string**|  | 
  **endTime** | **string**|  | 
  **isolatedSymbol** | **string**|  | 
  **current** | **int32**|  | 
  **size** | **int32**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterestHistoryUSERDATA**
> GetInterestHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Interest History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransferHistoryUSERDATA**
> GetTransferHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Transfer History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountBorrowMARGIN**
> MarginAccountBorrowMARGIN(ctx, asset, amount, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account Borrow (MARGIN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **amount** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountCancelOrderTRADE**
> MarginAccountCancelOrderTRADE(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account Cancel Order (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountCancelallOpenOrdersonaSymbol**
> MarginAccountCancelallOpenOrdersonaSymbol(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account Cancel all Open Orders on a Symbol

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountNewOrderTRADE**
> MarginAccountNewOrderTRADE(ctx, symbol, side, type_, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account New Order (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **side** | **string**|  | 
  **type_** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountRepayMARGIN**
> MarginAccountRepayMARGIN(ctx, asset, amount, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account Repay (MARGIN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **amount** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginAccountTransferMARGIN**
> MarginAccountTransferMARGIN(ctx, asset, amount, type_, timestamp, signature, contentType, xMBXAPIKEY)
Margin Account Transfer (MARGIN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **amount** | **string**|  | 
  **type_** | **string**| 1: transfer from main account to margin account 2: transfer from margin account to main account | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryLoanRecordUSERDATA**
> QueryLoanRecordUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Query Loan Record (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAccountDetailsUSERDATA**
> QueryMarginAccountDetailsUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Account Details (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAccountsAllOrderUSERDATA**
> QueryMarginAccountsAllOrderUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Account's All Order (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAccountsOpenOrderUSERDATA**
> QueryMarginAccountsOpenOrderUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Account's Open Order (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAccountsOrderUSERDATA**
> QueryMarginAccountsOrderUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Account's Order (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAccountsTradeListUSERDATA**
> QueryMarginAccountsTradeListUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Account's Trade List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginAssetMARKETDATA**
> QueryMarginAssetMARKETDATA(ctx, asset, contentType, xMBXAPIKEY)
Query Margin Asset (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginInterestRateHistoryUSERDATA**
> QueryMarginInterestRateHistoryUSERDATA(ctx, asset, recvWindow, timestamp, signature, contentType, xMBXAPIKEY)
Query Margin Interest Rate History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **recvWindow** | **string**| No more than 60000 | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginPairMARKETDATA**
> QueryMarginPairMARKETDATA(ctx, symbol, contentType, xMBXAPIKEY)
Query Margin Pair (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMarginPriceIndexMARKETDATA**
> QueryMarginPriceIndexMARKETDATA(ctx, symbol, contentType, xMBXAPIKEY)
Query Margin PriceIndex (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMaxBorrowUSERDATA**
> QueryMaxBorrowUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Query Max Borrow (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryMaxTransferOutAmountUSERDATA**
> QueryMaxTransferOutAmountUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Query Max Transfer-Out Amount (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRepayRecordUSERDATA**
> QueryRepayRecordUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Query Repay Record (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleBNBBurnOnSpotTradeAndMarginInterestUSERDATA**
> ToggleBNBBurnOnSpotTradeAndMarginInterestUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

