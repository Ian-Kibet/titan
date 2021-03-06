// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package vsock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemporaryNetErr(t *testing.T) {
	assert.True(t, isTemporaryNetErr(&ackError{cause: errors.New("ack")}))

	assert.False(t, isTemporaryNetErr(&connectMsgError{cause: errors.New("connect")}))
	assert.False(t, isTemporaryNetErr(errors.New("something else")))
	assert.False(t, isTemporaryNetErr(nil))
}
