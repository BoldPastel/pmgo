package pmgo

// handles simple merge methods
// can be used as a template for other methods

func mapStructAV(input []ApplicationVisibilities) map[string]ApplicationVisibilities {
	ace := map[string]ApplicationVisibilities{}
	for _, v := range input {
		ace[v.Application] = v
	}
	return ace
}

func mapStructCP(input []CustomPermissions) map[string]CustomPermissions {
	ace := map[string]CustomPermissions{}
	for _, v := range input {
		ace[v.Name] = v
	}
	return ace
}

func mapStructRT(input []RecordTypeVisibilities) map[string]RecordTypeVisibilities {
	ace := map[string]RecordTypeVisibilities{}
	for _, v := range input {
		ace[v.RecordType] = v
	}
	return ace
}

func mapStructCA(input []ClassAccesses) map[string]ClassAccesses {
	ace := map[string]ClassAccesses{}
	for _, v := range input {
		ace[v.ApexClass] = v
	}
	return ace
}

func mapStructFP(input []FieldPermissions) map[string]FieldPermissions {
	ace := map[string]FieldPermissions{}
	for _, v := range input {
		ace[v.Field] = v
	}
	return ace
}

func mapStructOP(input []ObjectPermissions) map[string]ObjectPermissions {
	ace := map[string]ObjectPermissions{}
	for _, v := range input {
		ace[v.Object] = v
	}
	return ace
}

func mapStructTS(input []TabSettings) map[string]TabSettings {
	ace := map[string]TabSettings{}
	for _, v := range input {
		ace[v.Tab] = v
	}
	return ace
}

func mapExternalDS(input []ExternalDataSourceAccesses) map[string]ExternalDataSourceAccesses {
	ace := map[string]ExternalDataSourceAccesses{}
	for _, v := range input {
		ace[v.ExternalDataSource] = v
	}
	return ace
}

func mapPageAccesses(input []PageAccesses) map[string]PageAccesses {
	ace := map[string]PageAccesses{}
	for _, v := range input {
		ace[v.ApexPage] = v
	}
	return ace
}

func mapUserPermissions(input []UserPermissions) map[string]UserPermissions {
	ace := map[string]UserPermissions{}
	for _, v := range input {
		ace[v.Name] = v
	}
	return ace
}
