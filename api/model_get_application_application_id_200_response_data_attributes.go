/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETApplicationApplicationId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETApplicationApplicationId200ResponseDataAttributes{}

// GETApplicationApplicationId200ResponseDataAttributes struct for GETApplicationApplicationId200ResponseDataAttributes
type GETApplicationApplicationId200ResponseDataAttributes struct {
	// The application's internal name.
	Name interface{} `json:"name,omitempty"`
	// The application's kind, can be one of: 'sales_channel', 'integration' and 'webapp'.
	Kind interface{} `json:"kind,omitempty"`
	// Indicates if the application has public access.
	PublicAccess interface{} `json:"public_access,omitempty"`
	// The application's redirect URI.
	RedirectUri interface{} `json:"redirect_uri,omitempty"`
	// The application's scopes.
	Scopes interface{} `json:"scopes,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETApplicationApplicationId200ResponseDataAttributes instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETApplicationApplicationId200ResponseDataAttributes() *GETApplicationApplicationId200ResponseDataAttributes {
	this := GETApplicationApplicationId200ResponseDataAttributes{}
	return &this
}

// NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults() *GETApplicationApplicationId200ResponseDataAttributes {
	this := GETApplicationApplicationId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKind() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKindOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return &o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasKind() bool {
	if o != nil && IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given interface{} and assigns it to the Kind field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetKind(v interface{}) {
	o.Kind = v
}

// GetPublicAccess returns the PublicAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccess() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PublicAccess
}

// GetPublicAccessOk returns a tuple with the PublicAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccessOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublicAccess) {
		return nil, false
	}
	return &o.PublicAccess, true
}

// HasPublicAccess returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasPublicAccess() bool {
	if o != nil && IsNil(o.PublicAccess) {
		return true
	}

	return false
}

// SetPublicAccess gets a reference to the given interface{} and assigns it to the PublicAccess field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetPublicAccess(v interface{}) {
	o.PublicAccess = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUri() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUriOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return &o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasRedirectUri() bool {
	if o != nil && IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given interface{} and assigns it to the RedirectUri field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetRedirectUri(v interface{}) {
	o.RedirectUri = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given interface{} and assigns it to the Scopes field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetScopes(v interface{}) {
	o.Scopes = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETApplicationApplicationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETApplicationApplicationId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.PublicAccess != nil {
		toSerialize["public_access"] = o.PublicAccess
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
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
	return toSerialize, nil
}

type NullableGETApplicationApplicationId200ResponseDataAttributes struct {
	value *GETApplicationApplicationId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) Get() *GETApplicationApplicationId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) Set(val *GETApplicationApplicationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETApplicationApplicationId200ResponseDataAttributes(val *GETApplicationApplicationId200ResponseDataAttributes) *NullableGETApplicationApplicationId200ResponseDataAttributes {
	return &NullableGETApplicationApplicationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
