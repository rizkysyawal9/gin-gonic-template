package apperror

import "errors"

var NoRecordFoundErr = errors.New("no record found")
var UnknownTableStatusErr = errors.New("unkonwn table status")
var CheckInErr = errors.New("failed to check in")
var TableOccupiedErr = errors.New("table is occupied")
var FieldRequiredErr = errors.New("fields required")
