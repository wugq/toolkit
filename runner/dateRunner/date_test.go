package dateRunner

import (
	"testing"
	"time"
)

func TestDiff_Positive(t *testing.T) {
	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 3, 2, 30, 15, 0, time.UTC)
	d := Diff(t1, t2)
	if d.Days != 2 || d.Hours != 2 || d.Minutes != 30 || d.Seconds != 15 {
		t.Errorf("unexpected diff: %+v", d)
	}
}

func TestDiff_Absolute(t *testing.T) {
	t1 := time.Date(2024, 1, 3, 2, 30, 15, 0, time.UTC)
	t2 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	d := Diff(t1, t2)
	if d.Days != 2 || d.Hours != 2 || d.Minutes != 30 || d.Seconds != 15 {
		t.Errorf("diff should be absolute, got: %+v", d)
	}
}

func TestDiff_Zero(t *testing.T) {
	t1 := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	d := Diff(t1, t1)
	if d.Days != 0 || d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0 {
		t.Errorf("expected zero diff, got: %+v", d)
	}
}

func TestParseDateText(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
		year    int
		month   time.Month
		day     int
	}{
		{"20240101", false, 2024, time.January, 1},
		{"20231215", false, 2023, time.December, 15},
		{"20240101120000", false, 2024, time.January, 1},
		{"2024", true, 0, 0, 0},
		{"invalid!", true, 0, 0, 0},
		{"", true, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseDateText(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseDateText(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr {
				if got.Year() != tt.year || got.Month() != tt.month || got.Day() != tt.day {
					t.Errorf("got %v, want %d-%s-%02d", got, tt.year, tt.month, tt.day)
				}
			}
		})
	}
}

func TestParseDateText_WithTime(t *testing.T) {
	got, err := ParseDateText("20240315143045")
	if err != nil {
		t.Fatal(err)
	}
	if got.Year() != 2024 || got.Month() != time.March || got.Day() != 15 {
		t.Errorf("unexpected date part: %v", got)
	}
	if got.Hour() != 14 || got.Minute() != 30 || got.Second() != 45 {
		t.Errorf("unexpected time part: %v", got)
	}
}

func TestAdd_YearMonthDay(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	got, err := Add(base, AddDate{Year: 1, Month: 2, Day: 3})
	if err != nil {
		t.Fatal(err)
	}
	want := time.Date(2025, 3, 4, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAdd_HourMinuteSecond(t *testing.T) {
	base := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	got, err := Add(base, AddDate{Hour: 4, Minute: 5, Second: 6})
	if err != nil {
		t.Fatal(err)
	}
	want := time.Date(2024, 6, 1, 4, 5, 6, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAdd_Zero(t *testing.T) {
	base := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	got, err := Add(base, AddDate{})
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(base) {
		t.Errorf("zero add should return same time, got %v", got)
	}
}

func TestAdd_Negative(t *testing.T) {
	base := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	got, err := Add(base, AddDate{Day: -5, Hour: -2})
	if err != nil {
		t.Fatal(err)
	}
	want := time.Date(2024, 6, 10, 10, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
