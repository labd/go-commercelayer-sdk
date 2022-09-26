# PackageResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockLocation** | Pointer to [**DeliveryLeadTimeResponseDataRelationshipsStockLocation**](DeliveryLeadTimeResponseDataRelationshipsStockLocation.md) |  | [optional] 
**Parcels** | Pointer to [**PackageResponseDataRelationshipsParcels**](PackageResponseDataRelationshipsParcels.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewPackageResponseDataRelationships

`func NewPackageResponseDataRelationships() *PackageResponseDataRelationships`

NewPackageResponseDataRelationships instantiates a new PackageResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageResponseDataRelationshipsWithDefaults

`func NewPackageResponseDataRelationshipsWithDefaults() *PackageResponseDataRelationships`

NewPackageResponseDataRelationshipsWithDefaults instantiates a new PackageResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockLocation

`func (o *PackageResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *PackageResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *PackageResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *PackageResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetParcels

`func (o *PackageResponseDataRelationships) GetParcels() PackageResponseDataRelationshipsParcels`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *PackageResponseDataRelationships) GetParcelsOk() (*PackageResponseDataRelationshipsParcels, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *PackageResponseDataRelationships) SetParcels(v PackageResponseDataRelationshipsParcels)`

SetParcels sets Parcels field to given value.

### HasParcels

`func (o *PackageResponseDataRelationships) HasParcels() bool`

HasParcels returns a boolean if a field has been set.

### GetAttachments

`func (o *PackageResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PackageResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PackageResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PackageResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


