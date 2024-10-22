# ParcelLineItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parcel** | Pointer to [**PackageDataRelationshipsParcels**](PackageDataRelationshipsParcels.md) |  | [optional] 
**StockLineItem** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewParcelLineItemDataRelationships

`func NewParcelLineItemDataRelationships() *ParcelLineItemDataRelationships`

NewParcelLineItemDataRelationships instantiates a new ParcelLineItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelLineItemDataRelationshipsWithDefaults

`func NewParcelLineItemDataRelationshipsWithDefaults() *ParcelLineItemDataRelationships`

NewParcelLineItemDataRelationshipsWithDefaults instantiates a new ParcelLineItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParcel

`func (o *ParcelLineItemDataRelationships) GetParcel() PackageDataRelationshipsParcels`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ParcelLineItemDataRelationships) GetParcelOk() (*PackageDataRelationshipsParcels, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ParcelLineItemDataRelationships) SetParcel(v PackageDataRelationshipsParcels)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ParcelLineItemDataRelationships) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetStockLineItem

`func (o *ParcelLineItemDataRelationships) GetStockLineItem() LineItemDataRelationshipsStockLineItems`

GetStockLineItem returns the StockLineItem field if non-nil, zero value otherwise.

### GetStockLineItemOk

`func (o *ParcelLineItemDataRelationships) GetStockLineItemOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemOk returns a tuple with the StockLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItem

`func (o *ParcelLineItemDataRelationships) SetStockLineItem(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItem sets StockLineItem field to given value.

### HasStockLineItem

`func (o *ParcelLineItemDataRelationships) HasStockLineItem() bool`

HasStockLineItem returns a boolean if a field has been set.

### GetVersions

`func (o *ParcelLineItemDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ParcelLineItemDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ParcelLineItemDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ParcelLineItemDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


