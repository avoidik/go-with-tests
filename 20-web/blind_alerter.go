package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(ScheduledAlert)
}

type BlindAlerterFunc func(ScheduledAlert)

func (a BlindAlerterFunc) ScheduleAlertAt(alert ScheduledAlert) {
	a(alert)
}

func StdOutAlerter(alert ScheduledAlert) {
	time.AfterFunc(alert.At, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", alert.Amount)
	})
}
