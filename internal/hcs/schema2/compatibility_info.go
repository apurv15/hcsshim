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

// An opaque VM compatibility data, which is primarily used in migration. This should be provided in MigrationInitializeOptions::CompatibilityData
type CompatibilityInfo struct {
	// Pattern: /^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$/
	Data string `json:"Data,omitempty"`
}
