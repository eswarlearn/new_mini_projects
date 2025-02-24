import React from 'react'
import { IoIosArrowDown } from "react-icons/io";
import { IoIosArrowForward } from "react-icons/io";
// max-lg:mt-[-160px]
const Navcom = ({clickeve , signinn}) => {
  return (
    <div className='flex flex-col lg:flex-row max-lg:p-6 '>
        <div>
            <button className='inline-flex text-lg rounded-md px-4 py-2 lg:hover:bg-black/30'>
                Use Cases
                <IoIosArrowDown className='mt-1 lg:ml-2 hidden lg:block'/>
                <IoIosArrowForward className='mt-1 ml-1 block lg:hidden text-md '/>               
            </button>
            <hr className='text-white w-full mt-2 block lg:hidden'/>
        </div>
        <a className='rounded-md text-lg px-4 py-2 lg:hover:bg-black/30'>
            Scanners
        </a>
        <hr className='text-white w-full mt-2 block lg:hidden'/>
        <a className='rounded-md text-lg px-4 py-2 lg:hover:bg-black/30'>
            Pricing
        </a>
        <hr className='text-white w-full mt-2 block lg:hidden'/>
        <div>
            <button className='inline-flex text-lg justify-center rounded-md px-4 py-2 lg:hover:bg-black/30'>
                Resources
                <IoIosArrowDown className='mt-1 lg:ml-2 hidden lg:block'/>
                <IoIosArrowForward className='mt-1 ml-1 block lg:hidden text-md '/>
            </button>
            <hr className='text-white w-full mt-2 block lg:hidden'/>
        </div>
        <a className='rounded-md text-lg px-4 py-2 lg:hover:bg-black/30'>
            Blog
        </a>
        <hr className='text-white w-full mt-2 block sm:hidden'/>
        <button onClick={clickeve} className='sm:hidden mt-5 text-left block border-white/50 bg-white text-black font-bold px-5 py-2 rounded-md'>Log in</button>
        <button onClick={signinn} className='sm:hidden mt-5 block text-left whitespace-nowrap border bg-transparent text-[#f8fafc] font-bold px-5 py-2 rounded-md'>Sign Up</button>   
    </div>
  )
}

export default Navcom
