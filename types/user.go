package types

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//User Contains Discord user information as well as warning and ban information.
type User struct {
	Member     discordgo.Member
	Warnings   int
	Banned     bool
	BannedDate time.Time
}

//Ban function calls the BanCreate api for discord and sets the warnning
//and banned information appropriatly
func (u *User) Ban(dg *discordgo.Session) (success bool, err error) {
	u.Banned = true
	u.BannedDate = time.Now()
	u.Warnings = 0
	err = dg.GuildBanCreate(u.Member.GuildID, u.Member.User.ID, 0)
	if err == nil {
		return true, nil
	}
	return false, nil
}

//UnBan function calls the BanDelete api for discord and sets the warning and
//banned information appropriatly
func (u *User) UnBan(dg *discordgo.Session) {
	u.Banned = false
	u.Warnings = 0
	dg.GuildBanCreate(u.Member.GuildID, u.Member.User.ID, 0)
}

//AddWarning Increments the number of warnings that the usere has recieved.
func (u *User) AddWarning() (warningCount int) {
	u.Warnings++
	return u.Warnings
}
