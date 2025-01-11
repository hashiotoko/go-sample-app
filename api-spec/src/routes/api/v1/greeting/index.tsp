import "@typespec/http";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@route("/greeting")
namespace GoSampleApp.Routes.Api.V1.Greeting {
  @summary("挨拶")
  @doc("挨拶メッセージを取得する")
  @get
  op greeting(): {
    @statusCode statusCode: 200;
    @body _: string;
  };
}
