/*
 * ICC PROJECT
 *
 * Welcome to use ICC product
 *
 * API version: 1.0.0
 * Contact: 
 * Generated by: lcu
 */

package service

type User struct {

	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	UserStatus int32 `json:"userStatus,omitempty"`

	Age int32 `json:"age,omitempty"`
}
