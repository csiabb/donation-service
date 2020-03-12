/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"io"

	"github.com/csiabb/donation-service/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadObject implement image upload interface
func UploadObject(name string, content io.Reader, aLiYunCfg config.ALiYunCfg) (err error) {
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}

	client, err := oss.New(aLiYunCfg.Endpoint, aLiYunCfg.AccessKeyID, aLiYunCfg.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(aLiYunCfg.BucketName)
	if err != nil {
		return err
	}

	if err = bucket.PutObject(name, content, options...); err != nil {
		return err
	}

	return
}
