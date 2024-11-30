package srvbook

import "context"

func (srv *Srv) Delete(ctx context.Context, id string) error {
	query := srv.repos.DB.Query()
	return query.DeleteBook(ctx, id)
}
