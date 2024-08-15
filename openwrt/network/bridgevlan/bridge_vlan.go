package bridgevlan

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/joneshf/terraform-provider-openwrt/lucirpc"
	"github.com/joneshf/terraform-provider-openwrt/openwrt/internal/lucirpcglue"
)

const (
	deviceAttribute            = "device"
	deviceAttributeDescription = "The switch to configure."
	deviceUCIOption            = "device"

	portsAttribute            = "ports"
	portsAttributeDescription = "A string of space-separated port indicies that should be associated with the VLAN. Adding the suffix `\"t\"` to a port indicates that egress packets should be tagged, for example `\"0 1 3t 5t\"`."
	portsUCIOption            = "ports"

	schemaDescription = "Legacy VLAN configuration"

	uciConfig = "network"
	uciType   = "bridge-vlan"

	vLanAttribute            = "vlan"
	vLanAttributeDescription = `The VLAN "table index" to configure. This index corresponds to the order on LuCI's UI`
	vLanUCIOption            = "vlan"
)

var (
	deviceSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       deviceAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetDevice, deviceAttribute, deviceUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetDevice, deviceAttribute, deviceUCIOption),
	}

	portsSchemaAttribute = lucirpcglue.ListStringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       portsAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionListString(modelSetPorts, portsAttribute, portsUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionListString(modelGetPorts, portsAttribute, portsUCIOption),
	}

	schemaAttributes = map[string]lucirpcglue.SchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		deviceAttribute:         deviceSchemaAttribute,
		lucirpcglue.IdAttribute: lucirpcglue.IdSchemaAttribute(modelGetId, modelSetId),
		portsAttribute:          portsSchemaAttribute,
		vLanAttribute:           vLanSchemaAttribute,
	}

	vLanSchemaAttribute = lucirpcglue.Int64SchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       vLanAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionInt64(modelSetVLan, vLanAttribute, vLanUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionInt64(modelGetVLan, vLanAttribute, vLanUCIOption),
		Validators: []validator.Int64{
			int64validator.Any(),
		},
	}
)

func NewDataSource() datasource.DataSource {
	return lucirpcglue.NewDataSource(
		modelGetId,
		schemaAttributes,
		schemaDescription,
		uciConfig,
		uciType,
	)
}

func NewResource() resource.Resource {
	return lucirpcglue.NewResource(
		modelGetId,
		schemaAttributes,
		schemaDescription,
		uciConfig,
		uciType,
	)
}

type model struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Ports  types.List   `tfsdk:"ports"`
	VLan   types.Int64  `tfsdk:"vlan"`
}

func modelGetDevice(m model) types.String { return m.Device }
func modelGetId(m model) types.String     { return m.Id }
func modelGetPorts(m model) types.List    { return m.Ports }
func modelGetVLan(m model) types.Int64    { return m.VLan }

func modelSetDevice(m *model, value types.String) { m.Device = value }
func modelSetId(m *model, value types.String)     { m.Id = value }
func modelSetPorts(m *model, value types.List)    { m.Ports = value }
func modelSetVLan(m *model, value types.Int64)    { m.VLan = value }
