/*
Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package operator

import (
	"testing"

	"k8s.io/client-go/pkg/api/v1"

	"github.com/rook/rook/pkg/operator/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateCluster(t *testing.T) {
	clientset := test.New(3)
	o := New("foo", "rook", nil, clientset)
	o.retryDelay = 0

	// fail to init k8s client since we're not actually inside k8s
	_, err := o.initResources()
	assert.NotNil(t, err)

	// create the tpr
	err = o.createTPR()
	assert.Nil(t, err)
	tpr, err := clientset.ExtensionsV1beta1().ThirdPartyResources().List(v1.ListOptions{})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tpr.Items))
	assert.Equal(t, "cluster.rook.io", tpr.Items[0].Name)
	assert.Equal(t, tprDescription, tpr.Items[0].Description)

	// TODO: Watch for a new Rook cluster and create it. Need a mocked http client to be working
}
