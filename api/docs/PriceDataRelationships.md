# PriceDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | Pointer to [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**PriceTiers** | Pointer to [**PriceDataRelationshipsPriceTiers**](PriceDataRelationshipsPriceTiers.md) |  | [optional] 
**PriceVolumeTiers** | Pointer to [**PriceDataRelationshipsPriceVolumeTiers**](PriceDataRelationshipsPriceVolumeTiers.md) |  | [optional] 
**PriceFrequencyTiers** | Pointer to [**PriceDataRelationshipsPriceFrequencyTiers**](PriceDataRelationshipsPriceFrequencyTiers.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**JwtCustomer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**JwtMarkets** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**JwtStockLocations** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewPriceDataRelationships

`func NewPriceDataRelationships() *PriceDataRelationships`

NewPriceDataRelationships instantiates a new PriceDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDataRelationshipsWithDefaults

`func NewPriceDataRelationshipsWithDefaults() *PriceDataRelationships`

NewPriceDataRelationshipsWithDefaults instantiates a new PriceDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceList

`func (o *PriceDataRelationships) GetPriceList() MarketDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *PriceDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *PriceDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *PriceDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetSku

`func (o *PriceDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PriceDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PriceDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *PriceDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPriceTiers

`func (o *PriceDataRelationships) GetPriceTiers() PriceDataRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *PriceDataRelationships) GetPriceTiersOk() (*PriceDataRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *PriceDataRelationships) SetPriceTiers(v PriceDataRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *PriceDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetPriceVolumeTiers

`func (o *PriceDataRelationships) GetPriceVolumeTiers() PriceDataRelationshipsPriceVolumeTiers`

GetPriceVolumeTiers returns the PriceVolumeTiers field if non-nil, zero value otherwise.

### GetPriceVolumeTiersOk

`func (o *PriceDataRelationships) GetPriceVolumeTiersOk() (*PriceDataRelationshipsPriceVolumeTiers, bool)`

GetPriceVolumeTiersOk returns a tuple with the PriceVolumeTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVolumeTiers

`func (o *PriceDataRelationships) SetPriceVolumeTiers(v PriceDataRelationshipsPriceVolumeTiers)`

SetPriceVolumeTiers sets PriceVolumeTiers field to given value.

### HasPriceVolumeTiers

`func (o *PriceDataRelationships) HasPriceVolumeTiers() bool`

HasPriceVolumeTiers returns a boolean if a field has been set.

### GetPriceFrequencyTiers

`func (o *PriceDataRelationships) GetPriceFrequencyTiers() PriceDataRelationshipsPriceFrequencyTiers`

GetPriceFrequencyTiers returns the PriceFrequencyTiers field if non-nil, zero value otherwise.

### GetPriceFrequencyTiersOk

`func (o *PriceDataRelationships) GetPriceFrequencyTiersOk() (*PriceDataRelationshipsPriceFrequencyTiers, bool)`

GetPriceFrequencyTiersOk returns a tuple with the PriceFrequencyTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFrequencyTiers

`func (o *PriceDataRelationships) SetPriceFrequencyTiers(v PriceDataRelationshipsPriceFrequencyTiers)`

SetPriceFrequencyTiers sets PriceFrequencyTiers field to given value.

### HasPriceFrequencyTiers

`func (o *PriceDataRelationships) HasPriceFrequencyTiers() bool`

HasPriceFrequencyTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *PriceDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PriceDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PriceDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PriceDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *PriceDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PriceDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PriceDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PriceDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetJwtCustomer

`func (o *PriceDataRelationships) GetJwtCustomer() CouponRecipientDataRelationshipsCustomer`

GetJwtCustomer returns the JwtCustomer field if non-nil, zero value otherwise.

### GetJwtCustomerOk

`func (o *PriceDataRelationships) GetJwtCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetJwtCustomerOk returns a tuple with the JwtCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtCustomer

`func (o *PriceDataRelationships) SetJwtCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetJwtCustomer sets JwtCustomer field to given value.

### HasJwtCustomer

`func (o *PriceDataRelationships) HasJwtCustomer() bool`

HasJwtCustomer returns a boolean if a field has been set.

### GetJwtMarkets

`func (o *PriceDataRelationships) GetJwtMarkets() AvalaraAccountDataRelationshipsMarkets`

GetJwtMarkets returns the JwtMarkets field if non-nil, zero value otherwise.

### GetJwtMarketsOk

`func (o *PriceDataRelationships) GetJwtMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetJwtMarketsOk returns a tuple with the JwtMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtMarkets

`func (o *PriceDataRelationships) SetJwtMarkets(v AvalaraAccountDataRelationshipsMarkets)`

SetJwtMarkets sets JwtMarkets field to given value.

### HasJwtMarkets

`func (o *PriceDataRelationships) HasJwtMarkets() bool`

HasJwtMarkets returns a boolean if a field has been set.

### GetJwtStockLocations

`func (o *PriceDataRelationships) GetJwtStockLocations() DeliveryLeadTimeDataRelationshipsStockLocation`

GetJwtStockLocations returns the JwtStockLocations field if non-nil, zero value otherwise.

### GetJwtStockLocationsOk

`func (o *PriceDataRelationships) GetJwtStockLocationsOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetJwtStockLocationsOk returns a tuple with the JwtStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtStockLocations

`func (o *PriceDataRelationships) SetJwtStockLocations(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetJwtStockLocations sets JwtStockLocations field to given value.

### HasJwtStockLocations

`func (o *PriceDataRelationships) HasJwtStockLocations() bool`

HasJwtStockLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


