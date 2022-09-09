# MarketCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | [**MarketDataRelationshipsMerchant**](MarketDataRelationshipsMerchant.md) |  | 
**PriceList** | [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | 
**InventoryModel** | [**InventoryReturnLocationDataRelationshipsInventoryModel**](InventoryReturnLocationDataRelationshipsInventoryModel.md) |  | 
**TaxCalculator** | Pointer to [**MarketDataRelationshipsTaxCalculator**](MarketDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**CustomerDataRelationshipsCustomerGroup**](CustomerDataRelationshipsCustomerGroup.md) |  | [optional] 

## Methods

### NewMarketCreateDataRelationships

`func NewMarketCreateDataRelationships(merchant MarketDataRelationshipsMerchant, priceList MarketDataRelationshipsPriceList, inventoryModel InventoryReturnLocationDataRelationshipsInventoryModel, ) *MarketCreateDataRelationships`

NewMarketCreateDataRelationships instantiates a new MarketCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketCreateDataRelationshipsWithDefaults

`func NewMarketCreateDataRelationshipsWithDefaults() *MarketCreateDataRelationships`

NewMarketCreateDataRelationshipsWithDefaults instantiates a new MarketCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *MarketCreateDataRelationships) GetMerchant() MarketDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MarketCreateDataRelationships) GetMerchantOk() (*MarketDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MarketCreateDataRelationships) SetMerchant(v MarketDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.


### GetPriceList

`func (o *MarketCreateDataRelationships) GetPriceList() MarketDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *MarketCreateDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *MarketCreateDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetInventoryModel

`func (o *MarketCreateDataRelationships) GetInventoryModel() InventoryReturnLocationDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *MarketCreateDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *MarketCreateDataRelationships) SetInventoryModel(v InventoryReturnLocationDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.


### GetTaxCalculator

`func (o *MarketCreateDataRelationships) GetTaxCalculator() MarketDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *MarketCreateDataRelationships) GetTaxCalculatorOk() (*MarketDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *MarketCreateDataRelationships) SetTaxCalculator(v MarketDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *MarketCreateDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *MarketCreateDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *MarketCreateDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *MarketCreateDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *MarketCreateDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


