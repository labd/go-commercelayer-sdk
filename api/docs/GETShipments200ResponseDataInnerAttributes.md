# GETShipments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Unique identifier for the shipment | [optional] 
**Status** | Pointer to **string** | The shipment status, one of &#39;draft&#39;, &#39;upcoming&#39;, &#39;cancelled&#39;, &#39;on_hold&#39;, &#39;picking&#39;, &#39;packing&#39;, &#39;ready_to_ship&#39;, or &#39;shipped&#39; | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the associated order. | [optional] 
**CostAmountCents** | Pointer to **int32** | The cost of this shipment from the selected carrier account, in cents. | [optional] 
**CostAmountFloat** | Pointer to **float32** | The cost of this shipment from the selected carrier account, float. | [optional] 
**FormattedCostAmount** | Pointer to **string** | The cost of this shipment from the selected carrier account, formatted. | [optional] 
**SkusCount** | Pointer to **int32** | The total number of SKUs in the shipment&#39;s line items. This can be useful to display a preview of the shipment content. | [optional] 
**SelectedRateId** | Pointer to **string** | The selected purchase rate from the available shipping rates. | [optional] 
**Rates** | Pointer to **[]map[string]interface{}** | The available shipping rates. | [optional] 
**PurchaseErrorCode** | Pointer to **string** | The shipping rate purchase error code, if any. | [optional] 
**PurchaseErrorMessage** | Pointer to **string** | The shipping rate purchase error message, if any. | [optional] 
**GetRatesErrors** | Pointer to **[]map[string]interface{}** | Any errors collected when fetching shipping rates. | [optional] 
**GetRatesStartedAt** | Pointer to **string** | Time at which the getting of the shipping rates started. | [optional] 
**GetRatesCompletedAt** | Pointer to **string** | Time at which the getting of the shipping rates completed. | [optional] 
**PurchaseStartedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate started. | [optional] 
**PurchaseCompletedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate completed. | [optional] 
**PurchaseFailedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate failed. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETShipments200ResponseDataInnerAttributes

`func NewGETShipments200ResponseDataInnerAttributes() *GETShipments200ResponseDataInnerAttributes`

NewGETShipments200ResponseDataInnerAttributes instantiates a new GETShipments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShipments200ResponseDataInnerAttributesWithDefaults

`func NewGETShipments200ResponseDataInnerAttributesWithDefaults() *GETShipments200ResponseDataInnerAttributes`

NewGETShipments200ResponseDataInnerAttributesWithDefaults instantiates a new GETShipments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETShipments200ResponseDataInnerAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETShipments200ResponseDataInnerAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETShipments200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *GETShipments200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETShipments200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETShipments200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCents() int32`

GetCostAmountCents returns the CostAmountCents field if non-nil, zero value otherwise.

### GetCostAmountCentsOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCentsOk() (*int32, bool)`

GetCostAmountCentsOk returns a tuple with the CostAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountCents(v int32)`

SetCostAmountCents sets CostAmountCents field to given value.

### HasCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountCents() bool`

HasCostAmountCents returns a boolean if a field has been set.

### GetCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloat() float32`

GetCostAmountFloat returns the CostAmountFloat field if non-nil, zero value otherwise.

### GetCostAmountFloatOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloatOk() (*float32, bool)`

GetCostAmountFloatOk returns a tuple with the CostAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountFloat(v float32)`

SetCostAmountFloat sets CostAmountFloat field to given value.

### HasCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountFloat() bool`

HasCostAmountFloat returns a boolean if a field has been set.

### GetFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmount() string`

GetFormattedCostAmount returns the FormattedCostAmount field if non-nil, zero value otherwise.

### GetFormattedCostAmountOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmountOk() (*string, bool)`

GetFormattedCostAmountOk returns a tuple with the FormattedCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) SetFormattedCostAmount(v string)`

SetFormattedCostAmount sets FormattedCostAmount field to given value.

### HasFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) HasFormattedCostAmount() bool`

HasFormattedCostAmount returns a boolean if a field has been set.

### GetSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateId() string`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateIdOk() (*string, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) SetSelectedRateId(v string)`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### GetRates

