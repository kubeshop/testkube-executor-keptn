package main

// Here is a generic example of an echo service: https://github.com/keptn-sandbox/echo-service
// It listens for all cloud events (see deploy/service.yaml: PUBSUB_TOPIC wildcard: "sh.keptn.>"") and automatically responds with .started and .finished events
import (
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2" // make sure to use v2 cloudevents here
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

/**
* Here are all the handler functions for the individual event
* See https://github.com/keptn/spec/blob/0.8.0-alpha/cloudevents.md for details on the payload
**/

// GenericLogKeptnCloudEventHandler is a generic handler for Keptn Cloud Events that logs the CloudEvent
func GenericLogKeptnCloudEventHandler(myKeptn *keptnv2.Keptn, incomingEvent cloudevents.Event, data interface{}) error {
	log.Printf("Handling %s Event: %s", incomingEvent.Type(), incomingEvent.Context.GetID())
	log.Printf("CloudEvent %T: %v", data, data)

	return nil
}

// HandleTestTriggeredEvent handles test.triggered events
// TODO: add in your handler code
func HandleTestTriggeredEvent(myKeptn *keptnv2.Keptn, incomingEvent cloudevents.Event, data *keptnv2.TestTriggeredEventData) error {
	log.Printf("Handling test.triggered Event: %s", incomingEvent.Context.GetID())

	// Keptn Service Lifecycle
	// -----------------------------------------------------
	// 1. Send Started Cloud-Event
	// -----------------------------------------------------
	myKeptn.SendTaskStartedEvent(data, ServiceName)

	// -----------------------------------------------------
	// 2. Do your work here...
	// -----------------------------------------------------
	time.Sleep(5 * time.Second) // Example: Wait 5 seconds. Maybe the problem fixes itself.

	// Optional: You might want to send preliminary results to Keptn part way through the work
	myKeptn.SendTaskStatusChangedEvent(data, ServiceName)

	// -----------------------------------------------------
	// 3. Send Action.Finished Cloud-Event
	// -----------------------------------------------------
	myKeptn.SendTaskFinishedEvent(&keptnv2.EventData{
		Status:  keptnv2.StatusSucceeded, // alternative: keptnv2.StatusErrored
		Result:  keptnv2.ResultPass,      // alternative: keptnv2.ResultFailed
		Message: "Successfully sleeped!",
	}, ServiceName)

	return nil
}
