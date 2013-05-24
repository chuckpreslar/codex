![librarian](http://i.imgur.com/lvQmuIY.png)

Installation
============

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/librarian

Usage
=====

To show a little sample of what using Libraian looks like, lets assume you have a table in your database (we'll call it `users`) and you want to select all the records it contains.  The SQL looks a little like this:

```sql
SELECT "users".* FROM "users"
```

Now using Librarian:

```go
import (
  /* Maybe some other imports. */
  l "github.com/chuckpreslar/librarian"
)

// ... lets assume `session` is your database session

users := l.Table("users")
records, err := l.Search(users).Run(session).Query()
```

Now that wasn't too bad, was it?

#### Projections

You can, of course, speed up your queries by only selecting the columns you need:

```go
// ...

users := l.Table("users")
records, err := l.Search(users).
                Select("id", "email", "first_name", "last_name").
                Run(session).
                Query()
```

#### Filtering

An example of how to search for records that meet a specified criteria:

```go
// ...

users := l.Table("users")
records, err := l.Search(users).
                Where(users("id").Eq(1).Or(users("email").Eq("test@example.com"))).
                Run(Session).
                Query()
```

#### Joins

Still with me?  Last but not least, an example of a `JOIN` operation:

```go
// ...

users := l.Table("users")
orders := l.Table("orders")
records, err := l.Search(users).
                Select("id", "email", "first_name", "last_name").
                Select(orders("id").As("order_id")).
                Join(orders.On("user_id").Eq(users("id"))).
                Run(session).
                Query()
```

Notes
=====

This project is still under heavy development, a lot of work still needs to be done.  Come back in a week or so ;)

##### MIT licensed
##### Release v 0.0.10
