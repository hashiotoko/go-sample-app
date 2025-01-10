schema "go_sample_app" {
    charset = "utf8mb4"
    collate = "utf8mb4_0900_ai_ci"
}

table "users" {
    schema  = schema.go_sample_app
    collate = "utf8mb4_unicode_ci"

    column "id" {
        null    = false
        type    = varchar(50)
        unsigned = true
        comment = "ユーザーID"
    }

    column "name" {
        null    = false
        type    = varchar(255)
        comment = "ユーザー名前"
    }

    column "created_at" {
        null    = false
        type    = timestamp
        default = sql("CURRENT_TIMESTAMP")
        comment = "作成日時"
    }
    column "updated_at" {
        null      = false
        type      = timestamp
        default   = sql("CURRENT_TIMESTAMP")
        comment   = "更新日時"
        on_update = sql("CURRENT_TIMESTAMP")
    }

    primary_key {
        columns = [column.id]
    }
}
