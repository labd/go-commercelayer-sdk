# POSTSkuLists201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**POSTCouponRecipients201ResponseDataRelationshipsCustomer**](POSTCouponRecipients201ResponseDataRelationshipsCustomer.md) |  | [optional] 
**Skus** | Pointer to [**POSTBundles201ResponseDataRelationshipsSkus**](POSTBundles201ResponseDataRelationshipsSkus.md) |  | [optional] 
**SkuListItems** | Pointer to [**POSTSkuLists201ResponseDataRelationshipsSkuListItems**](POSTSkuLists201ResponseDataRelationshipsSkuListItems.md) |  | [optional] 
**Bundles** | Pointer to [**POSTSkuLists201ResponseDataRelationshipsBundles**](POSTSkuLists201ResponseDataRelationshipsBundles.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Links** | Pointer to [**POSTOrders201ResponseDataRelationshipsLinks**](POSTOrders201ResponseDataRelationshipsLinks.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTSkuLists201ResponseDataRelationships

`func NewPOSTSkuLists201ResponseDataRelationships() *POSTSkuLists201ResponseDataRelationships`

NewPOSTSkuLists201ResponseDataRelationships instantiates a new POSTSkuLists201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuLists201ResponseDataRelationshipsWithDefaults

`func NewPOSTSkuLists201ResponseDataRelationshipsWithDefaults() *POSTSkuLists201ResponseDataRelationships`

NewPOSTSkuLists201ResponseDataRelationshipsWithDefaults instantiates a new POSTSkuLists201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *POSTSkuLists201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTSkuLists201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTSkuLists201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSkus

`func (o *POSTSkuLists201ResponseDataRelationships) GetSkus() POSTBundles201ResponseDataRelationshipsSkus`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetSkusOk() (*POSTBundles201ResponseDataRelationshipsSkus, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *POSTSkuLists201ResponseDataRelationships) SetSkus(v POSTBundles201ResponseDataRelationshipsSkus)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *POSTSkuLists201ResponseDataRelationships) HasSkus() bool`

HasSkus returns a boolean if a field has been set.

### GetSkuListItems

`func (o *POSTSkuLists201ResponseDataRelationships) GetSkuListItems() POSTSkuLists201ResponseDataRelationshipsSkuListItems`

GetSkuListItems returns the SkuListItems field if non-nil, zero value otherwise.

### GetSkuListItemsOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetSkuListItemsOk() (*POSTSkuLists201ResponseDataRelationshipsSkuListItems, bool)`

GetSkuListItemsOk returns a tuple with the SkuListItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListItems

`func (o *POSTSkuLists201ResponseDataRelationships) SetSkuListItems(v POSTSkuLists201ResponseDataRelationshipsSkuListItems)`

SetSkuListItems sets SkuListItems field to given value.

### HasSkuListItems

`func (o *POSTSkuLists201ResponseDataRelationships) HasSkuListItems() bool`

HasSkuListItems returns a boolean if a field has been set.

### GetBundles

`func (o *POSTSkuLists201ResponseDataRelationships) GetBundles() POSTSkuLists201ResponseDataRelationshipsBundles`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetBundlesOk() (*POSTSkuLists201ResponseDataRelationshipsBundles, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *POSTSkuLists201ResponseDataRelationships) SetBundles(v POSTSkuLists201ResponseDataRelationshipsBundles)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *POSTSkuLists201ResponseDataRelationships) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTSkuLists201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTSkuLists201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTSkuLists201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetLinks

`func (o *POSTSkuLists201ResponseDataRelationships) GetLinks() POSTOrders201ResponseDataRelationshipsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetLinksOk() (*POSTOrders201ResponseDataRelationshipsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTSkuLists201ResponseDataRelationships) SetLinks(v POSTOrders201ResponseDataRelationshipsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTSkuLists201ResponseDataRelationships) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVersions

`func (o *POSTSkuLists201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTSkuLists201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTSkuLists201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTSkuLists201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


