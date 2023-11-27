package types

type result struct {
  Result  []team `json:"result"`
  Success int    `json:"success"`
}
type team struct {
  Name string `json:"team_name"`
  Key int `json:"team_key"`
}