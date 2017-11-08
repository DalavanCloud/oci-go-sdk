// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateDbHomeDetails struct {
	Database *CreateDatabaseDetails `mandatory:"true" json:"database,omitempty"`

	// A valid Oracle database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"true" json:"dbVersion,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model CreateDbHomeDetails) String() string {
	return common.PointerString(model)
}