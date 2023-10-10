package generation

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"net/http"

	storage "github.com/sleepysonya/discordGoBot/util"
)

var (
	Token string = storage.GetEnvVar("COHERE_KEY")
)

func GenerateResponse(message string) string {

	url := "https://api.cohere.ai/v1/generate"
	payload := strings.NewReader("{\"max_tokens\":120,\"truncate\":\"END\",\"return_likelihoods\":\"NONE\",\"prompt\":\" " + message + "\"}")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+Token)
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		fmt.Println("Error: ", res.StatusCode)
		fmt.Println(err)
		return "Something went wrong"
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var finalResponse storage.CohereResponse
	json.Unmarshal([]byte(body), &finalResponse)

	return finalResponse.Generations[0].Text
}
