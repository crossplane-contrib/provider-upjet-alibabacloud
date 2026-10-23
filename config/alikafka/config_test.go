// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package alikafka

import (
	"errors"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// sdkError builds the error upstream's delete functions return for a failed
// RPC call, as it appeared when a consumer group was deleted after its
// instance had been released.
func sdkError(code, message string) error {
	return alicloud.WrapErrorf(tea.NewSDKError(map[string]any{
		"code":    code,
		"message": "code: 400, " + message + " request id: 01A16893-473D-8539-A427-683F5642F315",
		"data": map[string]any{
			"statusCode": 400,
			"Code":       code,
			"Message":    message,
		},
	}), alicloud.DefaultErrorMsg, "alikafka_post-cn-xpn50jm34001:crossplane-example-consumer-group", "DeleteConsumerGroup", alicloud.AlibabaCloudSdkGoERROR)
}

func TestGoneWithReleasedInstance(t *testing.T) {
	cases := map[string]struct {
		err     error
		wantNil bool
	}{
		"Success": {
			wantNil: true,
		},
		"InstanceReleased": {
			err:     sdkError("BIZ.INSTANCE.STATUS.ERROR", "instance is in RELEASED. Please check and try again later."),
			wantNil: true,
		},
		"InstanceReleasedUnderscoreCode": {
			err:     sdkError("BIZ_INSTANCE_STATUS_ERROR", "instance is in RELEASED. Please check and try again later."),
			wantNil: true,
		},
		// The instance still exists, so the child does too and its delete must
		// be retried rather than reported as done.
		"InstanceInAnotherState": {
			err: sdkError("BIZ.INSTANCE.STATUS.ERROR", "instance is in UPGRADING. Please check and try again later."),
		},
		"OtherError": {
			err: sdkError("ONS_SYSTEM_ERROR", "instance is in RELEASED. Please check and try again later."),
		},
		"NonSDKError": {
			err: errors.New("instance is in RELEASED"),
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			calls := 0
			del := goneWithReleasedInstance(func(*schema.ResourceData, any) error {
				calls++
				return tc.err
			})
			err := del(nil, nil)
			if calls != 1 {
				t.Fatalf("upstream delete called %d times, want 1", calls)
			}
			if tc.wantNil && err != nil {
				t.Fatalf("got error %v, want the delete to count as done", err)
			}
			if !tc.wantNil && !errors.Is(err, tc.err) {
				t.Fatalf("got error %v, want upstream's error %v", err, tc.err)
			}
		})
	}
}

func TestGoneWithReleasedInstanceNoDelete(t *testing.T) {
	if goneWithReleasedInstance(nil) != nil {
		t.Fatal("wrapping a nil delete function must leave it nil")
	}
}
