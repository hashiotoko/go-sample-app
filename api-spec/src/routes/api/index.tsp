import "@typespec/http";
import "@typespec/openapi3";

import "./v1/index.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@route("/api")
namespace GoSampleApp.Routes.Api;
