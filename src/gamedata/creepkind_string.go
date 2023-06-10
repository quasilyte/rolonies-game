// Code generated by "stringer -type=CreepKind -trimprefix=creep"; DO NOT EDIT.

package gamedata

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CreepPrimitiveWanderer-0]
	_ = x[CreepStunner-1]
	_ = x[CreepAssault-2]
	_ = x[CreepDominator-3]
	_ = x[CreepBuilder-4]
	_ = x[CreepTurret-5]
	_ = x[CreepTurretConstruction-6]
	_ = x[CreepCrawlerBaseConstruction-7]
	_ = x[CreepBase-8]
	_ = x[CreepCrawlerBase-9]
	_ = x[CreepCrawler-10]
	_ = x[CreepHowitzer-11]
	_ = x[CreepServant-12]
	_ = x[CreepUberBoss-13]
}

const _CreepKind_name = "CreepPrimitiveWandererCreepStunnerCreepAssaultCreepDominatorCreepBuilderCreepTurretCreepTurretConstructionCreepCrawlerBaseConstructionCreepBaseCreepCrawlerBaseCreepCrawlerCreepHowitzerCreepServantCreepUberBoss"

var _CreepKind_index = [...]uint8{0, 22, 34, 46, 60, 72, 83, 106, 134, 143, 159, 171, 184, 196, 209}

func (i CreepKind) String() string {
	if i < 0 || i >= CreepKind(len(_CreepKind_index)-1) {
		return "CreepKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CreepKind_name[_CreepKind_index[i]:_CreepKind_index[i+1]]
}