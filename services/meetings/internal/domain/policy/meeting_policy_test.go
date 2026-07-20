package policy_test

import (
	"testing"

	"prahari/services/meetings/internal/domain/attendance"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/policy"
)

func TestIsQuorumMet_SafetyCommittee(t *testing.T) {
	mtg := &meeting.Meeting{MeetingType: "SAFETY_COMMITTEE"}

	twoRecords := []*attendance.Record{{}, {}}
	if policy.IsQuorumMet(mtg, twoRecords) {
		t.Error("expected quorum NOT met with 2 attendees for SAFETY_COMMITTEE (needs 3)")
	}

	threeRecords := []*attendance.Record{{}, {}, {}}
	if !policy.IsQuorumMet(mtg, threeRecords) {
		t.Error("expected quorum met with 3 attendees for SAFETY_COMMITTEE")
	}
}

func TestIsQuorumMet_ToolboxTalk(t *testing.T) {
	mtg := &meeting.Meeting{MeetingType: "TOOLBOX_TALK"}

	oneRecord := []*attendance.Record{{}}
	if policy.IsQuorumMet(mtg, oneRecord) {
		t.Error("expected quorum NOT met with 1 attendee for TOOLBOX_TALK (needs 2)")
	}

	twoRecords := []*attendance.Record{{}, {}}
	if !policy.IsQuorumMet(mtg, twoRecords) {
		t.Error("expected quorum met with 2 attendees for TOOLBOX_TALK")
	}
}

func TestCanCloseWithoutMinutes(t *testing.T) {
	toolbox := &meeting.Meeting{MeetingType: "TOOLBOX_TALK"}
	if !policy.CanCloseWithoutMinutes(toolbox) {
		t.Error("expected TOOLBOX_TALK to close without formal minutes")
	}

	shiftBrief := &meeting.Meeting{MeetingType: "SHIFT_BRIEFING"}
	if !policy.CanCloseWithoutMinutes(shiftBrief) {
		t.Error("expected SHIFT_BRIEFING to close without formal minutes")
	}

	committee := &meeting.Meeting{MeetingType: "SAFETY_COMMITTEE"}
	if policy.CanCloseWithoutMinutes(committee) {
		t.Error("expected SAFETY_COMMITTEE to require formal minutes")
	}
}

func TestIsMeetingOverdue(t *testing.T) {
	scheduled := &meeting.Meeting{Status: "SCHEDULED", StartedAt: nil}
	if !policy.IsMeetingOverdue(scheduled) {
		t.Error("expected scheduled meeting without start to be overdue")
	}

	closed := &meeting.Meeting{Status: "CLOSED"}
	if policy.IsMeetingOverdue(closed) {
		t.Error("expected closed meeting to not be overdue")
	}
}
