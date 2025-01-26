namespace GoSampleApp.Responses.Api.V1.Users;

@doc("ユーザー作成成功のレスポンス")
model CreateUserResponse {
  @doc("ユーザー")
  user: Models.User;
}
