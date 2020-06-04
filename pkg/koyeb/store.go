package koyeb

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	stores "github.com/koyeb/koyeb-cli/pkg/kclient/client/stores"
	apimodel "github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

var (
	createStoreCommand = &cobra.Command{
		Use:     "stores [resource]",
		Aliases: []string{"ms", "store"},
		Short:   "Create stores",
		RunE:    createStores,
	}
	getStoreCommand = &cobra.Command{
		Use:     "stores [resource]",
		Aliases: []string{"ms", "store"},
		Short:   "Get stores",
		RunE:    getStores,
	}
	describeStoreCommand = &cobra.Command{
		Use:     "stores [resource]",
		Aliases: []string{"ms", "store"},
		Short:   "Describe stores",
		RunE:    getStores,
	}
	updateStoreCommand = &cobra.Command{
		Use:     "stores [resource]",
		Aliases: []string{"ms", "store"},
		Short:   "Update stores",
		RunE:    notImplemented,
	}
	deleteStoreCommand = &cobra.Command{
		Use:     "stores [resource]",
		Aliases: []string{"ms", "store"},
		Short:   "Delete stores",
		RunE:    deleteStores,
	}
)

type StorageStoresBody struct {
	Stores []StorageStore `json:"stores"`
}

func (a *StorageStoresBody) MarshalBinary() ([]byte, error) {
	if len(a.Stores) == 1 {
		return swag.WriteJSON(a.Stores[0])
	} else {
		return swag.WriteJSON(a)
	}
}

func (a *StorageStoresBody) GetHeaders() []string {
	return []string{"id", "name", "region", "status", "updated_at"}
}

func (a *StorageStoresBody) GetTableFields() [][]string {
	var data [][]string
	for _, item := range a.Stores {
		var fields []string
		for _, field := range a.GetHeaders() {
			fields = append(fields, item.GetField(field))
		}
		data = append(data, fields)
	}
	return data
}

func (a *StorageStoresBody) New() interface{} {
	return &apimodel.StorageStore{}
}

func (a *StorageStoresBody) Append(item interface{}) {
	store := item.(*apimodel.StorageStore)
	a.Stores = append(a.Stores, StorageStore{*store})
}

type StorageStore struct {
	apimodel.StorageStore
}

func (a StorageStore) GetNewBody() *apimodel.StorageNewStore {
	newBody := apimodel.StorageNewStore{}
	copier.Copy(&newBody, &a.StorageStore)
	return &newBody
}
func (a StorageStore) GetField(field string) string {

	type StorageStore struct {
		apimodel.StorageStore
	}

	return getField(a.StorageStore, field)
}

func displayStores(items []*apimodel.StorageStore, format string) {
	var stores StorageStoresBody

	for _, item := range items {
		stores.Stores = append(stores.Stores, StorageStore{*item})
	}
	render(&stores, format)
}

func getStores(cmd *cobra.Command, args []string) error {
	client := getApiClient()

	var all []*apimodel.StorageStore

	if len(args) > 0 {
		for _, arg := range args {
			p := stores.NewGetStoreParams()
			p.ID = arg
			resp, err := client.Stores.GetStore(p, getAuth())
			if err != nil {
				apiError(err)
				continue
			}
			log.Debugf("got response: %v", resp)
			all = append([]*apimodel.StorageStore{resp.GetPayload().Store}, all...)
		}

	} else {
		page := 0
		limit := 10
		offset := 0

		for {
			p := stores.NewListStoresParams()
			strLimit := fmt.Sprintf("%d", limit)
			p.SetLimit(&strLimit)
			strOffset := fmt.Sprintf("%d", offset)
			p.SetOffset(&strOffset)

			resp, err := client.Stores.ListStores(p, getAuth())
			if err != nil {
				apiError(err)
				er(err)
			}
			log.Debugf("got response: %v", resp)
			all = append(resp.GetPayload().Stores, all...)
			page += 1
			offset = page * limit
			if int64(offset) >= resp.GetPayload().Count {
				break
			}
		}

	}
	format := "table"
	if cmd.Parent().Name() == "describe" {
		format = "yaml"
	}
	displayStores(all, format)

	return nil
}

func createStores(cmd *cobra.Command, args []string) error {
	var all StorageStoresBody

	log.Debugf("Loading file %s", file)
	err := loadMultiple(file, &all, "managed_stores")
	if err != nil {
		er(err)
	}
	log.Debugf("Content loaded %v", all.Stores)

	client := getApiClient()
	for _, store := range all.Stores {
		p := stores.NewNewStoreParams()
		p.SetBody(store.GetNewBody())
		resp, err := client.Stores.NewStore(p, getAuth())
		if err != nil {
			apiError(err)
			continue
		}
		log.Debugf("got response: %v", resp)
	}
	return nil
}

func deleteStores(cmd *cobra.Command, args []string) error {
	client := getApiClient()

	if len(args) > 0 {
		for _, arg := range args {
			p := stores.NewDeleteStoreParams()
			p.ID = arg
			resp, err := client.Stores.DeleteStore(p, getAuth())
			if err != nil {
				apiError(err)
				continue
			}
			log.Debugf("got response: %v", resp)
			log.Infof("Store %s deleted", p.ID)
		}
	}
	return nil
}
