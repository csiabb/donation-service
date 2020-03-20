/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

import (
	"io"
)

//go:generate mockgen -destination=mock_backend/mock_backend.go -package=mock_backend github.com/csiabb/donation-service/components/aliyun IALiYunBackend

// IALiYunBackend aliyun services interface
type IALiYunBackend interface {
	IsExist(name string) (bool, error)
	UploadObject(name string, content io.Reader) error
}
