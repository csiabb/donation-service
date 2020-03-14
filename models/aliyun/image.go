/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadObject implement image upload interface
func (ai *BackendImpl) UploadObject(name string, content io.Reader) (err error) {
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}

	client, err := oss.New(ai.ALiYunConfig.Endpoint, ai.ALiYunConfig.AccessKeyID, ai.ALiYunConfig.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(ai.ALiYunConfig.BucketName)
	if err != nil {
		return err
	}

	if err = bucket.PutObject(name, content, options...); err != nil {
		return err
	}

	return
}
