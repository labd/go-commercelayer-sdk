# GETPaymentOptionsPaymentOptionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment option&#39;s name. Wehn blank is inherited by payment source type. | [optional] 
**PaymentSourceType** | Pointer to **interface{}** | The payment source type. One of &#39;adyen_payments&#39;, &#39;axerve_payments&#39;, &#39;braintree_payments&#39;, &#39;checkout_com_payments&#39;, &#39;credit_cards&#39;, &#39;external_payments&#39;, &#39;klarna_payments&#39;, &#39;paypal_payments&#39;, &#39;satispay_payments&#39;, &#39;stripe_payments&#39;, or &#39;wire_transfers&#39;. | [optional] 
**Data** | Pointer to **interface{}** | The payment options data to be added to the payment source payload. Check payment specific API for more details. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributes

`func NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributes() *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes`

NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributes instantiates a new GETPaymentOptionsPaymentOptionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributesWithDefaults

`func NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributesWithDefaults() *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes`

NewGETPaymentOptionsPaymentOptionId200ResponseDataAttributesWithDefaults instantiates a new GETPaymentOptionsPaymentOptionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPaymentSourceType

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.

### HasPaymentSourceType

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasPaymentSourceType() bool`

HasPaymentSourceType returns a boolean if a field has been set.

### SetPaymentSourceTypeNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetData

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCreatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


