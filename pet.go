/*
 * Swagger Petstore
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Pet struct {

	Id int64 `json:"id"`

	Name string `json:"name"`

	Tag string `json:"tag,omitempty"`
}
