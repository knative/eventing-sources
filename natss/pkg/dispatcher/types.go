/*
Copyright 2019 The Knative Authors

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

package dispatcher

import (
	"fmt"

	eventingduck "github.com/knative/eventing/pkg/apis/duck/v1alpha1"
)

type subscriptionReference struct {
	UID           string
	SubscriberURI string
	ReplyURI      string
}

func newSubscriptionReference(spec eventingduck.SubscriberSpec) subscriptionReference {
	return subscriptionReference{
		UID:           string(spec.UID),
		SubscriberURI: spec.SubscriberURI,
		ReplyURI:      spec.ReplyURI,
	}
}

func (r *subscriptionReference) String() string {
	return fmt.Sprintf("%s", r.UID)
}
