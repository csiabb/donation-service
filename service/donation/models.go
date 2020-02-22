/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package donation

// ListDefermentDonation ...
type ListDonationResponse struct {
	Compare int64 `json:"compare"`
}

// ListDonationResponses ...
type ListDonationResponses []ListDonationResponse

// Len ...
func (s ListDonationResponses) Len() int { return len(s) }

// Swap ...
func (s ListDonationResponses) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less ...
func (s ListDonationResponses) Less(i, j int) bool { return s[i].Compare < s[j].Compare }

// Data ...
func (s ListDonationResponses) Data() (donations []*ListDonationResponse) {
	donations = make([]*ListDonationResponse, 0)
	for i := 0; i < len(s); i++ {
		donations = append(donations, &s[i])
	}

	return
}
