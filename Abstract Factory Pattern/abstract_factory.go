package main

import "abstract-factory-pattern/factory"

type AbstractFactory struct{}

func (af *AbstractFactory) CreateFactory(factoryType string) factory.Factory {
	switch factoryType {
	case "Luxury":
		return &factory.LuxuryFactory{}
	case "Ordinary":
		return &factory.OrdinaryFactory{}
	default:
		return nil
	}
}
