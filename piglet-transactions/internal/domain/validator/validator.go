package validation

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"piglet-transactions-service/internal/domain/models"
)

var noCategoryExpUUID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
var noCategoryIncUUID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

const (
	debtTypeImCreditor = true
	debtTypeImDebtor   = false
	transTypeIncome    = 1
	transTypeExpense   = 2
	transTypeDebt      = 3
	transTypeTransfer  = 4
)

func TransValidator(
	id string,
	date *timestamppb.Timestamp,
	transType int32,
	sum float64,
	comment string,
	idCategory string,
	debtType bool,
	idBillTo string,
	idBillFrom string,
	person string,
	repeat bool,
) (trans models.Transaction, err error) {
	val := validator.New(validator.WithRequiredStructEnabled())

	if len(id) != 0 {
		if trans.Id, err = uuid.Parse(id); err != nil {
			return trans, err
		}
	}

	if err = simpleVal(val,
		ValTrans{
			Date:      date,
			TransType: transType,
			Sum:       sum,
			Comment:   comment,
		},
		&trans,
	); err != nil {
		return trans, err
	}

	switch transType {
	case transTypeIncome:
		if err = incomeValidator(
			val,
			ValIncome{
				IdCategory: idCategory,
				IdBillTo:   idBillTo,
				Sender:     person,
				Repeat:     repeat,
			},
			&trans,
		); err != nil {
			return trans, err
		}
	case transTypeExpense:
		if err = expenseValidator(
			val,
			ValExpense{
				IdCategory: idCategory,
				IdBillFrom: idBillFrom,
				Recipient:  person,
				Repeat:     repeat,
			},
			&trans,
		); err != nil {
			return trans, err
		}
	case transTypeDebt:
		if err = debtValidator(
			val,
			ValDebt{
				DebtType:       debtType,
				IdBillFrom:     idBillFrom,
				IdBillTo:       idBillTo,
				CreditorDebtor: person,
			},
			&trans,
		); err != nil {
			return trans, err
		}
	case transTypeTransfer:
		if err = transferValidator(
			val,
			ValTransfer{
				IdBillFrom: idBillFrom,
				IdBillTo:   idBillTo,
			},
			&trans,
		); err != nil {
			return trans, err
		}
	default:
		return trans, fmt.Errorf("invalid transaction type: %v", codes.InvalidArgument)
	}

	return trans, nil
}

func CategoryValidator(
	id string,
	categoryType bool,
	name string,
	mandatory bool,
) (category models.Category, err error) {
	val := validator.New(validator.WithRequiredStructEnabled())

	if len(id) != 0 {
		if category.Id, err = uuid.Parse(id); err != nil {
			return category, fmt.Errorf("invalid category creditals: %v", codes.InvalidArgument)
		}
	}

	if err = val.Struct(
		&ValCategory{
			CategoryType: categoryType,
			Name:         name,
			Mandatory:    mandatory,
		}); err != nil {
		return category, fmt.Errorf("invalid category creditals: %v", codes.InvalidArgument)
	}

	// HACK: обработка ошибок
	category.CategoryType = categoryType
	category.Name = name
	category.Mandatory = mandatory

	return category, nil
}

func BillValidator(id string) (uuidId uuid.UUID, err error) {
	uuidId, err = uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid category creditals: %v", codes.InvalidArgument)
	}
	return uuidId, nil
}

func simpleVal(
	val *validator.Validate,
	tr ValTrans,
	trans *models.Transaction,
) error {
	if err := val.Struct(tr); err != nil {
		return fmt.Errorf("invalid transaction creditals: %v", codes.InvalidArgument)
	}

	var date time.Time
	if tr.Date == nil {
		date = time.Now()
	} else {
		if err := tr.Date.CheckValid(); err != nil {
			return fmt.Errorf("invalid transaction creditals: %v", codes.InvalidArgument)
		}
		date = tr.Date.AsTime()
	}

	trans.Date = date
	trans.TransType = uint8(tr.TransType)
	trans.Sum = decimal.NewFromFloat(tr.Sum)
	trans.Comment = tr.Comment

	return nil
}

func incomeValidator(
	val *validator.Validate,
	tr ValIncome,
	trans *models.Transaction,
) (err error) {
	if err := val.Struct(tr); err != nil {
		return fmt.Errorf("invalid income creditals: %v", codes.InvalidArgument)
	}

	if len(tr.IdCategory) == 0 {
		trans.IdCategory = noCategoryIncUUID
	} else {
		if trans.IdCategory, err = uuid.Parse(tr.IdCategory); err != nil {
			return fmt.Errorf("invalid income creditals: %v", codes.InvalidArgument)
		}
	}

	if trans.IdBillTo, err = uuid.Parse(tr.IdBillTo); err != nil {
		return fmt.Errorf("invalid income creditals: %v", codes.InvalidArgument)
	}
	// HACK: обработка ошибок
	trans.Person = tr.Sender
	trans.Repeat = tr.Repeat

	return nil
}

func expenseValidator(
	val *validator.Validate,
	tr ValExpense,
	trans *models.Transaction,
) error {
	if err := val.Struct(tr); err != nil {
		return fmt.Errorf("invalid expense creditals: %v", codes.InvalidArgument)
	}

	if len(tr.IdCategory) == 0 {
		trans.IdCategory = noCategoryExpUUID
	}

	// HACK: обработка ошибок
	trans.IdCategory, _ = uuid.Parse(tr.IdCategory)
	trans.IdBillFrom, _ = uuid.Parse(tr.IdBillFrom)
	trans.Person = tr.Recipient
	trans.Repeat = tr.Repeat

	return nil
}

func debtValidator(
	val *validator.Validate,
	tr ValDebt,
	trans *models.Transaction,
) error {
	if err := val.Struct(tr); err != nil {
		return fmt.Errorf("invalid debt creditals: %v", codes.InvalidArgument)
	}

	if (tr.DebtType == debtTypeImCreditor && len(tr.IdBillFrom) == 0) ||
		(tr.DebtType == debtTypeImDebtor && len(tr.IdBillTo) == 0) {
		return fmt.Errorf("invalid bills creditals: %v", codes.InvalidArgument)
	}

	trans.DebtType = tr.DebtType
	// HACK: обработка ошибок
	trans.IdBillFrom, _ = uuid.Parse(tr.IdBillFrom)
	trans.IdBillTo, _ = uuid.Parse(tr.IdBillTo)
	trans.Person = tr.CreditorDebtor

	return nil
}

func transferValidator(
	val *validator.Validate,
	tr ValTransfer,
	trans *models.Transaction,
) error {
	if err := val.Struct(tr); err != nil {
		return fmt.Errorf("invalid transfer creditals: %v", codes.InvalidArgument)
	}

	// HACK: обработка ошибок
	trans.IdBillFrom, _ = uuid.Parse(tr.IdBillFrom)
	trans.IdBillTo, _ = uuid.Parse(tr.IdBillTo)

	return nil
}
