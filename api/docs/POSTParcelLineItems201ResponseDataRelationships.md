# POSTParcelLineItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parcel** | [**GETPackages200ResponseDataInnerRelationshipsParcels**](GETPackages200ResponseDataInnerRelationshipsParcels.md) |  | 
**StockLineItem** | [**GETLineItems200ResponseDataInnerRelationshipsStockLineItems**](GETLineItems200ResponseDataInnerRelationshipsStockLineItems.md) |  | 
**ShipmentLineItem** | Pointer to [**GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems**](GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems.md) |  | [optional] 

## Methods

### NewPOSTParcelLineItems201ResponseDataRelationships

`func NewPOSTParcelLineItems201ResponseDataRelationships(parcel GETPackages200ResponseDataInnerRelationshipsParcels, stockLineItem GETLineItems200ResponseDataInnerRelationshipsStockLineItems, ) *POSTParcelLineItems201ResponseDataRelationships`

NewPOSTParcelLineItems201ResponseDataRelationships instantiates a new POSTParcelLineItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTParcelLineItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTParcelLineItems201ResponseDataRelationshipsWithDefaults() *POSTParcelLineItems201ResponseDataRelationships`

NewPOSTParcelLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTParcelLineItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParcel

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetParcel() GETPackages200ResponseDataInnerRelationshipsParcels`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetParcelOk() (*GETPackages200ResponseDataInnerRelationshipsParcels, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *POSTParcelLineItems201ResponseDataRelationships) SetParcel(v GETPackages200ResponseDataInnerRelationshipsParcels)`

SetParcel sets Parcel field to given value.


### GetStockLineItem

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetStockLineItem() GETLineItems200ResponseDataInnerRelationshipsStockLineItems`

GetStockLineItem returns the StockLineItem field if non-nil, zero value otherwise.

### GetStockLineItemOk

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetStockLineItemOk() (*GETLineItems200ResponseDataInnerRelationshipsStockLineItems, bool)`

GetStockLineItemOk returns a tuple with the StockLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItem

`func (o *POSTParcelLineItems201ResponseDataRelationships) SetStockLineItem(v GETLineItems200ResponseDataInnerRelationshipsStockLineItems)`

SetStockLineItem sets StockLineItem field to given value.


### GetShipmentLineItem

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetShipmentLineItem() GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems`

GetShipmentLineItem returns the ShipmentLineItem field if non-nil, zero value otherwise.

### GetShipmentLineItemOk

`func (o *POSTParcelLineItems201ResponseDataRelationships) GetShipmentLineItemOk() (*GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems, bool)`

GetShipmentLineItemOk returns a tuple with the ShipmentLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLineItem

`func (o *POSTParcelLineItems201ResponseDataRelationships) SetShipmentLineItem(v GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems)`

SetShipmentLineItem sets ShipmentLineItem field to given value.

### HasShipmentLineItem

`func (o *POSTParcelLineItems201ResponseDataRelationships) HasShipmentLineItem() bool`

HasShipmentLineItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


