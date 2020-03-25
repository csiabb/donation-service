/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package bcadapter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/structs"

	"github.com/go-resty/resty/v2"
)

const (
	apiVersion = "v1"
)

const (
	urlBlockChainAccounts    = "blockchain/accounts"
	urlBlockChainPublicities = "blockchain/publicities"
)

const (
	retryCount = 3  // retry times
	timeOut    = 40 // seconds
)

// Register defines register on block chain
func (bc *BackendImpl) Register(accountID string) (*structs.RegisterResp, error) {
	logger.Info("got bc adapter register request")

	client := resty.New()

	body := structs.RegisterReq{AccountID: accountID}
	result := &structs.RegisterResp{}

	resp, err := client.R().
		SetHeader(rest.HeaderContentType, rest.HeaderApplicationJSON).
		SetBody(body).
		SetResult(result).
		Post(fmt.Sprintf("%s/api/%s/%s", bc.Config.Address, apiVersion, urlBlockChainAccounts))

	logger.Debug("request body : %v, result : %v", body, result)

	if err != nil {
		e := fmt.Errorf("register error, %v", err)
		logger.Error(e)
		return nil, e
	}

	if resp.StatusCode() != http.StatusOK {
		e := fmt.Errorf("register failed, code : %v, msg : %v", resp.StatusCode(), resp.Status())
		logger.Error(e)
		return nil, e
	}

	logger.Info("bc adapter register succeed")
	return result, nil
}

// Pubs defines publicity data on block chain
func (bc *BackendImpl) Pubs(bcID string, pubs []*string) ([]*structs.PubResp, error) {
	logger.Info("got bc adapter publicity request")

	results := make([]*structs.PubResp, 0)

	for _, pub := range pubs {
		client := resty.New()
		client.SetRetryCount(retryCount).SetTimeout(timeOut * time.Second)

		body := structs.PubReq{UID: bcID, Publicity: *pub}
		result := &structs.PubResp{}

		resp, err := client.R().
			SetHeader(rest.HeaderContentType, rest.HeaderApplicationJSON).
			SetBody(body).
			SetResult(result).
			Post(fmt.Sprintf("%s/api/%s/%s", bc.Config.Address, apiVersion, urlBlockChainPublicities))

		logger.Debug("request body : %v, result : %v", body, result)

		if err != nil {
			results = append(results, nil)
			e := fmt.Errorf("pub error, %v", err)
			logger.Error(e)
			continue
		}

		if resp.StatusCode() != http.StatusOK {
			results = append(results, nil)
			e := fmt.Errorf("pub failed, code : %v, msg : %v", resp.StatusCode(), resp.Status())
			logger.Error(e)
			continue
		}

		results = append(results, result)
	}

	logger.Info("bc adapter publicity succeed")
	return results, nil
}
