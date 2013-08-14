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
sql, err := users.InnerJoin(orders).On(orders("user_id").Eq(users("id"))).ToSql()

// SELECT "users".* FROM "users" INNER JOIN "orders" ON "orders"."user_id" = "users"."id"
```

## Insertions

```go
// ...

sql, err := users.Insert("Jon", "Doe", "jon@example.com").
    Into("first_name", "last_name", "email").ToSql()

// INSERT INTO "users" ("first_name", "last_name", "email") VALUES ('Jon', 'Doe', 'jon@example.com')
```

## Modifications

```go
// ...

sql, err := users.Set("first_name", "last_name", "email").
    To("Jon", "Doe", "jon@example.com").
    Where(users("id").Eq(1)).ToSql()

// UPDATE "users" SET "first_name" = 'Jon', "last_name" = 'Doe', "email" = 'jon@example.com'
// WHERE "users"."id" = 1
```

## Deletions

```go
// ...

sql, err := users.Delete(users("id").Eq(1)).ToSql()

// DELETE FROM "users" WHERE "users"."id" = 1
```

## Creations

The codex package currently provides a few common SQL data and constraint types as constants to use with creating and altering tables.

__Constraints__

* NOT_NULL
* UNIQUE
* PRIMARY_KEY
* FOREIGN_KEY
* CHECK
* DEFAULT

__Types__

* STRING
* TEXT
* BOOLEAN
* INTEGER
* FLOAT
* DECIMAL
* DATE
* TIME
* DATETIME
* TIMESTAMP

```go
// ...

users := codex.CreateTable("users").
  AddColumn("first_name", codex.STRING).
  AddColumn("last_name", codex.STRING).
  AddColumn("email", codex.STRING).
  AddColumn("id", codex.INTEGER).
  AddConstraint("first_name", codex.NOT_NULL).
  AddConstraint("id", codex.PRIMARY_KEY).
  AddConstraint("email", codex.UNIQUE, "users_uniq_email") // Optional last argument supplies index name.

sql, err := users.ToSql()

// CREATE TABLE "users" ();
// ALTER TABLE "users" ADD "first_name" character varying(255);
// ALTER TABLE "users" ADD "last_name" character varying(255);
// ALTER TABLE "users" ADD "email" character varying(255);
// ALTER TABLE "users" ADD "id" integer;
// ALTER TABLE "users" ALTER "first_name" SET NOT NULL;
// ALTER TABLE "users" ADD PRIMARY KEY("id");
// ALTER TABLE "users" ADD CONSTRAINT "users_email_uniq" UNIQUE("email");
```

## Alterations

Alterations work the same as Creations, only the the initializer function `AlterTable` is used in place of `CreateTable`.

## Documentation

View godoc or visit [godoc.org](http://godoc.org/github.com/chuckpreslar/codex).

    $ godoc codex

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
