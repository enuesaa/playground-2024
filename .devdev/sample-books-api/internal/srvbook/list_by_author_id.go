package srvbook

import "context"

func (srv *Srv) ListIdsByAuthorId(ctx context.Context, authorId string) ([]string, error) {
	list := make([]string, 0)

	// TODO: 実装
	list = append(list, "aaa", "bbb")

	return list, nil
}
