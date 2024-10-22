# POSTPaymentOptions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment option&#39;s name. Wehn blank is inherited by payment source type. | [optional] 
**PaymentSourceType** | **interface{}** | The payment source type. One of &#39;adyen_payments&#39;, &#39;axerve_payments&#39;, &#39;braintree_payments&#39;, &#39;checkout_com_payments&#39;, &#39;credit_cards&#39;, &#39;external_payments&#39;, &#39;klarna_payments&#39;, &#39;paypal_payments&#39;, &#39;satispay_payments&#39;, &#39;stripe_payments&#39;, or &#39;wire_transfers&#39;. | 
**Data** | **interface{}** | The payment options data to be added to the payment source payload. Check payment specific API for more details. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPaymentOptions201ResponseDataAttributes

`func NewPOSTPaymentOptions201ResponseDataAttributes(paymentSourceType interface{}, data interface{}, ) *POSTPaymentOptions201ResponseDataAttributes`

NewPOSTPaymentOptions201ResponseDataAttributes instantiates a new POSTPaymentOptions201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentOptions201ResponseDataAttributesWithDefaults

`func NewPOSTPaymentOptions201ResponseDataAttributesWithDefaults() *POSTPaymentOptions201ResponseDataAttributes`

NewPOSTPaymentOptions201ResponseDataAttributesWithDefaults instantiates a new POSTPaymentOptions201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *POSTPaymentOptions201ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPaymentSourceType

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.


### SetPaymentSourceTypeNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetData

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetReference

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPaymentOptions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPaymentOptions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPaymentOptions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPaymentOptions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPaymentOptions201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPaymentOptions201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


