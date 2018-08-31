package timebot

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kamaln7/timebot/config"
	"github.com/kamaln7/timebot/munge"

	"github.com/ejholmes/slash"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
)

type Timebot struct {
	Config *Config
	server *slash.Server
}

type Config struct {
	Host      string
	InChannel bool
	Timezones map[string]string
}

func readConfig() (*Config, error) {
	conf, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("couldn't read config: %v", err)
	}

	c := &Config{
		Host:      conf.Host,
		Timezones: conf.Timezones,
		InChannel: conf.InChannel,
	}

	return c, nil
}

func New() (*Timebot, error) {
	conf, err := readConfig()
	if err != nil {
		return nil, err
	}

	bot := &Timebot{
		Config: conf,
	}

	bot.Init()

	return bot, nil
}

func (t *Timebot) Init() {
	h := slash.HandlerFunc(t.handle)
	t.server = slash.NewServer(h)
}

func (t *Timebot) handle(ctx context.Context, r slash.Responder, command slash.Command) error {
	var (
		responseLines []string
		err           error
	)

	conf, err := readConfig()
	if err != nil {
		responseLines = append(responseLines, fmt.Sprintf("could not read config: %v", err))
		if t.Config != nil {
			responseLines = append(responseLines, "using latest valid config")
		}
	} else {
		t.Config = conf
	}

	if t.Config != nil {
		times, err := t.getTimes()
		if err != nil {
			responseLines = append(responseLines, err.Error())
		} else {
			responseLines = append(responseLines, times)
		}
	}

	err = r.Respond(slash.Response{
		InChannel: t.Config.InChannel,
		Text:      strings.Join(responseLines, "\n"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *Timebot) getTimes() (string, error) {
	tableOutput := new(bytes.Buffer)

	table := tablewriter.NewWriter(tableOutput)
	table.SetHeader([]string{"Name", "Time", "Timezone", "UTC Offset"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	currentTime := time.Now()
	for user, timezone := range t.Config.Timezones {
		localTime, err := time.LoadLocation(timezone)
		if err != nil {
			return "", err
		}
		userTime := currentTime.In(localTime)
		if t.Config.InChannel {
			user = munge.Munge(user)
		}
		table.Append([]string{user, userTime.Format("03:04 pm"), userTime.Format("MST"), userTime.Format("UTC-07")})
	}

	table.Render()

	return fmt.Sprintf("```\n%s```", tableOutput.String()), nil
}

func (t *Timebot) Listen() {
	log.Printf("timebot listening on %s\n", t.Config.Host)
	http.ListenAndServe(t.Config.Host, t.server)
}
