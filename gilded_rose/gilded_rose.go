package gilded_rose

import (
    "github.com/Harcyah/code-kata-go/maths"
    "strings"
)

type Item struct {
	name            string
	sellIn, quality int
}

func (i *Item) setQuality(quality int) {
	i.quality = maths.Max(maths.Min(quality, 50), 0)
}

func (i *Item) decrementSellIn() {
	i.sellIn--
}

// Quality strategies

type QualityUpdaterStrategy interface {
	Update(item *Item)
}

type LegendaryItemQualityUpdaterStrategy struct {
}

func (i LegendaryItemQualityUpdaterStrategy) Update(item *Item) {
	// do nothing
}

type CheeseItemQualityUpdaterStrategy struct {
}

func (i CheeseItemQualityUpdaterStrategy) Update(item *Item) {
	if item.sellIn <= 0 {
		item.setQuality((maths.Abs(item.sellIn) + 1) * 2)
	} else {
		item.setQuality(item.quality + 1)
	}
}

type ConcertItemQualityUpdaterStrategy struct {
}

func (i ConcertItemQualityUpdaterStrategy) Update(item *Item) {
	if item.sellIn < 0 {
		item.setQuality(0)
	} else if item.sellIn <= 4 {
		item.setQuality(item.quality + 3)
	} else if item.sellIn <= 9 {
		item.setQuality(item.quality + 2)
	} else {
		item.setQuality(item.quality + 1)
	}
}

type ConjuredItemQualityUpdaterStrategy struct {
}

func (i ConjuredItemQualityUpdaterStrategy) Update(item *Item) {
	var delta int
	if item.sellIn >= 0 {
		delta = 1
	} else {
		delta = 4
	}
	item.setQuality(item.quality - delta)
}

type QualityDecreasingItemQualityUpdaterStrategy struct {
}

func (i QualityDecreasingItemQualityUpdaterStrategy) Update(item *Item) {
	var delta int
	if item.sellIn >= 0 {
		delta = 1
	} else {
		delta = 2
	}
	item.setQuality(item.quality - delta)
}

// SellIn strategies

type SellInUpdaterStrategy interface {
	Update(item *Item)
}

type NoopSellInUpdaterStrategy struct {
}

func (s NoopSellInUpdaterStrategy) Update(item *Item) {
	// Do nothing
}

type DefaultSellInUpdaterStrategy struct {
}

func (s DefaultSellInUpdaterStrategy) Update(item *Item) {
	item.decrementSellIn()
}

// Main function

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		sellInStrategy, qualityStrategy := getStrategies(item)
		sellInStrategy.Update(item)
		qualityStrategy.Update(item)
	}
}

func getStrategies(item *Item) (SellInUpdaterStrategy, QualityUpdaterStrategy) {
	var name = item.name
	if name == "Sulfuras, Hand of Ragnaros" {
		return new(NoopSellInUpdaterStrategy), new(LegendaryItemQualityUpdaterStrategy)
	}
	if name == "Aged Brie" {
		return new(DefaultSellInUpdaterStrategy), new(CheeseItemQualityUpdaterStrategy)
	}
	if name == "Backstage passes to a TAFKAL80ETC concert" {
		return new(DefaultSellInUpdaterStrategy), new(ConcertItemQualityUpdaterStrategy)
	}
	if strings.Contains(name, "Conjured") {
		return new(DefaultSellInUpdaterStrategy), new(ConjuredItemQualityUpdaterStrategy)
	}
	return new(DefaultSellInUpdaterStrategy), new(QualityDecreasingItemQualityUpdaterStrategy)
}
