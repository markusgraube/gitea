// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullCompare(t *testing.T) {
	prepareTestEnv(t)

	session := loginUser(t, "user2", "password")
	req, err := http.NewRequest("GET", "/user2/repo1/pulls", nil)
	assert.NoError(t, err)
	resp := session.MakeRequest(t, req)
	assert.EqualValues(t, http.StatusOK, resp.HeaderCode)
	htmlDoc, err := NewHtmlParser(resp.Body)
	assert.NoError(t, err)
	link, exists := htmlDoc.doc.Find(".navbar").Find(".ui.green.button").Attr("href")
	assert.True(t, exists, "The template has changed")

	req, err = http.NewRequest("GET", link, nil)
	assert.NoError(t, err)
	resp = session.MakeRequest(t, req)
	assert.EqualValues(t, http.StatusOK, resp.HeaderCode)
}
