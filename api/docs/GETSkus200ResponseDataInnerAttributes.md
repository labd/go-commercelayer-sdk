# GETSkus200ResponseDataInnerAttributes

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
**Inventory** | Pointer to **map[string]interface{}** | Aggregated information about the SKU&#39;s inventory. Returned only when retrieving a single SKU. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETSkus200ResponseDataInnerAttributes

`func NewGETSkus200ResponseDataInnerAttributes() *GETSkus200ResponseDataInnerAttributes`

NewGETSkus200ResponseDataInnerAttributes instantiates a new GETSkus200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkus200ResponseDataInnerAttributesWithDefaults

`func NewGETSkus200ResponseDataInnerAttributesWithDefaults() *GETSkus200ResponseDataInnerAttributes`

NewGETSkus200ResponseDataInnerAttributesWithDefaults instantiates a new GETSkus200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GETSkus200ResponseDataInnerAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETSkus200ResponseDataInnerAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GETSkus200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GETSkus200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETSkus200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETSkus200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GETSkus200ResponseDataInnerAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETSkus200ResponseDataInnerAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETSkus200ResponseDataInnerAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *GETSkus200ResponseDataInnerAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETSkus200ResponseDataInnerAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETSkus200ResponseDataInnerAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetPiecesPerPack

`func (o *GETSkus200ResponseDataInnerAttributes) GetPiecesPerPack() int32`

GetPiecesPerPack returns the PiecesPerPack field if non-nil, zero value otherwise.

### GetPiecesPerPackOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetPiecesPerPackOk() (*int32, bool)`

GetPiecesPerPackOk returns a tuple with the PiecesPerPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiecesPerPack

`func (o *GETSkus200ResponseDataInnerAttributes) SetPiecesPerPack(v int32)`

SetPiecesPerPack sets PiecesPerPack field to given value.

### HasPiecesPerPack

`func (o *GETSkus200ResponseDataInnerAttributes) HasPiecesPerPack() bool`

HasPiecesPerPack returns a boolean if a field has been set.

### GetWeight

`func (o *GETSkus200ResponseDataInnerAttributes) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GETSkus200ResponseDataInnerAttributes) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GETSkus200ResponseDataInnerAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *GETSkus200ResponseDataInnerAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETSkus200ResponseDataInnerAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETSkus200ResponseDataInnerAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetHsTariffNumber

`func (o *GETSkus200ResponseDataInnerAttributes) GetHsTariffNumber() string`

GetHsTariffNumber returns the HsTariffNumber field if non-nil, zero value otherwise.

### GetHsTariffNumberOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetHsTariffNumberOk() (*string, bool)`

GetHsTariffNumberOk returns a tuple with the HsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsTariffNumber

`func (o *GETSkus200ResponseDataInnerAttributes) SetHsTariffNumber(v string)`

SetHsTariffNumber sets HsTariffNumber field to given value.

### HasHsTariffNumber

`func (o *GETSkus200ResponseDataInnerAttributes) HasHsTariffNumber() bool`

HasHsTariffNumber returns a boolean if a field has been set.

### GetDoNotShip

`func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotShip() bool`

GetDoNotShip returns the DoNotShip field if non-nil, zero value otherwise.

### GetDoNotShipOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotShipOk() (*bool, bool)`

GetDoNotShipOk returns a tuple with the DoNotShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShip

`func (o *GETSkus200ResponseDataInnerAttributes) SetDoNotShip(v bool)`

SetDoNotShip sets DoNotShip field to given value.

### HasDoNotShip

`func (o *GETSkus200ResponseDataInnerAttributes) HasDoNotShip() bool`

HasDoNotShip returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *GETSkus200ResponseDataInnerAttributes) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *GETSkus200ResponseDataInnerAttributes) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetInventory

`func (o *GETSkus200ResponseDataInnerAttributes) GetInventory() map[string]interface{}`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetInventoryOk() (*map[string]interface{}, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GETSkus200ResponseDataInnerAttributes) SetInventory(v map[string]interface{})`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *GETSkus200ResponseDataInnerAttributes) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETSkus200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETSkus200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETSkus200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETSkus200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETSkus200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETSkus200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETSkus200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETSkus200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETSkus200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETSkus200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


