# GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The price tier&#39;s name. | [optional] 
**Type** | Pointer to **interface{}** | The price tier&#39;s type. | [optional] 
**UpTo** | Pointer to **interface{}** | The tier upper limit, expressed as the line item quantity. When &#39;null&#39; it means infinity (useful to have an always matching tier). | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this price tier, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this price tier, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this price tier, formatted. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes

`func NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes() *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes`

NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes instantiates a new GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults

`func NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults() *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes`

NewGETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributesWithDefaults instantiates a new GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpTo

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpTo() interface{}`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpToOk() (*interface{}, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetUpTo(v interface{})`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### SetUpToNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetUpToNil(b bool)`

 SetUpToNil sets the value for UpTo to be an explicit nil

### UnsetUpTo
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetUpTo()`

UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
### GetPriceAmountCents

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


