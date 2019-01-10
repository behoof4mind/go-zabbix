package zabbix

import (
	"fmt"
)

// jHost is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/object
type jItem struct {
	HostID    string      `json:"hostid"`
	ItemID    string      `json:"itemid"`
	ItemName  string      `json:"name"`
	ItemDescr string      `json:"description"`
	LastClock int64       `json:"lastclock"`
	LastValue int         `json:"lastvalue"`
	Source    int         `json:"source"`
}

// Item returns a native Go Item struct mapped from the given JSON Item data.
func (c *jItem) Item() (*Item, error) {
	//var err error

	item := &Item{}
	item.HostID = c.HostID
	item.ItemID = c.ItemID
	item.ItemName = c.ItemName
	item.ItemDescr = c.ItemDescr
	item.LastClock = c.LastClock
	item.LastValue = c.LastValue
	item.Source = c.Source
	/*
		host.Source, err = strconv.Atoi(c.Flags)
		if err != nil {
			return nil, fmt.Errorf("Error parsing Host Flags: %v", err)
		}
	*/
	item.Source = c.Source
	return item, nil
}

// jItems is a slice of jItems structs.
type jItems []jItem

// Items returns a native Go slice of Items mapped from the given JSON ITEMS
// data.
func (c jItems) Items() ([]Item, error) {
	if c != nil {
		items := make([]Item, len(c))
		for i, jitem := range c {
			item, err := jitem.Item()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Item %d in JSON data: %v", i, err)
			}

			items[i] = *item
		}

		return items, nil
	}

	return nil, nil
}
