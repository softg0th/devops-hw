import { render, screen, fireEvent } from "@testing-library/react";
import { EditAnimalForm } from "../AnimalsPageController";
import { animalService } from "../../api/animalsService";
import { Provider } from "react-redux";
import { store } from "../../api/store";

jest.mock("../../api/animalsService", () => ({
  animalService: {
    useUpdateAnimalMutation: jest.fn(),
    useGetStoreAddressesQuery: jest.fn(),
  },
}));

describe("EditAnimalForm", () => {
  const mockAnimal = {
    id: 1,
    name: "Tiger",
    type: "Big Cat",
    color: "Orange",
    age: 5,
    price: 2000,
    store_address: "Zoo Store",
  };

  beforeEach(() => {
    (animalService.useUpdateAnimalMutation as jest.Mock).mockReturnValue([jest.fn(), { isLoading: false }]);
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [{ id: 1, address: "Zoo Store" }],
      isLoading: false,
    });
    jest.spyOn(window, "alert").mockImplementation(() => {});
  });

  test("рендерит форму с начальными значениями", () => {
    render(
      <Provider store={store}>
        <EditAnimalForm animal={mockAnimal} onClose={jest.fn()} />
      </Provider>
    );

    expect(screen.getByDisplayValue("Tiger")).toBeInTheDocument();
    expect(screen.getByDisplayValue("Big Cat")).toBeInTheDocument();
    expect(screen.getByDisplayValue("Orange")).toBeInTheDocument();
    expect(screen.getByDisplayValue("5")).toBeInTheDocument();
    expect(screen.getByDisplayValue("2000")).toBeInTheDocument();
  });

  test("не отправляет update, если не выбран магазин", () => {
    const updateMock = jest.fn();
    const onClose = jest.fn();

    (animalService.useUpdateAnimalMutation as jest.Mock).mockReturnValue([updateMock, { isLoading: false }]);
    (animalService.useGetStoreAddressesQuery as jest.Mock).mockReturnValue({
      data: [],
      isLoading: false,
    });

    render(
      <Provider store={store}>
        <EditAnimalForm animal={{ ...mockAnimal, store_address: "" }} onClose={onClose} />
      </Provider>
    );

    fireEvent.click(screen.getByText("Update"));
    expect(updateMock).not.toHaveBeenCalled();
    expect(window.alert).toHaveBeenCalledWith("Please select a store");
  });

  test("вызов update и onClose при валидных данных", () => {
    const updateMock = jest.fn().mockResolvedValue({});
    const onClose = jest.fn();

    (animalService.useUpdateAnimalMutation as jest.Mock).mockReturnValue([updateMock, { isLoading: false }]);

    render(
      <Provider store={store}>
        <EditAnimalForm animal={mockAnimal} onClose={onClose} />
      </Provider>
    );

    fireEvent.change(screen.getByLabelText("Name:"), { target: { value: "Lion" } });
    fireEvent.click(screen.getByText("Update"));

    expect(updateMock).toHaveBeenCalledWith({
      animalId: 1,
      name: "Lion",
      type: "Big Cat",
      color: "Orange",
      age: 5,
      price: 2000,
      storeID: 1,
    });
  });

  test("вызов onClose при cancel", () => {
    const onClose = jest.fn();

    render(
      <Provider store={store}>
        <EditAnimalForm animal={mockAnimal} onClose={onClose} />
      </Provider>
    );

    fireEvent.click(screen.getByText("Cancel"));
    expect(onClose).toHaveBeenCalled();
  });
});
