/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image/png"
	"strconv"
	"time"

	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"
)

var (
	logger = log.MustGetLogger("image-handler")
)

// RestHandler image handler
type RestHandler struct {
	srvcContext *context.Context
}

// NewRestHandler ...
func NewRestHandler(c *context.Context) (*RestHandler, error) {
	return &RestHandler{srvcContext: c}, nil
}

// GetContent define the prove content of donation items
func (h *RestHandler) GetContent(req *structs.DrawRequest) ([]string, error) {
	var err error
	var content []string
	if req.DonationType == rest.DonatedTypeSupplies {
		var supplies *models.SuppliesDetail
		if supplies, err = h.srvcContext.DBStorage.QuerySuppliesDetail(req.DonationID); err != nil {
			return nil, err
		}

		t := time.Unix(supplies.Supplies.BlockTime, 0)
		content = append(content, supplies.Supplies.DonorName,
			supplies.Supplies.TargetName,
			supplies.Supplies.Name+" x"+strconv.Itoa(int(supplies.Supplies.Number)),
			fmt.Sprintf("%d", supplies.Supplies.BlockHeight),
			supplies.Supplies.TxID,
			t.Format("2006-01-02 15:04:05"),
		)
	}

	if req.DonationType == rest.DonatedTypeFunds {
		var funds *models.FundsDetail
		if funds, err = h.srvcContext.DBStorage.QueryFundsDetail(req.DonationID); err != nil {
			return nil, err
		}

		t := time.Unix(funds.Funds.BlockTime, 0)
		amount, _ := funds.Funds.Amount.Float64()
		content = append(content, funds.Funds.DonorName,
			funds.Funds.TargetName,
			fmt.Sprintf("%0.2f", amount),
			fmt.Sprintf("%d", funds.Funds.BlockHeight),
			funds.Funds.TxID,
			t.Format("2006-01-02 15:04:05"),
		)
	}

	return content, err
}

// CreateDonationImage create image of donation prove items
func (h *RestHandler) CreateDonationImage(req *structs.DrawRequest, tag string) error {
	// get content
	content, err := h.GetContent(req)
	if err != nil {
		return err
	}

	// create donation image
	img, err := h.srvcContext.ImageBackend.CreateDonationImage(content, h.srvcContext.Config.WXCfg.AppID,
		h.srvcContext.Config.WXCfg.Secret, req.Scene, req.IsShare)
	if err != nil {
		return err
	}

	// upload
	var b bytes.Buffer
	if err = png.Encode(&b, img); err != nil {
		return err
	}
	r := bytes.NewReader(b.Bytes())

	if err = h.srvcContext.ALiYunBackend.UploadObject(tag+".png", r); err != nil {
		return err
	}

	return nil
}

// GetImageURL donation proof picture aliyun path
func (h *RestHandler) GetImageURL(req *structs.DrawRequest) (string, error) {
	var err error
	var isExist bool
	tag := h.URL(req, 0)
	imagURL := h.srvcContext.Config.LocalFileSystem + tag + ".png"
	if isExist, err = h.srvcContext.ALiYunBackend.IsExist(tag + ".png"); err != nil {
		return "", err
	}

	if isExist {
		return imagURL, nil
	}

	if err = h.CreateDonationImage(req, tag); err != nil {
		return "", err
	}

	return imagURL, nil
}

// URL create aliyun image name
func (h *RestHandler) URL(req *structs.DrawRequest, isShare int) string {
	if req.IsShare {
		isShare = 1
	}
	src := fmt.Sprintf("%s%s%s%d", req.DrawType, req.DonationType, req.DonationID, isShare)
	md5Inst := md5.New()
	md5Inst.Write([]byte(src))
	result := md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}
