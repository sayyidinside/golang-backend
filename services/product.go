package services

import (
	"golang/backend/helpers"
)

func FetchProducts() helpers.Response {
	var response helpers.Response

	response.Success = true
	response.Message = "Succesfully"

	return response
}
