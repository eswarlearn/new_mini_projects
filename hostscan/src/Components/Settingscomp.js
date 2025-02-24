import React, { useContext, useRef, useState,useEffect } from 'react'
import { DataContext } from './DataContext';
import { CgProfile } from "react-icons/cg";
import { MdOutlineVerifiedUser } from "react-icons/md";
import { CiCreditCard1 } from "react-icons/ci";
import { LiaUsersSolid } from "react-icons/lia";
import { IoTerminalOutline } from "react-icons/io5";
import { HiOutlineTag } from "react-icons/hi2";
import { TbLayoutGrid } from "react-icons/tb";
import { IoMdClose } from "react-icons/io";
import { BiError } from "react-icons/bi";

const Settingscomp = () => {
  var modalRef=useRef(null)
  var {hiddencontent,setHiddenContent}=useContext(DataContext)
  var {emailId, setEmailId}=useContext(DataContext)
  // var emailId = localStorage.getItem("email");
  var [deleteAcc,setDeleteAcc]=useState(true)
  var [textInput,setTextInput]=useState('')
  var num=0
  var handleInputChange = () => {
    setTextInput('')
    setDeleteAcc(!deleteAcc)
  };
  var buttontypes=[
    {id:1,name:'Account',icon:<CgProfile/>},
    {id:2,name:'Scanners',icon:<MdOutlineVerifiedUser/>},
    {id:3,name:'Billing',icon:<CiCreditCard1/>},
    {id:4,name:'Team',icon:<LiaUsersSolid/>},
    {id:5,name:'API & Webhooks',icon:<IoTerminalOutline/>},
    {id:6,name:'White label',icon:<HiOutlineTag/>},
    {id:7,name:'Workspaces',icon:<TbLayoutGrid/>},
  ]
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (modalRef.current && !modalRef.current.contains(event.target)) {
        setDeleteAcc(true);
        setTextInput('')
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);
      
  return (
    <div className={`mt-20 px-4 flex ${hiddencontent?'md:ml-[200px]':'md:ml-[90px]'} flex-col gap-8 md:mt-32 w-full`}>
      <div className={`fixed top-0 left-0 w-full h-full ${deleteAcc ? 'hidden' : 'block bg-gray-500 bg-opacity-50 z-50 overflow-auto'} rounded-lg flex justify-center items-center`}>
        <div ref={modalRef} className="flex flex-col items-center w-[550px] bg-white overflow-auto p-6 rounded-lg shadow-lg ">
          <div className="flex justify-between w-[850px] max-w-full mx-auto items-center mt-3">
            <div className='flex items-center'>
              <div className='rounded-[100%] bg-rose-200 p-2'><BiError className='text-xl text-red-500'/></div>
              <h2 className="text-lg ml-2">Delete Account</h2>
            </div>
            <button onClick={handleInputChange}><IoMdClose className="text-gray-400 text-2xl cursor-pointer" /></button>
          </div>
          <div className="mt-3 w-[850px] max-w-full sm:text-left px-4">
            <p className="text-gray-400 text-sm text-center sm:text-left">
              This will permanently delete this account and all associated data. Once deleted, this account cannot be recovered.
            </p>
            <p className='mt-5 text-sm font-medium'>To confirm, please type DELETE in the text box below.</p>
          </div>
          <div className='px-5 w-full'><input type="text" value={textInput} onChange={(e)=>setTextInput(e.target.value)} className="border py-2 px-2 w-full rounded-md" /></div>
          <div className="mt-5 flex justify-between w-[850px] max-w-full mx-auto">
            <button className="bg-red-400 py-3 rounded-md border border-red-700 text-white font-bold px-12">
              Permanently Delete
            </button>
            <button onClick={handleInputChange} className="bg-white py-3 rounded-md border border-black  font-bold px-20">
              Cancel
            </button>
          </div>
        </div>
      </div>


      <div className='gap-2 flex md:flex-row flex-col items-start justify-start md:items-center md:justify-between'>
        <div className=''>
          <h1 className='font-bold text-5xl'>Settings</h1>
        </div>
      </div>
      <div className='flex max-md:flex-col mt-6 gap-6'>
        <div className='flex flex-col xl:w-[330px] items-start gap-1'>
          {buttontypes.map((btn,index) => (
            <div key={index} className='max-md:w-full w-full flex'>
              <button className={`flex items-center ${num == index ? 'bg-gray-100' : 'hover:bg-gray-100'} px-3 w-full py-2 rounded-[5px] gap-3 group`}>
                <div className='text-lg text-gray-500 group-hover:text-green-500'>{btn.icon}</div>
                <span className='text-sm font-medium group-hover:text-blue-900'>{btn.name}</span>
                {index >= 4 ? (
                  <button className='text-white bg-emerald-500 px-2 rounded-[4px] h-5 text-xs'>Upgrade</button>
                ):(
                  <span></span>
                )}
              </button>
            </div>
          ))}
        </div>

        <div className={`w-full`}>
          <div className='border flex flex-col gap-2 p-6 rounded-[8px] max-xl:w-full w-[750px]'>
            <div className='flex flex-col items-start gap-2'>
              <h2 className='font-medium text-lg'>Account</h2>
              <p className='text-sm text-gray-500'>Change your account name.</p>
              <span className='mt-2 text-sm text-gray-800 font-medium'>Account Name</span>
              <input type='text' className='border w-full py-2 px-2 rounded-[5px] uppercase font-medium text-[12px]' value={emailId}/>
            </div>
            <div className='flex mt-7 items-end justify-end'>
              <button className='text-white rounded-[5px] bg-black px-5 py-[10px]'>Save</button>
            </div>
          </div>
          <div className='border mt-8 flex flex-col gap-2 p-6 rounded-[8px] max-xl:w-full w-[750px]'>
            <div className='flex flex-col items-start gap-2'>
              <h2 className='font-medium text-lg'>Permanently Delete Account & Data</h2>
              <p className='text-sm flex text-gray-500 text-left'>You can permanently delete your account here. Deleting your account will remove all 
                targets, scans, and risks for this account.
              </p>
            </div>
            <div className='flex mt-7 items-end justify-end'>
              <button onClick={()=>setDeleteAcc(!deleteAcc)} className='text-white rounded-[5px] bg-red-600 border border-stone-300 text-sm font-medium px-5 py-[12px]'>
                Delete Account
              </button>
            </div>
          </div>
        </div>
        {/* <div className='xl:w-full'></div> */}
      </div>
    </div>
  )
}

export default Settingscomp
