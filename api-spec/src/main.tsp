import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

import "./models/index.tsp";
import "./routes/index.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service({
  title: "go-sample-app",
})
@server("http://localhost:8888", "ローカル環境")
namespace GoSampleApp;
