import dataclasses
import os

import requests
import pytest
import psycopg2

API = 'http://localhost:9111'
DB_URL = os.getenv("DB_URL", "postgresql://postgres:12345@localhost:5432/warehouse")

@dataclasses.dataclass
class AniStore:
    animal_id: int
    store_id: int

ani_store = AniStore(0, 0)

def test_api_availability():
    response = requests.get(API)
    assert response.status_code == 404

@pytest.fixture(scope="session")
def db_session():
    conn = psycopg2.connect(DB_URL)
    cursor = conn.cursor()
    cursor.execute('TRUNCATE TABLE animals, stores RESTART IDENTITY CASCADE;')
    conn.commit()

    yield conn
    cursor.execute('TRUNCATE TABLE animals, stores RESTART IDENTITY CASCADE;')
    conn.commit()
    cursor.close()
    conn.close()

@pytest.fixture
def test_db_availability(db_session):
    return db_session

def test_insert_store(test_db_availability):
    conn = psycopg2.connect(DB_URL)
    cursor = conn.cursor()
    cursor.execute(
        "INSERT INTO stores (name, address) VALUES (%s, %s) RETURNING id;",
        ('test_store', 'pushkina')
    )
    conn.commit()
    store_id = cursor.fetchone()[0]
    assert type(store_id) == int
    ani_store.store_id = store_id

    cursor.execute(
        "INSERT INTO animals (name, type, color, store_id, age, price) VALUES (%s, %s, %s, %s, %s, %s) RETURNING id;",
        ('test_cat', 'cat', 'black', store_id, 1, 100)
    )
    conn.commit()
    animal_id = cursor.fetchone()[0]
    assert type(animal_id) == int
    ani_store.animal_id = animal_id

def test_stores_api():
    response = requests.get(f'{API}/api/stores')
    assert response.status_code == 200
    response = requests.post(f'{API}/api/stores', json={
        'Name': 'test_store_2',
        'Address': 'Kolotushkina'
    })
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 201
    
    temp_store_id = response.json()['id']
    response = requests.delete(f'{API}/api/stores?id={temp_store_id}')
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 200


def test_animals_api():
    response = requests.post(f'{API}/api/animals', json={
        'Name': 'test_dog',
        'Type': 'dog',
        'Color': 'white',
        'StoreID': ani_store.store_id,
        'Age': 2,
        'Price': 100
    })
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 201

    response = requests.get(f'{API}/api/animals')
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 200

    response = requests.put(f'{API}/api/animals', json={
        'ID': ani_store.animal_id,
        'Name': 'test_hamster',
        'Type': 'hamster',
        'Color': 'orange',
        'StoreID': ani_store.store_id,
        'Age': 2,
        'Price': 100
    })
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 200
    

    response = requests.delete(f'{API}/api/animals?id={ani_store.animal_id}')
    print("RESPONSE STATUS:", response.status_code)
    print("RESPONSE BODY:", response.text)
    assert response.status_code == 200
