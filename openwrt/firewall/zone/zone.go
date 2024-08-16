package zone

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/joneshf/terraform-provider-openwrt/lucirpc"
	"github.com/joneshf/terraform-provider-openwrt/openwrt/internal/lucirpcglue"
)

const (
	nameAttribute            = "name"
	nameAttributeDescription = "The name of the zone."
	nameUCIOption            = "name"

	forwardAttribute            = "forward"
	forwardAttributeDescription = "forward config"
	forwardUCIOption            = "forward"

	inputAttribute            = "input"
	inputAttributeDescription = "input config"
	inputUCIOption            = "input"

	outputAttribute            = "output"
	outputAttributeDescription = "output config"
	outputUCIOption            = "output"

	schemaDescription = "Legacy VLAN configuration"

	uciConfig = "firewall"
	uciType   = "zone"

	typeAccept = "ACCEPT"
	typeReject = "REJECT"
	typeDrop   = "DROP"
)

var (
	actionValidator = []validator.String{
		stringvalidator.OneOf(
			typeAccept,
			typeReject,
			typeDrop,
		),
	}

	nameSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       nameAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetName, nameAttribute, nameUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetName, nameAttribute, nameUCIOption),
	}

	forwardSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       forwardAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetForward, forwardAttribute, forwardUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetForward, forwardAttribute, forwardUCIOption),
		Validators:        actionValidator,
	}

	inputSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       inputAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetInput, inputAttribute, inputUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetInput, inputAttribute, inputUCIOption),
		Validators:        actionValidator,
	}

	outputSchemaAttribute = lucirpcglue.StringSchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		Description:       outputAttributeDescription,
		ReadResponse:      lucirpcglue.ReadResponseOptionString(modelSetOutput, outputAttribute, outputUCIOption),
		ResourceExistence: lucirpcglue.Required,
		UpsertRequest:     lucirpcglue.UpsertRequestOptionString(modelGetOutput, outputAttribute, outputUCIOption),
		Validators:        actionValidator,
	}

	schemaAttributes = map[string]lucirpcglue.SchemaAttribute[model, lucirpc.Options, lucirpc.Options]{
		forwardAttribute:        forwardSchemaAttribute,
		lucirpcglue.IdAttribute: lucirpcglue.IdSchemaAttribute(modelGetId, modelSetId),
		inputAttribute:          inputSchemaAttribute,
		outputAttribute:         outputSchemaAttribute,
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
	Id      types.String `tfsdk:"id"`
	Forward types.String `tfsdk:"forward"`
	Output  types.String `tfsdk:"output"`
	Input   types.String `tfsdk:"input"`
	Name    types.String `tfsdk:"name"`
}

func modelGetOutput(m model) types.String  { return m.Output }
func modelGetId(m model) types.String      { return m.Id }
func modelGetInput(m model) types.String   { return m.Input }
func modelGetName(m model) types.String    { return m.Name }
func modelGetForward(m model) types.String { return m.Forward }

func modelSetOutput(m *model, value types.String)  { m.Output = value }
func modelSetForward(m *model, value types.String) { m.Forward = value }
func modelSetInput(m *model, value types.String)   { m.Input = value }
func modelSetName(m *model, value types.String)    { m.Name = value }
func modelSetId(m *model, value types.String)      { m.Id = value }
