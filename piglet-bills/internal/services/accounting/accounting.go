package accounting

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	models "piglet-bills-service/internal/domain/model"
	"piglet-bills-service/internal/storage"
)

type Accounting struct {
	log          *slog.Logger
	billSaver    BillSaver
	billProvider BillProvider
}

type BillSaver interface {
	SaveBill(
		ctx context.Context,
		billType bool,
		billName string,
		goalSum decimal.Decimal,
		date time.Time,
		monthlyPayment decimal.Decimal,
	) (bill models.Bill, err error)
}

type BillProvider interface {
	BillReturner(
		ctx context.Context,
		billId string,
		billName string,
	) (bill models.Bill, err error)
	SomeBillsReturner(ctx context.Context, billType bool) (bills []*models.Bill, err error)
}

var (
	ErrBillExists   = errors.New("bill already exists")
	ErrBillNotFound = errors.New("bill not found")
)

// New returns a new intarface of the Accounting service
func New(
	log *slog.Logger,
	billSaver BillSaver,
	billProvider BillProvider,
) *Accounting {
	return &Accounting{
		log:          log,
		billSaver:    billSaver,
		billProvider: billProvider,
	}
}

// CreateBill create new bill in the system and returns bill
// If bill with given name already exists, returns error
func (a *Accounting) CreateBill(
	ctx context.Context,
	billType bool,
	billName string,
	goalSum decimal.Decimal,
	date time.Time,
) (bill models.Bill, err error) {
	const op = "pigletBills | accounting.saveBill"

	log := a.log.With(
		slog.String("op", op),
		// These may be things that are not profitable for business to log
		slog.String("billType", strconv.FormatBool(billType)),
		slog.String("billName", billName),
	)

	monthlyPayment := decimal.New(0, 0)
	if billType == false {
		if monthlyPayment, err = countPayment(date, goalSum); err != nil {
			log.Warn("something wrong with date", err)

			return models.Bill{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	log.Info("saving bill")

	bill, err = a.billSaver.SaveBill(
		ctx,
		billType,
		billName,
		goalSum,
		date,
		monthlyPayment,
	)
	if err != nil {
		if errors.Is(err, storage.ErrBillExists) {
			log.Warn("bill already exists", err)

			return models.Bill{}, fmt.Errorf("%s: %w", op, ErrBillExists)
		}

		log.Error("failed to save bill", err)

		return models.Bill{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("saved bill")
	return bill, nil
}

// GetSomeBills retrieves bills from the system by their types and returns them.
// If an error occurs, it returns an error.
func (a *Accounting) GetSomeBills(ctx context.Context, billType bool) (bills []*models.Bill, err error) {
	const op = "pigletBills | accounting.getSomeBills"
	log := a.log.With(slog.String("op", op))

	log.Info("searching bills")

	bills, err = a.billProvider.SomeBillsReturner(ctx, billType)
	if err != nil {
		log.Error("failed to search bill", err)

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("searched bill")
	return bills, nil
}

// GetBill retrieves a bill from the system by its name or uuid and returns it.
// If a bill with the given uuid or name does not exist, returns an error.
func (a *Accounting) GetBill(
	ctx context.Context,
	billId string,
	billName string,
) (bill models.Bill, err error) {
	const op = "pigletBills | accounting.getBill"

	log := a.log.With(
		slog.String("op", op),
		// These may be things that are not profitable for business to log
		slog.String("billId", billId),
		slog.String("billName", billName),
	)

	log.Info("searching bill")

	bill, err = a.billProvider.BillReturner(ctx, billId, billName)
	if err != nil {
		if errors.Is(err, storage.ErrBillNotFound) {
			log.Warn("bill not found", err)

			return models.Bill{}, fmt.Errorf("%s: %w", op, ErrBillExists)
		}

		log.Error("failed to search bill", err)

		return models.Bill{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("searched bill")
	return bill, nil
}

func countPayment(futureDate time.Time, sum decimal.Decimal) (monthlyPayment decimal.Decimal, err error) {
	// HACK: подумать над функцией поиска (или найти библиотеку)
	months := math.Ceil(time.Until(futureDate).Hours() / 24 / 30)

	if months == 0 {
		return sum, nil
	}

	monthlyPayment = sum.Div(decimal.New(int64(months), 0))

	return monthlyPayment.Round(0), nil
}
