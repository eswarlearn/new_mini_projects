import React, { useState, useEffect, useContext, useRef } from 'react'
import Iconimg from '../Images/download.png'
import { RxHamburgerMenu } from "react-icons/rx";
import { CgProfile } from "react-icons/cg";
import { IoIosArrowDropdownCircle } from "react-icons/io";
import { FaPlus } from "react-icons/fa6";
import DashNav from './DashNav';
import { DataContext } from './DataContext';
import { IoCloseSharp } from "react-icons/io5";
import { LuLogOut } from "react-icons/lu";
import { IoIosArrowDropupCircle } from "react-icons/io";
import axios from "axios";
import { useNavigate } from 'react-router-dom';

const DashHeader = () => {
    var headRef = useRef(null)
    var buttonRef = useRef(null);
    var [emailID, setEmailID] = useState('')
    var { hiddencontent, setHiddenContent } = useContext(DataContext)
    var [logout, setLogout] = useState(true)
    // const history = useNavigate();

    useEffect(() => {
        const email = localStorage.getItem("email"); // Get email from localStorage
        if (email) {
            setEmailID(email); // Set email to the state
        }
        const handleClickOutside = (event) => {
            if (
                headRef.current &&
                !headRef.current.contains(event.target) &&
                buttonRef.current &&
                !buttonRef.current.contains(event.target)
            ) {
                setLogout(true)
            }
        };
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);


    const loggOut = async () => {
        try {
            const token = axios.defaults.headers.common["Authorization"];
            const response = await axios.post(
                "http://localhost:5000/api/auth/logout",
                token,
                {
                    withCredentials: true,
                }
            );

            if (response.status === 200) {
                localStorage.removeItem("email");
                localStorage.removeItem("user_id");
                delete axios.defaults.headers.common["Authorization"];
                window.location.reload();
            }
        } catch (error) {
            console.error("Error sending email:", error);
            console.log("Something went wrong! Please try again.");
        }
    };

    // trial logout
    //  const loggOut = () => {
    //   localStorage.removeItem("email");

    // Optionally, clear Authorization header for Axios globally if set
    //     delete axios.defaults.headers.common["Authorization"];

    // Reload the page to reset the app state
    //    window.location.reload();
    // history.push('/');  // Assuming '/' is the route you want to "reload"  
    //    console.log("stopped");
    //   }
    return (
        <div className={` fixed`}>
            <header className='fixed top-0 z-20 w-full flex justify-center border-b-[1.5px] border-gray-300 backdrop-blur-md'
                style={{ backgroundColor: 'whitesmoke' }}>
                <nav className='flex w-full items-center justify-between py-2 px-4'>
                    <div className='flex gap-3'>
                        <img src={Iconimg} className='h-6 w-6' />
                        <h2 className='font-bold' style={{ fontFamily: 'inherit' }}>HostedScan</h2>
                    </div>
                    <div className={`w-full items-center justify-center lg:flex gap-1 md:hidden block`}>
                        {/* <DashNav /> */}
                    </div>
                    <div className="gap-3 flex flex-row items-center">
                        <div className="relative gap-3 flex">
                            <button className='bg-black text-white px-5 py-2 rounded-[4px] hidden md:block md:flex gap-2 w-auto md:w-max'>
                                <FaPlus className='mt-1' />
                                <span className='text-[15px]'>New Scan</span>
                            </button>
                            <button ref={buttonRef} className='md:border bg-white rounded-[4px] md:flex md:px-3 gap-2 md:py-2 md:items-center md:justify-between w-auto md:w-max'
                                onClick={() => setLogout(!logout)}>
                                <span className='uppercase text-slate-400 text-[12px] hidden md:block'>{emailID}'s account</span>
                                <CgProfile className="text-2xl font-light" />
                                <IoIosArrowDropdownCircle className={`max-md:absolute text-sm md:text-lg max-md:right-2 bg-white rounded-xl max-md:translate-x-3 
                            top-2 ${logout ? 'block' : 'hidden'}`} />
                                <IoIosArrowDropupCircle className={`max-md:absolute text-sm md:text-lg max-md:right-2 bg-white rounded-xl max-md:translate-x-3 
                            top-2 ${logout ? 'hidden' : 'block'}`} />
                            </button>
                        </div>
                        <RxHamburgerMenu className={`text-2xl font-bold sm:mt-2 md:hidden ${hiddencontent ? 'hidden' : 'block'}`} onClick={() => setHiddenContent(!hiddencontent)} />
                        <IoCloseSharp className={`md:hidden text-2xl font-bold ${hiddencontent ? 'block' : 'hidden'}`} onClick={() => setHiddenContent(!hiddencontent)} />
                    </div>
                </nav>
                <div ref={headRef} className={`mt-[46px] sm:mt-[55px] bg-white md:mt-[63px] border rounded-[8px] p-2 w-full md:max-w-[300px] md:right-5 absolute ${logout ? 'hidden' : 'block'}`}>
                    <div className='flex flex-col gap-4 justify-start w-full p-4 items-start'>
                        <div className='flex flex-col gap-2 justify-start w-full items-start'>
                            <span className='uppercase flex items-center w-full justify-between'>
                                organization <IoCloseSharp className='text-xl block md:hidden' onClick={() => setLogout(!logout)} />
                            </span>
                            <span className="bg-gray-300 px-4 py-2 w-full rounded-[5px] overflow-hidden text-ellipsis whitespace-nowrap block cursor-pointer group">
                                {emailID}

                                {/* Tooltip */}
                                <div className="absolute left-[20px] bottom-[40%] mb-1 w-max min-w-[250px] bg-black text-white text-sm rounded-md px-2 py-1 shadow-md opacity-0 group-hover:opacity-100 transition-opacity duration-300 z-10">
                                    {emailID}
                                </div>
                            </span>
                            <hr className='w-full border-gray-300 mb-2 mt-4' />
                        </div>
                        <button className='flex gap-2' onClick={loggOut}>
                            <LuLogOut className='text-xl mt-1' />
                            <span className='text-lg'>Logout</span>
                        </button>
                    </div>
                    <div className='block md:hidden'>
                        {/* <IoCloseSharp className='text-xl' onClick={()=>setLogout(!logout)}/> */}
                    </div>
                </div>
            </header>
            {/* <div className='px-4'> */}

            {/* </div> */}
            {/* <div className={`w-full ${hiddencontent ? 'max-w-[230px]' : 'max-w-[90px]'} ${hiddencontent ? 'max-md:max-w-full' : ''}`}>
            <div className={`mt-10 md:mt-14`}>
                <DashNav/>            
            </div>
        </div> */}
            {/* mt-10 sm:mt-12 md:mt-14*/}
        </div>
    )
}

export default DashHeader
