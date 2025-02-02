import React from "react";
import { useSelector } from "react-redux";
import { deleteCustomer } from "./slices/CustomerSlice"; //action creater
import { useDispatch } from "react-redux"; // used to dispach redux 

export default function CustomerView({}){
    // console.log('customer==> ',customer);
    
    const customer = useSelector((state)=>state.customer)

    const dispatch = useDispatch();

    function deleteHandler(index){
        dispatch(deleteCustomer(index));
    }

    return<div>
        <h3>Customer List</h3>
        <ul>
            {
                customer.map((customer, index)=><li>
                    {customer}
                    <button>edit</button>
                    <button onClick={()=> deleteHandler(index)}>del</button>
                    </li>)
            }
        </ul>
    </div>
}