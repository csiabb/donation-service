/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/context"
	"testing"
)

func TestImages(t *testing.T) {
	// init configure
	conf := &config.SrvcCfg{}
	con := &context.Context{Config: conf}

	imgImpl, err := NewImages(con)
	if err != nil {
		logger.Errorf("Failed initialization.")
		return
	}

	// read
	srcImage, err := imgImpl.Load("/Users/nancy/Downloads/test.png")
	if err != nil {
		t.Errorf("Failed to read the picture.")
		return
	}

	// create
	thumbnail, err := imgImpl.Compression(srcImage, 320)
	if err != nil {
		t.Errorf("Failed to create thumbnail.")
		return
	}

	// write
	err = imgImpl.Save(thumbnail, "/Users/nancy/Downloads/thumbnail.png")
	if err != nil {
		t.Errorf("Failed to write thumbnail.")
		return
	}
}
