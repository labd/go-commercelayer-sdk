# POSTGiftCardRecipients201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "gift_card_recipients"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTCouponRecipients201ResponseDataAttributes**](POSTCouponRecipients201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTCouponRecipients201ResponseDataRelationships**](POSTCouponRecipients201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTGiftCardRecipients201ResponseData

`func NewPOSTGiftCardRecipients201ResponseData() *POSTGiftCardRecipients201ResponseData`

NewPOSTGiftCardRecipients201ResponseData instantiates a new POSTGiftCardRecipients201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTGiftCardRecipients201ResponseDataWithDefaults

`func NewPOSTGiftCardRecipients201ResponseDataWithDefaults() *POSTGiftCardRecipients201ResponseData`

NewPOSTGiftCardRecipients201ResponseDataWithDefaults instantiates a new POSTGiftCardRecipients201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTGiftCardRecipients201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTGiftCardRecipients201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTGiftCardRecipients201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTGiftCardRecipients201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTGiftCardRecipients201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTGiftCardRecipients201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTGiftCardRecipients201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTGiftCardRecipients201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTGiftCardRecipients201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTGiftCardRecipients201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTGiftCardRecipients201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTGiftCardRecipients201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTGiftCardRecipients201ResponseData) GetAttributes() POSTCouponRecipients201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTGiftCardRecipients201ResponseData) GetAttributesOk() (*POSTCouponRecipients201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTGiftCardRecipients201ResponseData) SetAttributes(v POSTCouponRecipients201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTGiftCardRecipients201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTGiftCardRecipients201ResponseData) GetRelationships() POSTCouponRecipients201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTGiftCardRecipients201ResponseData) GetRelationshipsOk() (*POSTCouponRecipients201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTGiftCardRecipients201ResponseData) SetRelationships(v POSTCouponRecipients201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTGiftCardRecipients201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


