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

var bankIds = map[string]string{
	"commerz":  "688b791acb334cea8526f25ae9c116f6",
	"amex":     "e460bf056c8546f9add452e67d2b833c",
	"dinheiro": "bbc28caaefc8470f9d824e8139febd3d",
	"crypto":   "181a341facaf4939bae5392de14d5f36",
}

func buildPagePayload(dbid string, value float64, uuid string, bankSlug string, desc string) notion.CreatePageParams {
	emoji := "🟢"

	if value < 0 {
		emoji = "🔴"
	}

	title := []notion.RichText{
		{
			Text: &notion.Text{Content: uuid},
		},
	}

	description := []notion.RichText{
		{
			Text: &notion.Text{Content: desc},
		},
	}

	relation := []notion.Relation{
		{
			ID: bankIds[bankSlug],
		},
	}

	return notion.CreatePageParams{
		ParentType: notion.ParentTypeDatabase,
		ParentID:   dbid,
		Icon: &notion.Icon{
			Type:  "emoji",
			Emoji: &emoji,
		},
		DatabasePageProperties: &notion.DatabasePageProperties{
			"Ref": notion.DatabasePageProperty{
				Title: title,
			},
			"Desc": notion.DatabasePageProperty{
				RichText: description,
			},
			"Conta": notion.DatabasePageProperty{
				Relation: relation,
			},
			"Valor": notion.DatabasePageProperty{
				Number: &value,
			},
		},
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	notion_db := os.Getenv("NOTION_DB_ID")
	params := r.URL.Query()
	value, _ := strconv.ParseFloat(params["value"][0], 64)
	content := ""

	if r.Header.Get("Authorization") != os.Getenv("PERMISSION_TOKEN") {
		fmt.Fprint(w, "Not Allowed")
		return
	}

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
