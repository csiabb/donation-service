/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/csiabb/donation-service/structs"
)

// ConvertFundsDonation ...
func (funds *PubFunds) ConvertFundsDonation(images []*Image) (string, error) {
	if funds == nil {
		return "", errors.New("para m is nil")
	}

	fd := &structs.FundsDonation{
		ID:                funds.ID,
		UID:               funds.UID,
		DonorName:         funds.DonorName,
		Time:              time.Now().Unix(),
		Amount:            funds.Amount.String(),
		TargetName:        funds.TargetName,
		TargetBankCardNum: funds.TargetBankCardNum,
		DonationImages:    convertImages(images),
	}

	byte, err := json.Marshal(fd)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// ConvertFundsReceived ...
func (funds *PubFunds) ConvertFundsReceived(images []*Image) (string, error) {
	if funds == nil {
		return "", errors.New("para m is nil")
	}

	fd := &structs.FundsReceived{
		ID:                funds.ID,
		TargetUID:         funds.TargetUID,
		TargetName:        funds.TargetName,
		DonorName:         funds.DonorName,
		Time:              time.Now().Unix(),
		Amount:            funds.Amount.String(),
		TargetBankCardNum: funds.TargetBankCardNum,
		DonationImages:    convertImages(images),
	}

	byte, err := json.Marshal(fd)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// ConvertFundsDistributed ...
func (funds *PubFunds) ConvertFundsDistributed(images []*Image) (string, error) {
	if funds == nil {
		return "", errors.New("para m is nil")
	}

	fd := &structs.FundsDistributed{
		ID:                funds.ID,
		TargetUID:         funds.TargetUID,
		TargetName:        funds.TargetName,
		TargetBankCardNum: funds.TargetBankCardNum,
		AidName:           funds.AidName,
		Time:              time.Now().Unix(),
		Amount:            funds.Amount.String(),
		DonationImages:    convertImages(images),
	}

	byte, err := json.Marshal(fd)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// ConvertSuppliesDonation ...
func (supplies *PubSupplies) ConvertSuppliesDonation(billingAddr *Address, shippingAddr *Address, images []*Image) (string, error) {
	if supplies == nil {
		return "", errors.New("para m is nil")
	}

	sp := &structs.SuppliesDonation{
		ID:              supplies.ID,
		UID:             supplies.UID,
		DonorName:       supplies.DonorName,
		BillingAddress:  billingAddr.FullAddress(),
		Time:            time.Now().Unix(),
		Name:            supplies.Name,
		Number:          supplies.Number,
		Unit:            supplies.Unit,
		TargetName:      supplies.TargetName,
		ShippingAddress: shippingAddr.FullAddress(),
		WayBillNum:      supplies.WayBillNum,
		DonationImages:  convertImages(images),
	}

	byte, err := json.Marshal(sp)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// ConvertSuppliesReceived ...
func (supplies *PubSupplies) ConvertSuppliesReceived(billingAddr *Address, shippingAddr *Address, images []*Image) (string, error) {
	if supplies == nil {
		return "", errors.New("para m is nil")
	}

	sp := &structs.SuppliesReceived{
		ID:              supplies.ID,
		TargetUID:       supplies.TargetUID,
		DonorName:       supplies.DonorName,
		BillingAddress:  billingAddr.FullAddress(),
		Time:            time.Now().Unix(),
		Name:            supplies.Name,
		Number:          supplies.Number,
		Unit:            supplies.Unit,
		TargetName:      supplies.TargetName,
		ShippingAddress: shippingAddr.FullAddress(),
		WayBillNum:      supplies.WayBillNum,
		DonationImages:  convertImages(images),
	}

	byte, err := json.Marshal(sp)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// SuppliesDistributed ...
func (supplies *PubSupplies) SuppliesDistributed(billingAddr *Address, shippingAddr *Address, images []*Image) (string, error) {
	if supplies == nil {
		return "", errors.New("para m is nil")
	}

	sp := &structs.SuppliesDistributed{
		ID:              supplies.ID,
		TargetUID:       supplies.TargetUID,
		TargetName:      supplies.TargetName,
		BillingAddress:  billingAddr.FullAddress(),
		Name:            supplies.Name,
		Number:          supplies.Number,
		Unit:            supplies.Unit,
		Time:            time.Now().Unix(),
		AidName:         supplies.AidName,
		ShippingAddress: shippingAddr.FullAddress(),
		WayBillNum:      supplies.WayBillNum,
		DonationImages:  convertImages(images),
	}

	byte, err := json.Marshal(sp)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

// FullAddress ...
func (addr *Address) FullAddress() string {
	return addr.Country + addr.Province + addr.City + addr.District + addr.Address
}

func convertImages(images []*Image) []*structs.DonationImage {
	donaImages := make([]*structs.DonationImage, 0)
	for _, v := range images {
		donaImages = append(donaImages, &structs.DonationImage{
			URL:  v.URL + v.Index,
			Hash: v.Hash,
		})
	}
	return donaImages
}
