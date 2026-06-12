package sharedResponseDTO

// Common value types shared across the Amadeus hotel APIs. These are identical
// across the search, booking, and content schemas, so they live here and the
// per-module DTOs alias them instead of redefining them.

type (
	// MaxPersonCapacityResponse describes the occupancy capacity of a room.
	MaxPersonCapacityResponse struct {
		// Adults - maximum number of adults the room can accommodate.
		Adults int `json:"adults,omitempty"`
		// Children - maximum number of children the room can accommodate.
		Children int `json:"children,omitempty"`
		// Total - maximum total number of persons the room can accommodate.
		Total int `json:"total,omitempty"`
	}

	// MaxSleepFurnishingsResponse describes the extra sleeping furniture available.
	MaxSleepFurnishingsResponse struct {
		// Cribs - number of cribs available in the room.
		Cribs int `json:"cribs,omitempty"`
		// ExtraBeds - number of extra beds available in the room.
		ExtraBeds int `json:"extraBeds,omitempty"`
	}

	// RateFamilyEstimatedResponse describes the estimated rate family of an offer.
	RateFamilyEstimatedResponse struct {
		// Code - estimated rate family code. Examples: PRO, FAM, GOV. Pattern: [A-Z0-9]{3}.
		Code string `json:"code,omitempty"`
		// Type - type of the rate. P=public, N=negotiated, C=conditional. Pattern: [PNC].
		Type string `json:"type,omitempty"`
	}

	// WarningSourceResponse identifies the request element that triggered a warning.
	WarningSourceResponse struct {
		// Parameter - the key of the URI path or query parameter that caused the warning.
		Parameter string `json:"parameter,omitempty"`
		// Pointer - a JSON Pointer [RFC6901] to the associated entity in the request body.
		Pointer string `json:"pointer,omitempty"`
		// Example - a sample input to guide the user when resolving the issue.
		Example string `json:"example,omitempty"`
	}
)
