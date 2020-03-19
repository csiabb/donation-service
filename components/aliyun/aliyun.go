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
	Client       *oss.Client
	Bucket       *oss.Bucket
	Options      []oss.Option
}

// Init aliyun services init
func (c *Client) Init() error {
	client, err := oss.New(c.ALiYunConfig.Endpoint, c.ALiYunConfig.AccessKeyID, c.ALiYunConfig.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(c.ALiYunConfig.BucketName)
	if err != nil {
		return err
	}

	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}

	c.Client = client
	c.Bucket = bucket
	c.Options = options

	return err
}

// UploadObject implement image upload interface
func (c *Client) UploadObject(name string, content io.Reader) (err error) {
	return c.Bucket.PutObject(name, content, c.Options...)
}

// IsExist determines whether an object exists
func (c *Client) IsExist(name string) (bool, error) {
	return c.Bucket.IsObjectExist(name, c.Options...)
}
