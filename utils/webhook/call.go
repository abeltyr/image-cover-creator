package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func Call(
	id string,
	imageProcessingId string,
) {
	url := os.Getenv("WEBHOOK_URL")
	postBody, _ := json.Marshal(map[string]interface{}{
		"artId":             id,
		"imageProcessingId": imageProcessingId,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Println(err)
		return
	}

}
