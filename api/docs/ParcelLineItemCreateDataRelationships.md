# ParcelLineItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parcel** | [**PackageDataRelationshipsParcels**](PackageDataRelationshipsParcels.md) |  | 
**StockLineItem** | [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | 
**ShipmentLineItem** | Pointer to [**LineItemDataRelationshipsShipmentLineItems**](LineItemDataRelationshipsShipmentLineItems.md) |  | [optional] 

## Methods

### NewParcelLineItemCreateDataRelationships

`func NewParcelLineItemCreateDataRelationships(parcel PackageDataRelationshipsParcels, stockLineItem LineItemDataRelationshipsStockLineItems, ) *ParcelLineItemCreateDataRelationships`

NewParcelLineItemCreateDataRelationships instantiates a new ParcelLineItemCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelLineItemCreateDataRelationshipsWithDefaults

`func NewParcelLineItemCreateDataRelationshipsWithDefaults() *ParcelLineItemCreateDataRelationships`

NewParcelLineItemCreateDataRelationshipsWithDefaults instantiates a new ParcelLineItemCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParcel

`func (o *ParcelLineItemCreateDataRelationships) GetParcel() PackageDataRelationshipsParcels`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ParcelLineItemCreateDataRelationships) GetParcelOk() (*PackageDataRelationshipsParcels, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ParcelLineItemCreateDataRelationships) SetParcel(v PackageDataRelationshipsParcels)`

SetParcel sets Parcel field to given value.


### GetStockLineItem

`func (o *ParcelLineItemCreateDataRelationships) GetStockLineItem() LineItemDataRelationshipsStockLineItems`

GetStockLineItem returns the StockLineItem field if non-nil, zero value otherwise.

### GetStockLineItemOk

`func (o *ParcelLineItemCreateDataRelationships) GetStockLineItemOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemOk returns a tuple with the StockLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItem

`func (o *ParcelLineItemCreateDataRelationships) SetStockLineItem(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItem sets StockLineItem field to given value.


### GetShipmentLineItem

`func (o *ParcelLineItemCreateDataRelationships) GetShipmentLineItem() LineItemDataRelationshipsShipmentLineItems`

GetShipmentLineItem returns the ShipmentLineItem field if non-nil, zero value otherwise.

### GetShipmentLineItemOk

`func (o *ParcelLineItemCreateDataRelationships) GetShipmentLineItemOk() (*LineItemDataRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemOk returns a tuple with the ShipmentLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItem

`func (o *ParcelLineItemCreateDataRelationships) SetShipmentLineItem(v LineItemDataRelationshipsShipmentLineItems)`

SetShipmentLineItem sets ShipmentLineItem field to given value.

### HasShipmentLineItem

`func (o *ParcelLineItemCreateDataRelationships) HasShipmentLineItem() bool`

HasShipmentLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


