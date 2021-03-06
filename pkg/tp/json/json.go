package json

import (
	"bytes"
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

func (j *JSON) getDataAsJSONMap(readOnlyStateContext iface.IStateContext, address []byte) (map[string]interface{}, error) {
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

func (j *JSON) getDataAsBytes(readOnlyStateContext iface.IStateContext, address []byte) ([]byte, error) {
	entry, err := readOnlyStateContext.GetState(address)
	if err != nil {
		return nil, err
	}

	return entry.Payload, nil
}

func (j *JSON) Search(_ context.Context,
	_ iface.IStateContext,
	_ *pb.SearchInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (j *JSON) Iterate(_ context.Context,
	_ iface.IStateContext,
	_ *pb.IterateInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (j *JSON) Retrieve(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	retrieveInput *pb.RetrieveInput) (interface{}, error) {
	if len(retrieveInput.Addresses) == 0 {
		return nil, nil
	}

	dataList := make([][]byte, 0, len(retrieveInput.Addresses))
	for _, sa := range retrieveInput.Addresses {
		if !bytes.HasPrefix(sa, retrieveInput.Namespace) {
			return nil, x.ErrAccessViolation
		}

		dataAsBytes, err := j.getDataAsBytes(readOnlyStateContext, sa)
		if err != nil {
			dataList = append(dataList, nil)
			continue
		}

		dataList = append(dataList, dataAsBytes)
	}

	return dataList, nil
}

func (j *JSON) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {
	tpr := &pb.TransactionResponse{
		Status:        pb.Status_REJECTED,
		ErrorCode:     1,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	jsonPayload := &JSONPayload{}
	if err := proto.Unmarshal(transaction.Payload, jsonPayload); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 2
		tpr.ErrorText = err.Error()
		return tpr
	}

	jsonMap := jsonPayload.ToJSONMap()

	if jsonMap == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 3
		tpr.ErrorText = "json decode error"
		return tpr
	}

	id, found := jsonMap["id"]

	if !found {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 4
		tpr.ErrorText = "id key not found"
		return tpr
	}

	idString := ""
	if val, ok := id.(string); !ok {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 5
		tpr.ErrorText = "id is not string"
		return tpr
	} else {
		idString = val
	}

	if len(idString) == 0 {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 6
		tpr.ErrorText = "id can not be empty"
		return tpr
	}

	address := j.generateAddress(transaction.Namespace, idString)

	switch jsonPayload.Action {
	case Action_INSERT:
		if se, _ := stateContext.GetState(address); se != nil {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 7
			tpr.ErrorText = "state is already available"
			return tpr
		}
		entry := &pb.StateEntry{
			Namespace:     transaction.Namespace,
			FamilyName:    transaction.FamilyName,
			FamilyVersion: transaction.FamilyVersion,
		}

		if payload, err := json.Marshal(jsonMap); err == nil {
			entry.Payload = payload
		} else {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 8
			tpr.ErrorText = err.Error()
			return tpr
		}

		return j.upsert(tpr, stateContext, address, entry, jsonMap)
	case Action_MERGE:
		return j.merge(tpr, stateContext, address, jsonMap)
	case Action_UPDATE:
		return j.update(tpr, stateContext, address, jsonMap)
	case Action_UPSERT:
		entry := &pb.StateEntry{
			Namespace:     transaction.Namespace,
			FamilyName:    transaction.FamilyName,
			FamilyVersion: transaction.FamilyVersion,
		}
		if payload, err := json.Marshal(jsonMap); err == nil {
			entry.Payload = payload
		} else {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 9
			tpr.ErrorText = err.Error()
			return tpr
		}

		return j.upsert(tpr, stateContext, address, entry, jsonMap)
	case Action_DELETE:
		return j.delete(tpr, stateContext, address)
	default:
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 10
		tpr.ErrorText = "unknown action"
		return tpr
	}
}

func (j *JSON) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	entry *pb.StateEntry,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 11
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.UpsertIndex(crypto.StateAddressByteSliceToHexString(address), jsonMap); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 12
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = pb.Status_ACCEPTED
	tpr.ErrorCode = 0
	tpr.ErrorText = ""

	return tpr
}

func (j *JSON) delete(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte) *pb.TransactionResponse {

	if err := stateContext.DeleteState(address); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 13
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.DeleteIndex(crypto.StateAddressByteSliceToHexString(address)); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 14
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = pb.Status_ACCEPTED
	tpr.ErrorCode = 0
	tpr.ErrorText = ""
	return tpr
}

func (j *JSON) update(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	var stateEntry *pb.StateEntry

	if se, _ := stateContext.GetState(address); se == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 15
		tpr.ErrorText = "state is not available"
		return tpr
	} else {
		stateEntry = se
	}

	stateJsonMap := make(map[string]interface{})

	if err := json.Unmarshal(stateEntry.Payload, &stateJsonMap); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 16
		tpr.ErrorText = err.Error()
		return tpr
	}

	// NOTE: updating state json map. update logic
	for k, v := range jsonMap {
		if _, found := stateJsonMap[k]; found {
			stateJsonMap[k] = v
		} else {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 17
			tpr.ErrorText = fmt.Sprintf("json key [%s] not found", k)
			return tpr
		}
	}

	return j.processFinalData(tpr, stateContext, address, stateEntry, stateJsonMap)
}

func (j *JSON) merge(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	jsonMap map[string]interface{}) *pb.TransactionResponse {

	var stateEntry *pb.StateEntry

	if se, _ := stateContext.GetState(address); se == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 18
		tpr.ErrorText = "state is not available"
		return tpr
	} else {
		stateEntry = se
	}

	stateJsonMap := make(map[string]interface{})

	if err := json.Unmarshal(stateEntry.Payload, &stateJsonMap); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 19
		tpr.ErrorText = err.Error()
		return tpr
	}

	// NOTE: updating state json map. merge logic
	for k, v := range jsonMap {
		stateJsonMap[k] = v
	}
	return j.processFinalData(tpr, stateContext, address, stateEntry, stateJsonMap)
}

func (j *JSON) processFinalData(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	stateEntry *pb.StateEntry,
	stateJsonMap map[string]interface{}) *pb.TransactionResponse {

	if data, err := json.Marshal(stateJsonMap); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 20
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		stateEntry.Payload = data
	}

	if err := stateContext.UpsertState(address, stateEntry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 21
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.UpsertIndex(crypto.StateAddressByteSliceToHexString(address), stateJsonMap); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 22
		tpr.ErrorText = err.Error()
		return tpr
	}

	tpr.Status = pb.Status_ACCEPTED
	tpr.ErrorCode = 0
	tpr.ErrorText = ""
	return tpr
}
