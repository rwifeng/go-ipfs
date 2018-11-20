package provider

import (
	"context"
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
	"gx/ipfs/QmdURv6Sbob8TVW2tFFve9vcEWrSUgwPqeqnXyvYhLrkyd/go-merkledag"
)

func NewProvideAllStrategy(dag ipld.DAGService) Strategy {
	return func(ctx context.Context, cids chan cid.Cid, root cid.Cid) {
		cids <- root
		merkledag.EnumerateChildren(ctx, merkledag.GetLinksWithDAG(dag), root, func(cid cid.Cid) bool {
			cids <- cid
			return true
		})
	}
}