workflow "Test" {
  on = "push"
  resolves = ["test"]
}

action "fmt" {
  uses = "./.github/actions/golang"
  args = "fmt"
}

action "test" {
  needs = ["fmt"]
  uses = "./.github/actions/golang"
  args = "test"
}
