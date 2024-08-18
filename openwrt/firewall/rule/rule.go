package rule

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/joneshf/terraform-provider-openwrt/lucirpc"
	"github.com/joneshf/terraform-provider-openwrt/openwrt/firewall/zone"
	"github.com/joneshf/terraform-provider-openwrt/openwrt/internal/lucirpcglue"
)

const (
	nameAttribute            = "name"
	nameAttributeDescription = "rule name"
	nameUCIOption            = "name"

	targetAttribute            = "target"
	targetAttributeDescription = "rule target"
	targetUCIOption            = "target"

	destAttribute            = "dest"
	destAttributeDescription = "zone dest"
	destUCIOption            = "dest"

	srcAttribute            = "src"
	srcAttributeDescription = "zone src"
	srcUCIOption            = "src"

	schemaDescription = "Legacy VLAN configuration"

	uciConfig = "firewall"
	uciType   = "forwarding"
)

var (
	nameSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       nameAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetName, nameAttribute, nameUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetName, nameAttribute, nameUCIOption),
	}

	targetSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       targetAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetTarget, targetAttribute, targetUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetTarget, targetAttribute, targetUCIOption),
		Validators:        zone.TypeValidators,
	}

	destSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       destAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetDest, destAttribute, destUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetDest, destAttribute, destUCIOption),
	}

	srcSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       srcAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetSrc, srcAttribute, srcUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetSrc, srcAttribute, srcUCIOption),
	}

	schemaAttributes = map[string]lucirpcglue.SchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		srcAttribute:            srcSchemaAttribute,
		lucirpcglue.IdAttribute: lucirpcglue.IdSchemaAttribute(modelGetId, modelSetId),
		destAttribute:           destSchemaAttribute,
		targetAttribute:         targetSchemaAttribute,
		nameAttribute:           nameSchemaAttribute,
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
	Id     types.String `tfsdk:"id"`
	Src    types.String `tfsdk:"src"`
	Dest   types.String `tfsdk:"dest"`
	Target types.String `tfsdk:"target"`
	Name   types.String `tfsdk:"name"`
}

func modelGetTarget(m model) types.String { return m.Target }
func modelGetName(m model) types.String   { return m.Name }
func modelGetSrc(m model) types.String    { return m.Src }
func modelGetId(m model) types.String     { return m.Id }
func modelGetDest(m model) types.String   { return m.Dest }

func modelSetSrc(m *model, value types.String)    { m.Src = value }
func modelSetDest(m *model, value types.String)   { m.Dest = value }
func modelSetId(m *model, value types.String)     { m.Id = value }
func modelSetTarget(m *model, value types.String) { m.Target = value }
func modelSetName(m *model, value types.String)   { m.Name = value }
