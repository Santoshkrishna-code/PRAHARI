package policy

import (
	"prahari/services/meetings/internal/domain/attendance"
	"prahari/services/meetings/internal/domain/meeting"
)

// MinToolboxTalkDurationMin is the minimum duration for a valid toolbox talk.
const MinToolboxTalkDurationMin = 5

// IsQuorumMet checks if the attendance count meets the minimum quorum for the meeting type.
func IsQuorumMet(mtg *meeting.Meeting, records []*attendance.Record) bool {
	if mtg == nil {
		return false
	}
	count := len(records)
	switch mtg.MeetingType {
	case "SAFETY_COMMITTEE":
		return count >= 3
	case "TOOLBOX_TALK":
		return count >= 2
	case "SHIFT_BRIEFING":
		return count >= 1
	default:
		return count >= 2
	}
}

// CanCloseWithoutMinutes checks if meeting type allows closure without formal minutes.
func CanCloseWithoutMinutes(mtg *meeting.Meeting) bool {
	if mtg == nil {
		return false
	}
	return mtg.MeetingType == "TOOLBOX_TALK" || mtg.MeetingType == "SHIFT_BRIEFING"
}

// IsMeetingOverdue checks if a scheduled meeting was never started.
func IsMeetingOverdue(mtg *meeting.Meeting) bool {
	if mtg == nil {
		return false
	}
	return mtg.Status == "SCHEDULED" && mtg.StartedAt == nil
}
