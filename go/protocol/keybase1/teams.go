// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/teams.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TeamRole int

const (
	TeamRole_NONE   TeamRole = 0
	TeamRole_READER TeamRole = 1
	TeamRole_WRITER TeamRole = 2
	TeamRole_ADMIN  TeamRole = 3
	TeamRole_OWNER  TeamRole = 4
)

func (o TeamRole) DeepCopy() TeamRole { return o }

var TeamRoleMap = map[string]TeamRole{
	"NONE":   0,
	"READER": 1,
	"WRITER": 2,
	"ADMIN":  3,
	"OWNER":  4,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "READER",
	2: "WRITER",
	3: "ADMIN",
	4: "OWNER",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return ""
}

type TeamApplication int

const (
	TeamApplication_KBFS     TeamApplication = 1
	TeamApplication_CHAT     TeamApplication = 2
	TeamApplication_SALTPACK TeamApplication = 3
)

func (o TeamApplication) DeepCopy() TeamApplication { return o }

var TeamApplicationMap = map[string]TeamApplication{
	"KBFS":     1,
	"CHAT":     2,
	"SALTPACK": 3,
}

var TeamApplicationRevMap = map[TeamApplication]string{
	1: "KBFS",
	2: "CHAT",
	3: "SALTPACK",
}

func (e TeamApplication) String() string {
	if v, ok := TeamApplicationRevMap[e]; ok {
		return v
	}
	return ""
}

type PerTeamKeyGeneration int

func (o PerTeamKeyGeneration) DeepCopy() PerTeamKeyGeneration {
	return o
}

type TeamApplicationKey struct {
	Application   TeamApplication      `codec:"application" json:"application"`
	KeyGeneration PerTeamKeyGeneration `codec:"keyGeneration" json:"keyGeneration"`
	Key           Bytes32              `codec:"key" json:"key"`
}

func (o TeamApplicationKey) DeepCopy() TeamApplicationKey {
	return TeamApplicationKey{
		Application:   o.Application.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		Key:           o.Key.DeepCopy(),
	}
}

type MaskB64 []byte

func (o MaskB64) DeepCopy() MaskB64 {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte(nil), x...)
	})(o)
}

type TeamInviteID string

func (o TeamInviteID) DeepCopy() TeamInviteID {
	return o
}

type ReaderKeyMask struct {
	Application TeamApplication      `codec:"application" json:"application"`
	Generation  PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Mask        MaskB64              `codec:"mask" json:"mask"`
}

func (o ReaderKeyMask) DeepCopy() ReaderKeyMask {
	return ReaderKeyMask{
		Application: o.Application.DeepCopy(),
		Generation:  o.Generation.DeepCopy(),
		Mask:        o.Mask.DeepCopy(),
	}
}

type PerTeamKey struct {
	Gen    PerTeamKeyGeneration `codec:"gen" json:"gen"`
	Seqno  Seqno                `codec:"seqno" json:"seqno"`
	SigKID KID                  `codec:"sigKID" json:"sigKID"`
	EncKID KID                  `codec:"encKID" json:"encKID"`
}

func (o PerTeamKey) DeepCopy() PerTeamKey {
	return PerTeamKey{
		Gen:    o.Gen.DeepCopy(),
		Seqno:  o.Seqno.DeepCopy(),
		SigKID: o.SigKID.DeepCopy(),
		EncKID: o.EncKID.DeepCopy(),
	}
}

type PerTeamKeySeed [32]byte

func (o PerTeamKeySeed) DeepCopy() PerTeamKeySeed {
	var ret PerTeamKeySeed
	copy(ret[:], o[:])
	return ret
}

type PerTeamKeySeedItem struct {
	Seed       PerTeamKeySeed       `codec:"seed" json:"seed"`
	Generation PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Seqno      Seqno                `codec:"seqno" json:"seqno"`
}

