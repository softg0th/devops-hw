package queries

var InsertAnimalQuery = `
INSERT INTO animals(name, type, color, store_id, age, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id`

var SelectAllAnimalQuery = `
SELECT 
    a.id, 
    a.name, 
    a.type, 
    a.color, 
    a.store_id, 
    s.address AS store_address,
    a.age, 
    a.price
FROM animals a
LEFT JOIN stores s ON a.store_id = s.id
WHERE a.store_id IS NOT NULL;

`

var SelectAnimalById = `
SELECT name FROM animals WHERE id = $1
`

var DeleteAnimalQuery = `
DELETE FROM animals WHERE id = $1
RETURNING id
`

var SelectAnimalsByFilterQuery = `
SELECT 
    id, 
    name, 
    type, 
    color, 
    store_id, 
    age, 
    price 
FROM animals 
WHERE %s = $1
`

var UpdateAnimalQuery = `
UPDATE animals 
SET %s = $2 
WHERE id = $1 
RETURNING id
`
