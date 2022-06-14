/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
)

// AttachDiskOptions attach disk options
type AttachDiskOptions struct {
	cachingMode             compute.CachingTypes
	diskName                string
	diskEncryptionSetID     string
	writeAcceleratorEnabled bool
	lun                     int32
	lunCh                   chan int32 // channel for early return of lun value
}

func NewAttachDiskOptions(cachingMode compute.CachingTypes, diskName, diskEncryptionSetID string, writeAcceleratorEnabled bool, lun int32, lunCh chan int32) *AttachDiskOptions {
	return &AttachDiskOptions{
		cachingMode:             cachingMode,
		diskName:                diskName,
		diskEncryptionSetID:     diskEncryptionSetID,
		writeAcceleratorEnabled: writeAcceleratorEnabled,
		lun:                     lun,
		lunCh:                   lunCh,
	}
}

func (a *AttachDiskOptions) CachingMode() compute.CachingTypes {
	return a.cachingMode
}

func (a *AttachDiskOptions) DiskName() string {
	return a.diskName
}

func (a *AttachDiskOptions) DiskEncryptionSetID() string {
	return a.diskEncryptionSetID
}

func (a *AttachDiskOptions) WriteAcceleratorEnabled() bool {
	return a.writeAcceleratorEnabled
}

func (a *AttachDiskOptions) Lun() int32 {
	return a.lun
}

func (a *AttachDiskOptions) LunChannel() chan int32 {
	return a.lunCh
}

func (a *AttachDiskOptions) SetLun(lun int32) {
	a.lun = lun
}

func (a *AttachDiskOptions) SetDiskEncryptionSetID(diskEncryptionSetID string) {
	a.diskEncryptionSetID = diskEncryptionSetID
}

func (a *AttachDiskOptions) CleanUpLunChannel() {
	if a.lunCh != nil {
		close(a.lunCh)
	}
}

func (a *AttachDiskOptions) String() string {
	return fmt.Sprintf("AttachDiskOptions{diskName: %q, lun: %d}", a.diskName, a.lun)
}

type AttachDiskParams struct {
	diskURI string
	options *AttachDiskOptions
	async   bool
}

func NewAttachDiskParams(diskURI string, options *AttachDiskOptions, async bool) AttachDiskParams {
	return AttachDiskParams{
		diskURI: diskURI,
		options: options,
		async:   async,
	}
}

func (a *AttachDiskParams) DiskURI() string {
	return a.diskURI
}

func (a *AttachDiskParams) Options() *AttachDiskOptions {
	return a.options
}

func (a *AttachDiskParams) Async() bool {
	return a.async
}

type AttachDiskResult struct {
	lun int32
	err error
}

func NewAttachDiskResult(lun int32, err error) AttachDiskResult {
	return AttachDiskResult{
		lun: lun,
		err: err,
	}
}

func (a *AttachDiskResult) Lun() int32 {
	return a.lun
}

func (a *AttachDiskResult) Error() error {
	return a.err
}

type DetachDiskParams struct {
	diskName string
	diskURI  string
}

func NewDetachDiskParams(diskName, diskURI string) DetachDiskParams {
	return DetachDiskParams{
		diskName: diskName,
		diskURI:  diskURI,
	}
}

func (d *DetachDiskParams) DiskName() string {
	return d.diskName
}

func (d *DetachDiskParams) DiskURI() string {
	return d.diskURI
}
