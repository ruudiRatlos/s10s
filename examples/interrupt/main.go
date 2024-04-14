package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/phsym/console-slog"
	"github.com/ruudiRatlos/s10s"
	"github.com/ruudiRatlos/s10s/api"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {
	logger := slog.New(console.NewHandler(os.Stderr, &console.HandlerOptions{Level: slog.LevelDebug}))

	rr, err := RegisterAndSave("COSMIC", randomName(), "", "registration.json")
	if err != nil {
		log.Fatal(err)
	}
	client, err := s10s.NewClient(s10s.WithToken(rr.Token))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = client.ContextWithToken(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer logger.Info("clean exit")
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			fmt.Fprintln(os.Stderr)
			logger.Warn("received CTRL-C. Exiting")
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()

	logger.Info("now interrupt the program by pressing CTRL-C or send SIGINT")
	client.Sleep(ctx, 1*time.Minute)
}

func RegisterAndSave(faction, username, email, fName string) (*api.Register201ResponseData, error) {
	f, err := os.Open(fName)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()
	out := api.NewRegister201ResponseDataWithDefaults()
	if err == nil {
		dec := json.NewDecoder(f)
		err := dec.Decode(out)
		if err == nil {
			return out, nil
		}
	}

	f, err = os.Create(fName)
	if err != nil {
		return nil, err
	}
	out, err = s10s.Register(faction, username, email)
	if err != nil {
		return nil, err
	}
	enc := json.NewEncoder(f)
	err = enc.Encode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func randomName() string {
	first := strings.Split("Taylor Teagan Trinity Tatum Tessa Talia Thea Tiana Treasure "+
		"Thalia Theodora Teresa Tiffany Tinsley Taytum Tatiana Tori Tallulah Taliyah Tru "+
		"Tilly True Tara Tenley Tinley Tegan Tabitha Tala Theresa Tyler Tess Tia Tillie "+
		"Toni Temperance Tahlia Taelynn Taliah Taya Tianna Tania Tamara Taelyn Tiara Tamia "+
		"Taylin Theia Teigan Teegan Theadora Toby Taryn Tahani Tina Taylee Timber Terra "+
		"Taylynn Tayla Tasneem Thyri Tracy Theo Tamar Teyana Truth Talayah Taniyah Twyla "+
		"Tylee Truly Tziporah Tanvi Taraji Tova Talitha Tanner Teddi Therese Talulah Tyla "+
		"Torvi Tyra Tali Talya Tenzin Triniti Taleah Tehila Tirzah Tzipora Tailynn Tommie "+
		"Trisha Tanya Taylen Taylyn Taleen Tinslee Treazure Theodore Thomas Thiago Theo "+
		"Tyler Tucker Timothy Tristan Tobias Tate Tyson Tanner Titus Travis Tatum Troy "+
		"Tripp Trevor Tadeo Taylor Trey Trenton Tru Tony Ty Tommy Trace Titan Tomas Talon "+
		"Thatcher Thaddeus Truett Turner Terry Trent Tristen Terrance Toby Teo Terrell "+
		"Torin Teddy Taj Tiago Terrence Taylen Truth Tylan Tayden Tyrone Tariq Thorin "+
		"Tristian Truman Tyree Teagan Taysom Townes Todd Tyrell Trevon Tyce Thor Tevin "+
		"Tamir Triston Taylin Tobin Tahir Tiberius Tylen Treyson Tymir Tegan Tytus "+
		"Theodor Tayson Truitt Taha Tom Trae Thompson Tye Tyrus Tyshawn Tai Theron "+
		"Tayvion Theoden Tristin Tyrese Tenzin Treyton Terence Travon Tylin", " ")
	name := fmt.Sprintf("%s-%s", strings.ToUpper(first[rand.Intn(len(first))]), "TESTER")
	return name
}
