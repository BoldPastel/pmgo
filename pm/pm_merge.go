package pmgo

func MergeXML(target Permissions, operant Permissions) Permissions {
	var mergedXML Permissions

	//reminder for creative me -> need both maps coz i cant loop on perm struct; the comparison is at the node level

	oAV := mapStructAV(operant.ApplicationVisibilities)
	tAV := mapStructAV(target.ApplicationVisibilities)

	for k, v := range oAV {
		tAV[k] = v
	}
	for _, z := range tAV {
		mergedXML.ApplicationVisibilities = append(mergedXML.ApplicationVisibilities, z)
	}

	oCA := mapStructCA(operant.ClassAccesses)
	tCA := mapStructCA(target.ClassAccesses)

	for k, v := range oCA {
		tCA[k] = v
	}
	for _, z := range tCA {
		mergedXML.ClassAccesses = append(mergedXML.ClassAccesses, z)
	}

	oCP := mapStructCP(operant.CustomPermissions)
	tCP := mapStructCP(target.CustomPermissions)

	for k, v := range oCP {
		tCP[k] = v
	}
	for _, z := range tCP {
		mergedXML.CustomPermissions = append(mergedXML.CustomPermissions, z)
	}

	oEDA := mapExternalDS(operant.ExternalDataSourceAccesses)
	tEDA := mapExternalDS(target.ExternalDataSourceAccesses)

	for k, v := range oEDA {
		tEDA[k] = v
	}
	for _, z := range tEDA {
		mergedXML.ExternalDataSourceAccesses = append(mergedXML.ExternalDataSourceAccesses, z)
	}

	oFP := mapStructFP(operant.FieldPermissions)
	tFP := mapStructFP(target.FieldPermissions)

	for k, v := range oFP {
		tFP[k] = v
	}
	for _, z := range tFP {
		mergedXML.FieldPermissions = append(mergedXML.FieldPermissions, z)
	}

	oOP := mapStructOP(operant.ObjectPermissions)
	tOP := mapStructOP(target.ObjectPermissions)

	for k, v := range oOP {
		tOP[k] = v
	}
	for _, z := range tOP {
		mergedXML.ObjectPermissions = append(mergedXML.ObjectPermissions, z)
	}

	oPA := mapPageAccesses(operant.PageAccesses)
	tPA := mapPageAccesses(target.PageAccesses)

	for k, v := range oPA {
		tPA[k] = v
	}
	for _, z := range tPA {
		mergedXML.PageAccesses = append(mergedXML.PageAccesses, z)
	}

	oRT := mapStructRT(operant.RecordTypeVisibilities)
	tRT := mapStructRT(target.RecordTypeVisibilities)

	for k, v := range oRT {
		tRT[k] = v
	}
	for _, z := range tRT {
		mergedXML.RecordTypeVisibilities = append(mergedXML.RecordTypeVisibilities, z)
	}

	oTS := mapStructTS(operant.TabSettings)
	tTS := mapStructTS(target.TabSettings)

	for k, v := range oTS {
		tTS[k] = v
	}
	for _, z := range tTS {
		mergedXML.TabSettings = append(mergedXML.TabSettings, z)
	}

	oUP := mapUserPermissions(operant.UserPermissions)
	tUP := mapUserPermissions(target.UserPermissions)

	for k, v := range oUP {
		tUP[k] = v
	}
	for _, z := range tUP {
		mergedXML.UserPermissions = append(mergedXML.UserPermissions, z)
	}

	mergedXML.xmlns = operant.xmlns
	mergedXML.XMLName = operant.XMLName
	mergedXML.HasActivationRequired = operant.HasActivationRequired
	mergedXML.Label = target.Label + operant.Label

	return mergedXML
}
