package time

import (
	types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/v2/beacon-chain/state"
	"github.com/prysmaticlabs/prysm/v2/config/params"
	"github.com/prysmaticlabs/prysm/v2/runtime/version"
	"github.com/prysmaticlabs/prysm/v2/time/slots"
)

// CurrentEpoch returns the current epoch number calculated from
// the slot number stored in beacon state.
//
// Spec pseudocode definition:
//  def get_current_epoch(state: BeaconState) -> Epoch:
//    """
//    Return the current epoch.
//    """
//    return compute_epoch_at_slot(state.slot)
func CurrentEpoch(state state.ReadOnlyBeaconState) types.Epoch {
	return slots.ToEpoch(state.Slot())
}

// PrevEpoch returns the previous epoch number calculated from
// the slot number stored in beacon state. It also checks for
// underflow condition.
//
// Spec pseudocode definition:
//  def get_previous_epoch(state: BeaconState) -> Epoch:
//    """`
//    Return the previous epoch (unless the current epoch is ``GENESIS_EPOCH``).
//    """
//    current_epoch = get_current_epoch(state)
//    return GENESIS_EPOCH if current_epoch == GENESIS_EPOCH else Epoch(current_epoch - 1)
func PrevEpoch(state state.ReadOnlyBeaconState) types.Epoch {
	currentEpoch := CurrentEpoch(state)
	if currentEpoch == 0 {
		return 0
	}
	return currentEpoch - 1
}

// NextEpoch returns the next epoch number calculated from
// the slot number stored in beacon state.
func NextEpoch(state state.ReadOnlyBeaconState) types.Epoch {
	return slots.ToEpoch(state.Slot()) + 1
}

// AltairCompatible returns if the input state `s` is altair compatible and input epoch `e` is higher equal than fork epoch.
func AltairCompatible(s state.BeaconState, e types.Epoch) bool {
	return s.Version() == version.Altair && e >= params.BeaconConfig().AltairForkEpoch
}

// CanUpgradeToAltair returns true if the input `slot` can upgrade to Altair.
// Spec code:
// If state.slot % SLOTS_PER_EPOCH == 0 and compute_epoch_at_slot(state.slot) == ALTAIR_FORK_EPOCH
func CanUpgradeToAltair(slot types.Slot) bool {
	epochStart := slots.IsEpochStart(slot)
	altairEpoch := slots.ToEpoch(slot) == params.BeaconConfig().AltairForkEpoch
	return epochStart && altairEpoch
}