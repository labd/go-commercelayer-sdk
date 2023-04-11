# GETShipments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the shipment | [optional] 
**Status** | Pointer to **interface{}** | The shipment status, one of &#39;draft&#39;, &#39;upcoming&#39;, &#39;cancelled&#39;, &#39;on_hold&#39;, &#39;picking&#39;, &#39;packing&#39;, &#39;ready_to_ship&#39;, or &#39;shipped&#39; | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the associated order. | [optional] 
**CostAmountCents** | Pointer to **interface{}** | The cost of this shipment from the selected carrier account, in cents. | [optional] 
**CostAmountFloat** | Pointer to **interface{}** | The cost of this shipment from the selected carrier account, float. | [optional] 
**FormattedCostAmount** | Pointer to **interface{}** | The cost of this shipment from the selected carrier account, formatted. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the shipment&#39;s line items. This can be useful to display a preview of the shipment content. | [optional] 
**SelectedRateId** | Pointer to **interface{}** | The selected purchase rate from the available shipping rates. | [optional] 
**Rates** | Pointer to **interface{}** | The available shipping rates. | [optional] 
**PurchaseErrorCode** | Pointer to **interface{}** | The shipping rate purchase error code, if any. | [optional] 
**PurchaseErrorMessage** | Pointer to **interface{}** | The shipping rate purchase error message, if any. | [optional] 
**GetRatesErrors** | Pointer to **interface{}** | Any errors collected when fetching shipping rates. | [optional] 
**GetRatesStartedAt** | Pointer to **interface{}** | Time at which the getting of the shipping rates started. | [optional] 
**GetRatesCompletedAt** | Pointer to **interface{}** | Time at which the getting of the shipping rates completed. | [optional] 
**PurchaseStartedAt** | Pointer to **interface{}** | Time at which the purchasing of the shipping rate started. | [optional] 
**PurchaseCompletedAt** | Pointer to **interface{}** | Time at which the purchasing of the shipping rate completed. | [optional] 
**PurchaseFailedAt** | Pointer to **interface{}** | Time at which the purchasing of the shipping rate failed. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETShipments200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETShipments200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETShipments200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETShipments200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETShipments200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETShipments200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShipments200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCents() interface{}`

GetCostAmountCents returns the CostAmountCents field if non-nil, zero value otherwise.

### GetCostAmountCentsOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountCentsOk() (*interface{}, bool)`

GetCostAmountCentsOk returns a tuple with the CostAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountCents(v interface{})`

SetCostAmountCents sets CostAmountCents field to given value.

### HasCostAmountCents

`func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountCents() bool`

HasCostAmountCents returns a boolean if a field has been set.

### SetCostAmountCentsNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountCentsNil(b bool)`

 SetCostAmountCentsNil sets the value for CostAmountCents to be an explicit nil

### UnsetCostAmountCents
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetCostAmountCents()`

UnsetCostAmountCents ensures that no value is present for CostAmountCents, not even an explicit nil
### GetCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloat() interface{}`

GetCostAmountFloat returns the CostAmountFloat field if non-nil, zero value otherwise.

### GetCostAmountFloatOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCostAmountFloatOk() (*interface{}, bool)`

GetCostAmountFloatOk returns a tuple with the CostAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountFloat(v interface{})`

SetCostAmountFloat sets CostAmountFloat field to given value.

### HasCostAmountFloat

`func (o *GETShipments200ResponseDataInnerAttributes) HasCostAmountFloat() bool`

HasCostAmountFloat returns a boolean if a field has been set.

### SetCostAmountFloatNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetCostAmountFloatNil(b bool)`

 SetCostAmountFloatNil sets the value for CostAmountFloat to be an explicit nil

### UnsetCostAmountFloat
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetCostAmountFloat()`

UnsetCostAmountFloat ensures that no value is present for CostAmountFloat, not even an explicit nil
### GetFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmount() interface{}`

GetFormattedCostAmount returns the FormattedCostAmount field if non-nil, zero value otherwise.

### GetFormattedCostAmountOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetFormattedCostAmountOk() (*interface{}, bool)`

GetFormattedCostAmountOk returns a tuple with the FormattedCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) SetFormattedCostAmount(v interface{})`

SetFormattedCostAmount sets FormattedCostAmount field to given value.

### HasFormattedCostAmount

`func (o *GETShipments200ResponseDataInnerAttributes) HasFormattedCostAmount() bool`

HasFormattedCostAmount returns a boolean if a field has been set.

### SetFormattedCostAmountNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetFormattedCostAmountNil(b bool)`

 SetFormattedCostAmountNil sets the value for FormattedCostAmount to be an explicit nil

