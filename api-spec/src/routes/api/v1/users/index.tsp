import "@typespec/http";
import "@typespec/openapi3";

import "../../../../models/index.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@route("/users")
namespace GoSampleApp.Routes.Api.V1.Users {
  @summary("ユーザー一覧の取得")
  @doc("ユーザー一覧を取得する")
  @get
  op getUsers(): {
    @statusCode statusCode: 200;
    @body _: Models.User[];
  };

  @summary("特定のユーザーの取得")
  @doc("特定のユーザーを取得する")
  @get
  op getUser(@path userId: int32): {
    @statusCode statusCode: 200;
    @body _: Models.User;
  };
}
