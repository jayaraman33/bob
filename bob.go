package bob

import "strings"

const (
	question      = "Sure."
	yell          = "Whoa, chill out!"
	yell_question = "Calm down, I know what I'm doing!"
	empty         = "Fine. Be that way!"
	other         = "Whatever."
)

func Hey(remark string) string {
	answer := other

	remark = strings.TrimSpace(remark)
	switch {
	case isSilent(remark):
		answer = empty

	case isYellingQuestion(remark):
		answer = yell_question

	case isQuestion(remark):
		answer = question

	case isYelling(remark):
		answer = yell

	default:
		answer = other
	}

	return answer
}

func isYellingQuestion(remark string) bool {
	return isYelling(remark) && isQuestion(remark)
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isSilent(remark string) bool {
	return remark == ""
}

func isYelling(remark string) bool {
	hasLetters := remark != strings.ToLower(remark)
	isUpcase := remark == strings.ToUpper(remark)
	return hasLetters && isUpcase
}
