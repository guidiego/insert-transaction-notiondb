package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dstotijn/go-notion"
	"github.com/google/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	notion_db := os.Getenv("NOTION_DB_ID")
	params := r.URL.Query()
	value, _ := strconv.ParseFloat(params["value"][0], 64)
	content := ""

	if params["content"] != nil {
		content = params["content"][0]
	}

	if params["action"][0] == "rm" {
		value = value * -1
	}

	n := notion.NewClient(os.Getenv("NOTION_TOKEN"))
	p := buildPagePayload(notion_db, value, uuid.New().String(), params["account"][0], content)

	_, err := n.CreatePage(context.Background(), p)

	if err != nil {
		log.Fatal(err)
	}

	t, err2 := p.MarshalJSON()

	if err2 != nil {
		log.Fatal(err2)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(t))
}
