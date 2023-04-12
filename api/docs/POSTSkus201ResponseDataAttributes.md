# POSTSkus201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **interface{}** | The SKU code, that uniquely identifies the SKU within the organization. | 
**Name** | **interface{}** | The internal name of the SKU. | 
**Description** | Pointer to **interface{}** | An internal description of the SKU. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the SKU. | [optional] 
**PiecesPerPack** | Pointer to **interface{}** | The number of pieces that compose the SKU. This is useful to describe sets and bundles. | [optional] 
**Weight** | Pointer to **interface{}** | The weight of the SKU. If present, it will be used to calculate the shipping rates. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**HsTariffNumber** | Pointer to **interface{}** | The Harmonized System Code used by customs to identify the products shipped across international borders. | [optional] 
**DoNotShip** | Pointer to **interface{}** | Indicates if the SKU doesn&#39;t generate shipments. | [optional] 
**DoNotTrack** | Pointer to **interface{}** | Indicates if the SKU doesn&#39;t track the stock inventory. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTSkus201ResponseDataAttributes

`func NewPOSTSkus201ResponseDataAttributes(code interface{}, name interface{}, ) *POSTSkus201ResponseDataAttributes`

NewPOSTSkus201ResponseDataAttributes instantiates a new POSTSkus201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkus201ResponseDataAttributesWithDefaults

`func NewPOSTSkus201ResponseDataAttributesWithDefaults() *POSTSkus201ResponseDataAttributes`

NewPOSTSkus201ResponseDataAttributesWithDefaults instantiates a new POSTSkus201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTSkus201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTSkus201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTSkus201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *POSTSkus201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTSkus201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *POSTSkus201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTSkus201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTSkus201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTSkus201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTSkus201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *POSTSkus201ResponseDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTSkus201ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTSkus201ResponseDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTSkus201ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *POSTSkus201ResponseDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *POSTSkus201ResponseDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *POSTSkus201ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTSkus201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTSkus201ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTSkus201ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTSkus201ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTSkus201ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPiecesPerPack

`func (o *POSTSkus201ResponseDataAttributes) GetPiecesPerPack() interface{}`

GetPiecesPerPack returns the PiecesPerPack field if non-nil, zero value otherwise.

### GetPiecesPerPackOk

`func (o *POSTSkus201ResponseDataAttributes) GetPiecesPerPackOk() (*interface{}, bool)`

GetPiecesPerPackOk returns a tuple with the PiecesPerPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiecesPerPack

`func (o *POSTSkus201ResponseDataAttributes) SetPiecesPerPack(v interface{})`

SetPiecesPerPack sets PiecesPerPack field to given value.

### HasPiecesPerPack

`func (o *POSTSkus201ResponseDataAttributes) HasPiecesPerPack() bool`

HasPiecesPerPack returns a boolean if a field has been set.

### SetPiecesPerPackNil

`func (o *POSTSkus201ResponseDataAttributes) SetPiecesPerPackNil(b bool)`

 SetPiecesPerPackNil sets the value for PiecesPerPack to be an explicit nil

### UnsetPiecesPerPack
`func (o *POSTSkus201ResponseDataAttributes) UnsetPiecesPerPack()`

UnsetPiecesPerPack ensures that no value is present for PiecesPerPack, not even an explicit nil
### GetWeight

`func (o *POSTSkus201ResponseDataAttributes) GetWeight() interface{}`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *POSTSkus201ResponseDataAttributes) GetWeightOk() (*interface{}, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *POSTSkus201ResponseDataAttributes) SetWeight(v interface{})`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *POSTSkus201ResponseDataAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *POSTSkus201ResponseDataAttributes) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *POSTSkus201ResponseDataAttributes) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetUnitOfWeight

`func (o *POSTSkus201ResponseDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *POSTSkus201ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *POSTSkus201ResponseDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *POSTSkus201ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *POSTSkus201ResponseDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *POSTSkus201ResponseDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetHsTariffNumber

`func (o *POSTSkus201ResponseDataAttributes) GetHsTariffNumber() interface{}`

GetHsTariffNumber returns the HsTariffNumber field if non-nil, zero value otherwise.

### GetHsTariffNumberOk

`func (o *POSTSkus201ResponseDataAttributes) GetHsTariffNumberOk() (*interface{}, bool)`

GetHsTariffNumberOk returns a tuple with the HsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsTariffNumber

`func (o *POSTSkus201ResponseDataAttributes) SetHsTariffNumber(v interface{})`

SetHsTariffNumber sets HsTariffNumber field to given value.

### HasHsTariffNumber

`func (o *POSTSkus201ResponseDataAttributes) HasHsTariffNumber() bool`

HasHsTariffNumber returns a boolean if a field has been set.

### SetHsTariffNumberNil

`func (o *POSTSkus201ResponseDataAttributes) SetHsTariffNumberNil(b bool)`

 SetHsTariffNumberNil sets the value for HsTariffNumber to be an explicit nil

### UnsetHsTariffNumber
`func (o *POSTSkus201ResponseDataAttributes) UnsetHsTariffNumber()`

UnsetHsTariffNumber ensures that no value is present for HsTariffNumber, not even an explicit nil
### GetDoNotShip

`func (o *POSTSkus201ResponseDataAttributes) GetDoNotShip() interface{}`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *POSTSkus201ResponseDataAttributes) GetDoNotShipOk() (*interface{}, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *POSTSkus201ResponseDataAttributes) SetDoNotShip(v interface{})`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *POSTSkus201ResponseDataAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### SetDoNotShipNil

`func (o *POSTSkus201ResponseDataAttributes) SetDoNotShipNil(b bool)`

 SetDoNotShipNil sets the value for DoNotShip to be an explicit nil

### UnsetDoNotShip
`func (o *POSTSkus201ResponseDataAttributes) UnsetDoNotShip()`

UnsetDoNotShip ensures that no value is present for DoNotShip, not even an explicit nil
### GetDoNotTrack

`func (o *POSTSkus201ResponseDataAttributes) GetDoNotTrack() interface{}`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *POSTSkus201ResponseDataAttributes) GetDoNotTrackOk() (*interface{}, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *POSTSkus201ResponseDataAttributes) SetDoNotTrack(v interface{})`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *POSTSkus201ResponseDataAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### SetDoNotTrackNil

`func (o *POSTSkus201ResponseDataAttributes) SetDoNotTrackNil(b bool)`

 SetDoNotTrackNil sets the value for DoNotTrack to be an explicit nil

### UnsetDoNotTrack
`func (o *POSTSkus201ResponseDataAttributes) UnsetDoNotTrack()`

UnsetDoNotTrack ensures that no value is present for DoNotTrack, not even an explicit nil
### GetReference

`func (o *POSTSkus201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkus201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkus201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkus201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTSkus201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTSkus201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTSkus201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkus201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkus201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkus201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTSkus201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTSkus201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTSkus201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkus201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkus201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkus201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTSkus201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTSkus201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


