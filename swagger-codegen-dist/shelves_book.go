/*
 * shelves.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ShelvesBook struct {

	// Resource name of the book. It must have the format of \"shelves/_*_/books/_*\". For example: \"shelves/shelf1/books/book2\".
	Name string `json:"name,omitempty"`
}
