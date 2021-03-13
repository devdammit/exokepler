package megad2561

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type configPathType string

const (
	FirstConfigPath  = "1"
	SecondConfigPath = "2"
)

type Config struct {
	hostIP       HostIP
	pwd          Pwd
	srv          SRV
	srvType      SrvType
	megadID      MegadID
	mqttPassword MQTTPwd
	script       Script
	gw           GW
}

// SetAndUpdateConfig need for merge when BatchConfigure and ResetConfigure.
func (c *Config) SetAndUpdateConfig(bc BuilderConfig) error {
	newConfig := c

	if bc.srv != c.srv || bc.mqttPassword != c.mqttPassword || bc.srvType != c.srvType {
		newConfig.srv = bc.srv
		newConfig.mqttPassword = bc.mqttPassword
		newConfig.srvType = bc.srvType

		if newConfig.UpdateFirstPartConfig() != nil {
			return ErrUpdateConfig
		}
	}

	if bc.megadID != c.megadID {
		newConfig.megadID = bc.megadID

		if newConfig.UpdateSecondPartConfig() != nil {
			return ErrUpdateConfig
		}
	}

	c.megadID = newConfig.megadID
	c.srv = newConfig.srv
	c.srvType = newConfig.srvType
	c.mqttPassword = newConfig.mqttPassword

	return nil
}

// UpdateFirstPartConfig - публикует первую часть конфига.
// Конфиги разделены на страницы в контроллере megad2561, поэтому отправка по частям.
func (c *Config) UpdateFirstPartConfig() error {
	queries := url.Values{}
	queries.Add("eip", string(c.hostIP))
	queries.Add("pwd", string(c.pwd))
	queries.Add("gw", string(c.gw))
	queries.Add("sip", string(c.srv))
	queries.Add("sct", string(c.script))
	queries.Add("pr", "")
	queries.Add("gsm", "0")
	queries.Add("srvt", fmt.Sprintf("%v", int(c.srvType)))

	if c.srvType == MQTT {
		queries.Add("auth", string(c.mqttPassword))
	}

	return c.updateConfig(queries, FirstConfigPath)
}

func (c *Config) UpdateSecondPartConfig() error {
	queries := url.Values{}
	queries.Add("mdid", string(c.megadID))

	return c.updateConfig(queries, SecondConfigPath)
}

func (c *Config) HasID() bool {
	return len(c.megadID) > 0
}

func (c *Config) IsEnabledMqtt() bool {
	return c.srvType == MQTT
}

/**
 * Private Area
**/

// updateConfig обновляет конифгурицию по указанному пути.
func (c *Config) updateConfig(queries url.Values, configPath configPathType) error {
	queries.Add("cf", string(configPath))

	httpClient := http.DefaultClient
	fullHostURI := fmt.Sprintf("http://%v/%v/?%v", c.hostIP, c.pwd, queries.Encode())

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fullHostURI, nil)
	if err != nil {
		return ErrUpdateConfig
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return ErrUpdateConfig
	}

	defer resp.Body.Close()

	return nil
}

// syncFirstConfig worked with cf=1.
func (c *Config) syncFirstConfig() {
	c.syncConfig(FirstConfigPath, func(r *client.Response) {
		gw, ok := r.HTMLDoc.Find("form input[name=gw]").Attr("value")
		if ok {
			c.gw = GW(gw)
		}

		srv, ok := r.HTMLDoc.Find("form input[name=sip]").Attr("value")
		if ok {
			c.srv = SRV(srv)

			srvTypeElem := r.HTMLDoc.Find("form select[name=srvt]")
			srvTypeElem.ChildrenFiltered("option").Each(func(i int, s *goquery.Selection) {
				_, ok = s.Attr("selected")

				if ok {
					c.srvType = SrvType(i)
				}
			})
		}

		mqttPassword, ok := r.HTMLDoc.Find("form input[name=auth]").Attr("value")

		if ok {
			c.mqttPassword = MQTTPwd(mqttPassword)
		}
	})
}

// syncSecondConfig worked sync with /?cf=2.
func (c *Config) syncSecondConfig() {
	c.syncConfig(SecondConfigPath, func(r *client.Response) {
		megadID, ok := r.HTMLDoc.Find("form input[name=mdid]").Attr("value")

		if ok {
			c.megadID = MegadID(megadID)
		}
	})
}

func (c *Config) syncConfig(configNum configPathType, cb func(r *client.Response)) {
	queries := url.Values{
		"cf": []string{string(configNum)},
	}

	fullHostURI := fmt.Sprintf("http://%v/%v/?%v", c.hostIP, c.pwd, queries.Encode())

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{fullHostURI},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			cb(r)
		},
	}).Start()
}
