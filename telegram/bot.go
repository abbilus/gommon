package telegram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Bot struct {
	url     string
	apiKey  string
	chatID  string
	verbose bool
}

func (bot *Bot) addQuery(urlv *string, queryv map[string]string) string {
	values := url.Values{}
	for key, value := range queryv {
		values.Add(key, value)
	}
	var query string
	query = values.Encode()
	if query != "" {
		return *urlv + "?" + query
	} else {
		return *urlv
	}
}

func (bot *Bot) preReq() (*http.Request, error) {
	req, err := http.NewRequest("", "", nil)
	req.Header.Add("user-agent", "etabot 1.0")
	return req, err
}

func (bot *Bot) get(method string, params ...map[string]string) (respStatusCode int, respBody []byte, err error) {
	req, _ := bot.preReq()
	err = nil
	var queryParams map[string]string
	if len(params) > 0 {
		queryParams = params[0]
	}
	uri := fmt.Sprintf("%s/bot%s/%s", bot.url, bot.apiKey, method)
	uri = bot.addQuery(&uri, queryParams)
	url, _ := url.Parse(uri)
	req.URL = url
	req.Method = "GET"
	if err != nil {
		errors.WithStack(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respStatusCode = resp.StatusCode
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//New - prepare struct, params.apiKey, params.chatid posible values
func New(apiKey string, chatid string, verbose bool) (telegramBot Bot) {
	if apiKey != "" {
		telegramBot.apiKey = apiKey
	} else {
		panic(fmt.Errorf("Bot api key  is required"))
	}
	if chatid != "" {
		telegramBot.chatID = chatid
	}
	if verbose {
		telegramBot.verbose = verbose
	}
	telegramBot.url = "https://api.telegram.org"
	return
}

func (bot *Bot) SetURL(url string) {
	bot.url = url
}

func (bot *Bot) SetVerbose(verbose bool) {
	bot.verbose = verbose
}

func (bot Bot) SendToChat(message string, chatid string) (err error) {
	method := "sendMessage"
	params := map[string]string{
		"text": message,
	}
	if chatid != "" {
		params["chat_id"] = chatid
	} else {
		err = fmt.Errorf("chatid not set")
		if bot.verbose {
			err = errors.WithStack(err)
			return
		}
	}
	var status int
	var respBody []byte
	if err != nil {
		return
	}
	status, respBody, err = bot.get(method, params)
	if err != nil {
		if bot.verbose {
			err = errors.WithStack(err)
			fmt.Printf("Error: %+v", err)
		}
		return
	}
	if status != 200 {
		if bot.verbose {
			err = errors.WithStack(fmt.Errorf("Bad telegram status code: %d\n%s", status, string(respBody)))
			return
		}
	}
	return
}

func (bot Bot) Send(message string) (err error) {
	method := "sendMessage"
	params := map[string]string{
		"text": message,
	}
	if bot.chatID != "" {
		params["chat_id"] = bot.chatID
	} else {
		err = fmt.Errorf("chatid not set")
		if bot.verbose {
			err = errors.WithStack(err)
			return
		}
	}
	var status int
	var respBody []byte
	if err != nil {
		return
	}
	status, respBody, err = bot.get(method, params)
	if err != nil {
		if bot.verbose {
			err = errors.WithStack(err)
			fmt.Printf("Error: %+v", err)
		}
		return
	}
	if status != 200 {
		if bot.verbose {
			err = errors.WithStack(fmt.Errorf("Bad telegram status code: %d\n%s", status, string(respBody)))
			return
		}
	}
	return
}
