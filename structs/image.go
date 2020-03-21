/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// ImageUploadResp defines the response of image upload
type ImageUploadResp struct {
	ID string `json:"id"` // image id
}

// ShareRequest  defines the request of query share
type ShareRequest struct {
	ShareType    string `form:"share_type"`
	DonationType string `form:"donation_type"`
	DonationID   string `form:"donation_id"`
	Scene        string `form:"scene"`
}

// ShareResp defines the response of image share
type ShareResp struct {
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}
