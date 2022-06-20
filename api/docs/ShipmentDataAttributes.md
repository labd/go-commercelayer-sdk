# ShipmentDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Unique identifier for the shipment | [optional] 
**Status** | Pointer to **string** | The shipment status, one of &#39;draft&#39;, &#39;upcoming&#39;, &#39;cancelled&#39;, &#39;on_hold&#39;, &#39;picking&#39;, &#39;packing&#39;, &#39;ready_to_ship&#39;, or &#39;shipped&#39; | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the associated order. | [optional] 
**CostAmountCents** | Pointer to **int32** | The cost of this shipment from the selected carrier account, in cents. | [optional] 
**CostAmountFloat** | Pointer to **float32** | The cost of this shipment from the selected carrier account, float. | [optional] 
**FormattedCostAmount** | Pointer to **string** | The cost of this shipment from the selected carrier account, formatted. | [optional] 
**SkusCount** | Pointer to **int32** | The total number of skus in the shipment&#39;s line items. This can be useful to display a preview of the shipment content. | [optional] 
**SelectedRateId** | Pointer to **string** | The selected purchase rate from the available shipping rates. | [optional] 
**Rates** | Pointer to **[]map[string]interface{}** | The available shipping rates. | [optional] 
**PurchaseErrorCode** | Pointer to **string** | The shipping rate purchase error code, if any. | [optional] 
**PurchaseErrorMessage** | Pointer to **string** | The shipping rate purchase error message, if any. | [optional] 
**GetRatesStartedAt** | Pointer to **string** | Time at which the getting of the shipping rates started. | [optional] 
**GetRatesCompletedAt** | Pointer to **string** | Time at which the getting of the shipping rates completed. | [optional] 
**PurchaseStartedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate started. | [optional] 
**PurchaseCompletedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate completed. | [optional] 
**PurchaseFailedAt** | Pointer to **string** | Time at which the purchasing of the shipping rate failed. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShipmentDataAttributes

`func NewShipmentDataAttributes() *ShipmentDataAttributes`

NewShipmentDataAttributes instantiates a new ShipmentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDataAttributesWithDefaults

`func NewShipmentDataAttributesWithDefaults() *ShipmentDataAttributes`

NewShipmentDataAttributesWithDefaults instantiates a new ShipmentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ShipmentDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ShipmentDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ShipmentDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ShipmentDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *ShipmentDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipmentDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ShipmentDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ShipmentDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ShipmentDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ShipmentDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCostAmountCents

`func (o *ShipmentDataAttributes) GetCostAmountCents() int32`

GetCostAmountCents returns the CostAmountCents field if non-nil, zero value otherwise.

### GetCostAmountCentsOk

`func (o *ShipmentDataAttributes) GetCostAmountCentsOk() (*int32, bool)`

GetCostAmountCentsOk returns a tuple with the CostAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountCents

`func (o *ShipmentDataAttributes) SetCostAmountCents(v int32)`

SetCostAmountCents sets CostAmountCents field to given value.

### HasCostAmountCents

`func (o *ShipmentDataAttributes) HasCostAmountCents() bool`

HasCostAmountCents returns a boolean if a field has been set.

### GetCostAmountFloat

`func (o *ShipmentDataAttributes) GetCostAmountFloat() float32`

GetCostAmountFloat returns the CostAmountFloat field if non-nil, zero value otherwise.

### GetCostAmountFloatOk

`func (o *ShipmentDataAttributes) GetCostAmountFloatOk() (*float32, bool)`

GetCostAmountFloatOk returns a tuple with the CostAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAmountFloat

`func (o *ShipmentDataAttributes) SetCostAmountFloat(v float32)`

SetCostAmountFloat sets CostAmountFloat field to given value.

### HasCostAmountFloat

`func (o *ShipmentDataAttributes) HasCostAmountFloat() bool`

HasCostAmountFloat returns a boolean if a field has been set.

### GetFormattedCostAmount

`func (o *ShipmentDataAttributes) GetFormattedCostAmount() string`

GetFormattedCostAmount returns the FormattedCostAmount field if non-nil, zero value otherwise.

### GetFormattedCostAmountOk

`func (o *ShipmentDataAttributes) GetFormattedCostAmountOk() (*string, bool)`

GetFormattedCostAmountOk returns a tuple with the FormattedCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCostAmount

`func (o *ShipmentDataAttributes) SetFormattedCostAmount(v string)`

SetFormattedCostAmount sets FormattedCostAmount field to given value.

### HasFormattedCostAmount

`func (o *ShipmentDataAttributes) HasFormattedCostAmount() bool`

HasFormattedCostAmount returns a boolean if a field has been set.

### GetSkusCount

`func (o *ShipmentDataAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *ShipmentDataAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *ShipmentDataAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *ShipmentDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetSelectedRateId

`func (o *ShipmentDataAttributes) GetSelectedRateId() string`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *ShipmentDataAttributes) GetSelectedRateIdOk() (*string, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *ShipmentDataAttributes) SetSelectedRateId(v string)`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *ShipmentDataAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### GetRates

`func (o *ShipmentDataAttributes) GetRates() []map[string]interface{}`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ShipmentDataAttributes) GetRatesOk() (*[]map[string]interface{}, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ShipmentDataAttributes) SetRates(v []map[string]interface{})`

SetRates sets Rates field to given value.

### HasRates

`func (o *ShipmentDataAttributes) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetPurchaseErrorCode

`func (o *ShipmentDataAttributes) GetPurchaseErrorCode() string`

GetPurchaseErrorCode returns the PurchaseErrorCode field if non-nil, zero value otherwise.

### GetPurchaseErrorCodeOk

