// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// UpdateUserRequest wrapper for the UpdateUser operation
type UpdateUserRequest struct {

	// The OCID of the user.
	UserID string `mandatory:"true" contributesTo:"path" name:"userId"`

	// Request object for updating a user.
	UpdateUserDetails UpdateUserDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

// UpdateUserResponse wrapper for the UpdateUser operation
type UpdateUserResponse struct {

	// The User instance
	User

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
