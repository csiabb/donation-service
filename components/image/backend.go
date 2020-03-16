/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import "image"

// IImageBackend ...
type IImageBackend interface {
	Init() error
	CreateDonationImage(content []string, isShare bool) (*image.NRGBA, error)
}
