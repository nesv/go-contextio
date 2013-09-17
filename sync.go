package contextio

type SyncResponse map[string]map[string]SyncMailboxBlock

type SyncMailboxBlock struct {
	InitialImportFinished bool `json:"initial_import_finished"`
	LastExpunge           int  `json:"last_expunge"`
	LastSyncStart         int  `json:"last_sync_start"`
	LastSyncStop          int  `json:"last_sync_stop"`
}
