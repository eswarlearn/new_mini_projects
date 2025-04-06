import React, { useState, useEffect, useContext, useRef } from 'react'
import { IoMdClose } from "react-icons/io";
import { DataContext } from './DataContext';
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import axios from "axios";

// Extend Day.js with the plugins
dayjs.extend(utc);
dayjs.extend(timezone);

const AddTarget = () => {
    var modalRef = useRef(null)
    var inputRef = useRef(null);
    var { addTargets, setAddTargets, setTargets, userId } = useContext(DataContext);
    const [inputs, setInputs] = useState([
        { target_url: "", tag: [], label: "" },
        { target_url: "", tag: [], label: "" }
    ]);
    const [contextUserId, setContextUserId] = useState('');
    const user_id = localStorage.getItem("user_id");


    // console.log("User ID in other component: out of useState", userId);

    // useEffect(() => {
    //     console.log("User ID in other component:", userId);
    //     setContextUserId(userId)
    //   }, [userId]); 

    //   console.log(useContext(DataContext));


    useEffect(() => {

        const handleClickOutside = (event) => {
            if (modalRef.current && !modalRef.current.contains(event.target)) {
                setAddTargets(true);
            }
        };
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);



    function handleInputChange(index, value) {
        const newInputs = [...inputs];
        newInputs[index] = { ...newInputs[index], target_url: value }; // Ensure it's an object
        setInputs(newInputs);

        // Add a new input field only if all existing fields are filled
        if (newInputs.every(input => input.target_url.trim() !== "")) {
            setInputs([...newInputs, { target_url: "", tag: [], label: "" }]); // Append new object
        }
    }

    const handleTagKeyDown = (index, event) => {
        if (event.key === "Enter" && event.target.value.trim() !== "") {
            const newTag = event.target.value.trim();
            const updatedInputs = [...inputs];
            if (!updatedInputs[index].tag) {
                updatedInputs[index].tag = [];
            }
            if (updatedInputs[index].tag.length < 5) {
                updatedInputs[index].tag = [...updatedInputs[index].tag, newTag];
                setInputs(updatedInputs);
                event.target.value = ""; // Clear input after adding tag
            } else {
                alert("You can add up to 5 tags only.");
            }
        }
    };

    const handleLabelChange = (index, value) => {
        const updatedInputs = [...inputs];
        updatedInputs[index].label = value;
        setInputs(updatedInputs);
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        const newYorkTime = dayjs().tz("America/New_York").format('YYYY-MM-DD HH:mm:ss');


        const dataToSend = inputs
            .filter((input) => input.target_url.trim() !== "")
            .map((input) => ({
                target_url: input.target_url,
                tag: input.tag,
                label: input.label,
                source: "manual",
                target_created_date: newYorkTime,
                user_id: user_id,
            }));

        try {
            const token = axios.defaults.headers.common["Authorization"];
            const response = await fetch('http://localhost:5000/api/targets', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': token,
                },
                body: JSON.stringify(dataToSend),
                credentials: 'include',  // This is equivalent to withCredentials: true
            });


            if (response.ok) {
                // Handle success (e.g., show a success message or reset form)
                console.log('Form submitted successfully!');
                setTargets((prevTargets) => [...dataToSend, ...prevTargets]);
                setInputs([
                    { target_url: "", tag: [], label: "" },
                    { target_url: "", tag: [], label: "" }
                ]);
                setAddTargets(true);
            } else {
                // Handle 500 or other server errors
                let errorMessage = 'Error submitting form'; // Default message

                try {
                    const errorData = await response.json(); // Try to parse the error message
                    if (errorData.message) {
                        errorMessage = errorData.message; // Use server message if available
                    }
                } catch (parseError) {
                    console.error('Error parsing error response:', parseError);
                }

                console.error(errorMessage);
                alert(errorMessage); // Show the error message to the user
            }
        } catch (error) {
            console.error('Network Error:', error);
            alert(error.message || 'Network Error'); // Show network error if the server doesn't respond
        }
    };


    // Add a new tag when user presses Enter
    const handleKeyDown = (index, e) => {
        if ((e.key === "Enter" || e.key === ",") && e.target.value.trim() !== "") {
            e.preventDefault();
            const newTag = e.target.value.trim();

            setInputs((prevInputs) =>
                prevInputs.map((input, i) =>
                    i === index && input.tag.length < 5
                        ? { ...input, tag: [...input.tag, newTag] }
                        : input
                )
            );

            e.target.value = ""; // Clear input field
        }
    };

    const handleBlur = (index, e) => {
        const newTag = e.target.value.trim();

        if (e.target.value.trim() !== "")  {
            setInputs((prevInputs) =>
                prevInputs.map((input, i) =>
                    i === index && input.tag.length < 5
                        ? { ...input, tag: [...input.tag, newTag] }
                        : input
                )
            );

            e.target.value = ""; // Clear input field);
        }
    };



    // Remove a single tag
    const removeTag = (inputIndex, tagIndex) => {
        setInputs((prevInputs) =>
            prevInputs.map((input, i) =>
                i === inputIndex
                    ? { ...input, tag: input.tag.filter((_, j) => j !== tagIndex) }
                    : input
            )
        );
    };

    // Remove all tags from a specific input set
    const clearTags = (index) => {
        setInputs((prevInputs) =>
            prevInputs.map((input, i) =>
                i === index ? { ...input, tag: [] } : input
            )
        );
    };

    return (
        <div className=' p-6 rounded-lg w-full flex justify-center items-center'>
            <div ref={modalRef} className="flex flex-col items-center w-full max-w-[900px] bg-white p-6 rounded-lg shadow-lg">
                <div className="flex justify-between w-[850px] max-w-full mx-auto items-center mt-3">
                    <h2 className="text-2xl font-bold ml-2">Add Targets</h2>
                    <button onClick={() => setAddTargets(true)}><IoMdClose className="text-gray-400 text-2xl cursor-pointer" /></button>
                </div>
                <div className="mt-3 w-[850px] max-w-full px-4">
                    <p className="text-gray-400 text-sm text-center sm:text-left">
                        Add targets individually by IP, DNS Name, or URL.
                        Once created, you can configure advanced options for each target.
                    </p>
                </div>
                <form onSubmit={handleSubmit} onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); } }}>
                    {inputs.map((input, index) => (
                        <div className="flex flex-col items-center justify-center sm:flex-row gap-4 w-[850px] max-w-full mt-3" key={index}>
                            <div className="flex items-start flex-col w-full sm:w-auto">
                                <span>Target *</span>
                                <input
                                    type="text"
                                    placeholder="IPv4, Domain, URL or Public CIDR"
                                    className="border py-2 rounded-md px-3 w-full sm:min-w-[310px]"
                                    value={input.target_url}
                                    onChange={(e) => {
                                        let value = e.target.value;

                                        // Remove 'http' or 'https' (case-insensitive)
                                        value = value.replace(/https?:\/\//gi, '');

                                        // Remove special characters except dots, dashes, slashes, and colons
                                        value = value.replace(/[^a-zA-Z0-9./:-]/g, '');

                                        handleInputChange(index, value);
                                    }}
                                />
                            </div>
                            <div className="flex flex-wrap items-start w-full sm:w-auto">
                                <span>Tags</span>
                                <div className="relative w-full flex items-center border py-2 px-3 rounded-md md:w-[180px] cursor-text">
                                    <div className="flex flex-wrap w-[90%] gap-1 truncate">
                                        {input.tag.map((tag, tagIndex) => (
                                            <div key={tagIndex} className="bg-gray-200 px-2 py-1 rounded flex items-center gap-1 justify-around">
                                                <span className="truncate max-w-[80px]">{tag}</span>

                                                <button
                                                    type="button"
                                                    onClick={() => removeTag(index, tagIndex)}
                                                    className="text-red-500 font-bold"
                                                >
                                                    ✖
                                                </button>
                                            </div>
                                        ))}

                                        {input.tag.length < 5 && (
                                            <input ref={inputRef}
                                                type="text"
                                                placeholder="Up to 5 tags"
                                                className="outline-none flex-grow min-w-[80px]"
                                                onKeyDown={(e) => handleKeyDown(index, e)}
                                                onBlur={(e) => handleBlur(index, e)}
                                            />
                                        )}
                                    </div>
                                    <div className="flex flex-wrap w-[10%]">
                                        {input.tag.length > 0 && (
                                            <button
                                                type="button"
                                                onClick={() => clearTags(index)}
                                                className="absolute right-2 text-red-600 font-bold"
                                            >
                                                ✖
                                            </button>
                                        )}
                                    </div>
                                </div>
                            </div>

                            <div className="flex items-start flex-col w-full sm:w-auto">
                                <span>Label</span>
                                <input
                                    type="text"
                                    placeholder="e.g. my-example-server"
                                    className="border py-2 rounded-md px-3 w-full sm:min-w-[310px]"
                                    value={input.label}
                                    onChange={(e) => handleLabelChange(index, e.target.value)}
                                />
                            </div>
                            <hr className="hidden sm:block w-[850px] mt-5" />
                        </div>
                    ))}
                    <div className="mt-5 flex justify-between w-[850px] max-w-full mx-auto">
                        <button className="bg-blue-500 py-3 rounded-md border border-blue-700 text-white font-bold px-3">
                            Import CSV
                        </button>
                        <button
                            type="submit"
                            className="bg-blue-500 py-3 rounded-md border border-blue-700 text-white font-bold px-3"
                        >
                            Submit
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default AddTarget
