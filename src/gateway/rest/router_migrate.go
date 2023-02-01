package rest

type migrateRequestBody struct {
	HostnamesToMigrate []string `json:"hostnames_to_migrate"`
	DestinationDSeq    uint64   `json:"destination_dseq"`
	DestinationGSeq    uint32   `json:"destination_gseq"`
}
