package _struct

import "time"

type Computer struct {
	Kind             string    `json:"kind"`
	ID               string    `json:"id"`
	HostName         string    `json:"hostName"`
	HostComputerID   string    `json:"hostComputerId"`
	MostFrequentUser string    `json:"mostFrequentUser"`
	MostRecentUser   string    `json:"mostRecentUser"`
	Manufacturer     string    `json:"manufacturer"`
	IsPortable       bool      `json:"isPortable"`
	IsVirtual        bool      `json:"isVirtual"`
	IsServer         bool      `json:"isServer"`
	Model            string    `json:"model"`
	OperatingSystem  string    `json:"operatingSystem"`
	Vendor           string    `json:"vendor"`
	Domain           string    `json:"domain"`
	IPAddress        string    `json:"ipAddress"`
	IsVDI            bool      `json:"isVDI"`
	LastScanDate     time.Time `json:"lastScanDate"`
	Status           string    `json:"status"`
	OrganizationID   string    `json:"organizationId"`
	ProcessorCount   int       `json:"processorCount"`
	CoreCount        int       `json:"coreCount"`
}

type ComputerList struct {
	Values   []Computer `json:"values"`
	Count    int        `json:"count"`
	PrevPage string     `json:"prevPage"`
	NextPage string     `json:"nextPage"`
	Kind     string     `json:"kind"`
}
