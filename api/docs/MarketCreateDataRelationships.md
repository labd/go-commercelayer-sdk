# MarketCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | [**MarketCreateDataRelationshipsMerchant**](MarketCreateDataRelationshipsMerchant.md) |  | 
**PriceList** | [**MarketCreateDataRelationshipsPriceList**](MarketCreateDataRelationshipsPriceList.md) |  | 
**InventoryModel** | [**InventoryReturnLocationCreateDataRelationshipsInventoryModel**](InventoryReturnLocationCreateDataRelationshipsInventoryModel.md) |  | 
**SubscriptionModel** | Pointer to [**MarketCreateDataRelationshipsSubscriptionModel**](MarketCreateDataRelationshipsSubscriptionModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**MarketCreateDataRelationshipsTaxCalculator**](MarketCreateDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**CustomerCreateDataRelationshipsCustomerGroup**](CustomerCreateDataRelationshipsCustomerGroup.md) |  | [optional] 
**Geocoder** | Pointer to [**AddressCreateDataRelationshipsGeocoder**](AddressCreateDataRelationshipsGeocoder.md) |  | [optional] 

## Methods

### NewMarketCreateDataRelationships

`func NewMarketCreateDataRelationships(merchant MarketCreateDataRelationshipsMerchant, priceList MarketCreateDataRelationshipsPriceList, inventoryModel InventoryReturnLocationCreateDataRelationshipsInventoryModel, ) *MarketCreateDataRelationships`

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

`func (o *MarketCreateDataRelationships) GetMerchant() MarketCreateDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MarketCreateDataRelationships) GetMerchantOk() (*MarketCreateDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MarketCreateDataRelationships) SetMerchant(v MarketCreateDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.


### GetPriceList

`func (o *MarketCreateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *MarketCreateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *MarketCreateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetInventoryModel

`func (o *MarketCreateDataRelationships) GetInventoryModel() InventoryReturnLocationCreateDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *MarketCreateDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationCreateDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *MarketCreateDataRelationships) SetInventoryModel(v InventoryReturnLocationCreateDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.


### GetSubscriptionModel

`func (o *MarketCreateDataRelationships) GetSubscriptionModel() MarketCreateDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *MarketCreateDataRelationships) GetSubscriptionModelOk() (*MarketCreateDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *MarketCreateDataRelationships) SetSubscriptionModel(v MarketCreateDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *MarketCreateDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *MarketCreateDataRelationships) GetTaxCalculator() MarketCreateDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *MarketCreateDataRelationships) GetTaxCalculatorOk() (*MarketCreateDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *MarketCreateDataRelationships) SetTaxCalculator(v MarketCreateDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *MarketCreateDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *MarketCreateDataRelationships) GetCustomerGroup() CustomerCreateDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *MarketCreateDataRelationships) GetCustomerGroupOk() (*CustomerCreateDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *MarketCreateDataRelationships) SetCustomerGroup(v CustomerCreateDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *MarketCreateDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetGeocoder

`func (o *MarketCreateDataRelationships) GetGeocoder() AddressCreateDataRelationshipsGeocoder`

GetGeocoder returns the Geocoder field if non-nil, zero value otherwise.

### GetGeocoderOk

`func (o *MarketCreateDataRelationships) GetGeocoderOk() (*AddressCreateDataRelationshipsGeocoder, bool)`

GetGeocoderOk returns a tuple with the Geocoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoder

`func (o *MarketCreateDataRelationships) SetGeocoder(v AddressCreateDataRelationshipsGeocoder)`

SetGeocoder sets Geocoder field to given value.

### HasGeocoder

`func (o *MarketCreateDataRelationships) HasGeocoder() bool`

HasGeocoder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


