package appm

import "testing"

func TestHMSToSeconds(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"00:00:00", 0},
		{"00:00:01", 1},
		{"00:01:00", 60},
		{"01:00:00", 3600},
	}

	for _, c := range cases {
		got, _ := HMSToSeconds(c.in)
		if got != c.want {
			t.Errorf("ParseHMS(%q) == $%q, want %q", c.in, got, c.want)
		}
	}
}

func TestSecondsToHMS(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "00:00:00"},
		{3662, "01:01:02"},
		{360000, "100:00:00"},
	}

	for _, c := range cases {
		got := SecondsToHMS(c.in)
		if got != c.want {
			t.Errorf("ParseHMS(%i) == %s, want %s", c.in, got, c.want)
		}
	}
}
