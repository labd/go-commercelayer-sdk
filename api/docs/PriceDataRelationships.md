# PriceDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | Pointer to [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**PriceTiers** | Pointer to [**PriceDataRelationshipsPriceTiers**](PriceDataRelationshipsPriceTiers.md) |  | [optional] 
**PriceVolumeTiers** | Pointer to [**PriceDataRelationshipsPriceVolumeTiers**](PriceDataRelationshipsPriceVolumeTiers.md) |  | [optional] 
**PriceFrequencyTiers** | Pointer to [**PriceDataRelationshipsPriceFrequencyTiers**](PriceDataRelationshipsPriceFrequencyTiers.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 

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

`func (o *PriceDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PriceDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PriceDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PriceDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


