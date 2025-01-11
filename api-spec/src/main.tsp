import "@typespec/http";
import "@typespec/openapi";

import "./models/index.tsp";
import "./routes/index.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service({
  title: "go-sample-app",
})
@info({
  license: {
    name: "private",
  },
})
@server("http://localhost:8888", "ローカル環境")
namespace GoSampleApp;