`func (o *GETShipments200ResponseDataInnerAttributes) GetRates() []map[string]interface{}`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetRatesOk() (*[]map[string]interface{}, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GETShipments200ResponseDataInnerAttributes) SetRates(v []map[string]interface{})`

SetRates sets Rates field to given value.

### HasRates

`func (o *GETShipments200ResponseDataInnerAttributes) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCode() string`

GetPurchaseErrorCode returns the PurchaseErrorCode field if non-nil, zero value otherwise.

### GetPurchaseErrorCodeOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCodeOk() (*string, bool)`

GetPurchaseErrorCodeOk returns a tuple with the PurchaseErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorCode(v string)`

SetPurchaseErrorCode sets PurchaseErrorCode field to given value.

### HasPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorCode() bool`

HasPurchaseErrorCode returns a boolean if a field has been set.

### GetPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessage() string`

GetPurchaseErrorMessage returns the PurchaseErrorMessage field if non-nil, zero value otherwise.

### GetPurchaseErrorMessageOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessageOk() (*string, bool)`

GetPurchaseErrorMessageOk returns a tuple with the PurchaseErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorMessage(v string)`

SetPurchaseErrorMessage sets PurchaseErrorMessage field to given value.

### HasPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorMessage() bool`

HasPurchaseErrorMessage returns a boolean if a field has been set.

### GetGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrors() []map[string]interface{}`

GetGetRatesErrors returns the GetRatesErrors field if non-nil, zero value otherwise.

### GetGetRatesErrorsOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrorsOk() (*[]map[string]interface{}, bool)`

GetGetRatesErrorsOk returns a tuple with the GetRatesErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesErrors(v []map[string]interface{})`

SetGetRatesErrors sets GetRatesErrors field to given value.

### HasGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesErrors() bool`

HasGetRatesErrors returns a boolean if a field has been set.

### GetGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAt() string`

GetGetRatesStartedAt returns the GetRatesStartedAt field if non-nil, zero value otherwise.

### GetGetRatesStartedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAtOk() (*string, bool)`

GetGetRatesStartedAtOk returns a tuple with the GetRatesStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesStartedAt(v string)`

SetGetRatesStartedAt sets GetRatesStartedAt field to given value.

### HasGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesStartedAt() bool`

HasGetRatesStartedAt returns a boolean if a field has been set.

### GetGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAt() string`

GetGetRatesCompletedAt returns the GetRatesCompletedAt field if non-nil, zero value otherwise.

### GetGetRatesCompletedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAtOk() (*string, bool)`

GetGetRatesCompletedAtOk returns a tuple with the GetRatesCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesCompletedAt(v string)`

SetGetRatesCompletedAt sets GetRatesCompletedAt field to given value.

### HasGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesCompletedAt() bool`

HasGetRatesCompletedAt returns a boolean if a field has been set.

### GetPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAt() string`

GetPurchaseStartedAt returns the PurchaseStartedAt field if non-nil, zero value otherwise.

### GetPurchaseStartedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAtOk() (*string, bool)`

GetPurchaseStartedAtOk returns a tuple with the PurchaseStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseStartedAt(v string)`

SetPurchaseStartedAt sets PurchaseStartedAt field to given value.

### HasPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseStartedAt() bool`

HasPurchaseStartedAt returns a boolean if a field has been set.

### GetPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAt() string`

GetPurchaseCompletedAt returns the PurchaseCompletedAt field if non-nil, zero value otherwise.

### GetPurchaseCompletedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAtOk() (*string, bool)`

GetPurchaseCompletedAtOk returns a tuple with the PurchaseCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseCompletedAt(v string)`

SetPurchaseCompletedAt sets PurchaseCompletedAt field to given value.

### HasPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseCompletedAt() bool`

HasPurchaseCompletedAt returns a boolean if a field has been set.

### GetPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAt() string`

GetPurchaseFailedAt returns the PurchaseFailedAt field if non-nil, zero value otherwise.

### GetPurchaseFailedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAtOk() (*string, bool)`

GetPurchaseFailedAtOk returns a tuple with the PurchaseFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseFailedAt(v string)`

SetPurchaseFailedAt sets PurchaseFailedAt field to given value.

### HasPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseFailedAt() bool`

HasPurchaseFailedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETShipments200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShipments200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShipments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


