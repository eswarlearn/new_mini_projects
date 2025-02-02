import { createSlice } from "@reduxjs/toolkit";

const initialState =[];

const customerSlice = createSlice({
    name: 'customer',
    initialState,
    reducers:{
        addCustomer(state, action){
            state.push(action.payload)
        },
        deleteCustomer(st,act){
            const deleteIndex = act.payload;
            return st.filter((val,index)=> index !== deleteIndex);
        }
    }
})

export const {addCustomer, deleteCustomer} = customerSlice.actions
export default customerSlice.reducer