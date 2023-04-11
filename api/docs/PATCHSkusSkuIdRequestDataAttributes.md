# PATCHSkusSkuIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The SKU code, that uniquely identifies the SKU within the organization. | [optional] 
**Name** | Pointer to **interface{}** | The internal name of the SKU. | [optional] 
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

### NewPATCHSkusSkuIdRequestDataAttributes

`func NewPATCHSkusSkuIdRequestDataAttributes() *PATCHSkusSkuIdRequestDataAttributes`

NewPATCHSkusSkuIdRequestDataAttributes instantiates a new PATCHSkusSkuIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSkusSkuIdRequestDataAttributesWithDefaults

`func NewPATCHSkusSkuIdRequestDataAttributesWithDefaults() *PATCHSkusSkuIdRequestDataAttributes`

NewPATCHSkusSkuIdRequestDataAttributesWithDefaults instantiates a new PATCHSkusSkuIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPiecesPerPack

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetPiecesPerPack() interface{}`

GetPiecesPerPack returns the PiecesPerPack field if non-nil, zero value otherwise.

### GetPiecesPerPackOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetPiecesPerPackOk() (*interface{}, bool)`

GetPiecesPerPackOk returns a tuple with the PiecesPerPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiecesPerPack

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetPiecesPerPack(v interface{})`

SetPiecesPerPack sets PiecesPerPack field to given value.

### HasPiecesPerPack

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasPiecesPerPack() bool`

HasPiecesPerPack returns a boolean if a field has been set.

### SetPiecesPerPackNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetPiecesPerPackNil(b bool)`

 SetPiecesPerPackNil sets the value for PiecesPerPack to be an explicit nil

### UnsetPiecesPerPack
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetPiecesPerPack()`

UnsetPiecesPerPack ensures that no value is present for PiecesPerPack, not even an explicit nil
### GetWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetWeight() interface{}`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetWeightOk() (*interface{}, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetWeight(v interface{})`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetUnitOfWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetHsTariffNumber

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetHsTariffNumber() interface{}`

GetHsTariffNumber returns the HsTariffNumber field if non-nil, zero value otherwise.

### GetHsTariffNumberOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetHsTariffNumberOk() (*interface{}, bool)`

GetHsTariffNumberOk returns a tuple with the HsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsTariffNumber

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetHsTariffNumber(v interface{})`

SetHsTariffNumber sets HsTariffNumber field to given value.

### HasHsTariffNumber

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasHsTariffNumber() bool`

HasHsTariffNumber returns a boolean if a field has been set.

### SetHsTariffNumberNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetHsTariffNumberNil(b bool)`

 SetHsTariffNumberNil sets the value for HsTariffNumber to be an explicit nil

### UnsetHsTariffNumber
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetHsTariffNumber()`

UnsetHsTariffNumber ensures that no value is present for HsTariffNumber, not even an explicit nil
### GetDoNotShip

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDoNotShip() interface{}`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDoNotShipOk() (*interface{}, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDoNotShip(v interface{})`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### SetDoNotShipNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDoNotShipNil(b bool)`

 SetDoNotShipNil sets the value for DoNotShip to be an explicit nil

### UnsetDoNotShip
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetDoNotShip()`

UnsetDoNotShip ensures that no value is present for DoNotShip, not even an explicit nil
### GetDoNotTrack

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDoNotTrack() interface{}`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetDoNotTrackOk() (*interface{}, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDoNotTrack(v interface{})`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### SetDoNotTrackNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetDoNotTrackNil(b bool)`

 SetDoNotTrackNil sets the value for DoNotTrack to be an explicit nil

### UnsetDoNotTrack
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetDoNotTrack()`

UnsetDoNotTrack ensures that no value is present for DoNotTrack, not even an explicit nil
### GetReference

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHSkusSkuIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHSkusSkuIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHSkusSkuIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHSkusSkuIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


