package gilded_rose

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func items() []*Item {
	return []*Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}
}

func TestUpdateQualityDay1(t *testing.T) {
	var items = items()

	UpdateQuality(items)

	verifyItem(t, items[0], 9, 19)
	verifyItem(t, items[1], 1, 1)
	verifyItem(t, items[2], 4, 6)
	verifyItem(t, items[3], 0, 80)
	verifyItem(t, items[4], -1, 80)
	verifyItem(t, items[5], 14, 21)
	verifyItem(t, items[6], 9, 50)
	verifyItem(t, items[7], 4, 50)
	verifyItem(t, items[8], 2, 5)
}

func TestUpdateQualityDay2(t *testing.T) {
	var items = items()

	UpdateQuality(items)
	UpdateQuality(items)

	verifyItem(t, items[0], 8, 18)
	verifyItem(t, items[1], 0, 2)
	verifyItem(t, items[2], 3, 5)
	verifyItem(t, items[3], 0, 80)
	verifyItem(t, items[4], -1, 80)
	verifyItem(t, items[5], 13, 22)
	verifyItem(t, items[6], 8, 50)
	verifyItem(t, items[7], 3, 50)
	verifyItem(t, items[8], 1, 4)
}

func TestUpdateQualityDay3(t *testing.T) {
	var items = items()

	UpdateQuality(items)
	UpdateQuality(items)
	UpdateQuality(items)

	verifyItem(t, items[0], 7, 17)
	verifyItem(t, items[1], -1, 4)
	verifyItem(t, items[2], 2, 4)
	verifyItem(t, items[3], 0, 80)
	verifyItem(t, items[4], -1, 80)
	verifyItem(t, items[5], 12, 23)
	verifyItem(t, items[6], 7, 50)
	verifyItem(t, items[7], 2, 50)
	verifyItem(t, items[8], 0, 3)
}

func TestUpdateQualityDay12(t *testing.T) {
	var items = items()

	for i := 0; i < 12; i++ {
		UpdateQuality(items)
	}

	verifyItem(t, items[0], -2, 6)
	verifyItem(t, items[1], -10, 22)
	verifyItem(t, items[2], -7, 0)
	verifyItem(t, items[3], 0, 80)
	verifyItem(t, items[4], -1, 80)
	verifyItem(t, items[5], 3, 41)
	verifyItem(t, items[6], -2, 0)
	verifyItem(t, items[7], -7, 0)
	verifyItem(t, items[8], -9, 0)
}

func TestUpdateQualityDay30(t *testing.T) {
	var items = items()

	for i := 0; i < 30; i++ {
		UpdateQuality(items)
	}

	verifyItem(t, items[0], -20, 0)
	verifyItem(t, items[1], -28, 50)
	verifyItem(t, items[2], -25, 0)
	verifyItem(t, items[3], 0, 80)
	verifyItem(t, items[4], -1, 80)
	verifyItem(t, items[5], -15, 0)
	verifyItem(t, items[6], -20, 0)
	verifyItem(t, items[7], -25, 0)
	verifyItem(t, items[8], -27, 0)
}

func TestNewImplRespectsOldImpl(t *testing.T) {
	var newItems = items()
	var oldItems = items()

	for day := 1; day <= 30; day++ {
		UpdateQuality(newItems)
		UpdateQualityOriginal(oldItems)

		// 8 only, because we don't compare conjured items : old impl doesn't handle them
		for j := 0; j < 8; j++ {
			var newItem = newItems[j]
			var oldItem = oldItems[j]
			assert.Equal(t, oldItem.sellIn, newItem.sellIn, "Expected sellIn of item %d, day %d, to equal %d but %d was found", j, day, oldItem.sellIn, newItem.sellIn)
			assert.Equal(t, oldItem.quality, newItem.quality, "Expected quality of item %d, day %d, to equal %d but %d was found", j, day, oldItem.quality, newItem.quality)
		}
	}
}

func verifyItem(t *testing.T, item *Item, expectedSellIn int, expectedQuality int) {
	assert.Equal(t, expectedSellIn, item.sellIn, "Expected sellIn of %s to equal %d but %d was found", item.name, expectedSellIn, item.sellIn)
	assert.Equal(t, expectedQuality, item.quality, "Expected quality of %s to equal %d but %d was found", item.name, expectedQuality, item.quality)
}
