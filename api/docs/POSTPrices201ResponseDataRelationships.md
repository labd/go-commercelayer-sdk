# POSTPrices201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | [**GETMarkets200ResponseDataInnerRelationshipsPriceList**](GETMarkets200ResponseDataInnerRelationshipsPriceList.md) |  | 
**Sku** | Pointer to [**GETBundles200ResponseDataInnerRelationshipsSkus**](GETBundles200ResponseDataInnerRelationshipsSkus.md) |  | [optional] 
**PriceTiers** | Pointer to [**GETPrices200ResponseDataInnerRelationshipsPriceTiers**](GETPrices200ResponseDataInnerRelationshipsPriceTiers.md) |  | [optional] 

## Methods

### NewPOSTPrices201ResponseDataRelationships

`func NewPOSTPrices201ResponseDataRelationships(priceList GETMarkets200ResponseDataInnerRelationshipsPriceList, ) *POSTPrices201ResponseDataRelationships`

NewPOSTPrices201ResponseDataRelationships instantiates a new POSTPrices201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPrices201ResponseDataRelationshipsWithDefaults

`func NewPOSTPrices201ResponseDataRelationshipsWithDefaults() *POSTPrices201ResponseDataRelationships`

NewPOSTPrices201ResponseDataRelationshipsWithDefaults instantiates a new POSTPrices201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceList

`func (o *POSTPrices201ResponseDataRelationships) GetPriceList() GETMarkets200ResponseDataInnerRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceListOk() (*GETMarkets200ResponseDataInnerRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTPrices201ResponseDataRelationships) SetPriceList(v GETMarkets200ResponseDataInnerRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetSku

`func (o *POSTPrices201ResponseDataRelationships) GetSku() GETBundles200ResponseDataInnerRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTPrices201ResponseDataRelationships) GetSkuOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTPrices201ResponseDataRelationships) SetSku(v GETBundles200ResponseDataInnerRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTPrices201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) GetPriceTiers() GETPrices200ResponseDataInnerRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceTiersOk() (*GETPrices200ResponseDataInnerRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) SetPriceTiers(v GETPrices200ResponseDataInnerRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


