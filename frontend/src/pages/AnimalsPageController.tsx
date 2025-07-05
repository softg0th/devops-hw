import {useState} from "react";
import {Animal, animalService} from "../api/animalsService";
import './page.css'

export const AnimalsPageContainer = () => {
    return (
        <div className="content">
            <CreateAnimalForm />
            <AnimalsList />
        </div>
    );
};

export const CreateAnimalForm = () => {
    const [Name, setName] = useState('');
    const [Type, setType] = useState('');
    const [Color, setColor] = useState('');
    const [Age, setAge] = useState(0);
    const [Price, setPrice] = useState(0);
    const [StoreID, setStoreId] = useState<number | ''>('');

    const storeAddressesResult = animalService.useGetStoreAddressesQuery?.();
    const storeAddresses = storeAddressesResult?.data ?? [];
    console.log('storeAddresses:', storeAddresses);

    const result = animalService.useCreateAnimalMutation?.();
    const [createAnimal] = Array.isArray(result) ? result : [() => {}, {}];


    const onCreate = () => {
        if (StoreID === '') {
            alert("Please select a store");
            return;
        }

        createAnimal({
            Name,
            Type,
            Color,
            Age,
            Price,
            StoreID: Number(StoreID)
        });
    };

    return (
        <div className="form-wrapper">
            <div className="form-title">
                <h2>Add New Animal</h2>
            </div>
            <div className="form-content">
                <div className="form-field">
                    <label htmlFor="name">Name:</label>
                    <input id="name" type="text" value={Name} onChange={e => setName(e.target.value)}/>
                </div>
                <div className="form-field">
                    <label htmlFor="type">Type:</label>
                    <input id="type" value={Type} onChange={e => setType(e.target.value)}/>
                </div>
                <div className="form-field">
                    <label htmlFor="color">Color:</label>
                    <input id="color" value={Color} onChange={e => setColor(e.target.value)}/>
                </div>
                <div className="form-field">
                    <label htmlFor="age">Age:</label>
                    <input id="age" type="number" min="0" value={Age} onChange={e => setAge(Number(e.target.value))}/>
                </div>
                <div className="form-field">
                    <label htmlFor="price">Price:</label>
                    <input id="price" type="number" min="0" value={Price} onChange={e => setPrice(Number(e.target.value))}/>
                </div>
                <div className="form-field">
                    <label htmlFor="storeAddress">Store:</label>
                    <select
                        id="storeAddress"
                        className="select-address"
                        value={StoreID}
                        onChange={(e) => setStoreId(Number(e.target.value))}
                    >
                        <option value="">Select Store</option>
                        {storeAddresses?.map(({ id, address }) => (
                            <option key={id} value={id}>
                                {address}
                            </option>
                        ))}
                    </select>
                </div>
            </div>
            <button className="form-button" onClick={onCreate}>Create</button>
        </div>
    );
};



const EditAnimalForm = ({ animal, onClose }: { animal: Animal; onClose: () => void }) => {
    const [name, setName] = useState(animal.name);
    const [type, setType] = useState(animal.type);
    const [color, setColor] = useState(animal.color);
    const [age, setAge] = useState(animal.age);
    const [price, setPrice] = useState(animal.price);

    // Храним ID магазина вместо адреса
    const { data: storeAddresses } = animalService.useGetStoreAddressesQuery() || {};
    const initialStore = storeAddresses?.find(s => s.address === animal.store_address)?.id || '';
    const [storeID, setStoreID] = useState<number | ''>(initialStore);

    const [updateAnimal] = animalService.useUpdateAnimalMutation();

    const onUpdate = async () => {
        if (storeID === '') {
            alert("Please select a store");
            return;
        }

        await updateAnimal({
            animalId: animal.id,
            name,
            type,
            color,
            age,
            price,
            storeID: Number(storeID),
        });

        onClose();
    };

    return (
        <div className="form-wrapper">
            <div className="form-title">
                <h3>Edit Animal</h3>
            </div>

            <div className="form-field">
                <label htmlFor="edit-name">Name:</label>
                <input id="edit-name" value={name} onChange={(e) => setName(e.target.value)} />
            </div>

            <div className="form-field">
                <label htmlFor="edit-type">Type:</label>
                <input id="edit-type" value={type} onChange={(e) => setType(e.target.value)} />
            </div>

            <div className="form-field">
                <label htmlFor="edit-color">Color:</label>
                <input id="edit-color" value={color} onChange={(e) => setColor(e.target.value)} />
            </div>

            <div className="form-field">
                <label htmlFor="edit-age">Age:</label>
                <input id="edit-age" type="number" min="0" value={age} onChange={(e) => setAge(Number(e.target.value))} />
            </div>

            <div className="form-field">
                <label htmlFor="edit-price">Price:</label>
                <input id="edit-price" type="number" min="0" value={price} onChange={(e) => setPrice(Number(e.target.value))} />
            </div>

            <div className="form-field">
                <label htmlFor="edit-storeID">Store:</label>
                <select
                    id="edit-storeID"
                    className="select-address"
                    value={storeID}
                    onChange={(e) => setStoreID(Number(e.target.value))}
                >
                    <option value="">Select Store</option>
                    {storeAddresses?.map(({ id, address }) => (
                        <option key={id} value={id}>
                            {address}
                        </option>
                    ))}
                </select>
            </div>

            <div className="button-group">
                <button className="update-button" onClick={onUpdate}>Update</button>
                <button className="delete-button" onClick={onClose}>Cancel</button>
            </div>
        </div>
    );
};



export const AnimalsList = () => {
    const {data: animals} = animalService.useGetAnimalsQuery();
            const [selectedAnimal, setSelectedAnimal] = useState<Animal | null>(null);
            const [deleteAnimal] = animalService.useDeleteAnimalMutation();

            const handleDelete = async (id: number) => {
            if (window.confirm("Are you sure you want to delete this animal?")) {
            await deleteAnimal(id);
        }
        };

            return (
            <div className="sub-table-wrapper">
        {selectedAnimal && <EditAnimalForm animal={selectedAnimal} onClose={() => setSelectedAnimal(null)}/>
}

    <table className="sub-table">
                <thead>
                <tr>
                    <th>Animal Id</th>
                    <th>Name</th>
                    <th>Type</th>
                    <th>Color</th>
                    <th>Age</th>
                    <th>Price</th>
                    <th>Store Address</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {animals?.map(a => (
                    <tr key={a.id}>
                        <td>{a.id}</td>
                        <td>{a.name}</td>
                        <td>{a.type}</td>
                        <td>{a.color}</td>
                        <td>{a.age}</td>
                        <td>{a.price}</td>
                        <td>{a.store_address}</td>
                        <td>
                            <button className="update-button" onClick={() => setSelectedAnimal(a)}>Edit</button>
                            <button className="delete-button" onClick={() => handleDelete(a.id)}
                                    style={{marginLeft: "10px"}}>
                                Delete
                            </button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};


export { EditAnimalForm };