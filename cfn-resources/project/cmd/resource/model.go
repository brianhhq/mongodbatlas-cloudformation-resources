package resource

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	Name         *string           `json:",omitempty"`
	OrgId        *string           `json:",omitempty"`
	Id           *string           `json:",omitempty"`
	Created      *string           `json:",omitempty"`
	ClusterCount *int              `json:",omitempty"`
	ApiKeys      *ApiKeyDefinition `json:",omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PublicKey  *string `json:",omitempty"`
	PrivateKey *string `json:",omitempty"`
}
