import React, { useState } from "react";
import CustomerView from "./CustomerView";
import { addCustomer as addCustomerAction } from "./slices/CustomerSlice";
import { useDispatch } from "react-redux";

export default function CustomerAdd() {
    const [inputs, setInputs] = useState(["", ""]); // Start with two input fields
    const dispatch = useDispatch();

    function handleInputChange(index, value) {
        const newInputs = [...inputs];
        newInputs[index] = value;
        setInputs(newInputs);

        // Add a new input field only if ALL existing fields are filled
        if (newInputs.every(input => input.trim() !== "")) {
            setInputs([...newInputs, ""]); // Append an empty input
        }
    }

    function addAllCustomers() {
        const filledInputs = inputs.filter(input => input.trim() !== "");
        if (filledInputs.length === 0) return; // Prevent adding empty fields

        filledInputs.forEach(input => dispatch(addCustomerAction(input)));

        // Reset to two empty input fields
        setInputs(["", ""]);
    }

    return (
        <div>
            <h3>Add New Customer</h3>
            {inputs.map((input, index) => (
                <div key={index}>
                    <input
                        type="text"
                        value={input}
                        onChange={(e) => handleInputChange(index, e.target.value)}
                    />
                </div>
            ))}
            <button onClick={addAllCustomers}>Add All</button>
        </div>
    );
}
