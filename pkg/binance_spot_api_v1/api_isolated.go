
/*
 * Binance spot API
 *
 * Binance official supported Postman collections.<br/> - API documents: https://binance-docs.github.io/apidocs/spot/en/#change-log - Telegram: https://t.me/binance_api_english - Open Issue at: https://github.com/binance-exchange/binance-api-postman
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type IsolatedApiService service
/*
IsolatedApiService Create New Isolated Margin Account Listen Key
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) CreateNewIsolatedMarginAccountListenKey(ctx context.Context, symbol string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/userDataStream/isolated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Delete Isolated Margin Account Listen Key
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenKey
 * @param symbol
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) DeleteIsolatedMarginAccountListenKey(ctx context.Context, listenKey string, symbol string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/userDataStream/isolated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("listenKey", parameterToString(listenKey, ""))
	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Get All Isolated Margin Pairs (MARKET_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) GetAllIsolatedMarginPairsMARKETDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/allPairs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Get Isolated Margin Account Transfer History (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param asset
 * @param symbol
 * @param transFrom \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot;
 * @param transTo \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot;
 * @param startTime
 * @param endTime
 * @param current Currently querying page. Start from 1. Default:1
 * @param size Default:10 Max:100
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) GetIsolatedMarginAccountTransferHistoryUSERDATA(ctx context.Context, asset string, symbol string, transFrom string, transTo string, startTime string, endTime string, current int32, size int32, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/transfer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("asset", parameterToString(asset, ""))
	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("transFrom", parameterToString(transFrom, ""))
	localVarQueryParams.Add("transTo", parameterToString(transTo, ""))
	localVarQueryParams.Add("startTime", parameterToString(startTime, ""))
	localVarQueryParams.Add("endTime", parameterToString(endTime, ""))
	localVarQueryParams.Add("current", parameterToString(current, ""))
	localVarQueryParams.Add("size", parameterToString(size, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Isolated Margin Account Transfer (MARGIN)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param asset
 * @param symbol isolated symbol
 * @param transFrom \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot;
 * @param transTo \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot;
 * @param amount
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) IsolatedMarginAccountTransferMARGIN(ctx context.Context, asset string, symbol string, transFrom string, transTo string, amount string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/transfer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("asset", parameterToString(asset, ""))
	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("transFrom", parameterToString(transFrom, ""))
	localVarQueryParams.Add("transTo", parameterToString(transTo, ""))
	localVarQueryParams.Add("amount", parameterToString(amount, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Query Isolated Margin Account Info (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbols Max 5 symbols can be sent; separated by &#x27;,&#x27;. e.g. &#x27;BTCUSDT,BNBUSDT,ADAUSDT&#x27;
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) QueryIsolatedMarginAccountInfoUSERDATA(ctx context.Context, symbols string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbols", parameterToString(symbols, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Query Isolated Margin Pair (MARKET_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) QueryIsolatedMarginPairMARKETDATA(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/pair"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
IsolatedApiService Renew Isolated Margin Account Listen Key
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenKey
 * @param symbol
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *IsolatedApiService) RenewIsolatedMarginAccountListenKey(ctx context.Context, listenKey string, symbol string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/userDataStream/isolated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("listenKey", parameterToString(listenKey, ""))
	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
