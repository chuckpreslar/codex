# codex

Codex is **NOT** an ORM, but a relation algebra inspired by [Arel](http://www.github.com/rails/arel). Project is still under heavy development.

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
sql := users.ToSql()

```

Now that wasn't too bad, was it?

### Selections

#### Projections

You can, of course, speed up your queries by only selecting the columns you need:

```go

// ...

users := codex.Table("users")
sql := users.Project("id", "email", "first_name", "last_name").ToSql()

```

#### Filtering

An example of how to search for records that meet a specified criteria:

```go

// ...

users := codex.Table("users")
sql := users.Where(users("id").Eq(1).Or(users("email").Eq("test@example.com"))).ToSql()

```

#### Joins

```go

// ...

users := codex.Table("users")
orders := codex.Table("orders")
sql := users.InnerJoin(orders).On(users("order_id").Eq(orders("id"))).ToSql()

// SELECT "users".* FROM "users" INNER JOIN "orders" ON "orders"."id" = "users"."order_id"

```

### Insertions

```go
users := codex.Table("users")

sql := users.Insert("Jon", "Doe", "jon@example.com").Into("first_name", "last_name", "email").ToSql()

// OR

sql := users.Insert(codex.Values{ "first_name": "Jon", "last_name": "Doe", "email": "jon@example.com"}).ToSql()

// INSERT INTO "users" ("first_name", "last_name", "email") VALUES ('Jon', 'Doe', 'jon@example.com')
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
