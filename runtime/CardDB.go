package runtime

import (
	"arkham-go/card"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var DO_LEECH = false

type CardDB struct {
	cards map[string]card.ArkhamCard

	investigators map[string]card.Investigator
	treachery     map[string]card.Treachery
	asset         map[string]card.Asset
	event         map[string]card.Event
	enemy         map[string]card.Enemy
	skill         map[string]card.Skill
	scenario      map[string]card.Scenario
	agenda        map[string]card.Agenda
	act           map[string]card.Act
	location      map[string]card.Location
	story         map[string]card.Story
}

func NewCardDB() *CardDB {
	return &CardDB{
		cards:         make(map[string]card.ArkhamCard),
		investigators: make(map[string]card.Investigator),
		treachery:     make(map[string]card.Treachery),
		asset:         make(map[string]card.Asset),
		event:         make(map[string]card.Event),
		enemy:         make(map[string]card.Enemy),
		skill:         make(map[string]card.Skill),
		scenario:      make(map[string]card.Scenario),
		agenda:        make(map[string]card.Agenda),
		act:           make(map[string]card.Act),
		location:      make(map[string]card.Location),
		story:         make(map[string]card.Story),
	}
}

func (d *CardDB) Init(file string) error {
	r, err := os.Open(file)
	if err != nil {
		return err
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	cardList := make([]json.RawMessage, 0)
	err = json.Unmarshal(b, &cardList)
	if err != nil {
		return err
	}
	for _, v := range cardList {
		ct := new(card.CType)
		err = json.Unmarshal(v, ct)
		if err != nil {
			return err
		}
		crd, err := d.processCard(ct.CardType, v)
		if err != nil {
			return err
		}
		d.cards[crd.CardCode()] = crd
		leechAssets(crd)
	}
	return nil
}

func (d *CardDB) processCard(cardType string, v json.RawMessage) (card.ArkhamCard, error) {
	var crd card.ArkhamCard

	switch cardType {
	case card.InvestigatorType:
		crd = dec(v, new(card.Investigator))
		i, _ := crd.(*card.Investigator)
		d.investigators[crd.CardCode()] = *i
	case card.TreacheryType:
		crd = dec(v, new(card.Treachery))
		i, _ := crd.(*card.Treachery)
		d.treachery[crd.CardCode()] = *i
	case "asset":
		crd = new(card.Asset)
		crd = dec(v, new(card.Asset))
		i, _ := crd.(*card.Asset)
		d.asset[crd.CardCode()] = *i
	case "event":
		crd = new(card.Event)
		crd = dec(v, new(card.Event))
		i, _ := crd.(*card.Event)
		d.event[crd.CardCode()] = *i
	case "enemy":
		crd = new(card.Enemy)
		crd = dec(v, new(card.Enemy))
		i, _ := crd.(*card.Enemy)
		d.enemy[crd.CardCode()] = *i
	case "skill":
		crd = new(card.Skill)
		crd = dec(v, new(card.Skill))
		i, _ := crd.(*card.Skill)
		d.skill[crd.CardCode()] = *i
	case "scenario":
		crd = new(card.Scenario)
		crd = dec(v, new(card.Scenario))
		i, _ := crd.(*card.Scenario)
		d.scenario[crd.CardCode()] = *i
	case "agenda":
		crd = new(card.Agenda)
		crd = dec(v, new(card.Agenda))
		i, _ := crd.(*card.Agenda)
		d.agenda[crd.CardCode()] = *i
	case "act":
		crd = new(card.Act)
		crd = dec(v, new(card.Act))
		i, _ := crd.(*card.Act)
		d.act[crd.CardCode()] = *i
	case "location":
		crd = new(card.Location)
		crd = dec(v, new(card.Location))
		i, _ := crd.(*card.Location)
		d.location[crd.CardCode()] = *i
	case "story":
		crd = new(card.Story)
		crd = dec(v, new(card.Story))
		i, _ := crd.(*card.Story)
		d.story[crd.CardCode()] = *i
	default:
		return nil, fmt.Errorf("unknown Card %s", cardType)
	}

	return crd, nil

}

func dec(v json.RawMessage, c card.ArkhamCard) card.ArkhamCard {
	err := json.Unmarshal(v, c)
	if err != nil {
		panic(err)
	}
	return c
}

func (d *CardDB) FindCardByName(name string) card.ArkhamCard {
	for _, v := range d.cards {
		if v.Base().Name == name {
			//fixme find duplicates check for cards without name
			return v
		}
	}
	return nil
}

func (d *CardDB) Status() string {
	return fmt.Sprintf(`
Investigators		 %d
Treacery		 %d
Assets			 %d
Event 			 %d
Enemy 			 %d
Skill 			 %d
Scenario 		 %d
Agenda 			 %d
Act 			 %d
Location 		 %d
Story 			 %d
-------------------
Total 			 %d
`, len(d.investigators), len(d.treachery), len(d.asset), len(d.event), len(d.enemy), len(d.skill), len(d.scenario), len(d.agenda), len(d.act), len(d.location), len(d.story), len(d.cards))
}

func leechAssets(crd card.ArkhamCard) {

	if !DO_LEECH {
		return
	}

	if crd.Base().Image != "" {
		processImageAsset(crd, false)
	}
	if crd.Base().BackImage != "" {
		processImageAsset(crd, true)
	}

}

func processImageAsset(v card.ArkhamCard, backimage bool) {

	cardFile := fmt.Sprintf("leech-img/%s.png", v.CardCode())
	if backimage {
		cardFile = fmt.Sprintf("leech-img/%sb.png", v.CardCode())
	}
	fi, err := os.Stat(cardFile)
	if err != nil || fi.Size() == 0 {
		target := fmt.Sprintf("https://arkhamdb.com%s", v.Base().Image)
		if backimage {
			target = fmt.Sprintf("https://arkhamdb.com%s", v.Base().BackImage)
		}

		if strings.HasSuffix(target, ".png") {
			fmt.Printf("Downloading: %s\n", target)
			resp, err := http.Get(target)
			if err != nil {
				fmt.Printf("Err %v", err)
				return
			}
			if resp.StatusCode != 200 {
				fmt.Printf("Err: Status Code: %s for %s\n", resp.Status, target)
				return
			}
			of, err := os.Create(cardFile)
			defer of.Close()
			if err != nil {
				fmt.Printf("Err %v", err)
				return
			}
			io.Copy(of, resp.Body)
		} else {
			fmt.Printf("Skipping unknown Target image: %s\n", target)
		}

	}
}
