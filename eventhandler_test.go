package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/keptn/go-utils/pkg/lib/v0_2_0/fake"

	keptn "github.com/keptn/go-utils/pkg/lib/keptn"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"

	cloudevents "github.com/cloudevents/sdk-go/v2" // make sure to use v2 cloudevents here
)

/**
 * loads a cloud event from the passed test json file and initializes a keptn object with it
 */
func initializeTestObjects(eventFileName string) (*keptnv2.Keptn, *cloudevents.Event, error) {
	// load sample event
	eventFile, err := os.ReadFile(eventFileName)
	if err != nil {
		return nil, nil, fmt.Errorf("Cant load %s: %s", eventFileName, err.Error())
	}

	incomingEvent := &cloudevents.Event{}
	err = json.Unmarshal(eventFile, incomingEvent)
	if err != nil {
		return nil, nil, fmt.Errorf("Error parsing: %s", err.Error())
	}

	// Add a Fake EventSender to KeptnOptions
	var keptnOptions = keptn.KeptnOpts{
		EventSender: &fake.EventSender{},
	}
	keptnOptions.UseLocalFileSystem = true
	myKeptn, err := keptnv2.NewKeptn(incomingEvent, keptnOptions)

	return myKeptn, incomingEvent, err
}

// Tests HandleTestTriggeredEvent
// TODO: Add your test-code
func TestHandleTestTriggeredEvent(t *testing.T) {
	myKeptn, incomingEvent, err := initializeTestObjects("test-events/action.triggered.json")
	if err != nil {
		t.Error(err)
		return
	}

	specificEvent := &keptnv2.TestTriggeredEventData{}
	err = incomingEvent.DataAs(specificEvent)
	if err != nil {
		t.Errorf("Error getting keptn event data")
	}

	err = HandleTestTriggeredEvent(myKeptn, *incomingEvent, specificEvent)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	gotEvents := len(myKeptn.EventSender.(*fake.EventSender).SentEvents)

	// Verify that HandleTestTriggeredEvent has sent 3 cloudevents
	if gotEvents != 3 {
		t.Errorf("Expected two events to be sent, but got %v", gotEvents)
	}

	// Verify that the first CE sent is a .started event
	if keptnv2.GetStartedEventType(keptnv2.ActionTaskName) != myKeptn.EventSender.(*fake.EventSender).SentEvents[0].Type() {
		t.Errorf("Expected a action.started event type, got %s", myKeptn.EventSender.(*fake.EventSender).SentEvents[0].Type())
	}

	// Verify that the first CE sent is a .started event
	if keptnv2.GetStatusChangedEventType(keptnv2.ActionTaskName) != myKeptn.EventSender.(*fake.EventSender).SentEvents[1].Type() {
		t.Errorf("Expected a action.changed event type, got %s", myKeptn.EventSender.(*fake.EventSender).SentEvents[1].Type())
	}

	// Verify that the second CE sent is a .finished event
	if keptnv2.GetFinishedEventType(keptnv2.ActionTaskName) != myKeptn.EventSender.(*fake.EventSender).SentEvents[2].Type() {
		t.Errorf("Expected a action.finished event type, got %s", myKeptn.EventSender.(*fake.EventSender).SentEvents[2].Type())
	}
}
