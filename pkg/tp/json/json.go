package json

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
)

type JSON struct {
}

func (j *JSON) FamilyName() string {
	return Name
}

func (j *JSON) FamilyVersion() string {
	return Version
}

func (j *JSON) IndexObject(statePayload []byte) interface{} {
	m := make(map[string]interface{})

	if err := json.Unmarshal(statePayload, &m); err != nil {
		return nil
	}

	return m
}

func (j *JSON) generateAddress(namespace []byte, id string) []byte {
	return GetJSONFamilyStateAddress(namespace, j.FamilyName(), id)
}

func (j *JSON) getData(readOnlyStateContext iface.IStateContext, address []byte) (map[string]interface{}, error) {
	entry, err := readOnlyStateContext.GetState(address)
	if err != nil {
		return nil, err
	}
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal(entry.Payload, &jsonMap); err != nil {
		return nil, err
	} else {
		return jsonMap, nil
	}
}

func (j *JSON) Lookup(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	query interface{}) (interface{}, error) {

	switch v := query.(type) {
	case string:
		return j.getData(readOnlyStateContext,
			crypto.GetStateAddressFromHexString(v))
	case []byte:
		return j.getData(readOnlyStateContext,
			v)
	default:
		return nil, x.ErrInvalidLookupInput
	}
}

func (j *JSON) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        0,
		ErrorCode:     0,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	jsonPayload := &JSONPayload{}
	if err := proto.Unmarshal(transaction.Payload, jsonPayload); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	jsonMap := jsonPayload.ToJSONMap()

	if jsonMap == nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "json decode error"
		return tpr
	}

	id, found := jsonMap["id"]

	if !found {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "id key not found"
		return tpr
	}

	idString := ""
	if val, ok := id.(string); !ok {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "id is not string"
		return tpr
	} else {
		idString = val
	}

	if len(idString) == 0 {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "id can not be empty"
		return tpr
	}

	address := j.generateAddress(transaction.Namespace, idString)

	switch jsonPayload.Action {
	case Action_INSERT:
		if se, _ := stateContext.GetState(address); se != nil {
			tpr.Status = 0
			tpr.ErrorCode = 0
			tpr.ErrorText = "state is already available"
			return tpr
		}
		entry := &pb.StateEntry{
			Payload:       transaction.Payload,
			Namespace:     transaction.Namespace,
			FamilyName:    transaction.FamilyName,
			FamilyVersion: transaction.FamilyVersion,
		}
		return j.upsert(tpr, stateContext, address, entry, jsonMap)

	case Action_MERGE:
		return j.merge(tpr, stateContext, address, jsonMap)
	case Action_UPDATE:
		return j.update(tpr, stateContext, address, jsonMap)
	case Action_UPSERT:
		entry := &pb.StateEntry{
			Payload:       transaction.Payload,
			Namespace:     transaction.Namespace,
			FamilyName:    transaction.FamilyName,
			FamilyVersion: transaction.FamilyVersion,
		}
		return j.upsert(tpr, stateContext, address, entry, jsonMap)
	case Action_DELETE:
		return j.delete(tpr, stateContext, address)
	}

	return tpr
}

func (j *JSON) processFinalData(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	stateEntry *pb.StateEntry,
	stateJsonMap map[string]interface{}) *pb.TransactionResponse {

	if data, err := json.Marshal(stateJsonMap); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		stateEntry.Payload = data
	}

	if err := stateContext.UpsertState(address, stateEntry); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.UpsertIndex(crypto.StateAddressByteSliceToHexString(address), stateJsonMap); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = 1
	tpr.ErrorCode = 0
	tpr.ErrorText = ""
	return tpr
}

func (j *JSON) merge(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	var stateEntry *pb.StateEntry

	if se, _ := stateContext.GetState(address); se == nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "state is not available"
		return tpr
	} else {
		stateEntry = se
	}

	stateJsonMap := make(map[string]interface{})

	if err := json.Unmarshal(stateEntry.Payload, &stateJsonMap); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	// NOTE: updating state json map. merge logic
	for k, v := range jsonMap {
		stateJsonMap[k] = v
	}
	return j.processFinalData(tpr, stateContext, address, stateEntry, stateJsonMap)
}

func (j *JSON) update(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	var stateEntry *pb.StateEntry

	if se, _ := stateContext.GetState(address); se == nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "state is not available"
		return tpr
	} else {
		stateEntry = se
	}

	stateJsonMap := make(map[string]interface{})

	if err := json.Unmarshal(stateEntry.Payload, &stateJsonMap); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	// NOTE: updating state json map. update logic
	for k, v := range jsonMap {
		if _, found := stateJsonMap[k]; found {
			stateJsonMap[k] = v
		} else {
			tpr.Status = 0
			tpr.ErrorCode = 0
			tpr.ErrorText = fmt.Sprintf("json key [%s] not found", k)
			return tpr
		}
	}

	return j.processFinalData(tpr, stateContext, address, stateEntry, stateJsonMap)
}

func (j *JSON) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	entry *pb.StateEntry,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.UpsertIndex(crypto.StateAddressByteSliceToHexString(address), jsonMap); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = 1
	tpr.ErrorCode = 0
	tpr.ErrorText = ""

	return tpr
}

func (j *JSON) delete(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte) *pb.TransactionResponse {

	if err := stateContext.DeleteState(address); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.DeleteIndex(crypto.StateAddressByteSliceToHexString(address)); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = 1
	tpr.ErrorCode = 0
	tpr.ErrorText = ""
	return tpr
}
