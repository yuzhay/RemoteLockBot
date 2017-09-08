package remotelock

import "time"

type Response struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Type       string                 `json:"type"`
	Attributes ResponseDataAttributes `json:"attributes"`
	ID         string                 `json:"id"`
	Links      ResponseDataLink       `json:"links"`
}

type ResponseDataAttributes struct {
	Source                 string    `json:"source"`
	Status                 string    `json:"status"`
	TimeZone               string    `json:"time_zone"`
	OccurredAt             time.Time `json:"occurred_at"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	AssociatedResourceName string    `json:"associated_resource_name"`
	StatusInfo             string    `json:"status_info"`
	Mathod                 string    `json:"method"`
	Card                   string    `json:"card"`
	Pin                    string    `json:"pin"`
	SmartCardSerialNumber  string    `json:"smart_card_serial_number"`
	PublisherID            string    `json:"publisher_id"`
	PublishedType          string    `json:"publisher_type"`
	AssociatedResourceID   string    `json:"associated_resource_id"`
	AssociatedResourceType string    `json:"associated_resource_type"`
}

type ResponseDataLink struct {
	Self               string `json:"self"`
	Publisher          string `json:"publisher"`
	AssociatedResource string `json:"associated_resource"`
}
