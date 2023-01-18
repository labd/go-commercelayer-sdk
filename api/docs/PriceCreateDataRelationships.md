# PriceCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | [**MarketCreateDataRelationshipsPriceList**](MarketCreateDataRelationshipsPriceList.md) |  | 
**Sku** | [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | 
**PriceTiers** | Pointer to [**PriceCreateDataRelationshipsPriceTiers**](PriceCreateDataRelationshipsPriceTiers.md) |  | [optional] 

## Methods

### NewPriceCreateDataRelationships

`func NewPriceCreateDataRelationships(priceList MarketCreateDataRelationshipsPriceList, sku InStockSubscriptionCreateDataRelationshipsSku, ) *PriceCreateDataRelationships`

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

`func (o *PriceCreateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *PriceCreateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *PriceCreateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetSku

`func (o *PriceCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PriceCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PriceCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.


### GetPriceTiers

`func (o *PriceCreateDataRelationships) GetPriceTiers() PriceCreateDataRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *PriceCreateDataRelationships) GetPriceTiersOk() (*PriceCreateDataRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *PriceCreateDataRelationships) SetPriceTiers(v PriceCreateDataRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *PriceCreateDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


