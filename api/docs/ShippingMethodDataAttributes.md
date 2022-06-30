# ShippingMethodDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The shipping method&#39;s name | [optional] 
**Scheme** | Pointer to **string** | The shipping method&#39;s scheme, one of &#39;flat&#39; or &#39;weight_tiered&#39;. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**DisabledAt** | Pointer to **string** | Time at which the shipping method was disabled. | [optional] 
**PriceAmountCents** | Pointer to **int32** | The price of this shipping method, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **float32** | The price of this shipping method, float. | [optional] 
**FormattedPriceAmount** | Pointer to **string** | The price of this shipping method, formatted. | [optional] 
**FreeOverAmountCents** | Pointer to **int32** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**FreeOverAmountFloat** | Pointer to **float32** | Apply free shipping if the order amount is over this value, float. | [optional] 
**FormattedFreeOverAmount** | Pointer to **string** | Apply free shipping if the order amount is over this value, formatted. | [optional] 
**PriceAmountForShipmentCents** | Pointer to **int32** | The calculated price (zero or price amount) when associated to a shipment, in cents. | [optional] 
**PriceAmountForShipmentFloat** | Pointer to **float32** | The calculated price (zero or price amount) when associated to a shipment, float. | [optional] 
**FormattedPriceAmountForShipment** | Pointer to **string** | The calculated price (zero or price amount) when associated to a shipment, formatted. | [optional] 
**MinWeight** | Pointer to **float32** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **float32** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShippingMethodDataAttributes

`func NewShippingMethodDataAttributes() *ShippingMethodDataAttributes`

NewShippingMethodDataAttributes instantiates a new ShippingMethodDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodDataAttributesWithDefaults

`func NewShippingMethodDataAttributesWithDefaults() *ShippingMethodDataAttributes`

NewShippingMethodDataAttributesWithDefaults instantiates a new ShippingMethodDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingMethodDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethodDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethodDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingMethodDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheme

`func (o *ShippingMethodDataAttributes) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ShippingMethodDataAttributes) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ShippingMethodDataAttributes) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ShippingMethodDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ShippingMethodDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ShippingMethodDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ShippingMethodDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ShippingMethodDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDisabledAt

`func (o *ShippingMethodDataAttributes) GetDisabledAt() string`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *ShippingMethodDataAttributes) GetDisabledAtOk() (*string, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *ShippingMethodDataAttributes) SetDisabledAt(v string)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *ShippingMethodDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *ShippingMethodDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *ShippingMethodDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *ShippingMethodDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *ShippingMethodDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetPriceAmountFloat

`func (o *ShippingMethodDataAttributes) GetPriceAmountFloat() float32`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *ShippingMethodDataAttributes) GetPriceAmountFloatOk() (*float32, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *ShippingMethodDataAttributes) SetPriceAmountFloat(v float32)`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *ShippingMethodDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### GetFormattedPriceAmount

`func (o *ShippingMethodDataAttributes) GetFormattedPriceAmount() string`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountOk() (*string, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *ShippingMethodDataAttributes) SetFormattedPriceAmount(v string)`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *ShippingMethodDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### GetFreeOverAmountCents

`func (o *ShippingMethodDataAttributes) GetFreeOverAmountCents() int32`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *ShippingMethodDataAttributes) GetFreeOverAmountCentsOk() (*int32, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *ShippingMethodDataAttributes) SetFreeOverAmountCents(v int32)`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *ShippingMethodDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### GetFreeOverAmountFloat

`func (o *ShippingMethodDataAttributes) GetFreeOverAmountFloat() float32`

GetFreeOverAmountFloat returns the FreeOverAmountFloat field if non-nil, zero value otherwise.

### GetFreeOverAmountFloatOk

`func (o *ShippingMethodDataAttributes) GetFreeOverAmountFloatOk() (*float32, bool)`

GetFreeOverAmountFloatOk returns a tuple with the FreeOverAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountFloat

`func (o *ShippingMethodDataAttributes) SetFreeOverAmountFloat(v float32)`

SetFreeOverAmountFloat sets FreeOverAmountFloat field to given value.

### HasFreeOverAmountFloat

`func (o *ShippingMethodDataAttributes) HasFreeOverAmountFloat() bool`

HasFreeOverAmountFloat returns a boolean if a field has been set.

### GetFormattedFreeOverAmount

