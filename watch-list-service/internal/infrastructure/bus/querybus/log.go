package querybus

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"
	"go.uber.org/zap"
)

type log struct {
	logger *zap.Logger
	next   domain.QueryBus
}

func (l log) RegisterHandler(query domain.Query, handler domain.QueryHandler) (err error) {
	defer func() {
		if err != nil {
			l.logger.Error("failed to register query",
				zap.String("query", query.Id()),
			)
			return
		}

		l.logger.Info("registered query", zap.String("query", query.Id()))
	}()

	err = l.next.RegisterHandler(query, handler)
	return err
}

func (l log) Ask(ctx context.Context, query domain.Query) (res interface{}, err error) {
	defer func() {
		if err != nil {
			l.logger.Warn("failed to ask query",
				zap.String("query", query.Id()),
				zap.String("desc", ddderr.GetDescription(err)),
				zap.String("parent_desc", ddderr.GetParentDescription(err)),
				zap.Error(err),
			)
			return
		}

		l.logger.Info("asked query", zap.String("query", query.Id()))
	}()

	res, err = l.next.Ask(ctx, query)
	return res, err
}
