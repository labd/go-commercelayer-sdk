# PriceCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | 
**Sku** | [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | 
**PriceTiers** | Pointer to [**PriceDataRelationshipsPriceTiers**](PriceDataRelationshipsPriceTiers.md) |  | [optional] 

## Methods

### NewPriceCreateDataRelationships

`func NewPriceCreateDataRelationships(priceList MarketDataRelationshipsPriceList, sku BundleDataRelationshipsSkus, ) *PriceCreateDataRelationships`

NewPriceCreateDataRelationships instantiates a new PriceCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceCreateDataRelationshipsWithDefaults

`func NewPriceCreateDataRelationshipsWithDefaults() *PriceCreateDataRelationships`

NewPriceCreateDataRelationshipsWithDefaults instantiates a new PriceCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceList

`func (o *PriceCreateDataRelationships) GetPriceList() MarketDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *PriceCreateDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *PriceCreateDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetSku

`func (o *PriceCreateDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PriceCreateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PriceCreateDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.


### GetPriceTiers

`func (o *PriceCreateDataRelationships) GetPriceTiers() PriceDataRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *PriceCreateDataRelationships) GetPriceTiersOk() (*PriceDataRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *PriceCreateDataRelationships) SetPriceTiers(v PriceDataRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *PriceCreateDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