`func (o *ShippingMethodDataAttributes) GetFormattedFreeOverAmount() string`

GetFormattedFreeOverAmount returns the FormattedFreeOverAmount field if non-nil, zero value otherwise.

### GetFormattedFreeOverAmountOk

`func (o *ShippingMethodDataAttributes) GetFormattedFreeOverAmountOk() (*string, bool)`

GetFormattedFreeOverAmountOk returns a tuple with the FormattedFreeOverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFreeOverAmount

`func (o *ShippingMethodDataAttributes) SetFormattedFreeOverAmount(v string)`

SetFormattedFreeOverAmount sets FormattedFreeOverAmount field to given value.

### HasFormattedFreeOverAmount

`func (o *ShippingMethodDataAttributes) HasFormattedFreeOverAmount() bool`

HasFormattedFreeOverAmount returns a boolean if a field has been set.

### GetPriceAmountForShipmentCents

`func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentCents() int32`

GetPriceAmountForShipmentCents returns the PriceAmountForShipmentCents field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentCentsOk

`func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentCentsOk() (*int32, bool)`

GetPriceAmountForShipmentCentsOk returns a tuple with the PriceAmountForShipmentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentCents

`func (o *ShippingMethodDataAttributes) SetPriceAmountForShipmentCents(v int32)`

SetPriceAmountForShipmentCents sets PriceAmountForShipmentCents field to given value.

### HasPriceAmountForShipmentCents

`func (o *ShippingMethodDataAttributes) HasPriceAmountForShipmentCents() bool`

HasPriceAmountForShipmentCents returns a boolean if a field has been set.

### GetPriceAmountForShipmentFloat

`func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentFloat() float32`

GetPriceAmountForShipmentFloat returns the PriceAmountForShipmentFloat field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentFloatOk

`func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentFloatOk() (*float32, bool)`

GetPriceAmountForShipmentFloatOk returns a tuple with the PriceAmountForShipmentFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentFloat

`func (o *ShippingMethodDataAttributes) SetPriceAmountForShipmentFloat(v float32)`

SetPriceAmountForShipmentFloat sets PriceAmountForShipmentFloat field to given value.

### HasPriceAmountForShipmentFloat

`func (o *ShippingMethodDataAttributes) HasPriceAmountForShipmentFloat() bool`

HasPriceAmountForShipmentFloat returns a boolean if a field has been set.

### GetFormattedPriceAmountForShipment

`func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountForShipment() string`

GetFormattedPriceAmountForShipment returns the FormattedPriceAmountForShipment field if non-nil, zero value otherwise.

### GetFormattedPriceAmountForShipmentOk

`func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountForShipmentOk() (*string, bool)`

GetFormattedPriceAmountForShipmentOk returns a tuple with the FormattedPriceAmountForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmountForShipment

`func (o *ShippingMethodDataAttributes) SetFormattedPriceAmountForShipment(v string)`

SetFormattedPriceAmountForShipment sets FormattedPriceAmountForShipment field to given value.

### HasFormattedPriceAmountForShipment

`func (o *ShippingMethodDataAttributes) HasFormattedPriceAmountForShipment() bool`

HasFormattedPriceAmountForShipment returns a boolean if a field has been set.

### GetMinWeight

`func (o *ShippingMethodDataAttributes) GetMinWeight() float32`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *ShippingMethodDataAttributes) GetMinWeightOk() (*float32, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *ShippingMethodDataAttributes) SetMinWeight(v float32)`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *ShippingMethodDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### GetMaxWeight

`func (o *ShippingMethodDataAttributes) GetMaxWeight() float32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *ShippingMethodDataAttributes) GetMaxWeightOk() (*float32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *ShippingMethodDataAttributes) SetMaxWeight(v float32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *ShippingMethodDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *ShippingMethodDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *ShippingMethodDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *ShippingMethodDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *ShippingMethodDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetId

`func (o *ShippingMethodDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethodDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethodDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingMethodDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShippingMethodDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShippingMethodDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShippingMethodDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShippingMethodDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ShippingMethodDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShippingMethodDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShippingMethodDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ShippingMethodDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ShippingMethodDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShippingMethodDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShippingMethodDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShippingMethodDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShippingMethodDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShippingMethodDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShippingMethodDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShippingMethodDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingMethodDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingMethodDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingMethodDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingMethodDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


