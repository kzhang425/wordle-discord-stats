package store

import "fmt"

type WordleResult struct {
	GuildID   string
	UserID    string // Discord snowflake; empty if FixedNick is set
	FixedNick string // non-Discord player name; empty if UserID is set
	MessageID string
	Day       int
	Score     int
	Complete  bool
}

// PlayerKey returns the canonical scoring key for a result.
// FixedNick is already a display name; UserID requires nickcache resolution.
func PlayerKey(r WordleResult) string {
	if r.FixedNick != "" {
		return r.FixedNick
	}
	return r.UserID
}

type StatsResult struct {
	AvgScore float64
	Rank     int
}

type TopEntry struct {
	PlayerKey string // either a Discord snowflake or a fixed nick
	AvgScore  float64
}

type Store interface {
	Save(result WordleResult) (bool, error)
	QueryStats(playerKey string, sinceDay int) (StatsResult, error)
	QueryTop(k int, sinceDay int) ([]TopEntry, error)
}

// FormatTop formats a leaderboard. resolve maps a PlayerKey to a display name:
// pass nickcache.Get, which handles both snowflakes and fixed nicks transparently.
func FormatTop(entries []TopEntry, resolve func(string) string) string {
	msg := ""
	for i, e := range entries {
		msg += fmt.Sprintf("%d. %s — %.2f\n", i+1, resolve(e.PlayerKey), e.AvgScore)
	}
	return msg
}
