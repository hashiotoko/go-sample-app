import "./index.tsp";

namespace GoSampleApp.Models;

@example(#{ id: 1, name: "田中太郎" })
model User {
  @doc("ユーザーID")
  id: int32;

  @doc("ユーザー名")
  name: string;
}
