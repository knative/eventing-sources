/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
)

type Heartbeat struct {
	Sequence int    `json:"id"`
	Label    string `json:"label"`
}

func receive(event cloudevents.Event) {
	ec := event.Context.AsV02()
	hb := &Heartbeat{}
	if err := event.DataAs(hb); err != nil {
		fmt.Printf("got data error: %s\n", err.Error())
	}
	log.Printf("CloudEvent:\n%s", event)
<<<<<<< HEAD
	log.Printf("[%s] %s %s: ", ec.Time, event.DataContentType(), ec.Source.String())
=======
	log.Printf("[%s] %s %s: ", ec.Time, ec.ContentType, ec.Source.String())
>>>>>>> Update cloudevent sdk to include auto-generated uuid options.
	log.Printf("\t%d, %q", hb.Sequence, hb.Label)
}

func main() {
	ctx := context.TODO()

	if _, err := client.StartHTTPReceiver(ctx, receive); err != nil {
		log.Fatalf("failed to start receiver: %s", err.Error())
	}

	log.Printf("listening on port %d\n", 8080)
	<-ctx.Done()
}
