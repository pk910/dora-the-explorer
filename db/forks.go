package db

import (
	"github.com/ethpandaops/dora/dbtypes"
	"github.com/jmoiron/sqlx"
)

func InsertFork(fork *dbtypes.Fork, tx *sqlx.Tx) error {
	_, err := tx.Exec(EngineQuery(map[dbtypes.DBEngineType]string{
		dbtypes.DBEnginePgsql: `
			INSERT INTO forks (
				fork_id, base_slot, base_root, leaf_slot, leaf_root, parent_fork
			) VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (fork_id) DO UPDATE
			SET 
				base_slot = excluded.base_slot,
				base_root = excluded.base_root,
				leaf_slot = excluded.leaf_slot,
				leaf_root = excluded.leaf_root,
				parent_fork = excluded.parent_fork;`,
		dbtypes.DBEngineSqlite: `
			INSERT OR REPLACE INTO unfinalized_blocks (
				fork_id, root, slot, header_ver, header_ssz, block_ver, block_ssz, status, fork_id
			) VALUES ($1, $2, $3, $4, $5, $6)`,
	}),
		fork.ForkId, fork.BaseSlot, fork.BaseRoot, fork.LeafSlot, fork.LeafRoot, fork.ParentFork)
	if err != nil {
		return err
	}
	return nil
}

func GetUnfinalizedForks(finalizedSlot uint64) []*dbtypes.Fork {
	forks := []*dbtypes.Fork{}

	err := ReaderDb.Select(&forks, `SELECT fork_id, base_slot, base_root, leaf_slot, leaf_root, parent_fork
		FROM forks
		WHERE base_slot >= $1
		ORDER BY base_slot ASC
	`, finalizedSlot)
	if err != nil {
		logger.Errorf("Error while fetching unfinalized forks: %v", err)
		return nil
	}

	return forks
}

func DeleteUnfinalizedForks(finalizedSlot uint64, tx *sqlx.Tx) error {
	_, err := tx.Exec(`DELETE FROM forks WHERE base_slot < $1`, finalizedSlot)
	if err != nil {
		return err
	}
	return nil
}
