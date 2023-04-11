# PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The tax calculator&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Username** | Pointer to **interface{}** | The Avalara account username. | [optional] 
**Password** | Pointer to **interface{}** | The Avalara account password. | [optional] 
**CompanyCode** | Pointer to **interface{}** | The Avalara company code. | [optional] 
**CommitInvoice** | Pointer to **interface{}** | Indicates if the transaction will be recorded and visible on the Avalara website. | [optional] 
**Ddp** | Pointer to **interface{}** | Indicates if the seller is responsible for paying/remitting the customs duty &amp; import tax to the customs authorities. | [optional] 

## Methods

### NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes

`func NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes() *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes`

NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes instantiates a new PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributesWithDefaults

`func NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributesWithDefaults() *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes`

NewPATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributesWithDefaults instantiates a new PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetUsername(v interface{})`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetPassword(v interface{})`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetCompanyCode() interface{}`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetCompanyCodeOk() (*interface{}, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetCompanyCode(v interface{})`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### SetCompanyCodeNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetCompanyCodeNil(b bool)`

 SetCompanyCodeNil sets the value for CompanyCode to be an explicit nil

### UnsetCompanyCode
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetCompanyCode()`

UnsetCompanyCode ensures that no value is present for CompanyCode, not even an explicit nil
### GetCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetCommitInvoice() interface{}`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetCommitInvoiceOk() (*interface{}, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetCommitInvoice(v interface{})`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### SetCommitInvoiceNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetCommitInvoiceNil(b bool)`

 SetCommitInvoiceNil sets the value for CommitInvoice to be an explicit nil

### UnsetCommitInvoice
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetCommitInvoice()`

UnsetCommitInvoice ensures that no value is present for CommitInvoice, not even an explicit nil
### GetDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetDdp() interface{}`

GetDdp returns the Ddp field if non-nil, zero value otherwise.

### GetDdpOk

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) GetDdpOk() (*interface{}, bool)`

GetDdpOk returns a tuple with the Ddp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetDdp(v interface{})`

SetDdp sets Ddp field to given value.

### HasDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) HasDdp() bool`

HasDdp returns a boolean if a field has been set.

### SetDdpNil

`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) SetDdpNil(b bool)`

 SetDdpNil sets the value for Ddp to be an explicit nil

### UnsetDdp
`func (o *PATCHAvalaraAccountsAvalaraAccountIdRequestDataAttributes) UnsetDdp()`

UnsetDdp ensures that no value is present for Ddp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


