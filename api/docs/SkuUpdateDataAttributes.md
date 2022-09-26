# SkuUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The SKU code, that uniquely identifies the SKU within the organization. | [optional] 
**Name** | Pointer to **string** | The internal name of the SKU. | [optional] 
**Description** | Pointer to **string** | An internal description of the SKU. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the SKU. | [optional] 
**PiecesPerPack** | Pointer to **int32** | The number of pieces that compose the SKU. This is useful to describe sets and bundles. | [optional] 
**Weight** | Pointer to **float32** | The weight of the SKU. If present, it will be used to calculate the shipping rates. | [optional] 
**UnitOfWeight** | Pointer to **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**HsTariffNumber** | Pointer to **string** | The Harmonized System Code used by customs to identify the products shipped across international borders. | [optional] 
**DoNotShip** | Pointer to **bool** | Indicates if the SKU doesn&#39;t generate shipments. | [optional] 
**DoNotTrack** | Pointer to **bool** | Indicates if the SKU doesn&#39;t track the stock inventory. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewSkuUpdateDataAttributes

`func NewSkuUpdateDataAttributes() *SkuUpdateDataAttributes`

NewSkuUpdateDataAttributes instantiates a new SkuUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuUpdateDataAttributesWithDefaults

`func NewSkuUpdateDataAttributesWithDefaults() *SkuUpdateDataAttributes`

NewSkuUpdateDataAttributesWithDefaults instantiates a new SkuUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SkuUpdateDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SkuUpdateDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SkuUpdateDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SkuUpdateDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *SkuUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkuUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkuUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SkuUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SkuUpdateDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkuUpdateDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkuUpdateDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkuUpdateDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *SkuUpdateDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SkuUpdateDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SkuUpdateDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SkuUpdateDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetPiecesPerPack

`func (o *SkuUpdateDataAttributes) GetPiecesPerPack() int32`

GetPiecesPerPack returns the PiecesPerPack field if non-nil, zero value otherwise.

### GetPiecesPerPackOk

`func (o *SkuUpdateDataAttributes) GetPiecesPerPackOk() (*int32, bool)`

GetPiecesPerPackOk returns a tuple with the PiecesPerPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiecesPerPack

`func (o *SkuUpdateDataAttributes) SetPiecesPerPack(v int32)`

SetPiecesPerPack sets PiecesPerPack field to given value.

### HasPiecesPerPack

`func (o *SkuUpdateDataAttributes) HasPiecesPerPack() bool`

HasPiecesPerPack returns a boolean if a field has been set.

### GetWeight

`func (o *SkuUpdateDataAttributes) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SkuUpdateDataAttributes) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SkuUpdateDataAttributes) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SkuUpdateDataAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *SkuUpdateDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *SkuUpdateDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *SkuUpdateDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *SkuUpdateDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetHsTariffNumber

`func (o *SkuUpdateDataAttributes) GetHsTariffNumber() string`

GetHsTariffNumber returns the HsTariffNumber field if non-nil, zero value otherwise.

### GetHsTariffNumberOk

`func (o *SkuUpdateDataAttributes) GetHsTariffNumberOk() (*string, bool)`

GetHsTariffNumberOk returns a tuple with the HsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsTariffNumber

`func (o *SkuUpdateDataAttributes) SetHsTariffNumber(v string)`

SetHsTariffNumber sets HsTariffNumber field to given value.

### HasHsTariffNumber

`func (o *SkuUpdateDataAttributes) HasHsTariffNumber() bool`

HasHsTariffNumber returns a boolean if a field has been set.

### GetDoNotShip

`func (o *SkuUpdateDataAttributes) GetDoNotShip() bool`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *SkuUpdateDataAttributes) GetDoNotShipOk() (*bool, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *SkuUpdateDataAttributes) SetDoNotShip(v bool)`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *SkuUpdateDataAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *SkuUpdateDataAttributes) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *SkuUpdateDataAttributes) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *SkuUpdateDataAttributes) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *SkuUpdateDataAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetReference

`func (o *SkuUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SkuUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SkuUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SkuUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *SkuUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *SkuUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *SkuUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *SkuUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *SkuUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SkuUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SkuUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SkuUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


