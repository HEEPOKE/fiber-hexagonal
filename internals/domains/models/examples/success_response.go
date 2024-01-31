package examples

import "github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"

type SuccessAccountsGetAllResponse struct {
	Status  SuccessStatusMessage  `json:"status"`
	PayLoad []models.AccountModel `json:"payload"`
}
