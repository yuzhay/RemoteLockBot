package remotelock

import (
	"fmt"
)

const URL = "https://connect.devicewebmanager.com"

func UnlockedEventDecorator(eventData ResponseData) string {
	data := eventData
	return fmt.Sprintf(
		"Unlocked: %s\n"+
			"Status: %s\n"+
			"Pin: %s\n"+
			"User: %s ",
		lockPublisherURL(data.Attributes.PublisherID),
		data.Attributes.Status,
		data.Attributes.Pin,
		userURL(data.Attributes.AssociatedResourceID))
}

func LockedEventDecorator(eventData ResponseData) string {
	data := eventData
	return fmt.Sprintf(
		"Locked: %s\nStatus: %s",
		lockPublisherURL(data.Attributes.PublisherID),
		data.Attributes.Status)
}

func AccessDeniedEventDecorator(eventData ResponseData) string {
	data := eventData
	return fmt.Sprintf(
		"Access Denied: %s\n"+
			"Status: %s\n"+
			"Pin: %s\n"+
			"User: %s ",
		lockPublisherURL(data.Attributes.PublisherID),
		data.Attributes.Status,
		data.Attributes.Pin,
		userURL(data.Attributes.AssociatedResourceID))
}

func AccessPersonSyncedEventDecorator(eventData ResponseData) string {
	data := eventData
	return fmt.Sprintf(
		"Access Person Synced: %s\n"+
			"Status: %s\n"+
			"User: %s ",
		lockPublisherURL(data.Attributes.PublisherID),
		data.Attributes.Status,
		userURL(data.Attributes.AssociatedResourceID))
}

func lockPublisherURL(publisher string) string {
	return fmt.Sprintf("%s/locks/%s", URL, publisher)
}

func userURL(person string) string {
	return fmt.Sprintf("%s/access-persons/%s", URL, person)
}
