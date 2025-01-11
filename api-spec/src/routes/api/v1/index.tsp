import "@typespec/http";
import "@typespec/openapi3";

import "./greeting/index.tsp";
import "./users/index.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@route("/v1")
namespace GoSampleApp.Routes.Api.V1;
