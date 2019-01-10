package zabbix

import (
	"fmt"
)

// Item represents a Zabbix Item returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/item/object
type Item struct {
	// HostID is the unique ID of the Host.
	HostID string

	// ItemID is the unique ID of the Item.
	ItemID string

	// Itemname is the technical name of the Item.
	ItemName string

	// ItemDescr is the description of the Item.
	ItemDescr string

	// todo LastClock is the visible name of the Item.
	LastClock int64

	// todo DisplayName is the visible name of the Host.
	LastValue int

	// todo Source is the origin of the Host and must be one of the HostSource
	// constants.
	Source int
}

// todo ItemGetParams represent the parameters for a `item.get` API call.
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/item/get
type ItemGetParams struct {
	GetParameters
	// GroupIDs filters search results to hosts that are members of the given
	// Group IDs.
	GroupIDs []string `json:"groupids,omitempty"`

	// ApplicationIDs filters search results to hosts that have items in the
	// given Application IDs.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// DiscoveredServiceIDs filters search results to hosts that are related to
	// the given discovered service IDs.
	DiscoveredServiceIDs []string `json:"dserviceids,omitempty"`

	// GraphIDs filters search results to hosts that have the given graph IDs.
	GraphIDs []string `json:"graphids,omitempty"`

	// HostIDs filters search results to hosts that matched the given Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// WebCheckIDs filters search results to hosts with the given Web Check IDs.
	WebCheckIDs []string `json:"httptestids,omitempty"`

	// InterfaceIDs filters search results to hosts that use the given Interface
	// IDs.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// ItemIDs filters search results to hosts with the given Item IDs.
	ItemIDs []string `json:"itemids,omitempty"`

	// MaintenanceIDs filters search results to hosts that are affected by the
	// given Maintenance IDs
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`

	// MonitoredOnly filters search results to return only monitored hosts.
	MonitoredOnly bool `json:"monitored_hosts,omitempty"`

	// ProxyOnly filters search results to hosts which are Zabbix proxies.
	ProxiesOnly bool `json:"proxy_host,omitempty"`

	// ProxyIDs filters search results to hosts monitored by the given Proxy
	// IDs.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// IncludeTemplates extends search results to include Templates.
	IncludeTemplates bool `json:"templated_hosts,omitempty"`

	// SelectGroups causes the Host Groups that each Host belongs to to be
	// attached in the search results.
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`

	// SelectApplications causes the Applications from each Host to be attached
	// in the search results.
	SelectApplications SelectQuery `json:"selectApplications,omitempty"`

	// SelectDiscoveries causes the Low-Level Discoveries from each Host to be
	// attached in the search results.
	SelectDiscoveries SelectQuery `json:"selectDiscoveries,omitempty"`

	// SelectDiscoveryRule causes the Low-Level Discovery Rule that created each
	// Host to be attached in the search results.
	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`

	// SelectGraphs causes the Graphs from each Host to be attached in the
	// search results.
	SelectGraphs SelectQuery `json:"selectGraphs,omitempty"`

	SelectHostDiscovery SelectQuery `json:"selectHostDiscovery,omitempty"`

	SelectWebScenarios SelectQuery `json:"selectHttpTests,omitempty"`

	SelectInterfaces SelectQuery `json:"selectInterfaces,omitempty"`

	SelectInventory SelectQuery `json:"selectInventory,omitempty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`

	SelectMacros SelectQuery `json:"selectMacros,omitempty"`

	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectScreens         SelectQuery `json:"selectScreens,omitempty"`
	SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"`
}

// GetItems queries the Zabbix API for Items matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetItems(params ItemGetParams) ([]Item, error) {
	items := make([]jItem, 0)
	err := c.Get("item.get", params, &items)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Item, len(items))
	for i, jitem := range items {
		item, err := jitem.Item()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Item %d in response: %v", i, err)
		}

		out[i] = *item
	}

	return out, nil
}
