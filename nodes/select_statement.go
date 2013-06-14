package nodes

type SelectStatement struct {
  cores  []SelectCore
  limit  int
  offset int
}
