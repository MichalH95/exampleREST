package error_response

const InvalidClientType = "Invalid Client.ClientType, specify either \"Company\" or \"Person\""
const NegativeClientId = "Negative client id"
const NoClientWithThisId = "No client found with this ID"

type ErrorResponse struct {
	ErrorMsg string
}
