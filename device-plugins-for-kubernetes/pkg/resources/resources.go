/*
 *  Copyright (C) 2025 Intel Corporation
 *  SPDX-License-Identifier: Apache-2.0
 */
package resources

import pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

// Resource defines the interface for all device resources
type Resource interface {
	GetResourceName() string
	GetSocketPath() string
	ListDevices() []*pluginapi.Device
	Allocate(deviceIDs []string) []*pluginapi.ContainerAllocateResponse
}
