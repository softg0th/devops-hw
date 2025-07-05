package queries

var SelectAllStoreQuery = `
SELECT id, address FROM stores WHERE id IS NOT NULL
`

var InsertStoreQurey = `
INSERT INTO stores(name, address)
VALUES ($1, $2)
RETURNING id`

var SelectStoreById = `
SELECT id FROM stores WHERE id = $1
`

var DeleteStoreQurey = `
DELETE FROM stores WHERE id = $1 RETURNING id;
`