func (o PerTeamKeySeedItem) DeepCopy() PerTeamKeySeedItem {
	return PerTeamKeySeedItem{
		Seed:       o.Seed.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Seqno:      o.Seqno.DeepCopy(),
	}
}

type TeamMember struct {
	Uid         UID      `codec:"uid" json:"uid"`
	Role        TeamRole `codec:"role" json:"role"`
	EldestSeqno Seqno    `codec:"eldestSeqno" json:"eldestSeqno"`
}

func (o TeamMember) DeepCopy() TeamMember {
	return TeamMember{
		Uid:         o.Uid.DeepCopy(),
		Role:        o.Role.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
	}
}

type TeamMembers struct {
	Owners  []UserVersion `codec:"owners" json:"owners"`
	Admins  []UserVersion `codec:"admins" json:"admins"`
	Writers []UserVersion `codec:"writers" json:"writers"`
	Readers []UserVersion `codec:"readers" json:"readers"`
}

func (o TeamMembers) DeepCopy() TeamMembers {
	return TeamMembers{
		Owners: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
	}
}

type TeamMemberDetails struct {
	Uv       UserVersion `codec:"uv" json:"uv"`
	Username string      `codec:"username" json:"username"`
	Active   bool        `codec:"active" json:"active"`
}

func (o TeamMemberDetails) DeepCopy() TeamMemberDetails {
	return TeamMemberDetails{
		Uv:       o.Uv.DeepCopy(),
		Username: o.Username,
		Active:   o.Active,
	}
}

type TeamMembersDetails struct {
	Owners  []TeamMemberDetails `codec:"owners" json:"owners"`
	Admins  []TeamMemberDetails `codec:"admins" json:"admins"`
	Writers []TeamMemberDetails `codec:"writers" json:"writers"`
	Readers []TeamMemberDetails `codec:"readers" json:"readers"`
}

