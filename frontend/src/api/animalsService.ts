import {api} from "./api";
import {BaseQueryArg} from "@reduxjs/toolkit/dist/query/baseQueryTypes";
import axios from "axios";

export type Animal = {
    id: number,
    name: string,
    type: string,
    color: string,
    age: number,
    price: number,
    store_address: string
}

export const animalService = api.injectEndpoints({
    endpoints: builder => ({
        getStoreAddresses: builder.query<{ id: number, address: string }[], void>({
            query: () => ({
                url: '/stores',
            }),
            transformResponse: (response: { Id: number, Address: string }[]) =>
                response.map(store => ({ id: store.Id, address: store.Address })),
        }),



        getAllAnimals: builder.mutation({
            query: ( body ) => ({
                url: '/animals',
                method: 'GET',
                body,
            }),
            invalidatesTags: ['Animal']
        }),


        createAnimal: builder.mutation<Animal, {
            Name: string,
            Type: string,
            Color: string,
            Age: number,
            Price: number,
            StoreID: number
        }>({
            query: (body) => ({
                url: '/animals',
                method: 'POST',
                body,
            }),
            invalidatesTags: ['Animal']
        }),

        deleteAnimal: builder.mutation<void, number>({
            query: (id) => ({
                url: `/animals?id=${id}`,
                method: 'DELETE',
            }),
            invalidatesTags: ['Animal']
        }),

        updateAnimal: builder.mutation<Animal, {
            animalId: number,
            name: string,
            type: string,
            color: string,
            age: number,
            price: number,
            storeID: number
        }>({
            query: ({ animalId, ...body }) => ({
                url: `/animals`,
                method: 'PUT',
                body: {
                    id: animalId, // добавляем ID животного
                    ...body
                },
            }),
            invalidatesTags: ['Animal']
        }),


        getAnimals: builder.query<Animal[], void>({
            query: () => ({
                url: '/animals',
            }),
            providesTags: ['Animal']
        })
    })
})
