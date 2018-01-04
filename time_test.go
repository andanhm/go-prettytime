package prettytime

import (
	"testing"
	"time"
)

func TestTimeAgo(t *testing.T) {
	var now = time.Now()

	t.Log(now.Format(time.RFC3339), "->", Format(now))

	Time := now.AddDate(-2, 0, 0)
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))

	Time = now.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))

	Time = now.AddDate(50, 0, 0)
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))
}