### UnsetFormattedCostAmount
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetFormattedCostAmount()`

UnsetFormattedCostAmount ensures that no value is present for FormattedCostAmount, not even an explicit nil
### GetSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETShipments200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateId() interface{}`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetSelectedRateIdOk() (*interface{}, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) SetSelectedRateId(v interface{})`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *GETShipments200ResponseDataInnerAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### SetSelectedRateIdNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetSelectedRateIdNil(b bool)`

 SetSelectedRateIdNil sets the value for SelectedRateId to be an explicit nil

### UnsetSelectedRateId
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetSelectedRateId()`

UnsetSelectedRateId ensures that no value is present for SelectedRateId, not even an explicit nil
### GetRates

`func (o *GETShipments200ResponseDataInnerAttributes) GetRates() interface{}`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetRatesOk() (*interface{}, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GETShipments200ResponseDataInnerAttributes) SetRates(v interface{})`

SetRates sets Rates field to given value.

### HasRates

`func (o *GETShipments200ResponseDataInnerAttributes) HasRates() bool`

HasRates returns a boolean if a field has been set.

### SetRatesNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetRatesNil(b bool)`

 SetRatesNil sets the value for Rates to be an explicit nil

### UnsetRates
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetRates()`

UnsetRates ensures that no value is present for Rates, not even an explicit nil
### GetPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCode() interface{}`

GetPurchaseErrorCode returns the PurchaseErrorCode field if non-nil, zero value otherwise.

### GetPurchaseErrorCodeOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorCodeOk() (*interface{}, bool)`

GetPurchaseErrorCodeOk returns a tuple with the PurchaseErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorCode(v interface{})`

SetPurchaseErrorCode sets PurchaseErrorCode field to given value.

### HasPurchaseErrorCode

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorCode() bool`

HasPurchaseErrorCode returns a boolean if a field has been set.

### SetPurchaseErrorCodeNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorCodeNil(b bool)`

 SetPurchaseErrorCodeNil sets the value for PurchaseErrorCode to be an explicit nil

### UnsetPurchaseErrorCode
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetPurchaseErrorCode()`

UnsetPurchaseErrorCode ensures that no value is present for PurchaseErrorCode, not even an explicit nil
### GetPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessage() interface{}`

GetPurchaseErrorMessage returns the PurchaseErrorMessage field if non-nil, zero value otherwise.

### GetPurchaseErrorMessageOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseErrorMessageOk() (*interface{}, bool)`

GetPurchaseErrorMessageOk returns a tuple with the PurchaseErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorMessage(v interface{})`

SetPurchaseErrorMessage sets PurchaseErrorMessage field to given value.

### HasPurchaseErrorMessage

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseErrorMessage() bool`

HasPurchaseErrorMessage returns a boolean if a field has been set.

### SetPurchaseErrorMessageNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseErrorMessageNil(b bool)`

 SetPurchaseErrorMessageNil sets the value for PurchaseErrorMessage to be an explicit nil

### UnsetPurchaseErrorMessage
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetPurchaseErrorMessage()`

UnsetPurchaseErrorMessage ensures that no value is present for PurchaseErrorMessage, not even an explicit nil
### GetGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrors() interface{}`

GetGetRatesErrors returns the GetRatesErrors field if non-nil, zero value otherwise.

### GetGetRatesErrorsOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesErrorsOk() (*interface{}, bool)`

GetGetRatesErrorsOk returns a tuple with the GetRatesErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesErrors(v interface{})`

SetGetRatesErrors sets GetRatesErrors field to given value.

### HasGetRatesErrors

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesErrors() bool`

HasGetRatesErrors returns a boolean if a field has been set.

### SetGetRatesErrorsNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesErrorsNil(b bool)`

 SetGetRatesErrorsNil sets the value for GetRatesErrors to be an explicit nil

### UnsetGetRatesErrors
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetGetRatesErrors()`

UnsetGetRatesErrors ensures that no value is present for GetRatesErrors, not even an explicit nil
### GetGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAt() interface{}`

GetGetRatesStartedAt returns the GetRatesStartedAt field if non-nil, zero value otherwise.

### GetGetRatesStartedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesStartedAtOk() (*interface{}, bool)`

