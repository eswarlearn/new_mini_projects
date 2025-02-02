import { configureStore } from "@reduxjs/toolkit";
import customerReducer from "./slices/CustomerSlice"

export const store = configureStore ({
    devTools:false,
    reducer :{
        customer : customerReducer
    }
})