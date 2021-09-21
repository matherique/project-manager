package cmd

type Cmd interface {
  Exec(a []string)
}

