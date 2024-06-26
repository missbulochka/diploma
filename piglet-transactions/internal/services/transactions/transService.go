package transactions

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"piglet-transactions-service/internal/domain/models"
)

type Transactions struct {
	log              *slog.Logger
	transSaver       TransactionSaver
	transProvider    TransactionProvider
	categorySaver    CategorySaver
	categoryProvider CategoryProvider
	billUpdater      BillUpdater
}

type TransactionSaver interface {
	SaveTransaction(ctx context.Context, trans models.Transaction) (err error)
	UpdateTransaction(ctx context.Context, trans models.Transaction) (err error)
	DeleteTransaction(ctx context.Context, id uuid.UUID, transType uint8) (err error)
}

type TransactionProvider interface {
	DefaultTransInfo(ctx context.Context, id uuid.UUID) (
		date time.Time,
		transType uint8,
		sum decimal.Decimal,
		comment string,
		err error)
	GetTransaction(ctx context.Context, id uuid.UUID, trans *models.Transaction) (err error)
	GetSomeTransactions(
		ctx context.Context,
		trans *[]*models.Transaction,
		count uint8) (err error)
}

type CategorySaver interface {
	SaveCategory(ctx context.Context, cat models.Category) (err error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (err error)
	UpdateCategory(ctx context.Context, cat models.Category) (err error)
}

type CategoryProvider interface {
	GetCategory(ctx context.Context, search interface{}) (category models.Category, err error)
	GetAllCategories(ctx context.Context, cat *[]*models.Category) (err error)
}

type BillUpdater interface {
	SaveBill(ctx context.Context, id uuid.UUID, billStatus bool) (err error)
	UpdateBill(ctx context.Context, id uuid.UUID, billStatus bool) (err error)
	GetBill(ctx context.Context, id uuid.UUID) (status bool, err error)
	DeleteBill(ctx context.Context, id uuid.UUID) (err error)
}

// New returns a new interface of the Transactions service
func New(
	log *slog.Logger,
	transSaver TransactionSaver,
	transProvider TransactionProvider,
	categorySaver CategorySaver,
	categoryProvider CategoryProvider,
	billUpdater BillUpdater,
) *Transactions {
	return &Transactions{
		log:              log,
		transSaver:       transSaver,
		transProvider:    transProvider,
		categorySaver:    categorySaver,
		categoryProvider: categoryProvider,
		billUpdater:      billUpdater,
	}
}
