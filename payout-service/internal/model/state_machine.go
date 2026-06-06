package model

var AllowedTransitions = map[string][]string{
	PayoutCreated: {
		PayoutProcessing,
	},
	PayoutProcessing: {
		PayoutSuccess,
		PayoutFailed,
	},
}

func IsValidTransition(
	current string,
	next string,
) bool {

	allowedStates, exists := AllowedTransitions[current]
	if !exists {
		return false
	}

	for _, state := range allowedStates {
		if state == next {
			return true
		}
	}

	return false
}
