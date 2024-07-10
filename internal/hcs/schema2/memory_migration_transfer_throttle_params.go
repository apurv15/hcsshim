// Autogenerated code; DO NOT EDIT.

// Schema retrieved from branch 'rs_base2_hyp' and build '27645.1002.240617-2119'.

/*
 * Schema Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

// Settings for migration memory transfer throttling
type MemoryMigrationTransferThrottleParams struct {
	// A flag indicating if throttling should be skipped
	SkipThrottling bool `json:"SkipThrottling,omitempty"`
	// The scale of the throttling. The value is in percentage (1-100).
	ThrottlingScale float32 `json:"ThrottlingScale,omitempty"`
	// Minimum percentage value to which memory transfer can be throttled
	MinimumThrottlePercentage uint8 `json:"MinimumThrottlePercentage,omitempty"`
	// Number of memory transfer passes targetted before the VM enters blackout
	TargetNumberOfBrownoutTransferPasses uint32 `json:"TargetNumberOfBrownoutTransferPasses,omitempty"`
	// The starting transfer pass where throttling is starting
	StartingBrownoutPassNumberForThrottling uint32 `json:"StartingBrownoutPassNumberForThrottling,omitempty"`
	// Maximum number of memory transfer passes before forcing the VM to enter blackout
	MaximumNumberOfBrownoutTransferPasses uint32 `json:"MaximumNumberOfBrownoutTransferPasses,omitempty"`
	// Expected duration for blackout transfer time
	TargetBlackoutTransferTime uint32 `json:"TargetBlackoutTransferTime,omitempty"`
	// Threshold for blackout duration prior to cancelling migration
	BlackoutTimeThresholdForCancellingMigration uint32 `json:"BlackoutTimeThresholdForCancellingMigration,omitempty"`
}
