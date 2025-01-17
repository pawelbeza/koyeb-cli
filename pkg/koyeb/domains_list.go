package koyeb

import (
	"strconv"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *DomainHandler) List(cmd *cobra.Command, args []string) error {
	list := []koyeb.Domain{}

	page := int64(0)
	offset := int64(0)
	limit := int64(100)
	for {
		res, st, err := h.client.DomainsApi.ListDomains(h.ctx).
			Limit(strconv.FormatInt(limit, 10)).
			Offset(strconv.FormatInt(offset, 10)).
			Types([]string{string(koyeb.DOMAINTYPE_CUSTOM)}).
			Execute()
		if err != nil {
			fatalApiError(err, st)
		}
		list = append(list, res.GetDomains()...)

		page++
		offset = page * limit
		if offset >= res.GetCount() {
			break
		}
	}

	full := GetBoolFlags(cmd, "full")
	output := GetStringFlags(cmd, "output")
	listDomainsReply := NewListDomainsReply(h.mapper, &koyeb.ListDomainsReply{Domains: list}, full)

	return renderer.NewListRenderer(listDomainsReply).Render(output)
}

type ListDomainsReply struct {
	mapper *idmapper.Mapper
	value  *koyeb.ListDomainsReply
	full   bool
}

func NewListDomainsReply(mapper *idmapper.Mapper, value *koyeb.ListDomainsReply, full bool) *ListDomainsReply {
	return &ListDomainsReply{
		mapper: mapper,
		value:  value,
		full:   full,
	}
}

func (ListDomainsReply) Title() string {
	return "Domains"
}

func (r *ListDomainsReply) MarshalBinary() ([]byte, error) {
	return r.value.MarshalJSON()
}

func (r *ListDomainsReply) Headers() []string {
	return []string{"id", "name", "app", "status", "verified_at", "type", "created_at"}
}

func (r *ListDomainsReply) Fields() []map[string]string {
	items := r.value.GetDomains()
	resp := make([]map[string]string, 0, len(items))

	for _, item := range items {
		fields := map[string]string{
			"id":          renderer.FormatDomainID(r.mapper, item.GetId(), r.full),
			"name":        item.GetName(),
			"app":         renderer.FormatAppName(r.mapper, item.GetAppId(), r.full),
			"status":      string(item.GetStatus()),
			"verified_at": formatVerifiedAt(&item),
			"type":        string(item.GetType()),
			"created_at":  renderer.FormatTime(item.GetCreatedAt()),
		}
		resp = append(resp, fields)
	}

	return resp
}
