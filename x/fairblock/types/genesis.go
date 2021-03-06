package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ShareList:       []Share{},
		TargetList:      []Target{},
		CommitList:      []Commit{},
		EncryptedtxList: []Encryptedtx{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in share
	shareIndexMap := make(map[string]struct{})

	for _, elem := range gs.ShareList {
		index := string(ShareKey(elem.Index))
		if _, ok := shareIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for share")
		}
		shareIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in target
	targetIndexMap := make(map[string]struct{})

	for _, elem := range gs.TargetList {
		index := string(TargetKey(elem.Index))
		if _, ok := targetIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for target")
		}
		targetIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in commit
	commitIndexMap := make(map[string]struct{})

	for _, elem := range gs.CommitList {
		index := string(CommitKey(elem.Index))
		if _, ok := commitIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for commit")
		}
		commitIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in encryptedtx
	encryptedtxIndexMap := make(map[string]struct{})

	for _, elem := range gs.EncryptedtxList {
		index := string(EncryptedtxKey(elem.Index))
		if _, ok := encryptedtxIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for encryptedtx")
		}
		encryptedtxIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
