package postgres

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

func (r *Repository) UserInventory(ctx context.Context, userID string) ([]dto.InventoryItem, error) {
	const sql = `
		with purchasesGroup as (
			select p.merch_id, sum(p.quantity) as quantity from purchases p 
			where p.user_id = $1
			group by p.merch_id
		)
		select m.name, p.quantity from purchasesGroup p
		left join merch m on m.id = p.merch_id
	`
	rows, err := r.pool.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	var inventory []dto.InventoryItem
	for rows.Next() {
		var inventoryItem dto.InventoryItem
		if err = rows.Scan(&inventoryItem.Name, &inventoryItem.Quantity); err != nil {
			return nil, err
		}
		inventory = append(inventory, inventoryItem)
	}

	return inventory, nil
}
