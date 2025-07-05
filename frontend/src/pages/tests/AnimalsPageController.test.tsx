import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Provider } from "react-redux";
import { store } from "../../api/store";
import { animalService } from "../../api/animalsService";
import { CreateAnimalForm, AnimalsList } from "../AnimalsPageController";

jest.mock("../../api/animalsService", () => ({
  animalService: {
    useGetStoreAddressesQuery: jest.fn(),
    useCreateAnimalMutation: jest.fn(),
    useGetAnimalsQuery: jest.fn(),
    useDeleteAnimalMutation: jest.fn(),
    useUpdateAnimalMutation: jest.fn(),
  },
}));

beforeEach(() => {
  jest.clearAllMocks();
  jest.spyOn(window, "alert").mockImplementation(() => {});
  jest.spyOn(window, "confirm").mockReturnValue(true);
});

describe("CreateAnimalForm", () => {
  test("рендер формы", () => {
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [{ id: 1, address: "Test Address" }],
      isLoading: false,
    });

    render(
      <Provider store={store}>
        <CreateAnimalForm />
      </Provider>
    );

    expect(screen.getByText("Add New Animal")).toBeInTheDocument();
    expect(screen.getByLabelText("Name:")).toBeInTheDocument();
  });

  test("отправка формы с валидными данными", async () => {
    const createMock = jest.fn();
    (animalService.useCreateAnimalMutation as jest.Mock).mockReturnValue([createMock, { isLoading: false }]);
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [{ id: 1, address: "Test Address" }],
      isLoading: false,
    });

    render(
      <Provider store={store}>
        <CreateAnimalForm />
      </Provider>
    );

    await userEvent.type(screen.getByLabelText("Name:"), "Lion");
    await userEvent.type(screen.getByLabelText("Type:"), "Cat");
    await userEvent.type(screen.getByLabelText("Color:"), "Golden");
    await userEvent.type(screen.getByLabelText("Age:"), "3");
    await userEvent.type(screen.getByLabelText("Price:"), "999");
    await userEvent.selectOptions(screen.getByLabelText("Store:"), "1");

    fireEvent.click(screen.getByText("Create"));

    expect(createMock).toHaveBeenCalledWith({
      Name: "Lion",
      Type: "Cat",
      Color: "Golden",
      Age: 3,
      Price: 999,
      StoreID: 1,
    });
  });

  test("отказ при незаполненном магазине", () => {
    const createMock = jest.fn();
    (animalService.useCreateAnimalMutation as jest.Mock).mockReturnValue([createMock, { isLoading: false }]);
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [],
      isLoading: false,
    });

    render(
      <Provider store={store}>
        <CreateAnimalForm />
      </Provider>
    );

    fireEvent.click(screen.getByText("Create"));
    expect(window.alert).toHaveBeenCalledWith("Please select a store");
    expect(createMock).not.toHaveBeenCalled();
  });
});

describe("AnimalsList", () => {
  const mockAnimals = [
    {
      id: 1,
      name: "Tiger",
      type: "Big Cat",
      color: "Orange",
      age: 5,
      price: 2000,
      store_address: "Zoo Store",
    },
    {
      id: 2,
      name: "Parrot",
      type: "Bird",
      color: "Green",
      age: 2,
      price: 300,
      store_address: "Bird Shop",
    },
  ];

  beforeEach(() => {
    (animalService.useGetAnimalsQuery as jest.Mock).mockReturnValue({
      data: mockAnimals,
      isLoading: false,
    });

    (animalService.useDeleteAnimalMutation as jest.Mock).mockReturnValue([jest.fn(), { isLoading: false }]);
    (animalService.useUpdateAnimalMutation as jest.Mock).mockReturnValue([jest.fn(), { isLoading: false }]);
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [{ id: 1, address: "Zoo Store" }],
      isLoading: false,
    });
  });

  test("отображает список животных", () => {
    render(
      <Provider store={store}>
        <AnimalsList />
      </Provider>
    );

    expect(screen.getByText("Tiger")).toBeInTheDocument();
    expect(screen.getByText("Parrot")).toBeInTheDocument();
  });

  test("удаляет животное по нажатию Delete", async () => {
    const deleteMock = jest.fn();
    (animalService.useDeleteAnimalMutation as jest.Mock).mockReturnValue([deleteMock, { isLoading: false }]);

    render(
      <Provider store={store}>
        <AnimalsList />
      </Provider>
    );

    fireEvent.click(screen.getAllByText("Delete")[0]);
    await waitFor(() => expect(deleteMock).toHaveBeenCalledWith(1));
  });

  test("не удаляет животное, если confirm = false", async () => {
    const deleteMock = jest.fn();
    (animalService.useDeleteAnimalMutation as jest.Mock).mockReturnValue([deleteMock, { isLoading: false }]);
    (window.confirm as jest.Mock).mockReturnValue(false);

    render(
      <Provider store={store}>
        <AnimalsList />
      </Provider>
    );

    fireEvent.click(screen.getAllByText("Delete")[0]);
    await waitFor(() => expect(deleteMock).not.toHaveBeenCalled());
  });

  test("открывает форму редактирования", async () => {
    render(
      <Provider store={store}>
        <AnimalsList />
      </Provider>
    );

    fireEvent.click(screen.getAllByText("Edit")[0]);

    expect(await screen.findByText("Edit Animal")).toBeInTheDocument();
    expect(screen.getByDisplayValue("Tiger")).toBeInTheDocument();
  });
});