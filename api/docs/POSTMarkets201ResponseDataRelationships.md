# POSTMarkets201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | [**GETMarkets200ResponseDataInnerRelationshipsMerchant**](GETMarkets200ResponseDataInnerRelationshipsMerchant.md) |  | 
**PriceList** | [**GETMarkets200ResponseDataInnerRelationshipsPriceList**](GETMarkets200ResponseDataInnerRelationshipsPriceList.md) |  | 
**InventoryModel** | [**GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel**](GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel.md) |  | 
**TaxCalculator** | Pointer to [**GETMarkets200ResponseDataInnerRelationshipsTaxCalculator**](GETMarkets200ResponseDataInnerRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**GETCustomers200ResponseDataInnerRelationshipsCustomerGroup**](GETCustomers200ResponseDataInnerRelationshipsCustomerGroup.md) |  | [optional] 

## Methods

### NewPOSTMarkets201ResponseDataRelationships

`func NewPOSTMarkets201ResponseDataRelationships(merchant GETMarkets200ResponseDataInnerRelationshipsMerchant, priceList GETMarkets200ResponseDataInnerRelationshipsPriceList, inventoryModel GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel, ) *POSTMarkets201ResponseDataRelationships`

NewPOSTMarkets201ResponseDataRelationships instantiates a new POSTMarkets201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTMarkets201ResponseDataRelationshipsWithDefaults

`func NewPOSTMarkets201ResponseDataRelationshipsWithDefaults() *POSTMarkets201ResponseDataRelationships`

NewPOSTMarkets201ResponseDataRelationshipsWithDefaults instantiates a new POSTMarkets201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *POSTMarkets201ResponseDataRelationships) GetMerchant() GETMarkets200ResponseDataInnerRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *POSTMarkets201ResponseDataRelationships) GetMerchantOk() (*GETMarkets200ResponseDataInnerRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *POSTMarkets201ResponseDataRelationships) SetMerchant(v GETMarkets200ResponseDataInnerRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.


### GetPriceList

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceList() GETMarkets200ResponseDataInnerRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceListOk() (*GETMarkets200ResponseDataInnerRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTMarkets201ResponseDataRelationships) SetPriceList(v GETMarkets200ResponseDataInnerRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetInventoryModel

`func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModel() GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModelOk() (*GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *POSTMarkets201ResponseDataRelationships) SetInventoryModel(v GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.


### GetTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculator() GETMarkets200ResponseDataInnerRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculatorOk() (*GETMarkets200ResponseDataInnerRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) SetTaxCalculator(v GETMarkets200ResponseDataInnerRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroup() GETCustomers200ResponseDataInnerRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroupOk() (*GETCustomers200ResponseDataInnerRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) SetCustomerGroup(v GETCustomers200ResponseDataInnerRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


