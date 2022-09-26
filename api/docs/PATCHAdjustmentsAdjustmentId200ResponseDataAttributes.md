# PATCHAdjustmentsAdjustmentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The adjustment name | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**AmountCents** | Pointer to **int32** | The adjustment amount, in cents. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributes

`func NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributes() *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes`

NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributes instantiates a new PATCHAdjustmentsAdjustmentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributesWithDefaults

`func NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributesWithDefaults() *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes`

NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributesWithDefaults instantiates a new PATCHAdjustmentsAdjustmentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetReference

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


