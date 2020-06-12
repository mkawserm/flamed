package utility

import (
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/pb"
)

func BuildIndexMetaFromMap(data map[string]interface{}) *pb.IndexMeta {
	meta := &pb.IndexMeta{}

	if v, ok := data["namespace"].(string); ok {
		meta.Namespace = []byte(v)
	}

	if v, ok := data["version"].(*kind.UInt64); ok {
		meta.Version = v.Value()
	}

	if v, ok := data["enabled"].(bool); ok {
		meta.Enabled = v
	}

	if v, ok := data["default"].(bool); ok {
		meta.Default = v
	}

	if v, found := data["indexDynamic"]; found {
		v1, ok := v.(bool)
		if ok {
			meta.IndexDynamic = v1
		}
	}

	if v, found := data["storeDynamic"]; found {
		v1, ok := v.(bool)
		if ok {
			meta.StoreDynamic = v1
		}
	}

	if v, found := data["docValuesDynamic"]; found {
		v1, ok := v.(bool)
		if ok {
			meta.DocValuesDynamic = v1
		}
	}

	if v, found := data["defaultType"]; found {
		v1, ok := v.(string)
		if ok {
			meta.DefaultType = v1
		}
	}

	if v, found := data["defaultAnalyzer"]; found {
		v1, ok := v.(string)
		if ok {
			meta.DefaultAnalyzer = v1
		}
	}

	if v, found := data["defaultDateTimeParser"]; found {
		v1, ok := v.(string)
		if ok {
			meta.DefaultDateTimeParser = v1
		}
	}

	if v, found := data["defaultField"]; found {
		v1, ok := v.(string)
		if ok {
			meta.DefaultField = v1
		}
	}

	if v, found := data["typeField"]; found {
		v1, ok := v.(string)
		if ok {
			meta.TypeField = v1
		}
	}

	if v, found := data["customAnalysis"]; found {
		v1, ok := v.(string)
		if ok {
			meta.CustomAnalysis = v1
		}
	}

	if v, found := data["indexDocumentList"]; found {
		documentList, ok := v.([]interface{})
		if ok {
			for _, doc := range documentList {
				newDoc := makeIndexDocument(doc)
				if newDoc != nil {
					meta.IndexDocumentList = append(meta.IndexDocumentList, newDoc)
				}
			}
		}
	}

	return meta
}

func makeIndexDocument(document interface{}) *pb.IndexDocument {
	documentMap, ok := document.(map[string]interface{})
	if !ok {
		return nil
	}

	indexDoc := &pb.IndexDocument{}

	if v, ok := documentMap["name"].(string); ok {
		indexDoc.Name = v
	}

	if v, ok := documentMap["enabled"].(bool); ok {
		indexDoc.Enabled = v
	}

	if v, ok := documentMap["default"].(bool); ok {
		indexDoc.Default = v
	}

	if v, ok := documentMap["dynamic"].(bool); ok {
		indexDoc.Dynamic = v
	}

	if indexFieldListInterface, found := documentMap["indexFieldList"]; found {
		indexFieldList, ok := indexFieldListInterface.([]interface{})
		if ok {
			for _, indexFieldInterface := range indexFieldList {
				newIndexField := buildIndexField(indexFieldInterface)
				if newIndexField != nil {
					indexDoc.IndexFieldList = append(indexDoc.IndexFieldList, newIndexField)
				}
			}
		}
	}

	return indexDoc
}

func buildIndexField(indexFieldInterface interface{}) *pb.IndexField {
	indexFieldMap, ok := indexFieldInterface.(map[string]interface{})
	if !ok {
		return nil
	}

	indexField := &pb.IndexField{}
	if v, ok := indexFieldMap["indexFieldType"].(int); ok {
		indexField.IndexFieldType = pb.IndexFieldType(v)
	}

	if v, ok := indexFieldMap["name"].(string); ok {
		indexField.Name = v
	}

	if v, ok := indexFieldMap["analyzer"].(string); ok {
		indexField.Analyzer = v
	}

	if v, ok := indexFieldMap["enabled"].(bool); ok {
		indexField.Enabled = v
	}

	if v, ok := indexFieldMap["index"].(bool); ok {
		indexField.Index = v
	}

	if v, ok := indexFieldMap["store"].(bool); ok {
		indexField.Store = v
	}

	if v, ok := indexFieldMap["includeTermVectors"].(bool); ok {
		indexField.IncludeTermVectors = v
	}

	if v, ok := indexFieldMap["includeInAll"].(bool); ok {
		indexField.IncludeInAll = v
	}

	if v, ok := indexFieldMap["docValues"].(bool); ok {
		indexField.DocValues = v
	}

	return indexField
}
