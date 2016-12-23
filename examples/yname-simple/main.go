package main

import (
	"fmt"
	"github.com/udhos/yname"
	"gopkg.in/yaml.v2"
)

func main() {
	input := []byte(example)

	var doc interface{}
	errParse := yaml.Unmarshal(input, &doc)
	if errParse != nil {
		fmt.Printf("parse error: %v\n", errParse)
		return
	}

	show(doc, "bill-to/address/city")
	show(doc, "product/1/price")
	show(doc, "product/11/price")
	show(doc, "product/xx/price")
	show(doc, "bill-to/address")
	show(doc, "bill-to/address/galaxy")
	show(doc, "invoice/34843")
	show(doc, "issuer/John")
}

func show(d interface{}, path string) {
	child, err := yname.Get(d, path)
	if err != nil {
		fmt.Printf("path=[%s] get error: %v\n", path, err)
		return
	}
	fmt.Printf("path=[%s] FOUND:\n", path)

	buf, errMar := yaml.Marshal(child)
	if errMar != nil {
		fmt.Printf("path=[%s] show error: %v\n", path, err)
		return
	}

	fmt.Printf("%v\n", string(buf))
}

const example = `
issuer: John
invoice: 34843
date   : 2001-01-23
bill-to: &id001
    given  : Chris
    family : Dumars
    address:
        lines: |
            458 Walkman Dr.
            Suite #292
        city    : Royal Oak
        state   : MI
        postal  : 48046
ship-to: *id001
product:
    - sku         : BL394D
      quantity    : 4
      description : Basketball
      price       : 450.00
    - sku         : BL4438H
      quantity    : 1
      description : Super Hoop
      price       : 2392.00
tax  : 251.42
total: 4443.52
comments: >
    Late afternoon is best.
    Backup contact is Nancy
    Billsmer @ 338-4338.
`
