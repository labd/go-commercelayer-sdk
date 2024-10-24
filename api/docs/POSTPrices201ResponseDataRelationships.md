# POSTPrices201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | Pointer to [**POSTMarkets201ResponseDataRelationshipsPriceList**](POSTMarkets201ResponseDataRelationshipsPriceList.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**PriceTiers** | Pointer to [**POSTPrices201ResponseDataRelationshipsPriceTiers**](POSTPrices201ResponseDataRelationshipsPriceTiers.md) |  | [optional] 
**PriceVolumeTiers** | Pointer to [**POSTPrices201ResponseDataRelationshipsPriceVolumeTiers**](POSTPrices201ResponseDataRelationshipsPriceVolumeTiers.md) |  | [optional] 
**PriceFrequencyTiers** | Pointer to [**POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers**](POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 
**JwtCustomer** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtCustomer**](POSTPrices201ResponseDataRelationshipsJwtCustomer.md) |  | [optional] 
**JwtMarkets** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtMarkets**](POSTPrices201ResponseDataRelationshipsJwtMarkets.md) |  | [optional] 
**JwtStockLocations** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtStockLocations**](POSTPrices201ResponseDataRelationshipsJwtStockLocations.md) |  | [optional] 

## Methods

### NewPOSTPrices201ResponseDataRelationships

`func NewPOSTPrices201ResponseDataRelationships() *POSTPrices201ResponseDataRelationships`

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

`func (o *POSTPrices201ResponseDataRelationships) GetPriceList() POSTMarkets201ResponseDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceListOk() (*POSTMarkets201ResponseDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTPrices201ResponseDataRelationships) SetPriceList(v POSTMarkets201ResponseDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *POSTPrices201ResponseDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetSku

`func (o *POSTPrices201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTPrices201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTPrices201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTPrices201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) GetPriceTiers() POSTPrices201ResponseDataRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) SetPriceTiers(v POSTPrices201ResponseDataRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *POSTPrices201ResponseDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetPriceVolumeTiers

`func (o *POSTPrices201ResponseDataRelationships) GetPriceVolumeTiers() POSTPrices201ResponseDataRelationshipsPriceVolumeTiers`

GetPriceVolumeTiers returns the PriceVolumeTiers field if non-nil, zero value otherwise.

### GetPriceVolumeTiersOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceVolumeTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceVolumeTiers, bool)`

GetPriceVolumeTiersOk returns a tuple with the PriceVolumeTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVolumeTiers

`func (o *POSTPrices201ResponseDataRelationships) SetPriceVolumeTiers(v POSTPrices201ResponseDataRelationshipsPriceVolumeTiers)`

SetPriceVolumeTiers sets PriceVolumeTiers field to given value.

### HasPriceVolumeTiers

`func (o *POSTPrices201ResponseDataRelationships) HasPriceVolumeTiers() bool`

HasPriceVolumeTiers returns a boolean if a field has been set.

### GetPriceFrequencyTiers

`func (o *POSTPrices201ResponseDataRelationships) GetPriceFrequencyTiers() POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers`

GetPriceFrequencyTiers returns the PriceFrequencyTiers field if non-nil, zero value otherwise.

### GetPriceFrequencyTiersOk

`func (o *POSTPrices201ResponseDataRelationships) GetPriceFrequencyTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers, bool)`

GetPriceFrequencyTiersOk returns a tuple with the PriceFrequencyTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFrequencyTiers

`func (o *POSTPrices201ResponseDataRelationships) SetPriceFrequencyTiers(v POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers)`

SetPriceFrequencyTiers sets PriceFrequencyTiers field to given value.

### HasPriceFrequencyTiers

`func (o *POSTPrices201ResponseDataRelationships) HasPriceFrequencyTiers() bool`

HasPriceFrequencyTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTPrices201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTPrices201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTPrices201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTPrices201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTPrices201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTPrices201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTPrices201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTPrices201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetJwtCustomer

`func (o *POSTPrices201ResponseDataRelationships) GetJwtCustomer() POSTPrices201ResponseDataRelationshipsJwtCustomer`

GetJwtCustomer returns the JwtCustomer field if non-nil, zero value otherwise.

### GetJwtCustomerOk

`func (o *POSTPrices201ResponseDataRelationships) GetJwtCustomerOk() (*POSTPrices201ResponseDataRelationshipsJwtCustomer, bool)`

GetJwtCustomerOk returns a tuple with the JwtCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtCustomer

`func (o *POSTPrices201ResponseDataRelationships) SetJwtCustomer(v POSTPrices201ResponseDataRelationshipsJwtCustomer)`

SetJwtCustomer sets JwtCustomer field to given value.

### HasJwtCustomer

`func (o *POSTPrices201ResponseDataRelationships) HasJwtCustomer() bool`

HasJwtCustomer returns a boolean if a field has been set.

### GetJwtMarkets

`func (o *POSTPrices201ResponseDataRelationships) GetJwtMarkets() POSTPrices201ResponseDataRelationshipsJwtMarkets`

GetJwtMarkets returns the JwtMarkets field if non-nil, zero value otherwise.

### GetJwtMarketsOk

`func (o *POSTPrices201ResponseDataRelationships) GetJwtMarketsOk() (*POSTPrices201ResponseDataRelationshipsJwtMarkets, bool)`

GetJwtMarketsOk returns a tuple with the JwtMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtMarkets

`func (o *POSTPrices201ResponseDataRelationships) SetJwtMarkets(v POSTPrices201ResponseDataRelationshipsJwtMarkets)`

SetJwtMarkets sets JwtMarkets field to given value.

### HasJwtMarkets

`func (o *POSTPrices201ResponseDataRelationships) HasJwtMarkets() bool`

HasJwtMarkets returns a boolean if a field has been set.

### GetJwtStockLocations

`func (o *POSTPrices201ResponseDataRelationships) GetJwtStockLocations() POSTPrices201ResponseDataRelationshipsJwtStockLocations`

GetJwtStockLocations returns the JwtStockLocations field if non-nil, zero value otherwise.

### GetJwtStockLocationsOk

`func (o *POSTPrices201ResponseDataRelationships) GetJwtStockLocationsOk() (*POSTPrices201ResponseDataRelationshipsJwtStockLocations, bool)`

GetJwtStockLocationsOk returns a tuple with the JwtStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtStockLocations

`func (o *POSTPrices201ResponseDataRelationships) SetJwtStockLocations(v POSTPrices201ResponseDataRelationshipsJwtStockLocations)`

SetJwtStockLocations sets JwtStockLocations field to given value.

### HasJwtStockLocations

`func (o *POSTPrices201ResponseDataRelationships) HasJwtStockLocations() bool`

HasJwtStockLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


