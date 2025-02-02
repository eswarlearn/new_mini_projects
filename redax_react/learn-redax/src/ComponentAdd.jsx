import React, { useState } from "react";
import CustomerView from "./CustomerView";
import { addCustomer as addCustomerAction } from "./slices/CustomerSlice";
import { useDispatch } from "react-redux";

export default function CustomerAdd() {
    const[input, setInput] = useState("");
    // const [customer, setCustomer] = useState([])
    const dispatch = useDispatch();

    function addCustomer(){
        if(input){
            // setCustomer((preSt)=>[...preSt,input]) //add old data to new
            // console.log(customer);
            dispatch(addCustomerAction(input))
            setInput('');
            
        }
    }

    return <><div>
    <h3> Add New Customer</h3>
    <input type="text" value={input} onChange={(e)=>setInput(e.target.value)}/>
    <button onClick={addCustomer}>add</button>
</div>
{/* <CustomerView customer={customer}/> props passing */}
</>

}