GetGetRatesStartedAtOk returns a tuple with the GetRatesStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesStartedAt(v interface{})`

SetGetRatesStartedAt sets GetRatesStartedAt field to given value.

### HasGetRatesStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesStartedAt() bool`

HasGetRatesStartedAt returns a boolean if a field has been set.

### SetGetRatesStartedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesStartedAtNil(b bool)`

 SetGetRatesStartedAtNil sets the value for GetRatesStartedAt to be an explicit nil

### UnsetGetRatesStartedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetGetRatesStartedAt()`

UnsetGetRatesStartedAt ensures that no value is present for GetRatesStartedAt, not even an explicit nil
### GetGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAt() interface{}`

GetGetRatesCompletedAt returns the GetRatesCompletedAt field if non-nil, zero value otherwise.

### GetGetRatesCompletedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetGetRatesCompletedAtOk() (*interface{}, bool)`

GetGetRatesCompletedAtOk returns a tuple with the GetRatesCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesCompletedAt(v interface{})`

SetGetRatesCompletedAt sets GetRatesCompletedAt field to given value.

### HasGetRatesCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasGetRatesCompletedAt() bool`

HasGetRatesCompletedAt returns a boolean if a field has been set.

### SetGetRatesCompletedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetGetRatesCompletedAtNil(b bool)`

 SetGetRatesCompletedAtNil sets the value for GetRatesCompletedAt to be an explicit nil

### UnsetGetRatesCompletedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetGetRatesCompletedAt()`

UnsetGetRatesCompletedAt ensures that no value is present for GetRatesCompletedAt, not even an explicit nil
### GetPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAt() interface{}`

GetPurchaseStartedAt returns the PurchaseStartedAt field if non-nil, zero value otherwise.

### GetPurchaseStartedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseStartedAtOk() (*interface{}, bool)`

GetPurchaseStartedAtOk returns a tuple with the PurchaseStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseStartedAt(v interface{})`

SetPurchaseStartedAt sets PurchaseStartedAt field to given value.

### HasPurchaseStartedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseStartedAt() bool`

HasPurchaseStartedAt returns a boolean if a field has been set.

### SetPurchaseStartedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseStartedAtNil(b bool)`

 SetPurchaseStartedAtNil sets the value for PurchaseStartedAt to be an explicit nil

### UnsetPurchaseStartedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetPurchaseStartedAt()`

UnsetPurchaseStartedAt ensures that no value is present for PurchaseStartedAt, not even an explicit nil
### GetPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAt() interface{}`

GetPurchaseCompletedAt returns the PurchaseCompletedAt field if non-nil, zero value otherwise.

### GetPurchaseCompletedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseCompletedAtOk() (*interface{}, bool)`

GetPurchaseCompletedAtOk returns a tuple with the PurchaseCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseCompletedAt(v interface{})`

SetPurchaseCompletedAt sets PurchaseCompletedAt field to given value.

### HasPurchaseCompletedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseCompletedAt() bool`

HasPurchaseCompletedAt returns a boolean if a field has been set.

### SetPurchaseCompletedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseCompletedAtNil(b bool)`

 SetPurchaseCompletedAtNil sets the value for PurchaseCompletedAt to be an explicit nil

### UnsetPurchaseCompletedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetPurchaseCompletedAt()`

UnsetPurchaseCompletedAt ensures that no value is present for PurchaseCompletedAt, not even an explicit nil
### GetPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAt() interface{}`

GetPurchaseFailedAt returns the PurchaseFailedAt field if non-nil, zero value otherwise.

### GetPurchaseFailedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetPurchaseFailedAtOk() (*interface{}, bool)`

GetPurchaseFailedAtOk returns a tuple with the PurchaseFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseFailedAt(v interface{})`

SetPurchaseFailedAt sets PurchaseFailedAt field to given value.

### HasPurchaseFailedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasPurchaseFailedAt() bool`

HasPurchaseFailedAt returns a boolean if a field has been set.

### SetPurchaseFailedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetPurchaseFailedAtNil(b bool)`

 SetPurchaseFailedAtNil sets the value for PurchaseFailedAt to be an explicit nil

### UnsetPurchaseFailedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetPurchaseFailedAt()`

UnsetPurchaseFailedAt ensures that no value is present for PurchaseFailedAt, not even an explicit nil
### GetCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShipments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETShipments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShipments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShipments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShipments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShipments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShipments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETShipments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETShipments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


