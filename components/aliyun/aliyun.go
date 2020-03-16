/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Client aliyun services client
type Client struct {
	ALiYunConfig *Config
}

// UploadObject implement image upload interface
func (c *Client) UploadObject(name string, content io.Reader) (err error) {
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}

	client, err := oss.New(c.ALiYunConfig.Endpoint, c.ALiYunConfig.AccessKeyID, c.ALiYunConfig.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(c.ALiYunConfig.BucketName)
	if err != nil {
		return err
	}

	if err = bucket.PutObject(name, content, options...); err != nil {
		return err
	}

	return
}