`func (o *ShipmentDataAttributes) GetPurchaseErrorCodeOk() (*string, bool)`

GetPurchaseErrorCodeOk returns a tuple with the PurchaseErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorCode

`func (o *ShipmentDataAttributes) SetPurchaseErrorCode(v string)`

SetPurchaseErrorCode sets PurchaseErrorCode field to given value.

### HasPurchaseErrorCode

`func (o *ShipmentDataAttributes) HasPurchaseErrorCode() bool`

HasPurchaseErrorCode returns a boolean if a field has been set.

### GetPurchaseErrorMessage

`func (o *ShipmentDataAttributes) GetPurchaseErrorMessage() string`

GetPurchaseErrorMessage returns the PurchaseErrorMessage field if non-nil, zero value otherwise.

### GetPurchaseErrorMessageOk

`func (o *ShipmentDataAttributes) GetPurchaseErrorMessageOk() (*string, bool)`

GetPurchaseErrorMessageOk returns a tuple with the PurchaseErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseErrorMessage

`func (o *ShipmentDataAttributes) SetPurchaseErrorMessage(v string)`

SetPurchaseErrorMessage sets PurchaseErrorMessage field to given value.

### HasPurchaseErrorMessage

`func (o *ShipmentDataAttributes) HasPurchaseErrorMessage() bool`

HasPurchaseErrorMessage returns a boolean if a field has been set.

### GetGetRatesStartedAt

`func (o *ShipmentDataAttributes) GetGetRatesStartedAt() string`

GetGetRatesStartedAt returns the GetRatesStartedAt field if non-nil, zero value otherwise.

### GetGetRatesStartedAtOk

`func (o *ShipmentDataAttributes) GetGetRatesStartedAtOk() (*string, bool)`

GetGetRatesStartedAtOk returns a tuple with the GetRatesStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesStartedAt

`func (o *ShipmentDataAttributes) SetGetRatesStartedAt(v string)`

SetGetRatesStartedAt sets GetRatesStartedAt field to given value.

### HasGetRatesStartedAt

`func (o *ShipmentDataAttributes) HasGetRatesStartedAt() bool`

HasGetRatesStartedAt returns a boolean if a field has been set.

### GetGetRatesCompletedAt

`func (o *ShipmentDataAttributes) GetGetRatesCompletedAt() string`

GetGetRatesCompletedAt returns the GetRatesCompletedAt field if non-nil, zero value otherwise.

### GetGetRatesCompletedAtOk

`func (o *ShipmentDataAttributes) GetGetRatesCompletedAtOk() (*string, bool)`

GetGetRatesCompletedAtOk returns a tuple with the GetRatesCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRatesCompletedAt

`func (o *ShipmentDataAttributes) SetGetRatesCompletedAt(v string)`

SetGetRatesCompletedAt sets GetRatesCompletedAt field to given value.

### HasGetRatesCompletedAt

`func (o *ShipmentDataAttributes) HasGetRatesCompletedAt() bool`

HasGetRatesCompletedAt returns a boolean if a field has been set.

### GetPurchaseStartedAt

`func (o *ShipmentDataAttributes) GetPurchaseStartedAt() string`

GetPurchaseStartedAt returns the PurchaseStartedAt field if non-nil, zero value otherwise.

### GetPurchaseStartedAtOk

`func (o *ShipmentDataAttributes) GetPurchaseStartedAtOk() (*string, bool)`

GetPurchaseStartedAtOk returns a tuple with the PurchaseStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseStartedAt

`func (o *ShipmentDataAttributes) SetPurchaseStartedAt(v string)`

SetPurchaseStartedAt sets PurchaseStartedAt field to given value.

### HasPurchaseStartedAt

`func (o *ShipmentDataAttributes) HasPurchaseStartedAt() bool`

HasPurchaseStartedAt returns a boolean if a field has been set.

### GetPurchaseCompletedAt

`func (o *ShipmentDataAttributes) GetPurchaseCompletedAt() string`

GetPurchaseCompletedAt returns the PurchaseCompletedAt field if non-nil, zero value otherwise.

### GetPurchaseCompletedAtOk

`func (o *ShipmentDataAttributes) GetPurchaseCompletedAtOk() (*string, bool)`

GetPurchaseCompletedAtOk returns a tuple with the PurchaseCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCompletedAt

`func (o *ShipmentDataAttributes) SetPurchaseCompletedAt(v string)`

SetPurchaseCompletedAt sets PurchaseCompletedAt field to given value.

### HasPurchaseCompletedAt

`func (o *ShipmentDataAttributes) HasPurchaseCompletedAt() bool`

HasPurchaseCompletedAt returns a boolean if a field has been set.

### GetPurchaseFailedAt

`func (o *ShipmentDataAttributes) GetPurchaseFailedAt() string`

GetPurchaseFailedAt returns the PurchaseFailedAt field if non-nil, zero value otherwise.

### GetPurchaseFailedAtOk

`func (o *ShipmentDataAttributes) GetPurchaseFailedAtOk() (*string, bool)`

GetPurchaseFailedAtOk returns a tuple with the PurchaseFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseFailedAt

`func (o *ShipmentDataAttributes) SetPurchaseFailedAt(v string)`

SetPurchaseFailedAt sets PurchaseFailedAt field to given value.

### HasPurchaseFailedAt

`func (o *ShipmentDataAttributes) HasPurchaseFailedAt() bool`

HasPurchaseFailedAt returns a boolean if a field has been set.

### GetId

`func (o *ShipmentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShipmentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShipmentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShipmentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShipmentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ShipmentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShipmentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShipmentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ShipmentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ShipmentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShipmentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShipmentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShipmentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShipmentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShipmentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShipmentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShipmentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


