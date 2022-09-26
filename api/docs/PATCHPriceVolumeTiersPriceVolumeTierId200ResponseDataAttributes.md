# PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The price tier&#39;s name | [optional] 
**UpTo** | Pointer to **float32** | The tier upper limit. When &#39;null&#39; it means infinity (useful to have an always matching tier). | [optional] 
**PriceAmountCents** | Pointer to **int32** | The price of this price tier, in cents. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes

`func NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes() *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes`

NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes instantiates a new PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults

`func NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults() *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes`

NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults instantiates a new PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpTo

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpTo() float32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpToOk() (*float32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetUpTo(v float32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetReference

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


