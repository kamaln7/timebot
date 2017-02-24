package timebot

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

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

func New(config *Config) *Timebot {
	bot := &Timebot{
		Config: config,
	}

	bot.Init()

	return bot
}

func (t *Timebot) Init() {
	h := slash.HandlerFunc(t.handle)
	t.server = slash.NewServer(h)
}

func (t *Timebot) handle(ctx context.Context, r slash.Responder, command slash.Command) error {
	times, err := t.getTimes()
	if err != nil {
		return err
	}

	if err := r.Respond(slash.Response{
		InChannel: t.Config.InChannel,
		Text:      times,
	}); err != nil {
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
	http.ListenAndServe(t.Config.Host, t.server)
}
