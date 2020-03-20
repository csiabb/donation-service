/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import "image"

//go:generate mockgen -destination=mock_backend/mock_backend.go -package=mock_backend github.com/csiabb/donation-service/components/image IImageBackend

// IImageBackend ...
type IImageBackend interface {
	Init() error
	CreateDonationImage(content []string, appID string, secret string, scene string, isShare bool) (*image.NRGBA, error)
}
