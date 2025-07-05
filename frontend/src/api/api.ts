import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { Animal } from "./animalsService";

export const api = createApi({
    reducerPath: 'api',
    baseQuery: fetchBaseQuery({
        baseUrl: "http://localhost:9111/api"
    }),
    endpoints: (builder) => ({
        createAnimal: builder.mutation<Animal, {
            readonly id: number
            readonly name: string
            readonly type: string
            readonly color: string
            readonly age: number
            readonly price: number
            readonly storeAddress: string
        }>({
            query: ( body ) => ({
                url: '/animals',
                method: 'POST',
                body,
            }),
            invalidatesTags: ['Animal']
        }),
    }),
    tagTypes: ['Animal']
})

