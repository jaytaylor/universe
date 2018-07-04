// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: package.proto

/*
Package domain is a generated protocol buffer package.

It is generated from these files:
	package.proto
	pending_references.proto
	remote_crawler.proto
	tocrawlentry.proto

It has these top-level messages:
	Package
	PackageReferences
	PackageReference
	PackageCrawl
	PackageSnapshot
	SubPackage
	PendingReferences
	CrawlResult
	ToCrawlEntry
*/
package domain

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *Package) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.Data != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return go_proto_validators.FieldError("Data", err)
		}
	}
	for _, item := range this.History {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("History", err)
			}
		}
	}
	if this.FirstSeenAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.FirstSeenAt); err != nil {
			return go_proto_validators.FieldError("FirstSeenAt", err)
		}
	}
	return nil
}
func (this *PackageReferences) Validate() error {
	for _, item := range this.Refs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Refs", err)
			}
		}
	}
	return nil
}
func (this *PackageReference) Validate() error {
	if this.FirstSeenAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.FirstSeenAt); err != nil {
			return go_proto_validators.FieldError("FirstSeenAt", err)
		}
	}
	if this.LastSeenAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.LastSeenAt); err != nil {
			return go_proto_validators.FieldError("LastSeenAt", err)
		}
	}
	return nil
}
func (this *PackageCrawl) Validate() error {
	if this.JobStartedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.JobStartedAt); err != nil {
			return go_proto_validators.FieldError("JobStartedAt", err)
		}
	}
	if this.JobFinishedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.JobFinishedAt); err != nil {
			return go_proto_validators.FieldError("JobFinishedAt", err)
		}
	}
	if this.Data != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PackageSnapshot) Validate() error {
	if this.CreatedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SubPackage) Validate() error {
	if this.FirstSeenAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.FirstSeenAt); err != nil {
			return go_proto_validators.FieldError("FirstSeenAt", err)
		}
	}
	if this.LastSeenAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.LastSeenAt); err != nil {
			return go_proto_validators.FieldError("LastSeenAt", err)
		}
	}
	return nil
}
