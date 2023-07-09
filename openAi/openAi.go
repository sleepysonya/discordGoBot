package openai

import (
	storage "github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/util"
)

var (
	OpenAiToken string = storage.OpenAiToken
	OrgId       string = storage.GetEnvVar("ORG_ID")
)

func OpenAi() {

}
