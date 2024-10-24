/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes{}

// GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes struct for GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes
type GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes struct {
	// The tax calculator's internal name.
	Name interface{} `json:"name,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The Avalara account username.
	Username interface{} `json:"username,omitempty"`
	// The Avalara company code.
	CompanyCode interface{} `json:"company_code,omitempty"`
	// Indicates if the transaction will be recorded and visible on the Avalara website.
	CommitInvoice interface{} `json:"commit_invoice,omitempty"`
	// Indicates if the seller is responsible for paying/remitting the customs duty & import tax to the customs authorities.
	Ddp interface{} `json:"ddp,omitempty"`
}

// NewGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes instantiates a new GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes() *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes {
	this := GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes{}
	return &this
}

// NewGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributesWithDefaults instantiates a new GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributesWithDefaults() *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes {
	this := GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUsername() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUsernameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return &o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasUsername() bool {
	if o != nil && IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given interface{} and assigns it to the Username field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetUsername(v interface{}) {
	o.Username = v
}

// GetCompanyCode returns the CompanyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCompanyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompanyCode
}

// GetCompanyCodeOk returns a tuple with the CompanyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCompanyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompanyCode) {
		return nil, false
	}
	return &o.CompanyCode, true
}

// HasCompanyCode returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasCompanyCode() bool {
	if o != nil && IsNil(o.CompanyCode) {
		return true
	}

	return false
}

// SetCompanyCode gets a reference to the given interface{} and assigns it to the CompanyCode field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetCompanyCode(v interface{}) {
	o.CompanyCode = v
}

// GetCommitInvoice returns the CommitInvoice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCommitInvoice() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CommitInvoice
}

// GetCommitInvoiceOk returns a tuple with the CommitInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCommitInvoiceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CommitInvoice) {
		return nil, false
	}
	return &o.CommitInvoice, true
}

// HasCommitInvoice returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasCommitInvoice() bool {
	if o != nil && IsNil(o.CommitInvoice) {
		return true
	}

	return false
}

// SetCommitInvoice gets a reference to the given interface{} and assigns it to the CommitInvoice field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetCommitInvoice(v interface{}) {
	o.CommitInvoice = v
}

// GetDdp returns the Ddp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetDdp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Ddp
}

// GetDdpOk returns a tuple with the Ddp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetDdpOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Ddp) {
		return nil, false
	}
	return &o.Ddp, true
}

// HasDdp returns a boolean if a field has been set.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasDdp() bool {
	if o != nil && IsNil(o.Ddp) {
		return true
	}

	return false
}

// SetDdp gets a reference to the given interface{} and assigns it to the Ddp field.
func (o *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetDdp(v interface{}) {
	o.Ddp = v
}

func (o GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.CompanyCode != nil {
		toSerialize["company_code"] = o.CompanyCode
	}
	if o.CommitInvoice != nil {
		toSerialize["commit_invoice"] = o.CommitInvoice
	}
	if o.Ddp != nil {
		toSerialize["ddp"] = o.Ddp
	}
	return toSerialize, nil
}

type NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes struct {
	value *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) Get() *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) Set(val *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes(val *GETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) *NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes {
	return &NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}