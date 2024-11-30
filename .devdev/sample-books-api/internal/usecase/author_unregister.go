package usecase

import (
	"context"

	"github.com/enuesaa/sample-books-api/internal/repository"
	"github.com/enuesaa/sample-books-api/internal/srvauthor"
	"github.com/enuesaa/sample-books-api/internal/srvbook"
)

func AuthorUnregister(repos repository.Repos, id string) error {
	bookSrv := srvbook.New(repos)
	authorSrv := srvauthor.New(repos)

	repos.DB.Tx(context.Background(), func(ctx context.Context) error {
		bookIds, err := bookSrv.ListIdsByAuthorId(ctx, id)
		if err != nil {
			return err
		}
		for _, bookId := range bookIds {
			if err := bookSrv.Delete(ctx, bookId); err != nil {
				return err
			}
		}
		if err := authorSrv.Delete(ctx, id); err != nil {
			return err
		}

		return nil
	})

	return nil
}
