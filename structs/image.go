/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// ImageUploadResp defines the response of image upload
type ImageUploadResp struct {
	ID string `json:"id"` // image id
}

// DrawRequest  defines the request of query draw
type DrawRequest struct {
	DrawType     string `form:"draw_type"`
	DonationType string `form:"donation_type"`
	DonationID   string `form:"donation_id"`
	Scene        string `form:"scene"`
	IsShare      bool   `form:"is_share"`
}

// DrawResp defines the response of image draw
type DrawResp struct {
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}
