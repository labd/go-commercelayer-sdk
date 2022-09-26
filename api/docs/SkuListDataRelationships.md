# SkuListDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Skus** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**SkuListItems** | Pointer to [**SkuListDataRelationshipsSkuListItems**](SkuListDataRelationshipsSkuListItems.md) |  | [optional] 
**Bundles** | Pointer to [**OrderDataRelationshipsAvailableFreeBundles**](OrderDataRelationshipsAvailableFreeBundles.md) |  | [optional] 

## Methods

### NewSkuListDataRelationships

`func NewSkuListDataRelationships() *SkuListDataRelationships`

NewSkuListDataRelationships instantiates a new SkuListDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListDataRelationshipsWithDefaults

`func NewSkuListDataRelationshipsWithDefaults() *SkuListDataRelationships`

NewSkuListDataRelationshipsWithDefaults instantiates a new SkuListDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *SkuListDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *SkuListDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *SkuListDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *SkuListDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSkus

`func (o *SkuListDataRelationships) GetSkus() BundleDataRelationshipsSkus`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *SkuListDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *SkuListDataRelationships) SetSkus(v BundleDataRelationshipsSkus)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *SkuListDataRelationships) HasSkus() bool`

HasSkus returns a boolean if a field has been set.

### GetSkuListItems

`func (o *SkuListDataRelationships) GetSkuListItems() SkuListDataRelationshipsSkuListItems`

GetSkuListItems returns the SkuListItems field if non-nil, zero value otherwise.

### GetSkuListItemsOk

`func (o *SkuListDataRelationships) GetSkuListItemsOk() (*SkuListDataRelationshipsSkuListItems, bool)`

GetSkuListItemsOk returns a tuple with the SkuListItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListItems

`func (o *SkuListDataRelationships) SetSkuListItems(v SkuListDataRelationshipsSkuListItems)`

SetSkuListItems sets SkuListItems field to given value.

### HasSkuListItems

`func (o *SkuListDataRelationships) HasSkuListItems() bool`

HasSkuListItems returns a boolean if a field has been set.

### GetBundles

`func (o *SkuListDataRelationships) GetBundles() OrderDataRelationshipsAvailableFreeBundles`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *SkuListDataRelationships) GetBundlesOk() (*OrderDataRelationshipsAvailableFreeBundles, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *SkuListDataRelationships) SetBundles(v OrderDataRelationshipsAvailableFreeBundles)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *SkuListDataRelationships) HasBundles() bool`

HasBundles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


