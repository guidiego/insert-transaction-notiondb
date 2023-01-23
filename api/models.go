package handler

import "github.com/dstotijn/go-notion"

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
