package models

import (
  "github.com/jinzhu/gorm"
)

type Todo struct {
  gorm.Model
  Text string `json:"text"`
  Done bool   `json:"done"`
}
