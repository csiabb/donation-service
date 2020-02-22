/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package donation

import (
	"sort"

	"github.com/prometheus/common/log"

	"github.com/csiabb/donation-service/models"
)

// ListDonations ...
func (d *DonationsImpl) ListDonations(request *models.Request) (donations []*ListDonationResponse, err error) {
	switch request.DonationType {
	case "fund":
		logger.Debugf("List of donations for funds")
		return d.ListDonationsByFund(request)
	case "supply":
		logger.Debugf("Get a list of donations for supplies")
		return d.ListDonationsByFund(request)
	default:
		logger.Debugf("Gets the default donation list, donation type = %s", request.DonationType)
		return d.ListDefaultDonations(request)
	}

	return nil, nil
}

// ListDonationsBySupply ...
func (d *DonationsImpl) ListDonationsBySupply(request *models.Request) (donations []*ListDonationResponse, err error) {
	supplies, err := d.context.DBStrorage.ListDonationsBySupply(request)
	if err != nil {
		logger.Debugf("Wrong material list, err = %+v", err)
		return
	}

	donations = make([]*ListDonationResponse, 0)
	for i := 0; i < len(supplies); i++ {
		var donation = &ListDonationResponse{Compare: supplies[i].CreatedAt.UnixNano()}
		donations = append(donations, donation)
	}

	log.Debugf("The material list was successful")
	return
}

// ListDonationsByFund ...
func (d *DonationsImpl) ListDonationsByFund(request *models.Request) (donations []*ListDonationResponse, err error) {
	funds, err := d.context.DBStrorage.ListDonationsByFund(request)
	if err != nil {
		logger.Debugf("Wrong funds list, err =%+v", err)
		return
	}

	donations = make([]*ListDonationResponse, 0)
	for i := 0; i < len(funds); i++ {
		var donation = &ListDonationResponse{Compare: funds[i].CreatedAt.UnixNano()}
		donations = append(donations, donation)
	}

	log.Debugf("The fund list was successful")
	return
}

// Sort ...
func (d *DonationsImpl) Sort(funds []*ListDonationResponse, supplies []*ListDonationResponse) (donations []*ListDonationResponse, err error) {
	total := make([]ListDonationResponse, 0)
	for i := 0; i < len(funds) || i < len(supplies); i++ {
		if i <= len(funds)-1 {
			total = append(total, *funds[i])
		}

		if i <= len(supplies)-1 {
			total = append(total, *supplies[i])
		}
	}

	var sorts ListDonationResponses = total
	sort.Sort(sort.Reverse(sorts))
	donations = sorts.Data()

	log.Debugf("Sort successful. total = %+v", donations)

	return
}

// ListDefaultDonations ...
func (d *DonationsImpl) ListDefaultDonations(request *models.Request) (donations []*ListDonationResponse, err error) {
	// fund
	request.DonationType = "fund"
	funds, err := d.ListDonationsByFund(request)
	if err != nil {
		logger.Debugf("Gets the default donation list, donation type = %s", request.DonationType)
		return
	}

	// supply
	request.DonationType = "supply"
	supplies, err := d.ListDonationsBySupply(request)
	if err != nil {
		logger.Debugf("Gets the default donation list, donation type = %s", request.DonationType)
		return
	}

	// sort
	logger.Debugf("Sort the donation data of materials and funds. funds=%+v, supply=%+v", funds, supplies)
	return d.Sort(funds, supplies)
}
