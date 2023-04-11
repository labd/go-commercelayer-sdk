# GETPriceFrequencyTiers200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The price tier&#39;s name | [optional] 
**UpTo** | Pointer to **interface{}** | The tier upper limit, expressed as the line item frequency in days (or frequency label, ie &#39;monthly&#39;). When &#39;null&#39; it means infinity (useful to have an always matching tier). | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this price tier, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this price tier, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this price tier, formatted. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPriceFrequencyTiers200ResponseDataInnerAttributes

`func NewGETPriceFrequencyTiers200ResponseDataInnerAttributes() *GETPriceFrequencyTiers200ResponseDataInnerAttributes`

NewGETPriceFrequencyTiers200ResponseDataInnerAttributes instantiates a new GETPriceFrequencyTiers200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPriceFrequencyTiers200ResponseDataInnerAttributesWithDefaults

`func NewGETPriceFrequencyTiers200ResponseDataInnerAttributesWithDefaults() *GETPriceFrequencyTiers200ResponseDataInnerAttributes`

NewGETPriceFrequencyTiers200ResponseDataInnerAttributesWithDefaults instantiates a new GETPriceFrequencyTiers200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUpTo

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetUpTo() interface{}`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetUpToOk() (*interface{}, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetUpTo(v interface{})`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### SetUpToNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetUpToNil(b bool)`

 SetUpToNil sets the value for UpTo to be an explicit nil

### UnsetUpTo
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetUpTo()`

UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
### GetPriceAmountCents

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPriceFrequencyTiers200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


