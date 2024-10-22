# GETShipmentsShipmentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the shipment. Cannot be passed by sales channels. | [optional] 
**Status** | Pointer to **interface{}** | The shipment status. One of &#39;draft&#39; (default), &#39;upcoming&#39;, &#39;cancelled&#39;, &#39;on_hold&#39;, &#39;picking&#39;, &#39;packing&#39;, &#39;ready_to_ship&#39;, &#39;shipped&#39;, or &#39;delivered&#39;. | [optional] 
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
**OnHoldAt** | Pointer to **interface{}** | Time at which the shipment was put on hold. | [optional] 
**PickingAt** | Pointer to **interface{}** | Time at which the shipment was picking. | [optional] 
**PackingAt** | Pointer to **interface{}** | Time at which the shipment was packing. | [optional] 
**ReadyToShipAt** | Pointer to **interface{}** | Time at which the shipment was ready to ship. | [optional] 
**ShippedAt** | Pointer to **interface{}** | Time at which the shipment was shipped. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETShipmentsShipmentId200ResponseDataAttributes

`func NewGETShipmentsShipmentId200ResponseDataAttributes() *GETShipmentsShipmentId200ResponseDataAttributes`

NewGETShipmentsShipmentId200ResponseDataAttributes instantiates a new GETShipmentsShipmentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShipmentsShipmentId200ResponseDataAttributesWithDefaults

`func NewGETShipmentsShipmentId200ResponseDataAttributesWithDefaults() *GETShipmentsShipmentId200ResponseDataAttributes`

NewGETShipmentsShipmentId200ResponseDataAttributesWithDefaults instantiates a new GETShipmentsShipmentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCurrencyCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCostAmountCents

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCostAmountCents() interface{}`

GetCostAmountCents returns the CostAmountCents field if non-nil, zero value otherwise.

### GetCostAmountCentsOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCostAmountCentsOk() (*interface{}, bool)`

GetCostAmountCentsOk returns a tuple with the CostAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountCents

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCostAmountCents(v interface{})`

SetCostAmountCents sets CostAmountCents field to given value.

### HasCostAmountCents

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasCostAmountCents() bool`

HasCostAmountCents returns a boolean if a field has been set.

### SetCostAmountCentsNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCostAmountCentsNil(b bool)`

 SetCostAmountCentsNil sets the value for CostAmountCents to be an explicit nil

### UnsetCostAmountCents
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetCostAmountCents()`

UnsetCostAmountCents ensures that no value is present for CostAmountCents, not even an explicit nil
### GetCostAmountFloat

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCostAmountFloat() interface{}`

GetCostAmountFloat returns the CostAmountFloat field if non-nil, zero value otherwise.

### GetCostAmountFloatOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCostAmountFloatOk() (*interface{}, bool)`

GetCostAmountFloatOk returns a tuple with the CostAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountFloat

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCostAmountFloat(v interface{})`

SetCostAmountFloat sets CostAmountFloat field to given value.

### HasCostAmountFloat

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasCostAmountFloat() bool`

HasCostAmountFloat returns a boolean if a field has been set.

### SetCostAmountFloatNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCostAmountFloatNil(b bool)`

 SetCostAmountFloatNil sets the value for CostAmountFloat to be an explicit nil

### UnsetCostAmountFloat
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetCostAmountFloat()`

UnsetCostAmountFloat ensures that no value is present for CostAmountFloat, not even an explicit nil
### GetFormattedCostAmount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetFormattedCostAmount() interface{}`

GetFormattedCostAmount returns the FormattedCostAmount field if non-nil, zero value otherwise.

### GetFormattedCostAmountOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetFormattedCostAmountOk() (*interface{}, bool)`

GetFormattedCostAmountOk returns a tuple with the FormattedCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCostAmount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetFormattedCostAmount(v interface{})`

SetFormattedCostAmount sets FormattedCostAmount field to given value.

### HasFormattedCostAmount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasFormattedCostAmount() bool`

HasFormattedCostAmount returns a boolean if a field has been set.

### SetFormattedCostAmountNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetFormattedCostAmountNil(b bool)`

 SetFormattedCostAmountNil sets the value for FormattedCostAmount to be an explicit nil

### UnsetFormattedCostAmount
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetFormattedCostAmount()`

UnsetFormattedCostAmount ensures that no value is present for FormattedCostAmount, not even an explicit nil
### GetSkusCount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetSelectedRateId

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateId() interface{}`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateIdOk() (*interface{}, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateId(v interface{})`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### SetSelectedRateIdNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateIdNil(b bool)`

 SetSelectedRateIdNil sets the value for SelectedRateId to be an explicit nil

### UnsetSelectedRateId
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetSelectedRateId()`

UnsetSelectedRateId ensures that no value is present for SelectedRateId, not even an explicit nil
### GetRates

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetRates() interface{}`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetRatesOk() (*interface{}, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetRates(v interface{})`

SetRates sets Rates field to given value.

### HasRates

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasRates() bool`

HasRates returns a boolean if a field has been set.

### SetRatesNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetRatesNil(b bool)`

 SetRatesNil sets the value for Rates to be an explicit nil

### UnsetRates
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetRates()`

UnsetRates ensures that no value is present for Rates, not even an explicit nil
### GetPurchaseErrorCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseErrorCode() interface{}`

GetPurchaseErrorCode returns the PurchaseErrorCode field if non-nil, zero value otherwise.

### GetPurchaseErrorCodeOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseErrorCodeOk() (*interface{}, bool)`

GetPurchaseErrorCodeOk returns a tuple with the PurchaseErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseErrorCode(v interface{})`

SetPurchaseErrorCode sets PurchaseErrorCode field to given value.

### HasPurchaseErrorCode

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPurchaseErrorCode() bool`

HasPurchaseErrorCode returns a boolean if a field has been set.

### SetPurchaseErrorCodeNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseErrorCodeNil(b bool)`

 SetPurchaseErrorCodeNil sets the value for PurchaseErrorCode to be an explicit nil

### UnsetPurchaseErrorCode
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPurchaseErrorCode()`

UnsetPurchaseErrorCode ensures that no value is present for PurchaseErrorCode, not even an explicit nil
### GetPurchaseErrorMessage

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseErrorMessage() interface{}`

GetPurchaseErrorMessage returns the PurchaseErrorMessage field if non-nil, zero value otherwise.

### GetPurchaseErrorMessageOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseErrorMessageOk() (*interface{}, bool)`

GetPurchaseErrorMessageOk returns a tuple with the PurchaseErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorMessage

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseErrorMessage(v interface{})`

SetPurchaseErrorMessage sets PurchaseErrorMessage field to given value.

### HasPurchaseErrorMessage

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPurchaseErrorMessage() bool`

HasPurchaseErrorMessage returns a boolean if a field has been set.

### SetPurchaseErrorMessageNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseErrorMessageNil(b bool)`

 SetPurchaseErrorMessageNil sets the value for PurchaseErrorMessage to be an explicit nil

### UnsetPurchaseErrorMessage
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPurchaseErrorMessage()`

UnsetPurchaseErrorMessage ensures that no value is present for PurchaseErrorMessage, not even an explicit nil
### GetGetRatesErrors

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesErrors() interface{}`

GetGetRatesErrors returns the GetRatesErrors field if non-nil, zero value otherwise.

### GetGetRatesErrorsOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesErrorsOk() (*interface{}, bool)`

GetGetRatesErrorsOk returns a tuple with the GetRatesErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesErrors

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesErrors(v interface{})`

SetGetRatesErrors sets GetRatesErrors field to given value.

### HasGetRatesErrors

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasGetRatesErrors() bool`

HasGetRatesErrors returns a boolean if a field has been set.

### SetGetRatesErrorsNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesErrorsNil(b bool)`

 SetGetRatesErrorsNil sets the value for GetRatesErrors to be an explicit nil

### UnsetGetRatesErrors
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetGetRatesErrors()`

UnsetGetRatesErrors ensures that no value is present for GetRatesErrors, not even an explicit nil
### GetGetRatesStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesStartedAt() interface{}`

GetGetRatesStartedAt returns the GetRatesStartedAt field if non-nil, zero value otherwise.

### GetGetRatesStartedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesStartedAtOk() (*interface{}, bool)`

GetGetRatesStartedAtOk returns a tuple with the GetRatesStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesStartedAt(v interface{})`

SetGetRatesStartedAt sets GetRatesStartedAt field to given value.

### HasGetRatesStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasGetRatesStartedAt() bool`

HasGetRatesStartedAt returns a boolean if a field has been set.

### SetGetRatesStartedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesStartedAtNil(b bool)`

 SetGetRatesStartedAtNil sets the value for GetRatesStartedAt to be an explicit nil

### UnsetGetRatesStartedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetGetRatesStartedAt()`

UnsetGetRatesStartedAt ensures that no value is present for GetRatesStartedAt, not even an explicit nil
### GetGetRatesCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesCompletedAt() interface{}`

GetGetRatesCompletedAt returns the GetRatesCompletedAt field if non-nil, zero value otherwise.

### GetGetRatesCompletedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetGetRatesCompletedAtOk() (*interface{}, bool)`

GetGetRatesCompletedAtOk returns a tuple with the GetRatesCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesCompletedAt(v interface{})`

SetGetRatesCompletedAt sets GetRatesCompletedAt field to given value.

### HasGetRatesCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasGetRatesCompletedAt() bool`

HasGetRatesCompletedAt returns a boolean if a field has been set.

### SetGetRatesCompletedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetGetRatesCompletedAtNil(b bool)`

 SetGetRatesCompletedAtNil sets the value for GetRatesCompletedAt to be an explicit nil

### UnsetGetRatesCompletedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetGetRatesCompletedAt()`

UnsetGetRatesCompletedAt ensures that no value is present for GetRatesCompletedAt, not even an explicit nil
### GetPurchaseStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseStartedAt() interface{}`

GetPurchaseStartedAt returns the PurchaseStartedAt field if non-nil, zero value otherwise.

### GetPurchaseStartedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseStartedAtOk() (*interface{}, bool)`

GetPurchaseStartedAtOk returns a tuple with the PurchaseStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseStartedAt(v interface{})`

SetPurchaseStartedAt sets PurchaseStartedAt field to given value.

### HasPurchaseStartedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPurchaseStartedAt() bool`

HasPurchaseStartedAt returns a boolean if a field has been set.

### SetPurchaseStartedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseStartedAtNil(b bool)`

 SetPurchaseStartedAtNil sets the value for PurchaseStartedAt to be an explicit nil

### UnsetPurchaseStartedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPurchaseStartedAt()`

UnsetPurchaseStartedAt ensures that no value is present for PurchaseStartedAt, not even an explicit nil
### GetPurchaseCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseCompletedAt() interface{}`

GetPurchaseCompletedAt returns the PurchaseCompletedAt field if non-nil, zero value otherwise.

### GetPurchaseCompletedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseCompletedAtOk() (*interface{}, bool)`

GetPurchaseCompletedAtOk returns a tuple with the PurchaseCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseCompletedAt(v interface{})`

SetPurchaseCompletedAt sets PurchaseCompletedAt field to given value.

### HasPurchaseCompletedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPurchaseCompletedAt() bool`

HasPurchaseCompletedAt returns a boolean if a field has been set.

### SetPurchaseCompletedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseCompletedAtNil(b bool)`

 SetPurchaseCompletedAtNil sets the value for PurchaseCompletedAt to be an explicit nil

### UnsetPurchaseCompletedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPurchaseCompletedAt()`

UnsetPurchaseCompletedAt ensures that no value is present for PurchaseCompletedAt, not even an explicit nil
### GetPurchaseFailedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseFailedAt() interface{}`

GetPurchaseFailedAt returns the PurchaseFailedAt field if non-nil, zero value otherwise.

### GetPurchaseFailedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPurchaseFailedAtOk() (*interface{}, bool)`

GetPurchaseFailedAtOk returns a tuple with the PurchaseFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseFailedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseFailedAt(v interface{})`

SetPurchaseFailedAt sets PurchaseFailedAt field to given value.

### HasPurchaseFailedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPurchaseFailedAt() bool`

HasPurchaseFailedAt returns a boolean if a field has been set.

### SetPurchaseFailedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPurchaseFailedAtNil(b bool)`

 SetPurchaseFailedAtNil sets the value for PurchaseFailedAt to be an explicit nil

### UnsetPurchaseFailedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPurchaseFailedAt()`

UnsetPurchaseFailedAt ensures that no value is present for PurchaseFailedAt, not even an explicit nil
### GetOnHoldAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetOnHoldAt() interface{}`

GetOnHoldAt returns the OnHoldAt field if non-nil, zero value otherwise.

### GetOnHoldAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetOnHoldAtOk() (*interface{}, bool)`

GetOnHoldAtOk returns a tuple with the OnHoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHoldAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetOnHoldAt(v interface{})`

SetOnHoldAt sets OnHoldAt field to given value.

### HasOnHoldAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasOnHoldAt() bool`

HasOnHoldAt returns a boolean if a field has been set.

### SetOnHoldAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetOnHoldAtNil(b bool)`

 SetOnHoldAtNil sets the value for OnHoldAt to be an explicit nil

### UnsetOnHoldAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetOnHoldAt()`

UnsetOnHoldAt ensures that no value is present for OnHoldAt, not even an explicit nil
### GetPickingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPickingAt() interface{}`

GetPickingAt returns the PickingAt field if non-nil, zero value otherwise.

### GetPickingAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPickingAtOk() (*interface{}, bool)`

GetPickingAtOk returns a tuple with the PickingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPickingAt(v interface{})`

SetPickingAt sets PickingAt field to given value.

### HasPickingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPickingAt() bool`

HasPickingAt returns a boolean if a field has been set.

### SetPickingAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPickingAtNil(b bool)`

 SetPickingAtNil sets the value for PickingAt to be an explicit nil

### UnsetPickingAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPickingAt()`

UnsetPickingAt ensures that no value is present for PickingAt, not even an explicit nil
### GetPackingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPackingAt() interface{}`

GetPackingAt returns the PackingAt field if non-nil, zero value otherwise.

### GetPackingAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetPackingAtOk() (*interface{}, bool)`

GetPackingAtOk returns a tuple with the PackingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPackingAt(v interface{})`

SetPackingAt sets PackingAt field to given value.

### HasPackingAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasPackingAt() bool`

HasPackingAt returns a boolean if a field has been set.

### SetPackingAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetPackingAtNil(b bool)`

 SetPackingAtNil sets the value for PackingAt to be an explicit nil

### UnsetPackingAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetPackingAt()`

UnsetPackingAt ensures that no value is present for PackingAt, not even an explicit nil
### GetReadyToShipAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReadyToShipAt() interface{}`

GetReadyToShipAt returns the ReadyToShipAt field if non-nil, zero value otherwise.

### GetReadyToShipAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReadyToShipAtOk() (*interface{}, bool)`

GetReadyToShipAtOk returns a tuple with the ReadyToShipAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToShipAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReadyToShipAt(v interface{})`

SetReadyToShipAt sets ReadyToShipAt field to given value.

### HasReadyToShipAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasReadyToShipAt() bool`

HasReadyToShipAt returns a boolean if a field has been set.

### SetReadyToShipAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReadyToShipAtNil(b bool)`

 SetReadyToShipAtNil sets the value for ReadyToShipAt to be an explicit nil

### UnsetReadyToShipAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetReadyToShipAt()`

UnsetReadyToShipAt ensures that no value is present for ReadyToShipAt, not even an explicit nil
### GetShippedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetShippedAt() interface{}`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetShippedAtOk() (*interface{}, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetShippedAt(v interface{})`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetCreatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETShipmentsShipmentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETShipmentsShipmentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


