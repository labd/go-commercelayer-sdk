# GETShippingMethods200ResponseDataInnerAttributes

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

### NewGETShippingMethods200ResponseDataInnerAttributes

`func NewGETShippingMethods200ResponseDataInnerAttributes() *GETShippingMethods200ResponseDataInnerAttributes`

NewGETShippingMethods200ResponseDataInnerAttributes instantiates a new GETShippingMethods200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults

`func NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults() *GETShippingMethods200ResponseDataInnerAttributes`

NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetDisabledAt() string`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetDisabledAtOk() (*string, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetDisabledAt(v string)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountFloat() float32`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*float32, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountFloat(v float32)`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### GetFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmount() string`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*string, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmount(v string)`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### GetFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountCents() int32`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountCentsOk() (*int32, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountCents(v int32)`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### GetFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountFloat() float32`

GetFreeOverAmountFloat returns the FreeOverAmountFloat field if non-nil, zero value otherwise.

### GetFreeOverAmountFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountFloatOk() (*float32, bool)`

GetFreeOverAmountFloatOk returns a tuple with the FreeOverAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountFloat(v float32)`

SetFreeOverAmountFloat sets FreeOverAmountFloat field to given value.

### HasFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFreeOverAmountFloat() bool`

HasFreeOverAmountFloat returns a boolean if a field has been set.

### GetFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedFreeOverAmount() string`

GetFormattedFreeOverAmount returns the FormattedFreeOverAmount field if non-nil, zero value otherwise.

### GetFormattedFreeOverAmountOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedFreeOverAmountOk() (*string, bool)`

GetFormattedFreeOverAmountOk returns a tuple with the FormattedFreeOverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedFreeOverAmount(v string)`

SetFormattedFreeOverAmount sets FormattedFreeOverAmount field to given value.

### HasFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedFreeOverAmount() bool`

HasFormattedFreeOverAmount returns a boolean if a field has been set.

### GetPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentCents() int32`

GetPriceAmountForShipmentCents returns the PriceAmountForShipmentCents field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentCentsOk() (*int32, bool)`

GetPriceAmountForShipmentCentsOk returns a tuple with the PriceAmountForShipmentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentCents(v int32)`

SetPriceAmountForShipmentCents sets PriceAmountForShipmentCents field to given value.

### HasPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountForShipmentCents() bool`

HasPriceAmountForShipmentCents returns a boolean if a field has been set.

### GetPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentFloat() float32`

GetPriceAmountForShipmentFloat returns the PriceAmountForShipmentFloat field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentFloatOk() (*float32, bool)`

GetPriceAmountForShipmentFloatOk returns a tuple with the PriceAmountForShipmentFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentFloat(v float32)`

SetPriceAmountForShipmentFloat sets PriceAmountForShipmentFloat field to given value.

### HasPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountForShipmentFloat() bool`

HasPriceAmountForShipmentFloat returns a boolean if a field has been set.

### GetFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountForShipment() string`

GetFormattedPriceAmountForShipment returns the FormattedPriceAmountForShipment field if non-nil, zero value otherwise.

### GetFormattedPriceAmountForShipmentOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountForShipmentOk() (*string, bool)`

GetFormattedPriceAmountForShipmentOk returns a tuple with the FormattedPriceAmountForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmountForShipment(v string)`

SetFormattedPriceAmountForShipment sets FormattedPriceAmountForShipment field to given value.

### HasFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedPriceAmountForShipment() bool`

HasFormattedPriceAmountForShipment returns a boolean if a field has been set.

### GetMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMinWeight() float32`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMinWeightOk() (*float32, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMinWeight(v float32)`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### GetMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMaxWeight() float32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMaxWeightOk() (*float32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMaxWeight(v float32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetId

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