func (o TeamMembersDetails) DeepCopy() TeamMembersDetails {
	return TeamMembersDetails{
		Owners: (func(x []TeamMemberDetails) []TeamMemberDetails {
			var ret []TeamMemberDetails
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []TeamMemberDetails) []TeamMemberDetails {
			var ret []TeamMemberDetails
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			var ret []TeamMemberDetails
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			var ret []TeamMemberDetails
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
	}
}

type TeamDetails struct {
	Members       TeamMembersDetails   `codec:"members" json:"members"`
	KeyGeneration PerTeamKeyGeneration `codec:"keyGeneration" json:"keyGeneration"`
}

func (o TeamDetails) DeepCopy() TeamDetails {
	return TeamDetails{
		Members:       o.Members.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
	}
}

type TeamChangeReq struct {
	Owners  []UserVersion `codec:"owners" json:"owners"`
	Admins  []UserVersion `codec:"admins" json:"admins"`
	Writers []UserVersion `codec:"writers" json:"writers"`
	Readers []UserVersion `codec:"readers" json:"readers"`
	None    []UserVersion `codec:"none" json:"none"`
}

func (o TeamChangeReq) DeepCopy() TeamChangeReq {
	return TeamChangeReq{
		Owners: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
		None: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.None),
	}
}

type UserVersion struct {
	Uid         UID   `codec:"uid" json:"uid"`
	EldestSeqno Seqno `codec:"eldestSeqno" json:"eldestSeqno"`
}

func (o UserVersion) DeepCopy() UserVersion {
	return UserVersion{
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
	}
}

type TeamPlusApplicationKeys struct {
	Id              TeamID               `codec:"id" json:"id"`
	Name            string               `codec:"name" json:"name"`
	Application     TeamApplication      `codec:"application" json:"application"`
	Writers         []UserVersion        `codec:"writers" json:"writers"`
	OnlyReaders     []UserVersion        `codec:"onlyReaders" json:"onlyReaders"`
	ApplicationKeys []TeamApplicationKey `codec:"applicationKeys" json:"applicationKeys"`
}

func (o TeamPlusApplicationKeys) DeepCopy() TeamPlusApplicationKeys {
	return TeamPlusApplicationKeys{
		Id:          o.Id.DeepCopy(),
		Name:        o.Name,
		Application: o.Application.DeepCopy(),
		Writers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		OnlyReaders: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.OnlyReaders),
		ApplicationKeys: (func(x []TeamApplicationKey) []TeamApplicationKey {
			var ret []TeamApplicationKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ApplicationKeys),
	}
}

type TeamData struct {
	Secretless      bool                                                 `codec:"secretless" json:"secretless"`
	Chain           TeamSigChainState                                    `codec:"chain" json:"chain"`
	PerTeamKeySeeds map[PerTeamKeyGeneration]PerTeamKeySeedItem          `codec:"perTeamKeySeeds" json:"perTeamKeySeeds"`
	ReaderKeyMasks  map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 `codec:"readerKeyMasks" json:"readerKeyMasks"`
	CachedAt        Time                                                 `codec:"cachedAt" json:"cachedAt"`
}

func (o TeamData) DeepCopy() TeamData {
	return TeamData{
		Secretless: o.Secretless,
		Chain:      o.Chain.DeepCopy(),
		PerTeamKeySeeds: (func(x map[PerTeamKeyGeneration]PerTeamKeySeedItem) map[PerTeamKeyGeneration]PerTeamKeySeedItem {
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeedItem)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeeds),
		ReaderKeyMasks: (func(x map[TeamApplication]map[PerTeamKeyGeneration]MaskB64) map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 {
			ret := make(map[TeamApplication]map[PerTeamKeyGeneration]MaskB64)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x map[PerTeamKeyGeneration]MaskB64) map[PerTeamKeyGeneration]MaskB64 {
					ret := make(map[PerTeamKeyGeneration]MaskB64)
					for k, v := range x {
						kCopy := k.DeepCopy()
						vCopy := v.DeepCopy()
						ret[kCopy] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderKeyMasks),
		CachedAt: o.CachedAt.DeepCopy(),
	}
}

type TeamInviteCategory int

const (
	TeamInviteCategory_NONE    TeamInviteCategory = 0
	TeamInviteCategory_UNKNOWN TeamInviteCategory = 1
	TeamInviteCategory_KEYBASE TeamInviteCategory = 2
	TeamInviteCategory_EMAIL   TeamInviteCategory = 3
	TeamInviteCategory_SBS     TeamInviteCategory = 4
)

func (o TeamInviteCategory) DeepCopy() TeamInviteCategory { return o }

var TeamInviteCategoryMap = map[string]TeamInviteCategory{
	"NONE":    0,
	"UNKNOWN": 1,
	"KEYBASE": 2,
	"EMAIL":   3,
	"SBS":     4,
}

var TeamInviteCategoryRevMap = map[TeamInviteCategory]string{
	0: "NONE",
	1: "UNKNOWN",
	2: "KEYBASE",
	3: "EMAIL",
	4: "SBS",
}

func (e TeamInviteCategory) String() string {
	if v, ok := TeamInviteCategoryRevMap[e]; ok {
		return v
	}
	return ""
}

type TeamInviteType struct {
	C__       TeamInviteCategory       `codec:"c" json:"c"`
	Unknown__ *string                  `codec:"unknown,omitempty" json:"unknown,omitempty"`
	Sbs__     *TeamInviteSocialNetwork `codec:"sbs,omitempty" json:"sbs,omitempty"`
}

func (o *TeamInviteType) C() (ret TeamInviteCategory, err error) {
	switch o.C__ {
	case TeamInviteCategory_UNKNOWN:
		if o.Unknown__ == nil {
			err = errors.New("unexpected nil value for Unknown__")
			return ret, err
		}
	case TeamInviteCategory_SBS:
		if o.Sbs__ == nil {
			err = errors.New("unexpected nil value for Sbs__")
			return ret, err
		}
	}
	return o.C__, nil
}

func (o TeamInviteType) Unknown() (res string) {
	if o.C__ != TeamInviteCategory_UNKNOWN {
		panic("wrong case accessed")
	}
	if o.Unknown__ == nil {
		return
	}
	return *o.Unknown__
}

func (o TeamInviteType) Sbs() (res TeamInviteSocialNetwork) {
	if o.C__ != TeamInviteCategory_SBS {
		panic("wrong case accessed")
	}
	if o.Sbs__ == nil {
		return
	}
	return *o.Sbs__
}

func NewTeamInviteTypeWithUnknown(v string) TeamInviteType {
	return TeamInviteType{
		C__:       TeamInviteCategory_UNKNOWN,
		Unknown__: &v,
	}
}

func NewTeamInviteTypeWithSbs(v TeamInviteSocialNetwork) TeamInviteType {
	return TeamInviteType{
		C__:   TeamInviteCategory_SBS,
		Sbs__: &v,
	}
}

func NewTeamInviteTypeDefault(c TeamInviteCategory) TeamInviteType {
	return TeamInviteType{
		C__: c,
	}
}

func (o TeamInviteType) DeepCopy() TeamInviteType {
	return TeamInviteType{
		C__: o.C__.DeepCopy(),
		Unknown__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Unknown__),
		Sbs__: (func(x *TeamInviteSocialNetwork) *TeamInviteSocialNetwork {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Sbs__),
	}
}

type TeamInviteSocialNetwork string

func (o TeamInviteSocialNetwork) DeepCopy() TeamInviteSocialNetwork {
	return o
}

type TeamInviteName string

func (o TeamInviteName) DeepCopy() TeamInviteName {
	return o
}

type TeamInvite struct {
	Role TeamRole       `codec:"role" json:"role"`
	Id   TeamInviteID   `codec:"id" json:"id"`
	Type TeamInviteType `codec:"type" json:"type"`
	Name TeamInviteName `codec:"name" json:"name"`
}

func (o TeamInvite) DeepCopy() TeamInvite {
	return TeamInvite{
		Role: o.Role.DeepCopy(),
		Id:   o.Id.DeepCopy(),
		Type: o.Type.DeepCopy(),
		Name: o.Name.DeepCopy(),
	}
}

type TeamSigChainState struct {
	Reader        UserVersion                         `codec:"reader" json:"reader"`
	Id            TeamID                              `codec:"id" json:"id"`
	Name          TeamName                            `codec:"name" json:"name"`
	LastSeqno     Seqno                               `codec:"lastSeqno" json:"lastSeqno"`
	LastLinkID    LinkID                              `codec:"lastLinkID" json:"lastLinkID"`
	ParentID      *TeamID                             `codec:"parentID,omitempty" json:"parentID,omitempty"`
	UserLog       map[UserVersion][]UserLogPoint      `codec:"userLog" json:"userLog"`
	SubteamLog    map[TeamID][]SubteamLogPoint        `codec:"subteamLog" json:"subteamLog"`
	PerTeamKeys   map[PerTeamKeyGeneration]PerTeamKey `codec:"perTeamKeys" json:"perTeamKeys"`
	LinkIDs       map[Seqno]LinkID                    `codec:"linkIDs" json:"linkIDs"`
	StubbedLinks  map[Seqno]bool                      `codec:"stubbedLinks" json:"stubbedLinks"`
	ActiveInvites map[TeamInviteID]TeamInvite         `codec:"activeInvites" json:"activeInvites"`
}

func (o TeamSigChainState) DeepCopy() TeamSigChainState {
	return TeamSigChainState{
		Reader:     o.Reader.DeepCopy(),
		Id:         o.Id.DeepCopy(),
		Name:       o.Name.DeepCopy(),
		LastSeqno:  o.LastSeqno.DeepCopy(),
		LastLinkID: o.LastLinkID.DeepCopy(),
		ParentID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ParentID),
		UserLog: (func(x map[UserVersion][]UserLogPoint) map[UserVersion][]UserLogPoint {
			ret := make(map[UserVersion][]UserLogPoint)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []UserLogPoint) []UserLogPoint {
					var ret []UserLogPoint
					for _, v := range x {
						vCopy := v.DeepCopy()
						ret = append(ret, vCopy)
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.UserLog),
		SubteamLog: (func(x map[TeamID][]SubteamLogPoint) map[TeamID][]SubteamLogPoint {
			ret := make(map[TeamID][]SubteamLogPoint)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []SubteamLogPoint) []SubteamLogPoint {
					var ret []SubteamLogPoint
					for _, v := range x {
						vCopy := v.DeepCopy()
						ret = append(ret, vCopy)
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.SubteamLog),
		PerTeamKeys: (func(x map[PerTeamKeyGeneration]PerTeamKey) map[PerTeamKeyGeneration]PerTeamKey {
			ret := make(map[PerTeamKeyGeneration]PerTeamKey)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		LinkIDs: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			ret := make(map[Seqno]LinkID)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkIDs),
		StubbedLinks: (func(x map[Seqno]bool) map[Seqno]bool {
			ret := make(map[Seqno]bool)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.StubbedLinks),
		ActiveInvites: (func(x map[TeamInviteID]TeamInvite) map[TeamInviteID]TeamInvite {
			ret := make(map[TeamInviteID]TeamInvite)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ActiveInvites),
	}
}

type UserLogPoint struct {
	Role    TeamRole          `codec:"role" json:"role"`
	SigMeta SignatureMetadata `codec:"sigMeta" json:"sigMeta"`
}

func (o UserLogPoint) DeepCopy() UserLogPoint {
	return UserLogPoint{
		Role:    o.Role.DeepCopy(),
		SigMeta: o.SigMeta.DeepCopy(),
	}
}

type SubteamLogPoint struct {
	Name  TeamName `codec:"name" json:"name"`
	Seqno Seqno    `codec:"seqno" json:"seqno"`
}

func (o SubteamLogPoint) DeepCopy() SubteamLogPoint {
	return SubteamLogPoint{
		Name:  o.Name.DeepCopy(),
		Seqno: o.Seqno.DeepCopy(),
	}
}

type TeamNamePart string

func (o TeamNamePart) DeepCopy() TeamNamePart {
	return o
}

type TeamName struct {
	Parts []TeamNamePart `codec:"parts" json:"parts"`
}

func (o TeamName) DeepCopy() TeamName {
	return TeamName{
		Parts: (func(x []TeamNamePart) []TeamNamePart {
			var ret []TeamNamePart
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Parts),
	}
}

type TeamCLKRMsg struct {
	TeamID     TeamID               `codec:"teamID" json:"team_id"`
	Generation PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Score      int                  `codec:"score" json:"score"`
}

func (o TeamCLKRMsg) DeepCopy() TeamCLKRMsg {
	return TeamCLKRMsg{
		TeamID:     o.TeamID.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Score:      o.Score,
	}
}

type TeamChangeRow struct {
	Id                TeamID `codec:"id" json:"id"`
	Name              string `codec:"name" json:"name"`
	KeyRotated        bool   `codec:"keyRotated" json:"key_rotated"`
	MembershipChanged bool   `codec:"membershipChanged" json:"membership_changed"`
	LatestSeqno       Seqno  `codec:"latestSeqno" json:"latest_seqno"`
}

func (o TeamChangeRow) DeepCopy() TeamChangeRow {
	return TeamChangeRow{
		Id:                o.Id.DeepCopy(),
		Name:              o.Name,
		KeyRotated:        o.KeyRotated,
		MembershipChanged: o.MembershipChanged,
		LatestSeqno:       o.LatestSeqno.DeepCopy(),
	}
}

// * TeamRefreshData are needed or wanted data requirements that, if unmet, will cause
// * a refresh of the cached.
type TeamRefreshers struct {
	NeedKeyGeneration PerTeamKeyGeneration `codec:"needKeyGeneration" json:"needKeyGeneration"`
	WantMembers       []UserVersion        `codec:"wantMembers" json:"wantMembers"`
}

func (o TeamRefreshers) DeepCopy() TeamRefreshers {
	return TeamRefreshers{
		NeedKeyGeneration: o.NeedKeyGeneration.DeepCopy(),
		WantMembers: (func(x []UserVersion) []UserVersion {
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.WantMembers),
	}
}

type LoadTeamArg struct {
	ID              TeamID         `codec:"ID" json:"ID"`
	Name            string         `codec:"name" json:"name"`
	NeedAdmin       bool           `codec:"needAdmin" json:"needAdmin"`
	Refreshers      TeamRefreshers `codec:"refreshers" json:"refreshers"`
	ForceFullReload bool           `codec:"forceFullReload" json:"forceFullReload"`
	ForceRepoll     bool           `codec:"forceRepoll" json:"forceRepoll"`
	StaleOK         bool           `codec:"staleOK" json:"staleOK"`
}

func (o LoadTeamArg) DeepCopy() LoadTeamArg {
	return LoadTeamArg{
		ID:              o.ID.DeepCopy(),
		Name:            o.Name,
		NeedAdmin:       o.NeedAdmin,
		Refreshers:      o.Refreshers.DeepCopy(),
		ForceFullReload: o.ForceFullReload,
		ForceRepoll:     o.ForceRepoll,
		StaleOK:         o.StaleOK,
	}
}

type ImplicitRole struct {
	Role     TeamRole `codec:"role" json:"role"`
	Ancestor TeamID   `codec:"ancestor" json:"ancestor"`
}

func (o ImplicitRole) DeepCopy() ImplicitRole {
	return ImplicitRole{
		Role:     o.Role.DeepCopy(),
		Ancestor: o.Ancestor.DeepCopy(),
	}
}

type MemberInfo struct {
	TeamID   TeamID        `codec:"teamID" json:"team_id"`
	FqName   string        `codec:"fqName" json:"fq_name"`
	Role     TeamRole      `codec:"role" json:"role"`
	Implicit *ImplicitRole `codec:"implicit,omitempty" json:"implicit,omitempty"`
}

func (o MemberInfo) DeepCopy() MemberInfo {
	return MemberInfo{
		TeamID: o.TeamID.DeepCopy(),
		FqName: o.FqName,
		Role:   o.Role.DeepCopy(),
		Implicit: (func(x *ImplicitRole) *ImplicitRole {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Implicit),
	}
}

type TeamList struct {
	Uid      UID          `codec:"uid" json:"uid"`
	Username string       `codec:"username" json:"username"`
	FullName string       `codec:"fullName" json:"fullName"`
	Teams    []MemberInfo `codec:"teams" json:"teams"`
}

func (o TeamList) DeepCopy() TeamList {
	return TeamList{
		Uid:      o.Uid.DeepCopy(),
		Username: o.Username,
		FullName: o.FullName,
		Teams: (func(x []MemberInfo) []MemberInfo {
			var ret []MemberInfo
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Teams),
	}
}

type TeamAddMemberResult struct {
	Invited   bool  `codec:"invited" json:"invited"`
	User      *User `codec:"user,omitempty" json:"user,omitempty"`
	EmailSent bool  `codec:"emailSent" json:"emailSent"`
	ChatSent  bool  `codec:"chatSent" json:"chatSent"`
}

func (o TeamAddMemberResult) DeepCopy() TeamAddMemberResult {
	return TeamAddMemberResult{
		Invited: o.Invited,
		User: (func(x *User) *User {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.User),
		EmailSent: o.EmailSent,
		ChatSent:  o.ChatSent,
	}
}

type TeamCreateArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
}

func (o TeamCreateArg) DeepCopy() TeamCreateArg {
	return TeamCreateArg{
		SessionID: o.SessionID,
		Name:      o.Name,
	}
}

type TeamGetArg struct {
	SessionID   int    `codec:"sessionID" json:"sessionID"`
	Name        string `codec:"name" json:"name"`
	ForceRepoll bool   `codec:"forceRepoll" json:"forceRepoll"`
}

func (o TeamGetArg) DeepCopy() TeamGetArg {
	return TeamGetArg{
		SessionID:   o.SessionID,
		Name:        o.Name,
		ForceRepoll: o.ForceRepoll,
	}
}

type TeamListArg struct {
	SessionID     int    `codec:"sessionID" json:"sessionID"`
	UserAssertion string `codec:"userAssertion" json:"userAssertion"`
}

func (o TeamListArg) DeepCopy() TeamListArg {
	return TeamListArg{
		SessionID:     o.SessionID,
		UserAssertion: o.UserAssertion,
	}
}

type TeamChangeMembershipArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Name      string        `codec:"name" json:"name"`
	Req       TeamChangeReq `codec:"req" json:"req"`
}

func (o TeamChangeMembershipArg) DeepCopy() TeamChangeMembershipArg {
	return TeamChangeMembershipArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Req:       o.Req.DeepCopy(),
	}
}

type TeamAddMemberArg struct {
	SessionID            int      `codec:"sessionID" json:"sessionID"`
	Name                 string   `codec:"name" json:"name"`
	Email                string   `codec:"email" json:"email"`
	Username             string   `codec:"username" json:"username"`
	Role                 TeamRole `codec:"role" json:"role"`
	SendChatNotification bool     `codec:"sendChatNotification" json:"sendChatNotification"`
}

func (o TeamAddMemberArg) DeepCopy() TeamAddMemberArg {
	return TeamAddMemberArg{
		SessionID:            o.SessionID,
		Name:                 o.Name,
		Email:                o.Email,
		Username:             o.Username,
		Role:                 o.Role.DeepCopy(),
		SendChatNotification: o.SendChatNotification,
	}
}

type TeamRemoveMemberArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
	Username  string `codec:"username" json:"username"`
}

func (o TeamRemoveMemberArg) DeepCopy() TeamRemoveMemberArg {
	return TeamRemoveMemberArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Username:  o.Username,
	}
}

type TeamLeaveArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
	Permanent bool   `codec:"permanent" json:"permanent"`
}

func (o TeamLeaveArg) DeepCopy() TeamLeaveArg {
	return TeamLeaveArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Permanent: o.Permanent,
	}
}

type TeamEditMemberArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Name      string   `codec:"name" json:"name"`
	Username  string   `codec:"username" json:"username"`
	Role      TeamRole `codec:"role" json:"role"`
}

func (o TeamEditMemberArg) DeepCopy() TeamEditMemberArg {
	return TeamEditMemberArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Username:  o.Username,
		Role:      o.Role.DeepCopy(),
	}
}

type LoadTeamPlusApplicationKeysArg struct {
	SessionID   int             `codec:"sessionID" json:"sessionID"`
	Id          TeamID          `codec:"id" json:"id"`
	Application TeamApplication `codec:"application" json:"application"`
	Refreshers  TeamRefreshers  `codec:"refreshers" json:"refreshers"`
}

func (o LoadTeamPlusApplicationKeysArg) DeepCopy() LoadTeamPlusApplicationKeysArg {
	return LoadTeamPlusApplicationKeysArg{
		SessionID:   o.SessionID,
		Id:          o.Id.DeepCopy(),
		Application: o.Application.DeepCopy(),
		Refreshers:  o.Refreshers.DeepCopy(),
	}
}

type TeamsInterface interface {
	TeamCreate(context.Context, TeamCreateArg) error
	TeamGet(context.Context, TeamGetArg) (TeamDetails, error)
	TeamList(context.Context, TeamListArg) (TeamList, error)
	TeamChangeMembership(context.Context, TeamChangeMembershipArg) error
	TeamAddMember(context.Context, TeamAddMemberArg) (TeamAddMemberResult, error)
	TeamRemoveMember(context.Context, TeamRemoveMemberArg) error
	TeamLeave(context.Context, TeamLeaveArg) error
	TeamEditMember(context.Context, TeamEditMemberArg) error
	// * loadTeamPlusApplicationKeys loads team information for applications like KBFS and Chat.
	// * If refreshers are non-empty, then force a refresh of the cache if the requirements
	// * of the refreshers aren't met.
	LoadTeamPlusApplicationKeys(context.Context, LoadTeamPlusApplicationKeysArg) (TeamPlusApplicationKeys, error)
}

func TeamsProtocol(i TeamsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.teams",
		Methods: map[string]rpc.ServeHandlerDescription{
			"teamCreate": {
				MakeArg: func() interface{} {
					ret := make([]TeamCreateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamCreateArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamCreateArg)(nil), args)
						return
					}
					err = i.TeamCreate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamGet": {
				MakeArg: func() interface{} {
					ret := make([]TeamGetArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamGetArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamGetArg)(nil), args)
						return
					}
					ret, err = i.TeamGet(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamList": {
				MakeArg: func() interface{} {
					ret := make([]TeamListArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamListArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamListArg)(nil), args)
						return
					}
					ret, err = i.TeamList(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamChangeMembership": {
				MakeArg: func() interface{} {
					ret := make([]TeamChangeMembershipArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamChangeMembershipArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamChangeMembershipArg)(nil), args)
						return
					}
					err = i.TeamChangeMembership(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamAddMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamAddMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamAddMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamAddMemberArg)(nil), args)
						return
					}
					ret, err = i.TeamAddMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamRemoveMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamRemoveMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamRemoveMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamRemoveMemberArg)(nil), args)
						return
					}
					err = i.TeamRemoveMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamLeave": {
				MakeArg: func() interface{} {
					ret := make([]TeamLeaveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamLeaveArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamLeaveArg)(nil), args)
						return
					}
					err = i.TeamLeave(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamEditMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamEditMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamEditMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamEditMemberArg)(nil), args)
						return
					}
					err = i.TeamEditMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"loadTeamPlusApplicationKeys": {
				MakeArg: func() interface{} {
					ret := make([]LoadTeamPlusApplicationKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoadTeamPlusApplicationKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoadTeamPlusApplicationKeysArg)(nil), args)
						return
					}
					ret, err = i.LoadTeamPlusApplicationKeys(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TeamsClient struct {
	Cli rpc.GenericClient
}

func (c TeamsClient) TeamCreate(ctx context.Context, __arg TeamCreateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamCreate", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamGet(ctx context.Context, __arg TeamGetArg) (res TeamDetails, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamGet", []interface{}{__arg}, &res)
	return
}

func (c TeamsClient) TeamList(ctx context.Context, __arg TeamListArg) (res TeamList, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamList", []interface{}{__arg}, &res)
	return
}

func (c TeamsClient) TeamChangeMembership(ctx context.Context, __arg TeamChangeMembershipArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamChangeMembership", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamAddMember(ctx context.Context, __arg TeamAddMemberArg) (res TeamAddMemberResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamAddMember", []interface{}{__arg}, &res)
	return
}

func (c TeamsClient) TeamRemoveMember(ctx context.Context, __arg TeamRemoveMemberArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamRemoveMember", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamLeave(ctx context.Context, __arg TeamLeaveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamLeave", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamEditMember(ctx context.Context, __arg TeamEditMemberArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamEditMember", []interface{}{__arg}, nil)
	return
}

// * loadTeamPlusApplicationKeys loads team information for applications like KBFS and Chat.
// * If refreshers are non-empty, then force a refresh of the cache if the requirements
// * of the refreshers aren't met.
func (c TeamsClient) LoadTeamPlusApplicationKeys(ctx context.Context, __arg LoadTeamPlusApplicationKeysArg) (res TeamPlusApplicationKeys, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.loadTeamPlusApplicationKeys", []interface{}{__arg}, &res)
	return
}
