# codex

Codex is **NOT** an ORM, but a relation algebra inspired by [Arel](http://www.github.com/rails/arel). Project is still under heavy development.

[![Build Status](https://drone.io/github.com/chuckpreslar/codex/status.png)](https://drone.io/github.com/chuckpreslar/codex/latest)

## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/codex

## Usage

To show a little sample of what using Codex looks like, lets assume you have a table in your database (we'll call it `users`) and you want to select all the records it contains.  The SQL looks a little like this:

```sql
SELECT "users".* FROM "users"
```

Now using Codex:

```go
import (
  /* Maybe some other imports. */
  "github.com/chuckpreslar/codex"
)


users := codex.Table("users")
sql, err := users.ToSql()

```

Now that wasn't too bad, was it?

## Selections

#### Projections

```go

// ...

users := codex.Table("users")
sql, err := users.Project("id", "email", "first_name", "last_name").ToSql()

```

#### Filtering

```go

// ...

users := codex.Table("users")
sql, err := users.Where(users("id").Eq(1).Or(users("email").Eq("test@example.com"))).ToSql()

```

#### Joins

```go

// ...

users := codex.Table("users")
orders := codex.Table("orders")
sql, err := users.InnerJoin(orders).On(orders("user_id").Eq(user("id"))).ToSql()

// SELECT "users".* FROM "users" INNER JOIN "orders" ON "orders"."user_id" = "users"."id"

```

## Insertions

```go
users := codex.Table("users")

sql, err := users.Insert("Jon", "Doe", "jon@example.com").
            Into("first_name", "last_name", "email").ToSql()

// OR

sql, err := users.Insert(codex.Values{
                "first_name": "Jon",
                "last_name": "Doe",
                "email": "jon@example.com"
            }).ToSql()

// INSERT INTO "users" ("first_name", "last_name", "email") VALUES ('Jon', 'Doe', 'jon@example.com')

```

## Updates

```go
users := codex.Table("users")

sql, err := users.Set("first_name", "last_name", "email").
            To("Jon", "Doe", "jon@example.com").
            Where(users("id").Eq(1)).ToSql()

// OR

sql, err := users.Set(codex.Values{
                "first_name": "Jon",
                "last_name": "Doe",
                "email": "jon@example.com"
            }).Where(users("id").Eq(1)).ToSql()

// UPDATE "users" SET "first_name" = 'Jon', "last_name" = 'Doe', "email" = 'jon@example.com'
// WHERE "users"."id" = 1

```

## Removals

```go
users := codex.Table("users")

sql, err := users.Delete(users("id").Eq(1)).ToSql()

// DELETE FROM "users" WHERE "users"."id" = 1

```

## License

> The MIT License (MIT)

> Copyright (c) 2013 Chuck Preslar

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
