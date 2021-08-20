# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdjustCrossCollateralLTVHistoryUSERDATA**](FuturesApi.md#AdjustCrossCollateralLTVHistoryUSERDATA) | **Get** /sapi/v1/futures/loan/adjustCollateral/history | Adjust Cross-Collateral LTV History (USER_DATA)
[**AdjustCrossCollateralLTVTRADE**](FuturesApi.md#AdjustCrossCollateralLTVTRADE) | **Post** /sapi/v1/futures/loan/adjustCollateral | Adjust Cross-Collateral LTV (TRADE)
[**AdjustCrossCollateralLTVV2TRADE**](FuturesApi.md#AdjustCrossCollateralLTVV2TRADE) | **Post** /sapi/v2/futures/loan/adjustCollateral | Adjust Cross-Collateral LTV V2 (TRADE)
[**BorrowForCrossCollateralTRADE**](FuturesApi.md#BorrowForCrossCollateralTRADE) | **Post** /sapi/v1/futures/loan/borrow | Borrow For Cross-Collateral (TRADE)
[**CalculateRateAfterAdjustCrossCollateralLTVUSERDATA**](FuturesApi.md#CalculateRateAfterAdjustCrossCollateralLTVUSERDATA) | **Get** /sapi/v1/futures/loan/calcAdjustLevel | Calculate Rate After Adjust Cross-Collateral LTV (USER_DATA)
[**CalculateRateAfterAdjustCrossCollateralLTVV2USERDATA**](FuturesApi.md#CalculateRateAfterAdjustCrossCollateralLTVV2USERDATA) | **Get** /sapi/v2/futures/loan/calcAdjustLevel | Calculate Rate After Adjust Cross-Collateral LTV V2 (USER_DATA)
[**CheckCollateralRepayLimitUSERDATA**](FuturesApi.md#CheckCollateralRepayLimitUSERDATA) | **Get** /sapi/v1/futures/loan/collateralRepayLimit | Check Collateral Repay Limit (USER_DATA)
[**CollateralRepaymentResultUSERDATA**](FuturesApi.md#CollateralRepaymentResultUSERDATA) | **Get** /sapi/v1/futures/loan/collateralRepayResult | Collateral Repayment Result (USER_DATA)
[**CrossCollateralBorrowHistoryUSERDATA**](FuturesApi.md#CrossCollateralBorrowHistoryUSERDATA) | **Get** /sapi/v1/futures/loan/borrow/history | Cross-Collateral Borrow History (USER_DATA)
[**CrossCollateralInformationUSERDATA**](FuturesApi.md#CrossCollateralInformationUSERDATA) | **Get** /sapi/v1/futures/loan/configs | Cross-Collateral Information (USER_DATA)
[**CrossCollateralInformationV2USERDATA**](FuturesApi.md#CrossCollateralInformationV2USERDATA) | **Get** /sapi/v2/futures/loan/configs | Cross-Collateral Information V2 (USER_DATA)
[**CrossCollateralInterestHistoryUSERDATA**](FuturesApi.md#CrossCollateralInterestHistoryUSERDATA) | **Get** /sapi/v1/futures/loan/interestHistory | Cross-Collateral Interest History (USER_DATA)
[**CrossCollateralLiquidationHistoryUSERDATA**](FuturesApi.md#CrossCollateralLiquidationHistoryUSERDATA) | **Get** /sapi/v1/futures/loan/liquidationHistory | Cross-Collateral Liquidation History (USER_DATA)
[**CrossCollateralRepaymentHistoryUSERDATA**](FuturesApi.md#CrossCollateralRepaymentHistoryUSERDATA) | **Get** /sapi/v1/futures/loan/repay/history | Cross-Collateral Repayment History (USER_DATA)
[**CrossCollateralWalletUSERDATA**](FuturesApi.md#CrossCollateralWalletUSERDATA) | **Get** /sapi/v1/futures/loan/wallet | Cross-Collateral Wallet (USER_DATA)
[**CrossCollateralWalletV2USERDATA**](FuturesApi.md#CrossCollateralWalletV2USERDATA) | **Get** /sapi/v2/futures/loan/wallet | Cross-Collateral Wallet V2 (USER_DATA)
[**GetCollateralRepayQuoteUSERDATA**](FuturesApi.md#GetCollateralRepayQuoteUSERDATA) | **Get** /sapi/v1/futures/loan/collateralRepay | Get Collateral Repay Quote (USER_DATA)
[**GetFutureAccountTransactionHistoryList**](FuturesApi.md#GetFutureAccountTransactionHistoryList) | **Get** /sapi/v1/futures/transfer | Get Future Account Transaction History List
[**GetMaxAmountforAdjustCrossCollateralLTVUSERDATA**](FuturesApi.md#GetMaxAmountforAdjustCrossCollateralLTVUSERDATA) | **Get** /sapi/v1/futures/loan/calcMaxAdjustAmount | Get Max Amount for Adjust Cross-Collateral LTV (USER_DATA)
[**GetMaxAmountforAdjustCrossCollateralLTVV2USERDATA**](FuturesApi.md#GetMaxAmountforAdjustCrossCollateralLTVV2USERDATA) | **Get** /sapi/v2/futures/loan/calcMaxAdjustAmount | Get Max Amount for Adjust Cross-Collateral LTV V2 (USER_DATA)
[**NewFutureAccountTransfer**](FuturesApi.md#NewFutureAccountTransfer) | **Post** /sapi/v1/futures/transfer | New Future Account Transfer
[**RepayForCrossCollateralTRADE**](FuturesApi.md#RepayForCrossCollateralTRADE) | **Post** /sapi/v1/futures/loan/repay | Repay For Cross-Collateral (TRADE)
[**RepaywithCollateralUSERDATA**](FuturesApi.md#RepaywithCollateralUSERDATA) | **Post** /sapi/v1/futures/loan/collateralRepay | Repay with Collateral (USER_DATA)

# **AdjustCrossCollateralLTVHistoryUSERDATA**
> AdjustCrossCollateralLTVHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Adjust Cross-Collateral LTV History (USER_DATA)

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

# **AdjustCrossCollateralLTVTRADE**
> AdjustCrossCollateralLTVTRADE(ctx, collateralCoin, amount, direction, timestamp, signature, contentType, xMBXAPIKEY)
Adjust Cross-Collateral LTV (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collateralCoin** | **string**|  | 
  **amount** | **string**|  | 
  **direction** | **string**| \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot; | 
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

# **AdjustCrossCollateralLTVV2TRADE**
> AdjustCrossCollateralLTVV2TRADE(ctx, loanCoin, collateralCoin, amount, direction, timestamp, signature, contentType, xMBXAPIKEY)
Adjust Cross-Collateral LTV V2 (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanCoin** | **string**|  | 
  **collateralCoin** | **string**|  | 
  **amount** | **string**|  | 
  **direction** | **string**| \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot; | 
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

# **BorrowForCrossCollateralTRADE**
> BorrowForCrossCollateralTRADE(ctx, coin, amount, collateralCoin, collateralAmount, timestamp, signature, contentType, xMBXAPIKEY)
Borrow For Cross-Collateral (TRADE)

limit: once per second per account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
  **amount** | **string**|  | 
  **collateralCoin** | **string**|  | 
  **collateralAmount** | **string**|  | 
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

# **CalculateRateAfterAdjustCrossCollateralLTVUSERDATA**
> CalculateRateAfterAdjustCrossCollateralLTVUSERDATA(ctx, collateralCoin, amount, direction, timestamp, signature, contentType, xMBXAPIKEY)
Calculate Rate After Adjust Cross-Collateral LTV (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collateralCoin** | **string**|  | 
  **amount** | **string**|  | 
  **direction** | **string**| \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot; | 
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

# **CalculateRateAfterAdjustCrossCollateralLTVV2USERDATA**
> CalculateRateAfterAdjustCrossCollateralLTVV2USERDATA(ctx, loanCoin, collateralCoin, amount, direction, timestamp, signature, contentType, xMBXAPIKEY)
Calculate Rate After Adjust Cross-Collateral LTV V2 (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanCoin** | **string**|  | 
  **collateralCoin** | **string**|  | 
  **amount** | **string**|  | 
  **direction** | **string**| \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot; | 
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

# **CheckCollateralRepayLimitUSERDATA**
> CheckCollateralRepayLimitUSERDATA(ctx, coin, collateralCoin, timestamp, signature, contentType, xMBXAPIKEY)
Check Collateral Repay Limit (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
  **collateralCoin** | **string**|  | 
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

# **CollateralRepaymentResultUSERDATA**
> CollateralRepaymentResultUSERDATA(ctx, quoteId, timestamp, signature, contentType, xMBXAPIKEY)
Collateral Repayment Result (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**|  | 
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

# **CrossCollateralBorrowHistoryUSERDATA**
> CrossCollateralBorrowHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Borrow History (USER_DATA)

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

# **CrossCollateralInformationUSERDATA**
> CrossCollateralInformationUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Information (USER_DATA)

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

# **CrossCollateralInformationV2USERDATA**
> CrossCollateralInformationV2USERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Information V2 (USER_DATA)

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

# **CrossCollateralInterestHistoryUSERDATA**
> CrossCollateralInterestHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Interest History (USER_DATA)

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

# **CrossCollateralLiquidationHistoryUSERDATA**
> CrossCollateralLiquidationHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Liquidation History (USER_DATA)

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

# **CrossCollateralRepaymentHistoryUSERDATA**
> CrossCollateralRepaymentHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Repayment History (USER_DATA)

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

# **CrossCollateralWalletUSERDATA**
> CrossCollateralWalletUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Wallet (USER_DATA)

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

# **CrossCollateralWalletV2USERDATA**
> CrossCollateralWalletV2USERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Cross-Collateral Wallet V2 (USER_DATA)

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

# **GetCollateralRepayQuoteUSERDATA**
> GetCollateralRepayQuoteUSERDATA(ctx, coin, collateralCoin, amount, timestamp, signature, contentType, xMBXAPIKEY)
Get Collateral Repay Quote (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
  **collateralCoin** | **string**|  | 
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

# **GetFutureAccountTransactionHistoryList**
> GetFutureAccountTransactionHistoryList(ctx, asset, startTime, timestamp, signature, contentType, xMBXAPIKEY)
Get Future Account Transaction History List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **startTime** | **string**|  | 
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

# **GetMaxAmountforAdjustCrossCollateralLTVUSERDATA**
> GetMaxAmountforAdjustCrossCollateralLTVUSERDATA(ctx, collateralCoin, timestamp, signature, contentType, xMBXAPIKEY)
Get Max Amount for Adjust Cross-Collateral LTV (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collateralCoin** | **string**|  | 
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

# **GetMaxAmountforAdjustCrossCollateralLTVV2USERDATA**
> GetMaxAmountforAdjustCrossCollateralLTVV2USERDATA(ctx, loanCoin, collateralCoin, timestamp, signature, contentType, xMBXAPIKEY)
Get Max Amount for Adjust Cross-Collateral LTV V2 (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanCoin** | **string**|  | 
  **collateralCoin** | **string**|  | 
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

# **NewFutureAccountTransfer**
> NewFutureAccountTransfer(ctx, asset, amount, type_, timestamp, signature, contentType, xMBXAPIKEY)
New Future Account Transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **amount** | **string**|  | 
  **type_** | **string**| 1: transfer from spot main account to future account 2: transfer from future account to spot main account | 
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

# **RepayForCrossCollateralTRADE**
> RepayForCrossCollateralTRADE(ctx, coin, collateralCoin, amount, timestamp, signature, contentType, xMBXAPIKEY)
Repay For Cross-Collateral (TRADE)

limit: once per second per account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
  **collateralCoin** | **string**|  | 
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

# **RepaywithCollateralUSERDATA**
> RepaywithCollateralUSERDATA(ctx, quoteId, timestamp, signature, contentType, xMBXAPIKEY)
Repay with Collateral (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**|  | 
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

