import React, { useContext, useEffect, useState } from 'react'
import { IoIosArrowRoundForward } from "react-icons/io";
import { FcGoogle } from "react-icons/fc";
import { DataContext } from './DataContext';
import { VscCheck } from "react-icons/vsc";
import axios from "axios";

const LogSigncom = () => {

    var [email, setEmail] = useState('')
    var { inn, setInn } = useContext(DataContext)
    var [emailsend, setEmailSend] = useState(true)
    var [emailerror, setEmailError] = useState(true)

    const validateEmail = (email) => {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);  // Regex for valid email
    };
    const validate = () => {
        if (email.trim() !== "") {
            setEmailSend(false);  // Show text below when email is entered
        }
        setEmailError(validateEmail(email))
    };
    useEffect(() => {
        if (email.trim() === "") {
            setEmailError(true)
            setEmailSend(true);  // Hide text below when email is empty
        }
    }, [email]);

    const handleSubmit = async () => {
        if (() => validate()) {
            try {
                const response = await axios.post(
                    "http://localhost:5000/api/auth/register",
                    { email },
                    { withCredentials: true } // This enables sending cookies
                );

                if (response.status === 200) {
                    setEmailSend(email == '' ? true : false)
                    console.log("A link has been sent to your email, and the token is stored in a secure cookie.");
                }
            } catch (error) {
                console.error("Error sending email:", error);
                console.log("Something went wrong! Please try again.");
            }
        }
    };


    return (
        <div className=''>
            <div className='mt-40 lg:mt-12 gap-8 text-left px-6 lg:flex lg:px-20 lg:items-center lg:justify-between' >
                <div className='max-w-[500px]'>
                    <h2 className='text-4xl font-light lg:text-5xl'>See the Power of HostedScan</h2><br />
                    <span className='text-lg'>HostedScan enables companies to meet compliance and security goals.</span>
                </div>

                <div className='border rounded-md bg-stone-100 mt-20 lg:w-[40%]'>
                    <div className='flex flex-col p-6'>
                        <span className=''>Sign {inn == 'up' ? 'up by entering your' : 'in with'} email</span>
                        <input type='text' className='border rounded-[4px] mt-4 h-12 p-4' placeholder='Email' value={email} onChange={(e) => setEmail(e.target.value)} />
                        <span className={`text-red-500 flex mt-2 text-xs font-semibold justify-center items-center ${emailerror ? 'hidden' : 'block'}`}>
                            Invalid email
                        </span>
                        <button onClick={handleSubmit} className={`mt-2 border ${email == '' ? 'border-gray-400' : 'border-green-600'}}
                      rounded-[4px] h-12 font-bold`} style={{ backgroundColor: email == '' ? 'rgb(214, 218, 215)' : 'springgreen' }}>
                            {email != '' ? (
                                emailsend == false ? (
                                    emailerror ? (
                                        <div className='flex items-center justify-center gap-3'>
                                            {'Sent'} <VscCheck className='text-2xl' />
                                        </div>
                                    ) : (
                                        'send me a link'
                                    )

                                ) : (
                                    'send me a link'
                                )
                            ) : (
                                'send me a link'
                            )}
                        </button>
                        <button className={`mt-4 ${emailsend ? 'hidden' : 'block'} ${email == '' ? 'hidden' : 'block'} ${!emailerror ? 'hidden' : 'block'}
                        font-bold border border-green-600 rounded-[4px] flex py-4 gap-5 items-center justify-center bg-transparent text-green-600`}>
                            <VscCheck className='text-2xl' /> A sign in link has been sent to your email.
                        </button>
                        <span className='mt-8'>Or through:</span>
                        <button className='flex items-center justify-between rounded-[4px] bg-white h-13 p-3 mt-3 border'>
                            <div className='flex flex-row'>
                                <FcGoogle className='text-2xl' />
                                <span className='ml-2 text-sm font-semibold'>Sign {inn} with Google</span>
                            </div>
                            <IoIosArrowRoundForward className='text-xl ml-2 ' />
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LogSigncom
